package middleware

import (
	"net/http"
	"strings"

	"github.com/ervinismu/purplestore/internal/pkg/handler"
	"github.com/ervinismu/purplestore/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type AccessTokenVerifier interface {
	VerifyAccessToken(tokenString string) (string, error)
}

func AuthMiddleware(tokenMaker AccessTokenVerifier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get token from header
		accessToken := tokenFromHeader(ctx)
		if accessToken == "" {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.Unauthorized)
			ctx.Abort()
			return
		}

		// verify
		sub, err := tokenMaker.VerifyAccessToken(accessToken)
		if err != nil {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.Unauthorized)
			ctx.Abort()
			return
		}

		// attach to request
		ctx.Set("user_id", sub)

		// continue
		ctx.Next()
	}
}

func tokenFromHeader(ctx *gin.Context) string {
	var accessToken string

	// Bearer xxxyyzz
	bearerToken := ctx.Request.Header.Get("Authorization")
	fields := strings.Fields(bearerToken)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	}

	return accessToken
}
