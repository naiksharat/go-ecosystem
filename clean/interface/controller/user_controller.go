package controller

import (
	"net/http"

	"clean/domain/model"
	"clean/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c Context) error
} // TODO: Why??

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c Context) error {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
