package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/Odvin/go-commercial-order/config"
	"github.com/Odvin/go-commercial-proto/golang/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()

	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	log.Printf("starting order service on port %d ...", a.port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}
