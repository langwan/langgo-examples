package server

import (
	"context"
	"langwan/langgo-examples/Advanced/Grpc/Single-Service/server/pb"
)

func (s Server) Hello(ctx context.Context, request *pb.Empty) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Msg: "hello"}, nil
}
