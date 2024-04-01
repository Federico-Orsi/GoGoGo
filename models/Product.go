package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Title string `gorm:"not null;uniqueIndex" json:"title"`
    Description string `gorm:"not null" json:"description"`
	Price int `gorm:"not null" json:"price"`
	Category string `gorm:"not null" json:"category"`
	Stock int `gorm:"not null" json:"stock"`

}