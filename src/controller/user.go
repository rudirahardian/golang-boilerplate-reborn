package controller

import (
	"strconv"
	"sync"
	"github.com/gin-gonic/gin"
	"net/http"
	services "example_app/service"
	httpEntity "example_app/entity/http"
	"fmt"
)

type UserController struct {
	UserService services.UserServiceInterface
}

func (handler *UserController) TestFunction(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func (service *UserController) GetUserByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	result := service.UserService.GetUserByID(id, &sync.WaitGroup{})
	if result == nil {
		context.JSON(http.StatusOK, gin.H{})
		return
	}
	context.JSON(http.StatusOK, result)
}

type Limitofset struct{
	Limit int `form:"limit"`
	Offset int `form:"offset"`
}
func (service *UserController) GetUsers(context *gin.Context) {
	queryparam := Limitofset{}
	err := context.ShouldBindQuery(&queryparam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	result := service.UserService.GetAllUser(queryparam.Limit, queryparam.Offset)
	context.JSON(http.StatusOK, result)
}

func (service *UserController) UpdateUsersByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	payload := httpEntity.UserRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	status := service.UserService.UpdateUserByID(id,payload)
	if nil != status {
		context.JSON(http.StatusOK, gin.H{
			"message error": status,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"playload": payload,
		"id": id,
	})
} 

// delete user by id
func (service *UserController) DeleteUserByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	payload := httpEntity.UserRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	status := service.UserService.DeleteUserByID(id)

	if status == true {
		context.JSON(http.StatusOK, gin.H{
			"message": "deleted",
			"id": id,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "row already deleted",
		"id": id,
	})
	
} 

func (service *UserController) CreateUser(context *gin.Context) {
	payload := httpEntity.UserRequestNew{}
	if err := context.ShouldBind(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	result, err:= service.UserService.CreateUser(payload)
	fmt.Println("Ada: ", payload)
	if nil != err {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}