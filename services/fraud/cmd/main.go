package main

import (
	"hexabank/api/proto/fraud"
	"hexabank/services/fraud/adapters/grpc"
	"hexabank/services/fraud/domain/service"
	"log"
	"net"

	g "google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := g.NewServer()
	fraudService := &service.FraudService{}
	fraudGRPC := grpc.NewFraudGRPC(fraudService)
	fraud.RegisterFraudServiceServer(server, fraudGRPC)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
