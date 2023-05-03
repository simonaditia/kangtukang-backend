package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	// ID       uint `gorm:"primaryKey"`
	// ID       uint   `json:"id" gorm:"primaryKey"`
	Kategori string `json:"kategori"`
}
