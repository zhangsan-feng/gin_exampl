package middleware

import "github.com/gin-gonic/gin"

func BindMiddleware(r *gin.Engine) {
	loggerMiddleware(r)
	userVerificationMiddleware(r)
	corsMiddleware(r)
}
