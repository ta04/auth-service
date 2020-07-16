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

	_ "github.com/lib/pq"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/ta04/auth-service/delivery/rpc/handler"
	"github.com/ta04/auth-service/internal/config"
	"github.com/ta04/auth-service/internal/database"
	proto "github.com/ta04/auth-service/model/proto"
	"github.com/ta04/auth-service/repository/postgres"
	usecase "github.com/ta04/auth-service/usecase/v1"
)

func main() {
	name := config.MicroServiceName()
	port := config.MicroServicePort()

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

	p := postgres.NewPostgres(db)

	u := usecase.NewUsecase(p)

	h := handler.NewHandler(u)

	proto.RegisterAuthServiceHandler(s.Server(), h)

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
