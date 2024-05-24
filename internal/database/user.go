package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"<-:create" json:"createdAt"` // 让 GORM 自动处理
	UpdatedAt time.Time      `gorm:"<-:update" json:"updatedAt"` // 让 GORM 自动处理
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	// gorm.Model                // gorm.Model includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Email string `gorm:"size:255;not null;unique" json:"email"`
	Tags  []*Tag
	Items []*Item
}
