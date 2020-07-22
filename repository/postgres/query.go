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

package postgres

import (
	"database/sql"
	"fmt"

	proto "github.com/ta04/auth-service/model/proto"
)

// Postgres is the implementor of Repository interface
type Postgres struct {
	DB *sql.DB
}

// NewPostgres will create a new postgres instance
func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{
		DB: db,
	}
}

// GetOneByUsername will get an auth data by username
func (repo *Postgres) GetOneByUsername(auth2 *proto.Auth2) (*proto.Auth1, error) {
	var id int32
	var username, t, c string

	query := fmt.Sprintf("SELECT * FROM auth WHERE username = '%s' ORDER BY id DESC LIMIT 1", auth2.Username)
	err := repo.DB.QueryRow(query).Scan(&id, &username, &t, &c)
	if err != nil {
		return nil, err
	}

	return &proto.Auth1{
		Username: username,
		T:        t,
		C:        c,
	}, err
}
