package usecase

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	"github.com/alexgtn/go-linkshort/domain/link"
	pb "github.com/alexgtn/go-linkshort/proto"
)

type mockLinkRepo struct {
	maxID int
	links []*link.Link
}

func newMockRepo() *mockLinkRepo {
	return &mockLinkRepo{}
}

func newMockRepoWithSeed(links ...*link.Link) *mockLinkRepo {
	var lnks []*link.Link

	for _, l := range links {
		lnks = append(lnks, l)
	}

	return &mockLinkRepo{links: lnks}
}

func (m *mockLinkRepo) Create(ctx context.Context, long string) (*link.Link, error) {
	_, ok := lo.Find(m.links, func(l *link.Link) bool { return l.LongURL() == long })
	if ok {
		return nil, fmt.Errorf("link already exists %s", long)
	}

	m.maxID++
	l, _ := link.NewLink(m.maxID, long, time.Now())
	m.links = append(m.links, l)

	return l, nil
}

func (m *mockLinkRepo) GetOne(ctx context.Context, long string) (*link.Link, error) {
	l, ok := lo.Find(m.links, func(l *link.Link) bool { return l.LongURL() == long })
	if !ok {
		return nil, fmt.Errorf("link not found %s", long)
	}
	return l, nil
}

type linkService interface {
	Redirect(context.Context, *pb.RedirectRequest) (*pb.RedirectReply, error)
	Create(context.Context, *pb.CreateLinkRequest) (*pb.CreateLinkReply, error)
}

func createUriOverMaxLen(maxLen int, baseURL string) string {
	e := make([]rune, maxLen+1)
	for i, _ := range e {
		e[i] = 'e'
	}
	return baseURL + string(e)
}

var (
	baseURL         = "http://localhost/"
	long            = "https://jsonplaceholder.typicode.com/albums"
	short           = "abcde"
	existingLink, _ = link.NewLink(1, long, time.Now())
)

func TestService_Create(t *testing.T) {
	// individual link test
	invididualLinkTests := []struct {
		name    string
		svc     linkService
		long    string
		wantErr bool
	}{
		{
			name:    "create link",
			svc:     NewLinkService(newMockRepo(), baseURL),
			long:    long,
			wantErr: false,
		},
		{
			name:    "return existing link",
			svc:     NewLinkService(newMockRepoWithSeed(existingLink), baseURL),
			long:    long,
			wantErr: false,
		},
		{
			name:    "create link with whitespace",
			svc:     NewLinkService(newMockRepo(), baseURL),
			long:    fmt.Sprintf("   %s   ", long),
			wantErr: false,
		},
		{
			name:    "error empty link",
			svc:     NewLinkService(newMockRepo(), baseURL),
			long:    "",
			wantErr: true,
		},
		{
			name:    "error is not uri",
			svc:     NewLinkService(newMockRepo(), baseURL),
			long:    "isnoturi.com",
			wantErr: true,
		},
		{
			name:    "errror long is greater than max len",
			svc:     NewLinkService(newMockRepo(), baseURL),
			long:    createUriOverMaxLen(link.MaxLen, baseURL),
			wantErr: true,
		},
	}
	for _, tt := range invididualLinkTests {
		t.Run(tt.name, func(t *testing.T) {
			gotLink, err := tt.svc.Create(context.Background(), &pb.CreateLinkRequest{
				LongUri: tt.long,
			})
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			// correct prefix
			assert.True(t, strings.HasPrefix(gotLink.ShortUri, baseURL))
			// is alphanumeric
			short := strings.TrimPrefix(gotLink.ShortUri, baseURL)
			assert.NotEmpty(t, short)
			assert.True(t, govalidator.IsAlphanumeric(short))
			// redirects to original long uri
			gotInitialLong, err := tt.svc.Redirect(context.Background(), &pb.RedirectRequest{
				ShortPath: short,
			})
			assert.NoError(t, err)
			assert.Equal(t, long, gotInitialLong.LongUri)
		})
	}

	t.Run("multiple inserts", func(t *testing.T) {
		svc := NewLinkService(newMockRepo(), baseURL)
		_, err := svc.Create(context.Background(), &pb.CreateLinkRequest{
			LongUri: long + "1",
		})
		assert.NoError(t, err)
		_, err = svc.Create(context.Background(), &pb.CreateLinkRequest{
			LongUri: long + "2",
		})
		assert.NoError(t, err)
		_, err = svc.Create(context.Background(), &pb.CreateLinkRequest{
			LongUri: long + "3",
		})
		assert.NoError(t, err)
	})

	t.Run("return existing link", func(t *testing.T) {
		repoWithLink := newMockRepoWithSeed(existingLink)

		svc := NewLinkService(repoWithLink, baseURL)
		l, err := svc.Create(context.Background(), &pb.CreateLinkRequest{
			LongUri: long,
		})
		assert.NoError(t, err)
		assert.Equal(t, l.ShortUri, existingLink.ShortPath())
	})
}

func TestService_Redirect(t *testing.T) {
	repoWithLink := newMockRepoWithSeed(existingLink)

	tests := []struct {
		name    string
		svc     linkService
		short   string
		wantErr bool
	}{
		{
			name:    "redirect",
			svc:     NewLinkService(repoWithLink, baseURL),
			short:   short,
			wantErr: false,
		},
		{
			name:    "redirect with whitespace",
			svc:     NewLinkService(newMockRepo(), baseURL),
			short:   fmt.Sprintf("   %s   ", short),
			wantErr: false,
		},
		{
			name:    "error empty",
			svc:     NewLinkService(newMockRepo(), baseURL),
			short:   "",
			wantErr: true,
		},
		{
			name:    "error doesn't exist",
			svc:     NewLinkService(newMockRepo(), baseURL),
			short:   "asdasdad",
			wantErr: true,
		},
		{
			name:    "error is not alpha numeric",
			svc:     NewLinkService(repoWithLink, baseURL),
			short:   "short!#?",
			wantErr: true,
		},
		{
			name:    "errror long is greater than max len",
			svc:     NewLinkService(newMockRepo(), baseURL),
			short:   createUriOverMaxLen(link.MaxLen, baseURL),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLink, err := tt.svc.Create(context.Background(), &pb.CreateLinkRequest{
				LongUri: tt.short,
			})
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			// correct prefix
			assert.True(t, strings.HasPrefix(gotLink.ShortUri, baseURL))
			// is alphanumeric
			short := strings.TrimPrefix(gotLink.ShortUri, baseURL)
			assert.True(t, govalidator.IsAlphanumeric(short))
			// redirects to original long uri
			gotInitialLong, err := tt.svc.Redirect(context.Background(), &pb.RedirectRequest{
				ShortPath: short,
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.short, gotInitialLong.LongUri)
		})
	}
}
