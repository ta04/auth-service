// auth-service/repository/postgres/repository.go

package postgres

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/micro/go-micro/client"
	authPB "github.com/ta04/auth-service/proto"
	userPB "github.com/ta04/user-service/proto"
)

// Store will store a new auth
func (repo *Postgres) Store(auth1 *authPB.Auth1) (*authPB.C, error) {
	userClient := userPB.NewUserServiceClient("com.ta04.srv.user", client.DefaultClient)
	response, err := userClient.ShowUserByUsername(context.Background(), &userPB.User{Username: auth1.Username})
	if err != nil {
		return nil, err
	}
	user := response.User

	query := fmt.Sprintf("INSERT INTO auth (username, t) VALUES ('%s', '%s')",
		auth1.Username, auth1.T)
	_, err = repo.DB.Exec(query)

	rand.Seed(time.Now().UnixNano())
	parsedN, err := strconv.ParseInt(user.N, 10, 64)
	if err != nil {
		return nil, err
	}
	n := int(parsedN)

	return &authPB.C{
		C: fmt.Sprint(rand.Intn(n-1+1) + 1),
	}, err
}
