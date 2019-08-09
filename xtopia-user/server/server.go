package server

import (
	"net"

	"github.com/sherlockblaze/piontopia/xtopia-user/proto"
	"github.com/sherlockblaze/piontopia/xtopia-user/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address listening address of server
	Address = "127.0.0.1:50011"
)

// NewServer Start grpc server
func NewServer() {
	lis, err := net.Listen("tcp", Address)

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpclog.Printf("successfully listening on port: %d", 50011)

	// credentials.NewClientTLSFromFile("./keys/server.pem", "sherlockblaze")
	creds, err := credentials.NewServerTLSFromFile("./keys/server.pem", "./keys/server.key")

	if err != nil {
		grpclog.Fatalf("failed to generate credentials %v", err)
	}

	s := service.Server{}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	proto.RegisterUserServicesServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}
}
