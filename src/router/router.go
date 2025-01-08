package router

import "github.com/gin-gonic/gin"

func BandHttpRouter(r *gin.Engine) {
	r.GET()
	r.POST()
	r.PUT()
	r.DELETE()
	r.Group()
}
