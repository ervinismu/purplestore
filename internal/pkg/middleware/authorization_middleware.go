package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/ervinismu/purplestore/internal/pkg/handler"
	"github.com/ervinismu/purplestore/internal/pkg/reason"
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
			handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.Unauthorized)
			ctx.Abort()
			return
		}
	}
}
