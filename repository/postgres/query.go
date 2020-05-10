// auth-service/repository/postgres/query.go

package postgres

import (
	"database/sql"
	"fmt"

	authPB "github.com/ta04/auth-service/proto"
)

// Postgres is the implementor of Repository interface
type Postgres struct {
	DB *sql.DB
}

// ShowByUsername will show auth info of an active user by it's username
func (repo *Postgres) ShowByUsername(auth2 *authPB.Auth2) (*authPB.Auth1, error) {
	var id int32
	var username, t string

	query := fmt.Sprintf("SELECT * FROM auth WHERE username = '%s' ORDER BY id DESC LIMIT 1", auth2.Username)
	err := repo.DB.QueryRow(query).Scan(&id, &username, &t)
	if err != nil {
		return nil, err
	}

	return &authPB.Auth1{
		Username: username,
		T:        t,
	}, err
}
