package database

import (
	"log"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	if DB != nil {
		return
	}
	// 加载 .env 文件
	// 因为测试环境的工作目录与你的 .env 文件所在的目录不同导致的。
	// godotenv.Load() 默认会在当前工作目录下查找 .env 文件。如果你的测试运行在一个不同的目录，需要确保正确地指定 .env 文件的路径。

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

func TruncateTables(t *testing.T, tables []string) {
	// 禁用外键检查
	err := DB.Exec("SET FOREIGN_KEY_CHECKS=0;").Error
	if err != nil {
		t.Fatalf("Failed to disable foreign key checks: %v", err)
	}

	// 清空所有给定的表
	for _, table := range tables {
		if err = DB.Exec("TRUNCATE TABLE " + table + ";").Error; err != nil {
			t.Fatalf("Failed to truncate table %s: %v", table, err)
		}
	}

	// 重新启用外键检查
	err = DB.Exec("SET FOREIGN_KEY_CHECKS=1;").Error
	if err != nil {
		t.Fatalf("Failed to enable foreign key checks: %v", err)
	}
}
