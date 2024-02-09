package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "root/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 3001, "The server port")
)

type server struct {
	pb.UnimplementedPostsServiceServer
}

func (s *server) HelloWorld(ctx context.Context, in *pb.HelloWorldResponse) (*pb.HelloWorldResponse, error) {
	log.Printf("Received")
	return &pb.HelloWorldResponse{Message: in.Message}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	pb.RegisterPostsServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(lis); e != nil {
		panic(err)
	}
}
