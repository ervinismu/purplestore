package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		url := ctx.Request.URL.Path

		log.WithFields(log.Fields{
			"client_id": clientIP,
			"url":       url,
			"method":    method,
		}).Info("http request")

		ctx.Next()
	}
}
