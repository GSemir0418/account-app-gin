package database

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	// gorm.Model
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `gorm:"<-:create" json:"createdAt"` // 让 GORM 自动处理
	UpdatedAt  time.Time      `gorm:"<-:update" json:"updatedAt"` // 让 GORM 自动处理
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	UserID     uint           `gorm:"not null;index" json:"userId"`     // 使用 uint 反映自增长性质
	Tags       []*Tag         `gorm:"many2many:item_tags;" json:"tags"` // 多对多关系，通过 item_tags 表连接 Item 和 Tag
	Amount     int            `gorm:"not null" json:"amount"`           // 保持为 int 类型，不变
	Kind       string         `gorm:"size:100;not null" json:"kind"`    // varchar(100) 映射
	HappenedAt time.Time      `gorm:"not null" json:"happenedAt"`       // MySQL 支持 CURRENT_TIMESTAMP
	// User       *User     `gorm:"foreignKey:UserID"`             // 添加 User 字段，指定外键为 UserID
}
