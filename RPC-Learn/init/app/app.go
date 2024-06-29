package app

import (
	"RPC-Learn/config"
	"RPC-Learn/gRPC/client"
	"RPC-Learn/repository"
	"RPC-Learn/router"
	"RPC-Learn/service"
)

type App struct {
	config     *config.Config
	repository *repository.Repository
	service    *service.Service
	router     *router.Router
	gRPCClient *client.GRPCClient
}

func NewApp(config *config.Config) {
	app := &App{
		config: config,
	}

	var err error
	if app.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	} else if app.service, err = service.NewService(config, app.repository); err != nil {
		panic(err)
	} else if app.gRPCClient, err = client.NewGRPCClient(config); err != nil {
		panic(err)
	} else if app.router, err = router.NewRouter(config, app.service, app.gRPCClient); err != nil {
		panic(err)
	} else {
		if err = app.router.StartServer(); err != nil {
			panic(err)
		}
	}
}
