package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	if DB != nil {
		return
	}
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 从环境变量中获取数据库连接字符串
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN is not set")
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
		err := DB.AutoMigrate(&User{}, &Item{})
		if err != nil {
			log.Fatal("Fail to migrate database!", err)
		}
	}
}
