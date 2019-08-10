package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"piontopia/xtopia-services/etcdservice"
	pb "piontopia/xtopia-services/proto"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

var host = "127.0.0.1"

var (
	// ServiceName service name
	ServiceName = flag.String("ServiceName", "hello_service", "service name")
	// Port port
	Port = flag.Int("Port", 50001, "listening port")
	// EtcdAddr etcd address
	EtcdAddr = flag.String("EtcdAddr", "127.0.0.1:2379", "register etcd address")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *Port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	} else {
		fmt.Printf("listen at:%d\n", *Port)
	}
	defer lis.Close()

	s := grpc.NewServer()
	defer s.GracefulStop()

	pb.RegisterGreeterServer(s, &server{})
	addr := fmt.Sprintf("%s:%d", host, *Port)
	fmt.Printf("server addr:%s\n", addr)
	go etcdservice.Register(*EtcdAddr, *ServiceName, addr, 5)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		etcdservice.UnRegister(*ServiceName, addr)

		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}

	}()

	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %s", err)
		//log.Fatalf("failed to serve: %s", err)
	}
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("%v: Receive is %s\n", time.Now(), in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
