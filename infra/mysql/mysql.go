package mysql

import (
	"database/sql"
	"log"

	"github.com/ttakuya50/wire-example/domain/repository"

	"github.com/ttakuya50/wire-example/config"

	"github.com/google/wire"
)

// Set is a Wire provider set that produces a *sql.DB.
var Set = wire.NewSet(
	Open,
)

// DataSource specifies how to connect to a source database.
//type Config struct {
//	DriverName     string
//	DataSourceName string
//}

type Client struct {
	UserRepo repository.UserRepo
}

// Open opens a connection to a SQL database.
func Open(cfg *config.Config) (*Client, func(), error) {
	db, err := sql.Open(cfg.DBName, cfg.DBPassword)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Print(err)
		}
	}

	userDao := NewUserDao(db)

	return &Client{
		UserRepo: userDao,
	}, cleanup, nil
}
