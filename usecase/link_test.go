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

func (m *mockLinkRepo) Create(ctx context.Context, long string) (*link.Link, error) {
	_, ok := lo.Find(m.links, func(l *link.Link) bool { return l.Long() == long })
	if ok {
		return nil, fmt.Errorf("link already exists %s", long)
	}

	m.maxID++
	l, _ := link.NewLink(m.maxID, long, time.Now())
	m.links = append(m.links, l)

	return l, nil
}

func (m *mockLinkRepo) GetOne(ctx context.Context, long string) (*link.Link, error) {
	l, ok := lo.Find(m.links, func(l *link.Link) bool { return l.Long() == long })
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

func TestService_Create(t *testing.T) {
	baseURL := "http://localhost/"
	long := "https://jsonplaceholder.typicode.com/albums"
	existingLink, _ := link.NewLink(1, long, time.Now())

	tests := []struct {
		name    string
		svc     linkService
		long    string
		wantErr bool
	}{
		{
			name:    "create link",
			svc:     NewLinkService(&mockLinkRepo{}, baseURL),
			long:    long,
			wantErr: false,
		},
		{
			name:    "error already exists",
			svc:     NewLinkService(&mockLinkRepo{links: []*link.Link{existingLink}}, baseURL),
			long:    long,
			wantErr: true,
		},
		{
			name:    "error empty link",
			svc:     NewLinkService(&mockLinkRepo{}, baseURL),
			long:    "",
			wantErr: true,
		},
		{
			name:    "create link with whitespace",
			svc:     NewLinkService(&mockLinkRepo{}, baseURL),
			long:    fmt.Sprintf("   %s   ", long),
			wantErr: false,
		},
		{
			name:    "error is not uri",
			svc:     NewLinkService(&mockLinkRepo{}, baseURL),
			long:    "isnoturi.com",
			wantErr: false,
		},
		{
			name:    "errror long is greater than max len",
			svc:     NewLinkService(&mockLinkRepo{}, baseURL),
			long:    createUriOverMaxLen(link.MaxLen, baseURL),
			wantErr: false,
		},
	}
	for _, tt := range tests {
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
			assert.True(t, govalidator.IsAlphanumeric(short))
			// redirects to original long uri
			gotInitialLong, err := tt.svc.Redirect(context.Background(), &pb.RedirectRequest{
				ShortPath: short,
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.long, gotInitialLong.LongUri)
		})
	}
}
