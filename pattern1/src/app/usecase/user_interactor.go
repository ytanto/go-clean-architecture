package usecase

import "github.com/ytanto/go-clean-architecture/pattern1/src/app/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	var identifier int
	identifier, err = interactor.UserRepository.Store(u)
	if err != nil {
		return
	}

	user, err = interactor.UserRepository.FindById(identifier)
	return
}

func (interactor *UserInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}
