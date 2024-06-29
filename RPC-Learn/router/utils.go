package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (r *Router) verityLogin(c *gin.Context) {
	token := getAuthToken(c)

	if token == "" {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
	} else {
		if _, err := r.gRPCClient.VerifyAuth(token); err != nil {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func getAuthToken(c *gin.Context) string {
	var token string

	authToken := c.Request.Header.Get("Authorization")

	authSliced := strings.Split(authToken, " ")
	if len(authSliced) > 1 {
		token = authSliced[1]
	}

	return token
}
