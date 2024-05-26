package network

import (
	"CRUD_SERVER/service"
	"CRUD_SERVER/types"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router      *NetWork
	userService *service.UserService
}

func newUserRouter(router *NetWork, userService *service.UserService) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}

		router.registerGET("/", userRouterInstance.getAll)
		router.registerPOST("/", userRouterInstance.create)
		router.registerPUT("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)
	})

	return userRouterInstance
}

// CRUD

func (u *userRouter) getAll(c *gin.Context) {
	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		User:        u.userService.GetAll(),
	})
}

func (u *userRouter) create(c *gin.Context) {
	var req types.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류입니다.", -1, err.Error()),
		})
		return
	} else if err := u.userService.Create(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("Create 에러 입니다.", -1, err.Error()),
		})
		return
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}
}

func (u *userRouter) update(c *gin.Context) {
	var req types.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류입니다.", -1, err.Error()),
		})
		return
	} else if err := u.userService.Update(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("Update 에러 입니다.", -1, err.Error()),
		})
		return
	} else {
		u.router.okResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}
}

func (u *userRouter) delete(c *gin.Context) {
	var req types.DeleteUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류입니다.", -1, err.Error()),
		})
		return
	} else if err := u.userService.Delete(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("Delete 에러 입니다.", -1, err.Error()),
		})
		return
	} else {
		u.router.okResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}
}
