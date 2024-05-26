package service

import (
	"CRUD_SERVER/repository"
	"sync"
)

// Network, repository 간의 다리 역할

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	repository  *repository.Repository
	UserService *UserService
}

func CreateService(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}

		serviceInstance.UserService = newUserService(rep.UserRepository)
	})

	return serviceInstance
}
