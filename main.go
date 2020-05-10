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

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ta04/auth-service/database"
	"github.com/ta04/auth-service/handler/jwt"
	"github.com/ta04/auth-service/repository/postgres"

	_ "github.com/lib/pq"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/ta04/auth-service/handler"
	authPB "github.com/ta04/auth-service/proto"
)

func init() {
	// Load environment variables from .env into the system
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	name, exist := os.LookupEnv("MICRO_SERVICE_NAME")
	if !exist {
		log.Fatal("MICRO_SERVICE_NAME is not exists in .env file")
	}

	port, exist := os.LookupEnv("MICRO_SERVICE_PORT")
	if !exist {
		log.Fatal("MICRO_SERVICE_PORT is not exists in .env file")
	}

	registry := consul.NewRegistry()

	s := micro.NewService(
		micro.Name(name),
		micro.Address(port),
		micro.Registry(registry),
	)
	s.Init()

	db, err := database.OpenPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	h := handler.NewHandler(&postgres.Postgres{
		DB: db,
	}, &jwt.JWT{})
	authPB.RegisterAuthServiceHandler(s.Server(), h)

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
