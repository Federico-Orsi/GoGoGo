package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWT_middleware() func(*fiber.Ctx) error {

    return jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte("s3cr3t")},
    })
}