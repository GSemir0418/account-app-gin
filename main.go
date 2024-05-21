package main

import (
	"account-app-gin/cmd"
	"account-app-gin/config"
)

func main() {
	// 加载环境变量
	config.LoadEnv()
	cmd.Run()
}
