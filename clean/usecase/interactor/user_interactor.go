package interactor

import (
	"clean/domain/model"
	"clean/usecase/presenter"
	"clean/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
} //TODO: Why do we need this interface??

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}
