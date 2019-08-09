package server

import (
	"context"
	"net"

	"github.com/sherlockblaze/piontopia/xtopia-user/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address listening address of server
	Address = "127.0.0.1:50011"
)

// Server represents the gRPC server
type Server struct {
}

// CreateUser create a user
func (s *Server) CreateUser(ctx context.Context, user *proto.User) (*proto.User, error) {
	return nil, nil
}

// DeleteUser delete a user
func (s *Server) DeleteUser(ctx context.Context, userID *proto.UserID) (*proto.User, error) {
	return nil, nil
}

// CheckoutUser checkout a user's info
func (s *Server) CheckoutUser(ctx context.Context, userID *proto.UserID) (*proto.User, error) {
	return nil, nil
}

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

	s := Server{}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	proto.RegisterUserServicesServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}
}
