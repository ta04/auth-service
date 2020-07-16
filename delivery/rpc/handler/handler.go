/*
Dear Programmers,

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*                                                 *
*	This file belongs to Kevin Veros Hamonangan   *
*	and	Fandi Fladimir Dachi and is a part of     *
*	our	last project as the student of Del        *
*	Institute of Technology, Sitoluama.           *
*	Please contact us via Instagram:              *
*	sleepingnext and fandi_dachi                  *
*	before copying this file.                     *
*	Thank you, buddy. 😊                          *
*                                                 *
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*/

package handler

import (
	"context"
	"errors"

	proto "github.com/ta04/auth-service/model/proto"
	"github.com/ta04/auth-service/usecase"
)

// Handler is the handler of auth service
type Handler struct {
	Usecase usecase.Usecase
}

// NewHandler creates a new auth service handler
func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}

func (handler *Handler) AuthRPC1(ctx context.Context, req *proto.Auth1, res *proto.Response) error {
	c, err := handler.Usecase.Auth1(req)
	if err != nil {
		res.C = ""
		res.Error = err

		return errors.New(err.Message)
	}

	res.C = c

	return nil
}

func (handler *Handler) AuthRPC2(ctx context.Context, req *proto.Auth2, res *proto.Response) error {
	token, err := handler.Usecase.Auth2(req)
	if err != nil {
		res.Token = ""
		res.Error = err

		return errors.New(err.Message)
	}

	res.Token = token

	return nil
}
