package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
    // gorm por defecto al crear una nueva tabla, se basa en el nombre del struct pero lo coloca en plural!!
	// gorm por defecto crea automaticamente un nuevo ID por cada elemento creado. A este campo ID siempre lo considera por defecto como la Primary Key!!
	FirstName string `gorm:"not null" json:"first_name"`
    LastName string `gorm:"not null" json:"last_name"`
	Email string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Task []Task `json:"tasks"` //cuando hacemos referencia a otro modelo dentro de un struct, GORM automaticamente lo considera como Foreign Key e infiere el tipo de relación entre ambos structs o elementos!!
	Rol string `gorm:"not null" json:"rol"`

}




