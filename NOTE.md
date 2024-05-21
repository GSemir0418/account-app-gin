go mod init account-app-gin      
go get -u github.com/gin-gonic/gin

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

连接逻辑，连接实例，定义 model
数据库建库
env 保存连接字符串
go get github.com/joho/godotenv

命令行程序
go get -u github.com/spf13/cobra@latest

lsof -i :8080
kill -9 54603

go test ./...

refactor router

init controller interface

setup test

加载环境变量逻辑抽离 for test

TDD item create

抽离清空数据表的逻辑 for test

c.DefaultQuery 提供了一种方便的方式来直接获取参数值，并为不存在的参数提供默认值。这使得你的代码在处理可选查询参数时更加简洁。
c.Request.URL.Query().Get 这是标准 Go http 包的用法,这是标准 Go http 包的用法，它不自动处理默认值。如果你使用这种方式，你需要自己编写额外的逻辑来为参数提供默认值（如果这是你的需求）。

抽离 api 接口出入参类型



log 测试中的 log 和正常接口日志的log

接口文档