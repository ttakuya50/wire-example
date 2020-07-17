package main

import (
	"context"
	"log"

	"github.com/ttakuya50/wire-example/config"
	"github.com/ttakuya50/wire-example/infra/mysql"
)

func main() {
	_, cleanup, err := initializeApp(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

}

type app struct {
	//useCase usecase.UserUseCase
	//db       *sql.DB
	mysql  *mysql.Client
	config *config.Config
	//userRepo repository.UserRepo
}

func newApp(config *config.Config, mysqlClient *mysql.Client) *app {
	return &app{
		//useCase: useCase,
		//db:       db,
		config: config,
		mysql:  mysqlClient,
		//userRepo: userRepo,
	}
}
