package main

import (
	"fmt"

	accountapi "github.com/joaocarvalhodb1/arch-poc/src/app/account_service_api"
	ws "github.com/joaocarvalhodb1/arch-poc/src/shared/ws"

	"github.com/joaocarvalhodb1/arch-poc/src/shared/contracts/protog"
	"github.com/joaocarvalhodb1/arch-poc/src/shared/grpc"
	"github.com/joaocarvalhodb1/arch-poc/src/shared/helpers"
)

const (
	webPort     = "8080"
	serviceName = "account api"
	gRPCAddress = "5055"
)

func main() {
	log := helpers.NewLoggers(serviceName)
	accountServiceConn, err := grpc.NewgRPCConnection(fmt.Sprintf("0.0.0.0:%s", gRPCAddress))
	if err != nil {
		log.Panic("account service dial error", err)
	}
	log.Debug("Connected to account service gRPC")
	defer accountServiceConn.Close()

	accountServiceClient := protog.NewAccountServiceClient(accountServiceConn)
	accountAPI := accountapi.NewAccountAppAPI(accountServiceClient, log)

	server := ws.NewHttpServer(accountAPI.Routes(), webPort, log)
	err = server.Listen()
	if err != nil {
		log.Panic(err.Error())
	}

}
