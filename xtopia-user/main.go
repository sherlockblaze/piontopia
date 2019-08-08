package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sherlockblaze/piontopia/xtopia-user/proto"
	"github.com/sherlockblaze/piontopia/xtopia-user/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3821))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("successfully listening on port: %d", 3821)

	s := service.Server{}
	grpcServer := grpc.NewServer()

	proto.RegisterUserServicesServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
