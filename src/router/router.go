package router

import (
	"gin_exampl/src/middleware"
	"github.com/gin-gonic/gin"
)

func BandHttpRouter(r *gin.Engine) {
	middleware.BindMiddleware(r)
}
