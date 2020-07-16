package client

import (
	"github.com/micro/go-micro"
	proto "github.com/ta04/user-service/model/proto"
)

// NewUserSC creates a new user service client
func NewUserSC() proto.UserServiceClient {
	s := micro.NewService(
		micro.Name("com.ta04.api.user"),
	)
	s.Init()

	userServiceClient := proto.NewUserServiceClient("com.ta04.srv.user", s.Client())
	return userServiceClient
}
