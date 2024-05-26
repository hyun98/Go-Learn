package network

import (
	"CRUD_SERVER/service"
	"github.com/gin-gonic/gin"
)

type NetWork struct {
	engine *gin.Engine

	service *service.Service
}

func NewNetWork(service *service.Service) *NetWork {
	n := &NetWork{
		engine: gin.New(),
	}

	_ = newUserRouter(n, service.UserService)

	return n
}

func (n *NetWork) ServerStart(port string) (err error) {
	return n.engine.Run(port)
}
