package database

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model        // gorm.Model已经包含了ID, CreatedAt, UpdatedAt, DeletedAt字段，但要注意gorm.Model中的ID不是SERIAL类型，而是uint类型
	UserID     uint   `gorm:"not null;index"`    // 因为Gorm主键ID默认为uint类型，所以这里把user_id也定义为uint
	User       User   `gorm:"foreignKey:UserID"` // 用于软删除功能，如果不需要软删除，就不需要这个字段
	Name       string `gorm:"size:50;not null"`  // VARCHAR(50) NOT NULL
	Sign       string `gorm:"size:10;not null"`  // VARCHAR(10) NOT NULL
	Kind       string `gorm:"size:100;not null"` // VARCHAR(100) NOT NULL
}
