package service

import (
	"context"
	"log"

	"github.com/sherlockblaze/piontopia/xtopia-user/proto"
)

// Server represents the gRPC server
type Server struct {
}

// CreateUser create user
func (s *Server) CreateUser(ctx context.Context, in *proto.User) (*proto.UserID, error) {
	log.Printf("Receive message %s", in.Name)
	return &proto.UserID{ID: "007", Status: 0}, nil
}

// DeleteUser delete a user
func (s *Server) DeleteUser(ctx context.Context, in *proto.UserID) (*proto.UserID, error) {
	return nil, nil
}

// CheckoutUser checkout a user's info
func (s *Server) CheckoutUser(ctx context.Context, in *proto.UserID) (*proto.User, error) {
	return nil, nil
}
