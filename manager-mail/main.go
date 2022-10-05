package main

import (
	"log"
	"net"

	"github.com/Abeldlp/manager-mail/config"
	"github.com/Abeldlp/manager-mail/mailpb"
	"github.com/Abeldlp/manager-mail/server"
	"google.golang.org/grpc"
)

func main() {
	config.InitializeDatabase()

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	mailpb.RegisterMailServiceServer(s, &server.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
