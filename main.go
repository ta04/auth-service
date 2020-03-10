package main

import (
	"github.com/SleepingNext/auth-service/handler"
	authPB "github.com/SleepingNext/auth-service/proto"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	// Create a new service
	s := micro.NewService(
		micro.Name("com.ta04.srv.auth"),
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

