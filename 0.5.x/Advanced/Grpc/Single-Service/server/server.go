package main

import (
	"github.com/langwan/langgo"
	helper_grpc "github.com/langwan/langgo/helpers/grpc"
	"langwan/langgo-examples/Advanced/Grpc/Single-Service/server/pb"
	"langwan/langgo-examples/Advanced/Grpc/Single-Service/server/service/server"
)

const addr = "localhost:8000"

func main() {
	langgo.Run()
	cg := helper_grpc.NewServer(nil)
	cg.Use(helper_grpc.LogUnaryServerInterceptor())
	gs, err := cg.Server()
	if err != nil {
		panic(err)
	}
	pb.RegisterServerServer(gs, server.Server{})
	err = cg.Run(addr)
	if err != nil {
		panic(err)
	}
}
