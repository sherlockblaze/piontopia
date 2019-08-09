package service

import (
	"context"

	"github.com/sherlockblaze/piontopia/xtopia-user/proto"
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
