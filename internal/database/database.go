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

package database

import (
	"database/sql"
	"fmt"

	"github.com/ta04/auth-service/internal/config"
)

// OpenPostgresConnection opens a connection to postgres database
func OpenPostgresConnection() (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("host=%s port =%d user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost(), config.PostgresPort(), config.PostgresUser(),
		config.PostgresPassword(), config.PostgresDBName()))
}
