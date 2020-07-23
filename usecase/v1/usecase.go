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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/ta04/auth-service/helper"
	authPB "github.com/ta04/auth-service/model/proto"
	"github.com/ta04/auth-service/repository"
	userPB "github.com/ta04/user-service/model/proto"
)

// Usecase is the struct of order usecase
type Usecase struct {
	Repository repository.Repository
}

type getUserStruct struct {
	User userPB.User `json:"user"`
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

// Auth1 will authenticate the user
func (usecase *Usecase) Auth1(auth1 *authPB.Auth1) (string, *authPB.Error) {
	if auth1 == nil {
		log.Println("auth1 is nil")
		return "", badRequestError
	}

	getUserWithoutCredentialsRequestBody := map[string]interface{}{
		"username":         auth1.Username,
		"With_credentials": false,
	}

	marshaledGetUserWithoutCredentialsRequestBody, err := json.Marshal(getUserWithoutCredentialsRequestBody)
	if err != nil {
		log.Fatalln(err)
	}

	getUserWithoutCredentialsResponse, err := http.Post("http://localhost:50057/user/show", "application/json", bytes.NewBuffer(marshaledGetUserWithoutCredentialsRequestBody))
	if err != nil {
		log.Println(1, err)
		return "", internalServerError
	}

	defer getUserWithoutCredentialsResponse.Body.Close()

	body, err := ioutil.ReadAll(getUserWithoutCredentialsResponse.Body)
	if err != nil {
		log.Println(err)
		return "", internalServerError
	}

	var getUser getUserStruct
	err = json.Unmarshal(body, &getUser)
	if err != nil {
		log.Println(err)
		return "", internalServerError
	}

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := int(getUser.User.PrimeNumber)
	c := fmt.Sprint(rand.Intn(max-min+1) + min)

	auth1.C = c

	err = usecase.Repository.CreateOne(auth1)
	if err != nil {
		log.Println(4, err)
		return "", internalServerError
	}

	return c, nil
}

// Auth2 will authenticate the user
func (usecase *Usecase) Auth2(auth2 *authPB.Auth2) (string, *authPB.Error) {
	if auth2 == nil {
		log.Println("auth2 is nil")
		return "", badRequestError
	}

	auth1, err := usecase.Repository.GetOneByUsername(auth2)
	if err != nil {
		log.Println(err)
		return "", internalServerError
	}

	getUserWithCredentialsRequestBody := map[string]interface{}{
		"username":         auth2.Username,
		"With_credentials": true,
	}

	marshaledGetUserWithCredentialsRequestBody, err := json.Marshal(getUserWithCredentialsRequestBody)
	if err != nil {
		log.Fatalln(err)
	}

	getUserWithCredentialsResponse, err := http.Post("http://localhost:50057/user/show", "application/json", bytes.NewBuffer(marshaledGetUserWithCredentialsRequestBody))
	if err != nil {
		log.Println(1, err)
		return "", internalServerError
	}

	defer getUserWithCredentialsResponse.Body.Close()

	body, err := ioutil.ReadAll(getUserWithCredentialsResponse.Body)
	if err != nil {
		log.Println(err)
		return "", internalServerError
	}

	var getUserWithCredential getUserStruct
	err = json.Unmarshal(body, &getUserWithCredential)
	if err != nil {
		log.Println(err)
		return "", internalServerError
	}

	userWithCredentials := getUserWithCredential.User

	getUserWithoutCredentialsRequestBody := map[string]interface{}{
		"username":         auth1.Username,
		"With_credentials": false,
	}

	marshaledGetUserWithoutCredentialsRequestBody, err := json.Marshal(getUserWithoutCredentialsRequestBody)
	if err != nil {
		log.Fatalln(err)
	}

	getUserWithoutCredentialsResponse, err := http.Post("http://localhost:50057/user/show", "application/json", bytes.NewBuffer(marshaledGetUserWithoutCredentialsRequestBody))
	if err != nil {
		log.Println(1, err)
		return "", internalServerError
	}

	defer getUserWithoutCredentialsResponse.Body.Close()

	body, err = ioutil.ReadAll(getUserWithoutCredentialsResponse.Body)
	if err != nil {
		log.Println(err)
		return "", internalServerError
	}

	var getUserWithoutCredential getUserStruct
	err = json.Unmarshal(body, &getUserWithoutCredential)
	if err != nil {
		log.Println(err)
		return "", internalServerError
	}

	userWithoutCredentials := getUserWithoutCredential.User

	// Calculate the result and compare to T
	var result string
	parsedR, err := strconv.ParseInt(auth2.R, 10, 64)
	if err != nil {
		log.Println(err)
		return "", internalServerError
	}

	rInt := int(parsedR)
	if rInt < 0 {
		// With inverse mod
		calculateResultURL := fmt.Sprintf("http://localhost:5000/calculateResult?g=%d&r=%d&n=%d&y=%d&c=%s",
			userWithoutCredentials.GeneratorValue, rInt, userWithoutCredentials.PrimeNumber, userWithCredentials.Password, auth1.C)
		calculateResultRes, err := http.Get(calculateResultURL)
		if err != nil {
			log.Println(err)
			return "", internalServerError
		}

		defer calculateResultRes.Body.Close()

		body, err := ioutil.ReadAll(calculateResultRes.Body)
		if err != nil {
			log.Println(err)
			return "", internalServerError
		}

		type ResultStruct struct {
			Result int64 `json:"result"`
		}
		var unmarshalledBody ResultStruct
		err = json.Unmarshal(body, &unmarshalledBody)
		if err != nil {
			log.Println(err)
			return "", internalServerError
		}

		result = strconv.FormatInt(unmarshalledBody.Result, 10)
	} else {
		// Without inverse mod
		g := float64(userWithoutCredentials.GeneratorValue)
		n := int64(userWithoutCredentials.PrimeNumber)
		y := float64(userWithCredentials.Password)
		c, r, err := parseCalculateResultParams(auth1.C, auth2.R)
		if err != nil {
			log.Println(err)
			return "", internalServerError
		}

		resultTemp := ((int64(math.Pow(g, r)) % n) * (int64(math.Pow(y, c)) % n)) % n
		result = strconv.FormatInt(resultTemp, 10)
	}

	if result == auth1.T {
		token, err := helper.Encode(&userWithoutCredentials)
		if err != nil {
			log.Println(err)
			return "", internalServerError
		}

		err = usecase.Repository.DeleteByUsername(auth1)
		if err != nil {
			log.Println(err)
			return "", internalServerError
		}

		return token, nil
	}

	return "", &authPB.Error{
		Code:    http.StatusUnauthorized,
		Message: http.StatusText(http.StatusUnauthorized),
	}
}

func parseCalculateResultParams(cString string, rString string) (float64, float64, error) {
	var err2 error
	c, err := strconv.ParseFloat(cString, 64)
	if err != nil {
		err2 = err
	}

	r, err := strconv.ParseFloat(rString, 64)
	if err != nil {
		err2 = err
	}

	return c, r, err2
}
