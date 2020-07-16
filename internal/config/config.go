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

package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env into the system
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

// MicroServiceName will return MICRO_SERVICE_NAME from .env file
func MicroServiceName() string {
	return lookupEnv("MICRO_SERVICE_NAME")
}

// MicroServicePort will return MICRO_SERVICE_PORT from .env file
func MicroServicePort() string {
	return lookupEnv("MICRO_SERVICE_PORT")
}

// PostgresHost will return POSTGRES_HOST from .env file
func PostgresHost() string {
	return lookupEnv("POSTGRES_HOST")
}

// PostgresPort will return POSTGRES_PORT from .env file
func PostgresPort() int32 {
	postgresPort, err := strconv.ParseInt(lookupEnv("POSTGRES_PORT"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int32(postgresPort)
}

// PostgresUser will return POSTGRES_USER from .env file
func PostgresUser() string {
	return lookupEnv("POSTGRES_USER")
}

// PostgresPassword will return POSTGRES_PASSWORD from .env file
func PostgresPassword() string {
	return lookupEnv("POSTGRES_PASSWORD")
}

// PostgresDBName will return POSTGRES_DBNAME from .env file
func PostgresDBName() string {
	return lookupEnv("POSTGRES_DBNAME")
}

func lookupEnv(key string) string {
	env, exist := os.LookupEnv(key)
	if !exist {
		log.Fatal(key, "is not exists in .env file")
	}

	return env
}
