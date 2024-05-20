package database

import "gorm.io/gorm"

type User struct {
	gorm.Model        // gorm.Model includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Email      string `gorm:"size:255;not null;unique"`
}
