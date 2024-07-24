package main

import (
	"hex_arch_go/internal/adapters/app/api"
	"hex_arch_go/internal/adapters/core/arithmetic"
	gRPC "hex_arch_go/internal/adapters/framework/left/grpc"
	"hex_arch_go/internal/adapters/framework/right/db"
	"hex_arch_go/internal/ports"
	"log"
	"os"
)

func main() {
	// ports
	var dbaseAdapter ports.DBPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dbaseSource := os.Getenv("DB_SOURCE")

	dbaseAdapter, err := db.NewAdapter(dbaseDriver, dbaseSource)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDBConn()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
