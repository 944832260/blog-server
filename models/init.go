package models

import (
	"blog-server/config"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"fmt"
)

var Config config.Config
var db *gorm.DB
var err error

func init() {
	Config, _ = config.Init()
	dbConfig := Config.DataBase

	dbURI := []string {
		dbConfig.Username,
		":",
		dbConfig.Password,
		"@",
		dbConfig.Server,
		"/",
		dbConfig.DBname,
		"?charset=utf8&parseTime=True&loc=Local",
	}

	db, err = gorm.Open("mysql",strings.Join(dbURI, ""))
	//不限制最大打开连接数
	db.DB().SetMaxIdleConns(0)
	//db.CreateTable(
	//	&User{},
	//	&Article{},
	//)
	fmt.Printf("5555")
}

func GetDB() *gorm.DB {
	return db
}