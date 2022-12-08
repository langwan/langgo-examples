package main

import (
	"flag"
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
	"langwan/langgo-examples/Advanced/Grpc/ETCD-And-GRPC/server/service/server"

	"langwan/langgo-examples/Advanced/Grpc/ETCD-And-GRPC/server/pb"

	helper_grpc "github.com/langwan/langgo/helpers/grpc"
	"os"

	"syscall"
)

const (
	etcdHost    = "localhost:2379"
	serviceName = "langgo/server"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8001, "port")
	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", port)
	langgo.Run()
	langgo.SignalHandlers(func(sig os.Signal) {
		helper_grpc.EtcdUnRegister(serviceName, addr)
		os.Exit(int(syscall.SIGINT))
	}, syscall.SIGINT)
	defer func() {
		core.DeferRun()
	}()

	helper_grpc.EtcdRegister(etcdHost, serviceName, addr, 50)

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
