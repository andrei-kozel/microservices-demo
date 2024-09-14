package main

import (
	"log"

	"github.com/andrei-kozel/microservices-demo/config"
	"github.com/andrei-kozel/microservices-demo/internal/adapters/db"
	"github.com/andrei-kozel/microservices-demo/internal/adapters/grpc"
	"github.com/andrei-kozel/microservices-demo/internal/adapters/payment"
	"github.com/andrei-kozel/microservices-demo/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to create db adapter: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("failed to create payment adapter: %v", err)
	}

	application := api.NewApplication(
		dbAdapter,
		paymentAdapter,
	)

	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
