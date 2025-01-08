package main

import (
	"gin_exampl/src/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	engine := gin.New()
	router.BandHttpRouter(engine)

	engine.Run("0.0.0.0:8000")

	/*

		https://github.com/gin-gonic/gin
		https://github.com/gogf/gf

		https://github.com/duke-git/lancet
		https://github.com/go-resty/resty
		https://github.com/samber/lo

		https://github.com/go-gorm/gorm
		https://github.com/qax-os/excelize
		https://github.com/robfig/cron
		https://github.com/dgraph-io/badger
		https://github.com/casbin/casbin
		https://github.com/ent/ent
		https://github.com/spf13/viper
		https://github.com/sirupsen/logrus
		https://github.com/uber-go/zap
		https://github.com/brianvoe/gofakeit
		https://github.com/jinzhu/copier
		https://github.com/dtm-labs/dtm
		https://github.com/apache/incubator-seata-go

	*/
}
