package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func initDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GlobalConfig.GetString("mysql.user"),
		GlobalConfig.GetString("mysql.pass"),
		GlobalConfig.GetString("mysql.host"),
		GlobalConfig.GetInt64("mysql.port"),
		GlobalConfig.GetString("mysql.dbname"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("connecting mysql: %v", err)
	}
	//DB = db.Debug()
	DB = db
}
