package main

import (
	"context"
	"fmt"
	helper_grpc "github.com/langwan/langgo/helpers/grpc"
	"google.golang.org/grpc"
	"langwan/langgo-examples/Advanced/Grpc/Single-Service/client/pb"
)

func main() {

	conn, err := helper_grpc.NewClient(nil, "127.0.0.1:8000", grpc.WithInsecure())

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	ServerClient := pb.NewServerClient(conn)

	helloResponse, err := ServerClient.Hello(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	fmt.Println(helloResponse)
}
