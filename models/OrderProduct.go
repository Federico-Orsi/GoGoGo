package models

import "gorm.io/gorm"

type OrderProduct struct {
	gorm.Model

	CartID uint `gorm:"not null" json:"cartID"`
	Product_id uint `gorm:"not null" json:"product_id"`
	Qty int `gorm:"not null" json:"qty"`
    Product Product
}