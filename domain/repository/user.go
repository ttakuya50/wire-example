package repository

import (
	"github.com/ttakuya50/wire-example/domain/model"
)

type UserRepo interface {
	Store(user *model.User) error
}
