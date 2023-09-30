package grpcserver

import (
	"fmt"
	"log"
	"net"

	"github.com/fbriansyah/micro-auth-service/internal/port"
	"github.com/fbriansyah/micro-payment-proto/protogen/go/auth"
	"google.golang.org/grpc"
)

type GrpcServerAdapter struct {
	server   *grpc.Server
	service  port.ServicePort
	grpcPort int

	auth.UnimplementedAuthServiceServer
}

// NewGrpcServerAdapter create GrpcServerAdapter server
func NewGrpcServerAdapter(service port.ServicePort, grpcPort int) *GrpcServerAdapter {
	return &GrpcServerAdapter{
		grpcPort: grpcPort,
		service:  service,
	}
}

// Run grpc server
func (a *GrpcServerAdapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf("%d", a.grpcPort))

	if err != nil {
		log.Fatalf("failed to listen on port %d: %v\n", a.grpcPort, err)
	}
	log.Printf("Server listen on port %d \n", a.grpcPort)
	grpcServer := grpc.NewServer()

	a.server = grpcServer

	auth.RegisterAuthServiceServer(grpcServer, a)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to server grpc on port %d: %v\n", a.grpcPort, err)
	}
}

// Stop the grpc server
func (a *GrpcServerAdapter) Stop() {
	a.server.Stop()
}
