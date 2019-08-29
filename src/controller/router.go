package controller

import (
	"github.com/gin-gonic/gin"
	srv "example_app/service"
)

func LoadRouter(routers *gin.Engine) {
	router := &UserRouterLoader{}
	router.UserRouter(routers)
}

type UserRouterLoader struct{
}

func (rLoader *UserRouterLoader) UserRouter(router *gin.Engine) {
	handler := &UserController{
		UserService: srv.UserServiceHandler(),
	}
	rLoader.routerDefinition(router,handler)
}

func (rLoader *UserRouterLoader) routerDefinition(router *gin.Engine,handler *UserController) {
	group := router.Group("v1/users")
	group.GET("", handler.GetUsers)
	group.GET(":id", handler.GetUserByID)
	group.PUT(":id", handler.UpdateUsersByID)
	group.DELETE(":id", handler.DeleteUserByID)
	group.POST("", handler.CreateUser)
}