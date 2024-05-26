package repository

import (
	"CRUD_SERVER/types"
	"CRUD_SERVER/types/errors"
)

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{},
	}
}

func (u *UserRepository) GetAll() []*types.User {
	return u.userMap
}

func (u *UserRepository) Create(user *types.User) error {
	u.userMap = append(u.userMap, user)
	return nil
}

func (u *UserRepository) Update(updateUser *types.User) error {
	isExisted := false

	for _, user := range u.userMap {
		if user.Name == updateUser.Name {
			user.Age = updateUser.Age
			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFountUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Delete(name string) error {
	isExisted := false

	for i, user := range u.userMap {
		if user.Name == name {
			u.userMap = append(u.userMap[:i], u.userMap[i+1:]...)
			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFountUser, nil)
	} else {
		return nil
	}
}
