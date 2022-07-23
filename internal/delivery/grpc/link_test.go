package grpc

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexgtn/go-linkshort/internal/domain/link"
	pb "github.com/alexgtn/go-linkshort/tools/proto"
)

type linkDelivery interface {
	Redirect(context.Context, *pb.RedirectRequest) (*pb.RedirectReply, error)
	CreateLink(context.Context, *pb.CreateLinkRequest) (*pb.CreateLinkReply, error)
}

type linkServiceMock struct {
}

func (l *linkServiceMock) Redirect(ctx context.Context, shortPath string) (string, error) {
	return "http://test.com/long-link", nil
}

func (l *linkServiceMock) CreateLink(ctx context.Context, longURL string) (string, error) {
	return "shortPath", nil
}

func newLinkServiceMock() *linkServiceMock {
	return &linkServiceMock{}
}

func createUriOverMaxLen(maxLen int, baseURL string) string {
	e := make([]rune, maxLen+1)
	for i, _ := range e {
		e[i] = 'e'
	}
	return baseURL + string(e)
}

var (
	baseURL = "http://localhost"
	long    = "https://jsonplaceholder.typicode.com/albums"
)

func TestService_Create(t *testing.T) {
	// individual link test
	invididualLinkTests := []struct {
		name    string
		svc     linkDelivery
		long    string
		wantErr bool
	}{
		{
			name:    "create link with whitespace",
			svc:     NewLinkDeliveryGrpc(newLinkServiceMock()),
			long:    fmt.Sprintf("   %s   ", long),
			wantErr: true,
		},
		{
			name:    "error empty link",
			svc:     NewLinkDeliveryGrpc(newLinkServiceMock()),
			long:    "",
			wantErr: true,
		},
		{
			name:    "error is not uri",
			svc:     NewLinkDeliveryGrpc(newLinkServiceMock()),
			long:    "isnoturi.com",
			wantErr: true,
		},
		{
			name:    "errror long is greater than max len",
			svc:     NewLinkDeliveryGrpc(newLinkServiceMock()),
			long:    createUriOverMaxLen(link.MaxLen, baseURL),
			wantErr: true,
		},
	}
	for _, tt := range invididualLinkTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.svc.CreateLink(context.Background(), &pb.CreateLinkRequest{
				LongUri: tt.long,
			})
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestService_Redirect(t *testing.T) {
	tests := []struct {
		name    string
		svc     linkDelivery
		short   string
		wantErr bool
	}{
		{
			name:    "error empty",
			svc:     NewLinkDeliveryGrpc(newLinkServiceMock()),
			short:   "",
			wantErr: true,
		},
		{
			name:    "error is not alpha numeric",
			svc:     NewLinkDeliveryGrpc(newLinkServiceMock()),
			short:   "short!#?",
			wantErr: true,
		},
		{
			name:    "errror long is greater than max len",
			svc:     NewLinkDeliveryGrpc(newLinkServiceMock()),
			short:   createUriOverMaxLen(link.MaxLen, baseURL),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLink, err := tt.svc.Redirect(context.Background(), &pb.RedirectRequest{
				ShortPath: tt.short,
			})
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			// returned original link
			assert.Equal(t, gotLink.LongUri, long)
		})
	}
}
