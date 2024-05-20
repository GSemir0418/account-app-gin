package database

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	UserID     uint      `gorm:"not null"`                           // 使用 uint 反映自增长性质
	Amount     int       `gorm:"not null"`                           // 保持为 int 类型，不变
	TagIDs     string    `gorm:"type:text;not null"`                 // 将数组改为字符串，可以存储为逗号分隔的数值
	Kind       string    `gorm:"size:100;not null"`                  // varchar(100) 映射
	HappenedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"` // MySQL 支持 CURRENT_TIMESTAMP
	CreatedAt  time.Time `gorm:"<-:create"`                          // 让 GORM 自动处理
	UpdatedAt  time.Time `gorm:"<-:update"`                          // 让 GORM 自动处理
}
