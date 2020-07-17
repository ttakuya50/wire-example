// wire-example/wire.go
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/ttakuya50/wire-example/config"
	"github.com/ttakuya50/wire-example/infra/mysql"
)

func initializeApp(ctx context.Context) (*app, func(), error) {
	wire.Build(
		//appSet,
		//server.Set,
		//usecase.Set,
		mysql.Set,
		config.Set,
		newApp,
	)
	return nil, nil, nil
}
