package main

import (
	"log"
	"os"

	"github.com/SleepingNext/auth-service/handler"
	authPB "github.com/SleepingNext/auth-service/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// Take or set the port
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":50055"
	}

	// Create a new registry
	registry := consul.NewRegistry()

	// Create a new service
	s := micro.NewService(
		micro.Name("com.ta04.srv.auth"),
		micro.Address(port),
		micro.Registry(registry),
	)

	// Initialize the service
	s.Init()

	// Create a new handler
	h := handler.NewHandler(&handler.TokenService{})

	// Register the handler
	authPB.RegisterAuthServiceHandler(s.Server(), h)

	// Run the server
	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
