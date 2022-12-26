package presenter

import "clean/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
