package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"

	http2 "github.com/alexgtn/go-linkshort/internal/delivery/http"
	"github.com/alexgtn/go-linkshort/tools/proto"
)

var httpPort = flag.Int("http-port", 8080, "The HTTP server port")

// httpCmd starts the gRPC gateway
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "gRPC HTTP gateway",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		// Register gRPC server endpoint
		mux := runtime.NewServeMux(
			runtime.WithForwardResponseOption(redirectResponseModifier),
		)
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		// assuming here the gRPC server is running on the same host, hence localhost
		err := link.RegisterLinkshortServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", *grpcPort), opts)
		if err != nil {
			log.Fatalf("failed to register gRPC gateway: %v", err)
		}

		// serve documentation
		err = mux.HandlePath("GET", "/api/docs", Handle(http2.ServeDocs))
		if err != nil {
			log.Fatalf("failed to register docs handler: %v", err)
		}

		// Start HTTP server (and proxy calls to gRPC server endpoint)
		err = http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), mux)
		if err != nil {
			log.Fatalf("failed to start gRPC gateway: %v", err)
		}
	},
}

// Handle takes basic http handle and returns gateway handle
func Handle(handlerFunc http.HandlerFunc) runtime.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		handlerFunc(w, r)
	}
}

// redirectResponseModifier sets HTTP redirect headers for the redirect response
func redirectResponseModifier(_ context.Context, w http.ResponseWriter, p proto.Message) error {
	redirectReply, ok := p.(*link.RedirectReply)
	if ok {
		w.Header().Set("Location", hexEscapeNonASCII(redirectReply.LongUri))
		w.WriteHeader(http.StatusMovedPermanently)

		_, err := w.Write([]byte(redirectReply.LongUri))
		if err != nil {
			log.Printf("error writing redirect response: %v", err)
		}
	}

	return nil
}

// hexEscapeNonASCII ripped from http.Redirect implementation :P
func hexEscapeNonASCII(s string) string {
	newLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			newLen += 3
		} else {
			newLen++
		}
	}
	if newLen == len(s) {
		return s
	}
	b := make([]byte, 0, newLen)
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			b = append(b, '%')
			b = strconv.AppendInt(b, int64(s[i]), 16)
		} else {
			b = append(b, s[i])
		}
	}
	return string(b)
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
