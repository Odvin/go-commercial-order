package main

import (
	"log"

	"github.com/Odvin/go-commercial-order/adapter/db"
	"github.com/Odvin/go-commercial-order/adapter/grpc"
	"github.com/Odvin/go-commercial-order/adapter/payment"
	"github.com/Odvin/go-commercial-order/application/core"
	"github.com/Odvin/go-commercial-order/config"
)

func main() {
	dbAdapter, err := db.InitAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.InitAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := core.InitApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.InitAdapter(application, config.GetApplicationPort())

	grpcAdapter.Run()
}
