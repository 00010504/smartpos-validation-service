package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Invan2/invan_validation_service/config"
	"google.golang.org/grpc"
)

func main() {

	cfg := config.Load()

	server := grpc.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.HttpPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		server.GracefulStop()
	}()

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
		return
	}
}
