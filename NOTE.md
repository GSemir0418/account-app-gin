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

TDD tag crud

基本逻辑打通了，下面开始数据库设计与实现
一个 user 可以创建多个 tag 和 item（一对多）
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
一个 item 可以属于多个 tag（多对多）
使用连接表来定义多对多的关系
这种关系声明在一个表的一个字段就可以
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
	Tags       []Tag     `gorm:"many2many:item_tags;"`
  
}
```
// Tags   []Tag `gorm:"many2many:item_tags;foreignKey:ID;joinForeignKey:ItemID;References:ID;joinReferences:TagID"`
foreignKey:ID指的是Item模型的ID字段作为连接表（item_tags）的外键。
joinForeignKey:ItemID指的是在连接表中用于指向Item记录的字段名。
References:ID指的是Tag模型中的ID字段作为参照。
joinReferences:TagID指的是在连接表中用于指向Tag记录的字段名。
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
// 创建一个 tag
	tag := &database.Tag{
		UserID: user.ID,
		Sign:   "⌚️",
		Name:   "电子产品",
		Kind:   "expenses",
	}
	tx = database.DB.Create(tag)
	if tx.Error != nil {
		t.Fatal("Create tag failed:", tx.Error)
	}

	// 创建 item
	item := &database.Item{
		Amount:     100,
		Tags:       []database.Tag{tag},
		UserID:     user.ID,
		Kind:       "in_come",
		HappenedAt: time.Now(),
	} 报错：cannot use tag (variable of type *database.Tag) as database.Tag value in array or slice literalcompilerIncompatibleAssign
  
通常，在处理数据库和关联关系时，推荐使用指针类型作为切片的元素类型，因为这可以更好地与ORM工作，并方便处理没有值（nil）的情形。
推荐将item中的Tags定义为指向Tag的指针的切片，

log 测试中的 log 和正常接口日志的log

接口文档