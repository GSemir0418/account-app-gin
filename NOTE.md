### 1 项目初始化

```bash
go mod init account-app-gin      
go get -u github.com/gin-gonic/gin
# 下载
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

> **Gin 的 log**
>
> - Panic/Panicf/Panicln
>
> 当程序运行到一个无法恢复的状态时，你可能想要立即中止程序，并将错误详情记录下来，同时让上层调用者有机会通过recover捕获panic，进行某些清理工作。
>
> - Print/Printf/Println
>
> 用于记录程序的运行状态或者调试信息，这类信息通常不会导致程序中断。
>
> - Fatal/Fatalf/Fatalln
>
> 启动时检查必需的配置文件不存在，无法继续执行程序。

### 2 连接数据库

#### 2.1 连接数据库

在 database package 暴露数据库实例 DB，声明 ConnectDB 方法，使用 gorm 连接数据库，并给 DB 赋值

#### 2.2 数据库设计

- 一对多

一个 user 可以创建多个 tag 和 item
使用外键约束，关联 tag 和 user 的关系即可

```go
type User struct {
	gorm.Model
	Email      string `gorm:"size:255;not null;unique"`
	Tags       []Tag
	Items      []Item
}
type Tag struct {
	gorm.Model
	UserID     uint       `gorm:"not null;index"`
	// ...
  // 只要遵守约定 模型名+ID 那么
  // GORM 能够自动识别这种外键关系，不需要显式地使用 foreignKey 标签指定。
	User       User       `gorm:"foreignKey:UserID"`
}
```

- 多对多

一个 item 可以属于多个 tag，同时一个 tag 下也可以有很多 item

使用连接表来定义多对多的关系，这种关系声明在一个表的一个字段就可以

会自动生成一个连接表

```go
type Tag struct {
	gorm.Model
	UserID     uint       `gorm:"not null;index"`
	User       User       `gorm:"foreignKey:UserID"`
}
type Item struct {
	gorm.Model
	UserID     uint      `gorm:"not null;index"`
	User       User      `gorm:"foreignKey:UserID"`
  // item_tags 多对多连接表的结构通常是由 GORM 自动生成的
  // 它会包含两个字段：item_id 和 tag_id 分别作为外键指向 items 表和 tags 表的主键
  // 通常，在处理数据库和关联关系时，推荐使用指针类型作为切片的元素类型，因为这可以更好地与ORM工作，并方便处理没有值（nil）的情形。因此推荐将 item 中的 Tags 定义为指向 Tag 的指针的切片
	Tags       []*Tag     `gorm:"many2many:item_tags;"`
}
```

若要重写外键，可以使用标签`foreignKey`、`references`、`joinforeignKey`、`joinReferences`。当然，您不需要使用全部的标签，你可以仅使用其中的一个重写部分的外键、引用，例如

```
Tags []Tag `gorm:"many2many:item_tags;foreignKey:ID;joinForeignKey:ItemID;References:ID;joinReferences:TagID"`
```

- `foreignKey:ID` 指的是Item模型的ID字段作为连接表（item_tags）的外键。

- `joinForeignKey:ItemID` 指的是在连接表中用于指向Item记录的字段名。
- `References:ID` 指的是Tag模型中的ID字段作为参照。
- `joinReferences:TagID` 指的是在连接表中用于指向Tag记录的字段名。

相当于如下 sql

```sql
CREATE TABLE `items` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `user_id` BIGINT,
    `created_at` DATETIME,
    `updated_at` DATETIME,
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);
CREATE TABLE `item_tags` (
    `item_id` BIGINT,
    `tag_id` BIGINT,
    PRIMARY KEY (`item_id`, `tag_id`),
    FOREIGN KEY (`item_id`) REFERENCES `items`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`tag_id`) REFERENCES `tags`(`id`) ON DELETE CASCADE
);
```

### 3 环境变量

将数据库连接字符串放入 .env 中，使用 godotenv 库辅助读取这个文件

```bash
go get github.com/joho/godotenv
```

抽离加载 env 文件到 os 中的逻辑，方便正式环境与测试环境调用

```go
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
```

### 4 命令行程序

使用 cobra 构建命令行程序，使开发测试流程实现自动化

```bash
go get -u github.com/spf13/cobra@latest
```

>`lsof -i :8080` 查看当前端口的占用
>
>`kill -9 54603` 终止进程

### 5 单元测试

> t 的常见用法
>
> 1. Error/Errorf - 在测试日志中记录一个错误信息，但不会终止当前的测试用例。
> 2. Fatal/Fatalf - 记录一个错误信息，并终止当前的测试用例
> 3. Log/Logf - 在测试日志中输出信息，但不会引起测试失败。
> 4. Fail/FailNow - 通过调用`Fail`标记测试函数为失败状态，继续执行当前测试用例中的剩余代码；通过调用`FailNow`，立即终止当前测试用例。
> 5. Skip/Skipf/SkipNow - 跳过当前测试用例。

#### 5.1 单元测试初始化

将加载环境变量、连接数据库、初始化 gin 应用与上下文对象、清空数据等工作抽离到 setupTestCase 函数中，每个测试用例执行时调用

下面是清空数据的函数

```go
// 清空 User 表 TRUNCATE 相当于先 DROP 在 CREATE
// 存在外键约束的话，为了保证数据完整性，不能使用 TRUNCATE
// 使用 DELETE 性能很差，所以这里使用更激进的策略，先关闭外键约束检查，清空后再恢复
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
```

#### 5.2 测试流程

调用初始化方法 => 注册路由 => 初始化 w => 准备接口入参 => 发起请求 => 断言响应

### 6 重构

#### 6.1 router

#### 6.2 controller

### 7 TDD

7.1 获取请求参数

c.DefaultQuery 提供了一种方便的方式来直接获取参数值，并为不存在的参数提供默认值。这使得你的代码在处理可选查询参数时更加简洁。
c.Request.URL.Query().Get 这是标准 Go http 包的用法,这是标准 Go http 包的用法，它不自动处理默认值。如果你使用这种方式，你需要自己编写额外的逻辑来为参数提供默认值（如果这是你的需求）。

7.2 抽离 api 接口出入参类型

7.3 整理状态码 整理接口出入参类型

TDD tag crud

7.4 更新接口
使用指针定义更新接口入参。保证灵活性，然后再赋值给 tag 对象
使用 model.updates 而不是 save
在模拟请求时, 使用 model 作为类型的话会自动补充空值，所以要使用更通用的 map[string]string

8 接口文档

9 部署


单测一定要使用前端真实传入的 json 参数，而不是构造结构体

请求体绑定的结构体类型不能是 model 的类型，要自定义一个 request struct，然后自定义校验规则也放在这里

值和地址！晕了已经

删除 tag，首先要删除关联表中的记录，再删除 tag 本身
所以使用事务来保证删除操作的原子性
db.Exec
db.Raw.Scan

新项目
安装 golang 1.21.0
vscode 插件 aldijav.golangwithdidi
配置代理
go env -w GOPROXY=https://goproxy.cn,direct
安装依赖 
go mod tidy
同步环境变量
DB_DSN=user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
同步数据库
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go;./main.exe db migrate:create
启动项目
go run main.go server

windows mysql 启动与停止
win+r services.msc

- docker 配置 proxy
sudo vi /etc/docker/daemon.json
```json
{
  "registry-mirrors": ["https://docker.xxx"]
}
```
重启 docker
sudo systemctl restart docker

- 用户系统
验证码
	建表 同步数据库 写controller
	
发送邮件
	go get gopkg.in/gomail.v2
登录
	session controller
jwt
