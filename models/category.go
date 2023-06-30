package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	// ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique"`
	Users []User `gorm:"many2many:user_categories;"`
}
