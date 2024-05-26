package repository

import (
	"sync"
)

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	UserRepository *UserRepository
}

func CreateRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			UserRepository: newUserRepository(),
		}
	})

	return repositoryInstance
}
