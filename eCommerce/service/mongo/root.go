package mongo

import (
	"eCommerce/repository"
)

type MService struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *MService {
	m := &MService{
		repository: repository,
	}

	return m
}
