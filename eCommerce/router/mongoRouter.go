package router

import (
	"context"
	"eCommerce/service/mongo"
	. "eCommerce/types"
	. "eCommerce/types/err"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	pM "go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoRouter struct {
	router  *Router
	service *mongo.MService
}

func NewMongoRouter(router *Router, mService *mongo.MService) {
	mr := &MongoRouter{
		router:  router,
		service: mService,
	}

	baseUri := "/mongo"
	mr.router.GET(baseUri+"/health", mr.Health)

	mr.router.GET(baseUri+"/bucket", mr.userBucket)                     // 장바구니에 대한 정보
	mr.router.GET(baseUri+"/content", mr.content)                       // 상품 정보를 가져오는 정보
	mr.router.GET(baseUri+"/user-bucket-history", mr.userBucketHistory) // 유저의 구매 이력 정보

}

func (mr *MongoRouter) userBucket(c *gin.Context) {
	var req BucketRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		mr.router.ResponseErr(c, ErrorMsg(BindingFailure, err))
		return
	} else if u, err := mr.service.GetUserBucket(req.User); err != nil {
		if errors.Is(err, pM.ErrNoDocuments) {
			mr.router.ResponseErr(c, ErrorMsg(NoDocument, err))
		} else {
			mr.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		}
		return
	} else {
		mr.router.ResponseOK(c, u)
	}
}

func (mr *MongoRouter) content(c *gin.Context) {
	var req ContentRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		mr.router.ResponseErr(c, ErrorMsg(BindingFailure, err))
		return
	} else if u, err := mr.service.GetContent(req.Name); err != nil {
		if errors.Is(err, pM.ErrNoDocuments) {
			mr.router.ResponseErr(c, ErrorMsg(NoDocument, err))
		} else {
			mr.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		}
		return
	} else {
		mr.router.ResponseOK(c, u)
	}
}

func (mr *MongoRouter) userBucketHistory(c *gin.Context) {
	var req BucketRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		mr.router.ResponseErr(c, ErrorMsg(BindingFailure, err))
		return
	} else if u, err := mr.service.GetUserBucketHistory(req.User); err != nil {
		if errors.Is(err, pM.ErrNoDocuments) {
			mr.router.ResponseErr(c, ErrorMsg(NoDocument, err))
		} else {
			mr.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		}
		return
	} else {
		mr.router.ResponseOK(c, u)
	}
}

func (mr *MongoRouter) Health(c *gin.Context) {
	time.Sleep(1 * time.Second)

	if errors.Is(c.Request.Context().Err(), context.DeadlineExceeded) {
		fmt.Println("에러 발생")
	} else {
		fmt.Println("Comming")
	}

	if !c.Writer.Written() {
		mr.router.ResponseOK(c, "test")
	}
}
