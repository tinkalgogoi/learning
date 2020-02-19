package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"learning/context/grpc/pb"
	"fmt"
	"time"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context has been cancelled from client...", ctx.Err())
		default:
			// say long running code
			for i:= 0 ; i < 10 ; i++ {
				fmt.Println("second elapsed: ", i)
				time.Sleep(1 * time.Second)
			}
			fmt.Println("received: ", in.Name)
		}
	}

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
