package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type UserUsecase struct {
	userRepository repository.UserRpository
}

func NewUserUsecase(repository repository.UserRpository) UserUsecase {
	return UserUsecase{
		userRepository: repository,
	}
}

func (u *UserUsecase) GetAllUsers() ([]model.User, error) {
	users, err := u.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}