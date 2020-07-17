package mysql

import (
	"database/sql"

	"github.com/ttakuya50/wire-example/domain/repository"

	"github.com/ttakuya50/wire-example/domain/model"
)

type userDto struct {
	Id   int64
	Name string
}

type userDao struct {
	db *sql.DB
}

func NewUserDao(db *sql.DB) repository.UserRepo {
	return &userDao{
		db: db,
	}
}

func (d *userDao) Store(user *model.User) error {
	return nil
}
