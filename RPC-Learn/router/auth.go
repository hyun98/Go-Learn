package router

import (
	"RPC-Learn/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) login(c *gin.Context) {
	var req types.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else if res, err := r.gRPCClient.CreateAuth(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (r *Router) verify(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
}
