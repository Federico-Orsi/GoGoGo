package models

import "gorm.io/gorm"

type Pet struct {
	gorm.Model

	Animal string `gorm:"not nul" json:"animal"`
    Name string `gorm:"not nul" json:"name"`


}
