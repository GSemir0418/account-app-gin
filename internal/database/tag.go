package database

import (
	"time"

	"gorm.io/gorm"
)

// type Tag struct {
// 	ID        uint           `gorm:"primaryKey" json:"id"`
// 	CreatedAt time.Time      `gorm:"<-:create" json:"createdAt"` // 让 GORM 自动处理
// 	UpdatedAt time.Time      `gorm:"<-:update" json:"updatedAt"` // 让 GORM 自动处理
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
// 	// gorm.Model                // gorm.Model已经包含了ID, CreatedAt, UpdatedAt, DeletedAt字段，但要注意gorm.Model中的ID不是SERIAL类型，而是uint类型
// 	UserID uint `gorm:"not null;index" json:"userId"` // 因为Gorm主键ID默认为uint类型，所以这里把user_id也定义为uint
// 	// User       *User  `gorm:"foreignKey:UserID"`             // 用于软删除功能，如果不需要软删除，就不需要这个字段
// 	Name string `gorm:"size:50;not null" json:"name"`  // VARCHAR(50) NOT NULL
// 	Sign string `gorm:"size:10;not null" json:"sign"`  // VARCHAR(10) NOT NULL
// 	Kind string `gorm:"size:100;not null" json:"kind"` // VARCHAR(100) NOT NULL
// }

// Tag 与 Item 改为一对多的关系

type Tag struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"<-:create" json:"createdAt"` // 让 GORM 自动处理
	UpdatedAt time.Time      `gorm:"<-:update" json:"updatedAt"` // 让 GORM 自动处理
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	Name string `gorm:"size:50;not null" json:"name"`  // VARCHAR(50) NOT NULL
	Sign string `gorm:"size:10;not null" json:"sign"`  // VARCHAR(10) NOT NULL
	Kind string `gorm:"size:100;not null" json:"kind"` // VARCHAR(100) NOT NULL

	Items  []*Item `json:"items"`
	UserID uint    `gorm:"not null;index" json:"userId"` // 因为Gorm主键ID默认为uint类型，所以这里把user_id也定义为uint
}
