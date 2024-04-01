package jwt_local

import (
	"fmt"
	"time"

	"github.com/Federico-Orsi/Go/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(user models.User)(string, int64, error) {

token := jwt.New(jwt.SigningMethodHS256)
exp := time.Now().Add(time.Second * 60).Unix()
claims := token.Claims.(jwt.MapClaims)

claims["user_id"] = fmt.Sprint(user.ID)
claims["exp"] = exp

t, err := token.SignedString([]byte("s3cr3t"))

if err != nil {
	return "", 0, err
}

return t, exp, nil
}