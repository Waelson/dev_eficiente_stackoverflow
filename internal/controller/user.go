package controller

import (
	"github.com/Waelson/internal/model"
	"github.com/Waelson/internal/service"
	"github.com/Waelson/internal/util/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController interface {
	Save(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func (c *userController) Save(ctx *gin.Context) {
	user := model.User{}

	if !middleware.ValidateStruct(ctx, &user) {
		return
	}

	_, err := c.userService.Save(ctx.Request.Context(), &user)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func (c *userController) Login(ctx *gin.Context) {
	auth := model.Login{}

	if !middleware.ValidateStruct(ctx, &auth) {
		return
	}

	token, err := c.userService.Login(ctx.Request.Context(), auth.Login, auth.Password)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(200, gin.H{"token": token})
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}
