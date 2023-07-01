package handler

import "github.com/gin-gonic/gin"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// response success
func ResponseSuccess(
	ctx *gin.Context,
	statusCode int,
	message string,
	data interface{},
) {
	resp := Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	ctx.JSON(statusCode, resp)
}

// response error
func ResponseError(
	ctx *gin.Context,
	statusCode int,
	message string,
) {
	resp := Response{
		Status:  "error",
		Message: message,
	}

	ctx.JSON(statusCode, resp)
}
