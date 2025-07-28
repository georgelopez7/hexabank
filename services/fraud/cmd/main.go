package main

import (
	"context"
	"hexabank/api/proto/fraud"
	"hexabank/internal/observability/metrics"
	"hexabank/internal/observability/tracing"
	"hexabank/services/fraud/adapters/grpc"
	"hexabank/services/fraud/domain/service"
	"log"
	"net"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	g "google.golang.org/grpc"
)

func main() {
	// TRACING
	tracerName := "fraud-service"
	shutdown := tracing.NewTracer(tracerName)
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Fatalf("Failed to shut down tracer: %v", err)
		}
	}()

	// METRICS
	metrics.NewtMetricsEndpoint()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// GRPC SERVER
	server := g.NewServer(
		g.StatsHandler(otelgrpc.NewServerHandler()),
	)

	fraudService := service.NewFraudService()
	fraudGRPC := grpc.NewFraudGRPC(fraudService)
	fraud.RegisterFraudServiceServer(server, fraudGRPC)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
