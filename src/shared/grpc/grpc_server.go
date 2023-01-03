package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/joaocarvalhodb1/arch-poc/src/shared/helpers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	Server *grpc.Server
	port   string
	log    *helpers.Loggers
}

func NewgRPCServer(port string, log *helpers.Loggers) (*GRPCServer, error) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	server := &GRPCServer{
		Server: grpcServer,
		port:   port,
		log:    log,
	}
	return server, nil
}

func (s *GRPCServer) ListenAndServe() {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", s.port))
	if err != nil {
		log.Panic("Error on listen TCP port", err)
	}
	s.log.Info("GRPC Server started on port:", s.port)
	if err := s.Server.Serve(listen); err != nil {
		s.log.Panic("Failed to listen for gRPC:", err)
	}
}

// func (s *AccountAppService) GRPCRegister(gRPCPort string) {
// 	grpcServer := grpc.NewServer()
// 	reflection.Register(grpcServer)

// 	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", gRPCPort))
// 	if err != nil {
// 		s.log.Panic("Error on listen TCP port", err)
// 	}

// 	protog.RegisterAccountServiceServer(grpcServer, s)
// 	s.log.Info("GRPC Server started on port:", gRPCPort)
// 	if err := grpcServer.Serve(listen); err != nil {
// 		s.log.Panic("Failed to listen for gRPC:", err)
// 	}
// }
