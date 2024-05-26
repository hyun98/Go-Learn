package service

import (
	"CRUD_SERVER/repository"
	"CRUD_SERVER/types"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) GetAll() []*types.User {

	if us == nil {
		return []*types.User{}
	}
	return us.userRepository.GetAll()
}

func (us *UserService) Create(newUser *types.User) error {
	return us.userRepository.Create(newUser)
}

func (us *UserService) Update(updateUser *types.User) error {
	return us.userRepository.Update(updateUser)
}

func (us *UserService) Delete(user *types.User) error {
	return us.userRepository.Delete(user.Name)
}
