package routes

import (
	"context"
	"log"

	"github.com/Federico-Orsi/Go/proto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Grpc_Route(app *fiber.App) {

	app.Get("/go", func(c *fiber.Ctx) error {

		clientServer, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Fatalf("No se pudo concectar: %v", err)
		}

		defer clientServer.Close()

		client := proto.NewBridgeServiceClient(clientServer)

		nombre := c.Query("nombre")
		apellido := c.Query("apellido")


		request := &proto.PersonRequest{

			Person: &proto.Person{
				Nombre:   nombre,
				Apellido: apellido,
			},
		}


		response, err := client.CallPerson(context.Background(), request)

		if err != nil {
			log.Fatalf("Falló doCallPerson(): %v", err)
		}

		// log.Printf("Respuesta de CallPerson(): %v\n", response.GetHi())

		return c.SendString(response.GetHi())
	})


}