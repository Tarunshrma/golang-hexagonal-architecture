package main

import (
	"github.com/Tarunshrma/golang-hexagonal-architecture/internal/adapters/app/api"
	"github.com/Tarunshrma/golang-hexagonal-architecture/internal/adapters/core/arithmetic"
	"github.com/Tarunshrma/golang-hexagonal-architecture/internal/adapters/framework/left/grpc"
	"github.com/Tarunshrma/golang-hexagonal-architecture/internal/adapters/framework/right/db"
	"log"
	"os"
)

func main() {

	var err error

	//Get the environment variables
	dbDriver := os.Getenv("DB_DRIVER")
	dbSourceName := os.Getenv("DS_NAME")

	//DB Adapter
	dbAdapter, err := db.NewAdapter(dbDriver, dbSourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbAdapter.CloseDBConnection()

	//Core Adapter
	coreAdapter := arithmetic.NewAdapter()

	//flow of direction from grpAdapter--> App Port --> Core Port
	//									   App Port --> DB Port

	//
	appAdapter := api.NewAdapter(coreAdapter, dbAdapter)
	gRPCAdapter := grpc.NewAdapter(appAdapter)

	gRPCAdapter.Run()

}
