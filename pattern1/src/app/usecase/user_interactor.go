package usecase

import "app/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err := interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) Users() (users domain.User, err error) {
	users, err := interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err := interactor.UserRepository.FindById(identifier)
	return
}
