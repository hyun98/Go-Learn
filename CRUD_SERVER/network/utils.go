package network

import "github.com/gin-gonic/gin"

// register 유틸 함수들

func (n *NetWork) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)
}

func (n *NetWork) registerPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handler...)
}

func (n *NetWork) registerPUT(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handler...)
}

func (n *NetWork) registerDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handler...)
}

// Response 형태를 맞추기 위한 유틸 함수

func (n *NetWork) okResponse(c *gin.Context, result interface{}) {
	c.JSON(200, result)
}

func (n *NetWork) failedResponse(c *gin.Context, result interface{}) {
	c.JSON(400, result)
}
