package main

import (
	"context"
	"google.golang.org/grpc"
	pb "learn_gRPC/hello-server/proto"
	"net"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "Hello " + req.GetRequestName()}, nil
}

func main() {
	listen, _ := net.Listen("tcp", ":9090")
	grpc := grpc.NewServer()
	pb.RegisterSayHelloServer(grpc, &server{})
	err := grpc.Serve(listen)
	if err != nil {
		panic(err)
	}
}
