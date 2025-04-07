package database

import (
	"admin_backend/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

type Database struct{}

func (Database) CreatePostgres(r *gin.Context) {
	Database{}.createDatabase("manager").createTable()
	r.JSON(200, gin.H{"data": "success"})
}

func (Database) TruncateTable(ctx *gin.Context) {
	db_name := utils.GetParams(ctx, "db_name").(string)
	table_name := utils.GetParams(ctx, "table_name").(string)
	address := "host=192.168.56.39 user=postgres password=root dbname='" + db_name + "' port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	conn, connErr := sql.Open("postgres", address)
	if connErr != nil {
		log.Println("connection postgresql error", connErr)
	}
	conn.Exec("drop table" + table_name)
	conn.Close()
}

func (Database) DropDatabase(ctx *gin.Context) {
	db_name := utils.GetParams(ctx, "db_name").(string)

	address := "host=192.168.56.39 user=postgres password=root dbname='" + db_name + "' port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	conn, connErr := sql.Open("postgres", address)
	if connErr != nil {
		log.Println("connection postgresql error", connErr)
	}
	conn.Exec("drop database" + db_name)
	conn.Close()
}
