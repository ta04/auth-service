// auth-service/repository/postgres/query.go

package postgres

import (
	"database/sql"
	"fmt"

	authPB "github.com/SleepingNext/auth-service/proto"
)

type Repository struct {
	DB *sql.DB
}

// ShowByUsername will show an active user by it's username
func (repo *Repository) ShowByUsername(auth2 *authPB.Auth2) (*authPB.Auth1, error) {
	var id int32
	var username, t string

	query := fmt.Sprintf("SELECT * FROM auth WHERE username = '%s' ORDER BY id DESC LIMIT 1", auth2.Username)
	err := repo.DB.QueryRow(query).Scan(&id, &username, &t)
	if err != nil {
		return nil, err
	}

	return &authPB.Auth1{
		Id:                     id,
		Username:               username,
		T:                      t,
	}, err
}
