package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/Federico-Orsi/Go/db"
	"github.com/Federico-Orsi/Go/models"
	"github.com/Federico-Orsi/Go/routes"
	"github.com/Federico-Orsi/Go/third_party"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)



func main() {

	err := godotenv.Load(); if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	PORT := os.Getenv("PORT")
	fmt.Println("El tipo de dato de Port es:", reflect.TypeOf(PORT))

	db.DB_Connection()

    //db.DB.Migrator().AddColumn(&models.User{}, "Rol")
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.Pet{})
	db.DB.AutoMigrate(models.Product{})
	db.DB.AutoMigrate(models.OrderProduct{})
	db.DB.AutoMigrate(models.Cart{})



	app := fiber.New()

	routes.TestRoute(app)
	routes.Users_Route(app)
	routes.Grpc_Route(app)
	routes.Go_routines_Route(app)
	routes.Kafka_Route(app)
	routes.Decorator_Route(app)
	routes.Adapter_Route(app)
	routes.Builder_Route(app)
	routes.Observer_Route(app)
	routes.Tasks_Route(app)
	routes.Products_Route(app)
	routes.Orders_Route(app)
	routes.Carts_Route(app)


	// random_id := uuid.New().String()

	var hello string = "Come onnnnnn mothafuckeeeerrrr!!"
	log.Println(hello)
	fmt.Printf("fmt printf %v\n", hello)
	fmt.Println(hello)
	// fmt.Println(random_id)

	fmt.Println("Hello World!! Go Go Gooooo for it!! It's fucking time!!")
	fmt.Println(`Go Go Gooooo get it Mannn!!`)

	basic_goku := third_party.BasicGoku{}
	third_party.ShowGoku(basic_goku) // Se le pasa el struct que implementa la interfaz!!


	// fmt.Printf("New id: %v.", random_id)
    // fmt.Printf("Port listening in port %v\n", port)



    //app.Listen(":4300") //Muy importante en GO, poner el app.Listen() debajo de las rutas!! Sino no funciona!!
	//en Fiber todo lo que se coloca por debajo del app.Listen() queda ignorado, entonces para poder imprimir en pantalla el mje de "Esuchando en puerto tal..." tuve que utilizar una goRutine
	go func() {
        err := app.Listen(":" + PORT)
        if err != nil {
            // Manejamos cualquier error que ocurra al iniciar el servidor
            fmt.Printf("Error al iniciar el servidor en el puerto %s: %v\n", PORT, err)
        }
    }()

    fmt.Printf("Servidor escuchando en el puerto %s\n", PORT)

    // con select{} Esperamos a que la aplicación termine (esto evitará que el programa principal termine antes de que la aplicación Fiber)
    // el select{} se usa principalmente para esperar la comunicacion entre canales, pero si no hay canales existentes bloquea en este caso la func main() indefinidamente hasta que la goRutine finalice!!
	select {}  // otras formas de lograr esto sería usando WaitGroups o timeSleep()

}
