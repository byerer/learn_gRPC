package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "learn_gRPC/hello-server/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// Create a client
	client := pb.NewSayHelloClient(conn)
	// Call the remote method
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "gRPC"})

	fmt.Println(resp.GetResponseMsg())
}
