package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerSuccessResponse(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(
		http.StatusOK, NewRespBodyData("Success", message, data))
}

func HandlerErrorResponse(ctx *gin.Context, message string) {
	ctx.JSON(
		http.StatusOK, NewRespMessage("Error", message))
}
