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

package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/ta04/auth-service/client"
	"github.com/ta04/auth-service/helper"
	authPB "github.com/ta04/auth-service/model/proto"
	"github.com/ta04/auth-service/repository"
	userPB "github.com/ta04/user-service/model/proto"
)

// usecase is the struct of order usecase
type Usecase struct {
	Repository repository.Repository
}

// NewUsecase will create a new order usecase
func NewUsecase(repository repository.Repository) *Usecase {
	return &Usecase{
		Repository: repository,
	}
}

var badRequestError = &authPB.Error{
	Code:    http.StatusBadRequest,
	Message: http.StatusText(http.StatusBadRequest),
}

var internalServerError = &authPB.Error{
	Code:    http.StatusInternalServerError,
	Message: http.StatusText(http.StatusInternalServerError),
}

func (usecase *Usecase) Auth1(auth1 *authPB.Auth1) (string, *authPB.Error) {
	if auth1 == nil {
		return "", badRequestError
	}

	err := usecase.Repository.CreateOne(auth1)
	if err != nil {
		return "", internalServerError
	}

	userClient := client.NewUserSC()
	response, err := userClient.GetOneUser(context.Background(), &userPB.GetOneUserRequest{Username: auth1.Username})
	if err != nil {
		return "", internalServerError
	}

	user := response.User

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := int(user.PrimeNumber)

	return fmt.Sprint(rand.Intn(max-min+1) + min), nil
}

func (usecase *Usecase) Auth2(auth2 *authPB.Auth2) (string, *authPB.Error) {
	if auth2 == nil {
		return "", badRequestError
	}

	auth1, err := usecase.Repository.GetOneByUsername(auth2)
	if err != nil {
		return "", internalServerError
	}

	userClient := client.NewUserSC()
	response, err := userClient.GetOneUser(context.Background(), &userPB.GetOneUserRequest{Username: auth2.Username})
	if err != nil {
		return "", internalServerError
	}

	user := response.User

	// Calculate the result and compare to T
	var result string
	parsedR, err := strconv.ParseInt(auth2.R, 10, 64)
	if err != nil {
		return "", internalServerError
	}

	rInt := int(parsedR)
	if rInt < 0 {
		// With inverse mod
		calculateResultUrl := fmt.Sprintf("http://localhost:5000/calculateResult?g=%d&r=%d&n=%d&y=%d&c=%s",
			user.GeneratorValue, rInt, user.PrimeNumber, user.Password, auth2.C)
		calculateResultRes, err := http.Get(calculateResultUrl)
		if err != nil {
			return "", internalServerError
		}

		defer calculateResultRes.Body.Close()

		body, err := ioutil.ReadAll(calculateResultRes.Body)
		if err != nil {
			return "", internalServerError
		}

		type ResultStruct struct {
			Result int64 `json:"result"`
		}
		var unmarshalledBody ResultStruct
		err = json.Unmarshal(body, &unmarshalledBody)
		if err != nil {
			return "", internalServerError
		}

		result = strconv.FormatInt(unmarshalledBody.Result, 10)
	} else {
		// Without inverse mod
		g := float64(user.GeneratorValue)
		n := int64(user.PrimeNumber)
		y := float64(user.Password)
		r, c, err := parseCalculateResultParams(auth2)
		if err != nil {
			return "", internalServerError
		}

		resultTemp := ((int64(math.Pow(g, r)) % n) * (int64(math.Pow(y, c)) % n)) % n
		result = strconv.FormatInt(resultTemp, 10)
	}

	if result == auth1.T {
		token, err := helper.Encode(user)
		if err != nil {
			return "", internalServerError
		}
		return token, nil
	}

	return "", &authPB.Error{
		Code:    http.StatusUnauthorized,
		Message: http.StatusText(http.StatusUnauthorized),
	}
}

func parseCalculateResultParams(auth2 *authPB.Auth2) (float64, float64, error) {
	var err2 error
	r, err := strconv.ParseFloat(auth2.R, 64)
	if err != nil {
		err2 = err
	}

	c, err := strconv.ParseFloat(auth2.C, 64)
	if err != nil {
		err2 = err
	}

	return r, c, err2
}
