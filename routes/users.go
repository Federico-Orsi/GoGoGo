package routes

import (
	"fmt"
	"log"
	"reflect"

	"github.com/Federico-Orsi/Go/controllers"
	"github.com/Federico-Orsi/Go/db"
	"github.com/Federico-Orsi/Go/jwt_local"
	"github.com/Federico-Orsi/Go/middlewares"
	"github.com/Federico-Orsi/Go/models"
	"github.com/gofiber/fiber/v2"

	"golang.org/x/crypto/bcrypt"
)




func Users_Route(app *fiber.App) {

	app.Get("/user/:id", func(c *fiber.Ctx) error {
      //hardcodeo y muestro id del user!!
		id := c.Params("id")
		name := c.Query("name")

		log.Println(name)


		return c.JSON("User id: " + id)
	})

	app.Get("/registered_user", middlewares.JWT_middleware(), middlewares.Auth,controllers.Registered_user)


	app.Get("/findUser/:id", func(c *fiber.Ctx) error {

	// Obtener el token de la cabecera de autorización
	// tokenString := c.Get("Authorization")

	// token, err := (tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	return secretKey, nil
	// })

	// if err != nil || !token.Valid {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"mensaje": "Token no válido"})
	// }


		id := c.Params("id")
        var user models.User

		db.DB.First(&user, id)

		if user.ID == 0 {
			//log.Fatalf("User not found!!") Ojo que el log.Fatalf() interrumpe la ejecución del programa!!
			return c.Status(fiber.StatusNotFound).JSON("Usuario no encontrado!!")
		}

		return c.JSON(user)
	})


	app.Get("/findUserWithTasks/:id", func(c *fiber.Ctx) error {


			id := c.Params("id")
			var user models.User

			db.DB.First(&user, id)

			if user.ID == 0 {
				return c.Status(fiber.StatusNotFound).JSON("Usuario no encontrado!!")
			}
            //aquí Asociamos ambos modelos (structs) y the fantastic GORM infiere las relaciones automaticamente!!
			db.DB.Model(&user).Association("Task").Find(&user.Task)

			return c.JSON(user)
		})



	app.Get("/users", func(c *fiber.Ctx) error {

		var users []models.User

        db.DB.Find(&users)

		return c.JSON(users)//Importante saber que Go es distinto con respecto a la respuesta que se manda en el JSON!! Se coloca directamente la variable creada (var). En caso de querer almacenar el resultado de la query en memoria, solamente esto lo hariamos para chequear y manejar un posible error en la consulta a la DB!!
	})

	app.Post("/login_con_map", func(c *fiber.Ctx) error {
    // OJO!! este enfoque puede ser menos seguro y menos claro que utilizar una estructura definida, especialmente si tienes más propiedades o si necesitas realizar verificaciones más complejas.

    // Parsear el cuerpo JSON a un mapa
			var requestData map[string]interface{}
			if err := c.BodyParser(&requestData); err != nil {
				return err
			}

			// Acceder a las propiedades específicas
			email, emailExists := requestData["email"].(string)
			password, passwordExists := requestData["password"].(string)

			// Verificar si las propiedades existen y son del tipo correcto
			if !emailExists || !passwordExists {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"mensaje": "Email y/o Password no proporcionados correctamente"})
			}

			// Hacer algo con las propiedades específicas (en este caso, imprimir)
			println("Email:", email)
			println("Password:", password)

			// Aquí puedes realizar la lógica de autenticación basada en Email y Password

			return c.JSON(requestData)


	})


	app.Post("/login", func(c *fiber.Ctx) error {

		var login models.Login

		if err := c.BodyParser(&login); err != nil {
			return err
		}

		user := models.User{}
        if err := db.DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"mensaje": "Credenciales incorrectas"})
		}


        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"mensaje": "Credenciales incorrectas"})
		}


		fmt.Println("tipo del UserID", reflect.TypeOf(user.ID))
		token, exp, err := jwt_local.CreateJWT(user)

		if err != nil{
			return err
		}

		//check here user´s rol and creat Cart if necessary:
        if user.Rol == "User" || user.Rol == "Premium" {
		//Let's check here if the user already have a Cart created:
			var existingCart models.Cart

		    if err := db.DB.Where("user_id = ?", user.ID).First(&existingCart).Error; err != nil {

				new_cart := models.NewCart(user.ID)

				db.DB.Create(&new_cart)
				if new_cart.ID == 0 {

				return c.Status(fiber.StatusBadRequest).JSON("Some error with the cart")
			}
				return c.JSON(fiber.Map{"User":user, "Token":token, "Exp":exp, "Cart": new_cart})
			}

		}


        return c.JSON(fiber.Map{"User":user, "Token":token, "Exp":exp})
	})

	app.Post("/register", func(c *fiber.Ctx) error {

		user := models.User{}

		if err := c.BodyParser(&user); err != nil {
			return err
		}

        //Cuando buscas un item de esta manera en gorm, te devuelve un error != nil cuando no lo encuentra en la db!!
		// en cambio si lo encuentra, el error es nil !!
		if err := db.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost )
		if err != nil{
			return err
		}

		user.Password = string(hash)

        log.Println(hash)

		 log.Println(user.Password) // para acceder a los valores de una instancia que hemos creado, ya sea con & o new , se realiza con "." , no hace falta colocar .value como haciamos en JS!!

        db.DB.Create(&user)

		return c.JSON(user)
		}
        return c.Status(fiber.StatusUnauthorized).JSON("Usuario ya registrado!! Para poder registrarse nuevamente debe hacerlo con un nuevo e-mail.")

	})


	app.Delete("/user/:id", func(c *fiber.Ctx) error {

		user := models.User{}
		// si el userName := user.Email lo colocas aquí antes del Find() se almacena en Memoria con el campo vacío!!
        id := c.Params("id")
        db.DB.Find(&user, id)
		userName := user.Email // es importante si queremos acceder a un campo del struct que lo coloquemos luego de haber hecho el Find() en la db!!

		result := db.DB.Delete(&user)
		if result.Error != nil {
			return c.JSON("Usuario no encontrado!!")
			//panic("Error al eliminar usuario: " + result.Error.Error())
		} else {
			//return c.JSON("El usuario: " + userName + ", ha sido eliminado correctamente.")
            return c.JSON(fiber.Map{"Message": fmt.Sprintf(`El usuario: %s, ha sido eliminado correctamente.`, userName)})
		}

        // Verificar si se eliminó correctamente (esto lo propuso chatGpt)
		// if result.RowsAffected == 0 {
		// 	fmt.Println("")

		// }
	})

}