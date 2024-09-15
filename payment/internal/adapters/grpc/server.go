package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/andrei-kozel/microservices-demo-proto/golang/payment"
	"github.com/andrei-kozel/microservices-demo/payment/config"
	"github.com/andrei-kozel/microservices-demo/payment/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	payment.UnimplementedPaymentServer
	api    ports.APIPort
	server *grpc.Server
	port   int
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	payment.RegisterPaymentServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
