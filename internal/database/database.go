package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dsn string) {
	if DB != nil {
		return
	}
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Fail to connect database!", err)
	}

	DB = database
}

func Migrate() {
	if DB != nil {
		// 迁移数据库
		DB.AutoMigrate(&Item{})
	}
}
