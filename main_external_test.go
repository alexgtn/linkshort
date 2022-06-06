package main_test

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	link "github.com/alexgtn/go-linkshort/proto"
)

// Burst-test suite. Launches 100x requests (create, redirect, create & redirect) concurrently.
func Test_Main(t *testing.T) {
	flag.Parse()

	if testing.Short() {
		log.Printf("Skipping tests because of Short flag")
		return
	}

	clientReqCount := 100

	long := fmt.Sprintf("https://test.com/%s", uuid.NewString())
	log.Printf("Long: %s", long)

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	t.Run("create and redirect", func(t *testing.T) {
		c := link.NewLinkshortServiceClient(conn)

		var wg sync.WaitGroup

		wg.Add(clientReqCount)

		for i := 0; i < clientReqCount; i++ {
			go func(long string, i int) {
				defer wg.Done()

				// Create link and redirect
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()

				l, err := c.CreateLink(ctx, &link.CreateLinkRequest{LongUri: long})
				if err != nil {
					log.Fatalf("%d could not create link: %v", i, err)
				}
				shortPath := strings.TrimPrefix(l.GetShortUri(), "http://localhost:8080/")

				r, err := c.Redirect(ctx, &link.RedirectRequest{ShortPath: shortPath})
				if err != nil {
					log.Fatalf("%d could not create link: %v", i, err)
				}

				assert.Equal(t, long, r.LongUri)
				log.Printf("%d short Link: %s, redirected successfully", i, l.GetShortUri())
			}(long, i)
		}
		wg.Wait()
	})

	t.Run("redirect", func(t *testing.T) {
		c := link.NewLinkshortServiceClient(conn)
		// Create link and redirect
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		l, err := c.CreateLink(ctx, &link.CreateLinkRequest{LongUri: long})
		if err != nil {
			log.Fatalf("could not create link: %v", err)
		}
		shortPath := strings.TrimPrefix(l.GetShortUri(), "http://localhost:8080/")

		var wg sync.WaitGroup

		wg.Add(clientReqCount)

		for i := 0; i < clientReqCount; i++ {
			go func(long string, i int) {
				defer wg.Done()

				// Redirect
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()

				r, err := c.Redirect(ctx, &link.RedirectRequest{ShortPath: shortPath})
				if err != nil {
					log.Fatalf("%d could not create link: %v", i, err)
				}

				assert.Equal(t, long, r.LongUri)
				log.Printf("%d short Link: %s, redirected successfully", i, l.GetShortUri())
			}(long, i)
		}
		wg.Wait()
	})

	t.Run("create", func(t *testing.T) {
		c := link.NewLinkshortServiceClient(conn)

		var wg sync.WaitGroup

		wg.Add(clientReqCount)

		for i := 0; i < clientReqCount; i++ {
			go func(long string, i int) {
				defer wg.Done()

				// Create link and redirect
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()

				l, err := c.CreateLink(ctx, &link.CreateLinkRequest{LongUri: long})
				if err != nil {
					log.Fatalf("%d could not create link: %v", i, err)
				}
				log.Printf("%d short Link: %s", i, l.GetShortUri())
			}(long, i)
		}
		wg.Wait()
	})
}
