package repository

import "clean/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
