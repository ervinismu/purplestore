package controllers

import (
	"github.com/ervinismu/purplestore/handler"
	"github.com/ervinismu/purplestore/schema"
	"github.com/ervinismu/purplestore/service"
	"github.com/gin-gonic/gin"
)

type userController struct {
	service service.UserService
}

func NewUserController(us service.UserService) userController {
	return userController{service: us}
}

func (uc userController) ListAllUser(ctx *gin.Context) {
	users, err := uc.service.ListAllUser()
	if err != nil {
		handler.HandlerErrorResponse(ctx, "Cannot get data users")
		return
	}

	handler.HandlerSuccessResponse(ctx, "Success get users", users )
}

func (uc userController) ShowUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uc.service.ShowUser(id)
	if err != nil {
		handler.HandlerErrorResponse(ctx, "Record not found")
		return
	}

	handler.HandlerSuccessResponse(ctx, "Success show user", user)
}

func (uc userController) CreateUser(ctx *gin.Context) {
	var input schema.CreateUserReq
	if err := ctx.ShouldBindJSON(&input); err != nil {
		handler.HandlerErrorResponse(ctx, err.Error())
		return
	}
	user, err := uc.service.CreateUser(input)
	if err != nil {
		handler.HandlerErrorResponse(ctx, err.Error())
		return
	}

	handler.HandlerSuccessResponse(ctx, "Success create user", user)
}

func (uc userController) UpdateUser(ctx *gin.Context) {
	var input schema.UpdateUserReq
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&input); err != nil {
		handler.HandlerErrorResponse(ctx, "Record not found")
		return
	}

	user, err := uc.service.UpdateUser(id, input)
	if err != nil {
		handler.HandlerErrorResponse(ctx, "Record not found")
		return
	}

	handler.HandlerSuccessResponse(ctx, "Success update user", user)
}

func (uc userController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uc.service.DeleteUser(id)
	if err != nil {
		handler.HandlerErrorResponse(ctx, "Record not found")
		return
	}

	handler.HandlerSuccessResponse(ctx, "Success delete user", user)
}
