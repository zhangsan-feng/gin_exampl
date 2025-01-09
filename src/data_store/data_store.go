package data_store

import (
	"github.com/gogf/gf/v2/database/gdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DataStore struct {
	DBName     string
	DataSource string
}

func InitGormStore() *gorm.DB {
	db, err := gorm.Open(mysql.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接数据库: %v", err)
	}
	return db
}

func InitEntStore() *DataStore {

	return nil
}

func InitGfStore() {
	err := gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Type: "mysql",
				Link: "sqlite::@file(/tmp/gf.db)",
			},
		},
	})
	if err != nil {
		return
	}

}
