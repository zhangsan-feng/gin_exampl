package router

import (
	"admin_backend/api"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"log"
)

func cors(r *gin.Context) {
	r.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	r.Writer.Header().Set("Access-Control-Allow-Credentials", "*")
	r.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	r.Writer.Header().Set("Access-Control-Allow-Methods", "*")

	if r.Request.Method == "OPTIONS" {
		r.AbortWithStatus(204)
		return
	}

	r.Next()
}

func record(r *gin.Engine) {
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {

			logFormat := map[string]interface{}{
				"path":       params.Path,
				"method":     params.Method,
				"status":     params.StatusCode,
				"ip":         params.ClientIP,
				"user_agent": params.Request.UserAgent(),
			}
			jsonData, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(logFormat)
			if err != nil {
				log.Println("gin log format error")
			}
			log.Println(string(jsonData))

			return ""
		},
	}))
}

func userAuth(ctx *gin.Context) {

}

func NewHttpRouter(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(func(context *gin.Context) { cors(context) })
	r.Use(func(context *gin.Context) { userAuth(context) })
	record(r)

	r.GET("/test", func(context *gin.Context) {
		context.JSON(200, "")
	})
	r.Any("/create_table", api.Database.CreatePostgres)

}
