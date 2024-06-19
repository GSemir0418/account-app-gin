package database

import (
	"time"

	"gorm.io/gorm"
)

type ValidationCode struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"size:255;not null" json:"email"`
	Code      string         `gorm:"size:255;not null" json:"code"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	UpdatedAt time.Time      `gorm:"<-:update" json:"updatedAt"`
	CreatedAt time.Time      `gorm:"<-:create" json:"createdAt"`
	UsedAt    *time.Time     `json:"usedAt"`
}
