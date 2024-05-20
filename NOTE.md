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