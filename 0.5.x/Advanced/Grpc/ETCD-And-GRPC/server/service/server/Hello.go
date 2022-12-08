package server

import (
	"context"
	"langwan/langgo-examples/Advanced/Grpc/ETCD-And-GRPC/server/pb"
)

func (s Server) Hello(ctx context.Context, empty *pb.Empty) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Msg: "hello"}, nil
}
