package main

import (
	"context"

	app "github.com/joaocarvalhodb1/arch-poc/src/app/account_service"
	domain "github.com/joaocarvalhodb1/arch-poc/src/domain/account_service"
	infra "github.com/joaocarvalhodb1/arch-poc/src/infra/database/account_service"
	"github.com/joaocarvalhodb1/arch-poc/src/shared/grpc"
	"github.com/joaocarvalhodb1/arch-poc/src/shared/helpers"
	"github.com/joaocarvalhodb1/arch-poc/src/shared/sql/driver"
)

const (
	serviceName = "account service"
	portService = "5055"
	DSN         = "host=localhost port=5432 user=postgres password=postgres dbname=db_account sslmode=disable"
)

func main() {
	log := helpers.NewLoggers(serviceName)
	ctx := context.Background()
	db, err := driver.NewPostgresSQLDriver(DSN, log, ctx)
	if err != nil {
		log.Fatal("Error in the connection: ", err)
	}
	defer db.Close()
	accountRepo := infra.NewAccountRepoPostgres(db, log)
	creditLimitService := domain.NewApplyCreditLimit(accountRepo)
	accountAppService := app.NewAccountAppService(accountRepo, creditLimitService, log)
	gRPCServer, err := grpc.NewgRPCServer(portService, log)
	if err != nil {
		log.Fatal("Error create gRPC server: ", err)
	}
	accountAppService.RegisterServiceServer(gRPCServer.Server)
	gRPCServer.ListenAndServe()
}
