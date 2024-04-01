package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	User_id uint `gorm:"not null" json:"user_id"`
	OrderProducts []OrderProduct `gorm:"foreignKey:CartID" json:"OrderProducts"`
	User User
}


func NewCart(userID uint) *Cart {
	return &Cart{
		User_id:   userID,
		OrderProducts: []OrderProduct{}, // Inicializar con un slice vacío
	}
}