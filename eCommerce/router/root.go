package router

import (
	"context"
	"eCommerce/config"
	"eCommerce/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Router struct {
	config  *config.Config
	engine  *gin.Engine
	service *service.Service
}

func NewRouter(config *config.Config, service *service.Service) (*Router, error) {
	r := &Router{
		config:  config,
		engine:  gin.New(),
		service: service,
	}

	r.engine.Use(requestTimeOutMiddleWare(5 * time.Second))

	NewMongoRouter(r, r.service.MService)

	return r, r.engine.Run(config.ServerInfo.Port)
}

func requestTimeOutMiddleWare(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 요청의 컨텍스트에 타임아웃 설정
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// 2. 요청에 새로운 컨텍스트 설정
		c.Request = c.Request.WithContext(ctx)

		// 3. 비동기 작업 완료 신호를 받을 채널 생성
		done := make(chan struct{})

		// 4. 비동기 작업 시작
		go func() {
			defer close(done)
			c.Next()
		}()

		// 5. 요청 처리 완료 또는 타임아웃 발생 여부를 대기
		select {
		case <-done:
			// 작업이 완료되면 아무것도 하지 않고 반환
			return
		case <-ctx.Done():
			// 타임아웃이 발생하면 요청을 중단하고 타임아웃 응답을 반환
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				c.AbortWithStatusJSON(http.StatusRequestTimeout, gin.H{"error": "Request Time Out"})
				return
			}
		}
	}
}
