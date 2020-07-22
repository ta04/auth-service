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
	"fmt"

	proto "github.com/ta04/auth-service/model/proto"
)

// CreateOne will create a new auth data
func (repo *Postgres) CreateOne(auth1 *proto.Auth1) error {
	query := fmt.Sprintf("INSERT INTO auth (username, t, c) VALUES ('%s', '%s', '%s')",
		auth1.Username, auth1.T, auth1.C)
	_, err := repo.DB.Exec(query)

	return err
}

// DeleteByUsername will delete auth data by username
func (repo *Postgres) DeleteByUsername(auth1 *proto.Auth1) error {
	query := fmt.Sprintf("DELETE FROM auth WHERE username='%s'", auth1.Username)
	_, err := repo.DB.Exec(query)

	return err
}
