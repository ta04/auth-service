// auth-service/handler/handler.go

package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"

	authRepo "github.com/SleepingNext/auth-service/repository"

	userPB "github.com/G0tYou/user-service/proto"
	authPB "github.com/SleepingNext/auth-service/proto"
	"github.com/micro/go-micro/client"
)

type handler struct {
	repository   authRepo.Repository
	tokenService Authable
}

func NewHandler(repo authRepo.Repository, service Authable) *handler {
	return &handler{
		repository:   repo,
		tokenService: service,
	}
}

func (h *handler) AuthRPC2(ctx context.Context, req *authPB.Auth2, res *authPB.Token) error {
	var result string
	userClient := userPB.NewUserServiceClient("com.ta04.srv.user", client.DefaultClient)
	response, err := userClient.ShowUserByUsername(context.Background(), &userPB.User{Username: req.Username})
	if err != nil {
		return err
	}

	user := response.User

	auth1, err := h.repository.ShowByUsername(req)
	if err != nil {
		return err
	}

	// Calculate the result and compare to T
	parsedR, err := strconv.ParseInt(req.R, 10, 64)
	if err != nil {
		return err
	}
	rInt := int(parsedR)
	if rInt < 0 {
		calculateResultUrl := "http://ec2-54-169-178-240.ap-southeast-1.compute.amazonaws.com/calculateResult?g=" + user.G + "&r=" + strconv.FormatInt(int64(rInt), 10) + "&n=" + user.N + "&y=" + user.Password + "&c=" + req.C
		calculateResultRes, err := http.Get(calculateResultUrl)
		if err != nil {
			return err
		}
		body, err := ioutil.ReadAll(calculateResultRes.Body)
		defer calculateResultRes.Body.Close()
		if err != nil {
			return err
		}

		type Result struct{ result string }
		var unmarshalledBody *Result
		err = json.Unmarshal(body, &unmarshalledBody)
		if err != nil {
			return err
		}

		result = unmarshalledBody.result
	} else {
		g, err := strconv.ParseFloat(user.G, 64)
		if err != nil {
			return err
		}

		rFloat, err := strconv.ParseFloat(req.R, 64)
		if err != nil {
			return err
		}

		n, err := strconv.ParseInt(user.N, 10, 64)
		if err != nil {
			return err
		}

		y, err := strconv.ParseFloat(user.Password, 64)
		if err != nil {
			return err
		}

		c, err := strconv.ParseFloat(req.C, 64)
		if err != nil {
			return err
		}

		resultTemp := ((int64(math.Pow(g, rFloat)) % n) * (int64(math.Pow(y, c)) % n)) % n
		result = strconv.FormatInt(resultTemp, 10)
	}

	if result == auth1.T {
		token, err := h.tokenService.Encode(user)
		if err != nil {
			return err
		}
		res.Token = token
	}
	return nil
}

func (h *handler) AuthRPC1(ctx context.Context, req *authPB.Auth1, res *authPB.C) error {
	c, err := h.repository.Store(req)
	if err != nil {
		return err
	}

	res.C = c.C

	return err
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
