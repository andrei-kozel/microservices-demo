package main

import (
	"log"

	"github.com/andrei-kozel/microservices-demo/payment/config"
	"github.com/andrei-kozel/microservices-demo/payment/internal/adapters/db"
	"github.com/andrei-kozel/microservices-demo/payment/internal/adapters/grpc"
	"github.com/andrei-kozel/microservices-demo/payment/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to create db adapter: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
