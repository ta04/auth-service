// user-service/database/database.go

package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "skit-auth-service.c6pg0ooh5mrq.ap-southeast-1.rds.amazonaws.com"
	port     = 5432
	user     = "sleepingnext"
	password = "kevin99123"
	dbname   = "auth"
)

// OpenPostgresConnection is to connect to postgres database
func OpenPostgresConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port =%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)

	return db, err
}
