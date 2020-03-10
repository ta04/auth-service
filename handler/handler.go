// auth-service/handler/handler.go

package handler

import (
	"context"
	"errors"
	authPB "github.com/SleepingNext/auth-service/proto"
	userPB "github.com/G0tYou/user-service/proto"
	"github.com/micro/go-micro/client"
)

type handler struct {
	tokenService Authable
}

func NewHandler(service Authable) *handler {
	return &handler{
		tokenService: service,
	}
}

func (h *handler) Auth(ctx context.Context, req *authPB.User, res *authPB.Token) error {
	userClient := userPB.NewUserServiceClient("com.ta04.h.user", client.DefaultClient)
	response, err := userClient.ShowUserByUsername(context.Background(), (*userPB.User)(req))
	if err != nil {
		return err
	}

	// Compare the given password and the password stored in database
	user := response.User
	if user.Password != req.Password {
		return errors.New("wrong credentials provided")
	}

	token, err := h.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *authPB.Token, res *authPB.Token) error {
	// Decode token
	_, err := h.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	res.Valid = true
	return nil
}
