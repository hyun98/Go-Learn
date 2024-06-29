package router

import (
	"RPC-Learn/config"
	"RPC-Learn/gRPC/client"
	"RPC-Learn/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	config     *config.Config
	service    *service.Service
	gRPCClient *client.GRPCClient
	engine     *gin.Engine
}

func NewRouter(config *config.Config, service *service.Service, gRPCClient *client.GRPCClient) (*Router, error) {
	r := &Router{
		config:     config,
		service:    service,
		gRPCClient: gRPCClient,
		engine:     gin.New(),
	}

	// 1. token 생성하는 api
	r.engine.POST("/login", r.login)

	// 2. token 검증하는 api
	r.engine.GET("/verify", r.verityLogin, r.verify)

	return r, nil
}

func (r *Router) StartServer() error {
	return r.engine.Run(r.config.ServerInfo.Port)
}
