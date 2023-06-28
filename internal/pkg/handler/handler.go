package handler

import (
	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Handle response error
func ResponseError(ctx *gin.Context, statusCode int, message string) {
	resp := ResponseBody{
		Status:  "error",
		Message: message,
	}
	ctx.JSON(statusCode, resp)
}

// Handle response success
func ResponseSuccess(ctx *gin.Context, statusCode int, message string, data interface{}) {
	resp := ResponseBody{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	ctx.JSON(statusCode, resp)
}
