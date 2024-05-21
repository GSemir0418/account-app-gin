package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// 获取当前文件的路径
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	// 加载距离当前文件上级的 .env 文件
	err := godotenv.Load(filepath.Join(basepath, "..", ".env"))
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
