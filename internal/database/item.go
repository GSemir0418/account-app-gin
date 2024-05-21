package database

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	UserID     uint      `gorm:"not null;index"`       // 使用 uint 反映自增长性质
	User       User      `gorm:"foreignKey:UserID"`    // 添加 User 字段，指定外键为 UserID
	Tags       []*Tag    `gorm:"many2many:item_tags;"` // 多对多关系，通过 item_tags 表连接 Item 和 Tag
	Amount     int       `gorm:"not null"`             // 保持为 int 类型，不变
	Kind       string    `gorm:"size:100;not null"`    // varchar(100) 映射
	HappenedAt time.Time `gorm:"not null"`             // MySQL 支持 CURRENT_TIMESTAMP
	CreatedAt  time.Time `gorm:"<-:create"`            // 让 GORM 自动处理
	UpdatedAt  time.Time `gorm:"<-:update"`            // 让 GORM 自动处理
}
