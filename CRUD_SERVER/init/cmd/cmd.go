package cmd

import (
	"CRUD_SERVER/config"
	"CRUD_SERVER/network"
	"CRUD_SERVER/repository"
	"CRUD_SERVER/service"
)

type Cmd struct {
	config     *config.Config
	network    *network.NetWork
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filepath),
	}

	c.repository = repository.CreateRepository()
	c.service = service.CreateService(c.repository)
	c.network = network.NewNetWork(c.service)

	if err := c.network.ServerStart(c.config.Server.Port); err != nil {
		panic(err)
	}

	return c
}
