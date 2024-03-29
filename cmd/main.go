package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"

	delivery "github.com/alexgtn/go-linkshort/internal/delivery/grpc"
	"github.com/alexgtn/go-linkshort/internal/infra/repository/link"
	"github.com/alexgtn/go-linkshort/internal/infra/sqlite"
	"github.com/alexgtn/go-linkshort/internal/usecase"
	ent "github.com/alexgtn/go-linkshort/tools/ent/codegen"
	pb "github.com/alexgtn/go-linkshort/tools/proto"
)

var grpcPort = flag.Int("port", 50051, "The server port")

// mainCmd starts the gRPC server
var mainCmd = &cobra.Command{
	Use:   "main",
	Short: "gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("main called")

		client := sqlite.OpenEnt(cfg.DatabaseURL)
		defer func(client *ent.Client) {
			err := client.Close()
			if err != nil {
				log.Fatal("error closing client")
			}
		}(client)
		flag.Parse()

		// logging
		zapconfig := zap.NewProductionConfig()
		zapconfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		logger, err := zapconfig.Build()
		if err != nil {
			log.Fatal(err)
		}

		// always log req/res payload
		alwaysLoggingDeciderServer := func(ctx context.Context, fullMethodName string, servingObject interface{}) bool { return true }

		// grpc server with middleware
		s := grpc.NewServer(grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(logger),
				grpc_zap.PayloadUnaryServerInterceptor(logger, alwaysLoggingDeciderServer),
			),
		))

		// dependency injection
		linkRepo := link.NewLinkRepo(client)
		linkUsecase := usecase.NewLinkService(linkRepo, cfg.BaseURL)
		linkDelivery := delivery.NewLinkDeliveryGrpc(linkUsecase)
		pb.RegisterLinkshortServiceServer(s, linkDelivery)

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Printf("server listening at %v", lis.Addr())

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(mainCmd)
}
