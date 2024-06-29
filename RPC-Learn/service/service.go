package service

import (
	"RPC-Learn/config"
	"RPC-Learn/repository"
)

type Service struct {
	config     *config.Config
	repository *repository.Repository
}

func NewService(config *config.Config, repository *repository.Repository) (*Service, error) {
	s := &Service{
		config:     config,
		repository: repository,
	}

	return s, nil
}
