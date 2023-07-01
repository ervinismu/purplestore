package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginationMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		page, err := strconv.Atoi(ctx.Query("page"))
		if err != nil {
			page = 1
		}

		pageSize, err := strconv.Atoi(ctx.Query("pageSize"))
		if err != nil {
			pageSize = 10
		}

		ctx.Set("page", page)
		ctx.Set("pageSize", pageSize)

		ctx.Next()
	}
}
