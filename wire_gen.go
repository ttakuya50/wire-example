// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
	"github.com/ttakuya50/wire-example/config"
	"github.com/ttakuya50/wire-example/infra/mysql"
)

// Injectors from wire.go:

func initializeApp(ctx context.Context) (*app, func(), error) {
	configConfig := config.NewConfig()
	client, cleanup, err := mysql.Open(configConfig)
	if err != nil {
		return nil, nil, err
	}
	mainApp := newApp(configConfig, client)
	return mainApp, func() {
		cleanup()
	}, nil
}
