package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=postgres password=eldiego10 dbname=test port=5432" //DSN(data source name. Es otra forma de llamar al string de conexion)
var DB *gorm.DB  // ojo acá va sin el = de asignación!! Luego revisar bien el motivo...

func DB_Connection() {
var err error

DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

if err != nil {
	log.Fatal(err)
	} else {
   log.Println("Db connected!!")
	}

}