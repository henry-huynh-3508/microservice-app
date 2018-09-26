package main

import (
	"context"
	"fmt"
	"os"

	pb "microservice-app/gRPC"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	args := os.Args
	conn, err := grpc.Dial("192.168.99.100:5300", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewReverseClient(conn)
	request := &pb.Request{
		Message: args[1],
	}
	response, err := client.Do(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	fmt.Println(response.Message)
}
