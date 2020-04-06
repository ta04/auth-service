// auth-service/repository/postgres/repository.go

package postgres

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	userPB "github.com/G0tYou/user-service/proto"
	authPB "github.com/SleepingNext/auth-service/proto"
	"github.com/micro/go-micro/client"
)

// Store will store a new auth
func (repo *Repository) Store(auth1 *authPB.Auth1) (*authPB.C, error) {
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
	parsedN, _ := strconv.ParseInt(user.N, 10, 64)
	n := int(parsedN)

	return &authPB.C{
		C: fmt.Sprint(rand.Intn(n-1+1) + 1),
	}, err
}
