package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// 获取当前文件的绝对路径
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	// 获取环境变量 GIN_ENV，默认值为 "dev"
	env := os.Getenv("GIN_ENV")
	if env == "" {
		env = "dev"
	}

	// 确定要加载的 .env 文件
	envFile := filepath.Join(basepath, "..", ".env."+env)
	// 加载环境文件
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file for environment %s: %v", env, err)
	}
}
