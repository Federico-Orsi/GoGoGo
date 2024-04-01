package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title string `gorm:"type:varchar(100);not nul;unique_index" json:"title"`
	Description  string `gorm:"not nul" json:"description"`
	Done     bool `gorm:"default:false" json:"done"`
	User_id  uint `gorm:"not nul" json:"user_id"`
}