package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(sub string, obj string, action string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// sub : the user want to access resource
		// obj : resource that going to be accesses
		// action : operation that the user will perform to the resource

		res, _ := enforcer.Enforce(sub, obj, action)
		if res {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "you are not authorized to perform this action."})
			ctx.Abort()
			return
		}
	}
}
