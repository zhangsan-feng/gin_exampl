package middleware

import "github.com/gin-gonic/gin"

func Middleware(r *gin.Engine)  {
	loggerMiddleware(r)
	userVerificationMiddleware(r)
}
