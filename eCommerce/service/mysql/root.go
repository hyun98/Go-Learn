package mysql

import (
	"eCommerce/repository"
)

type MySqlService struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *MySqlService {
	m := &MySqlService{
		repository: repository,
	}

	return m
}
