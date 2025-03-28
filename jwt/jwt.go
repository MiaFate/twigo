package jwt

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/miafate/twigo/models"
)

func GenerateJWT(t models.Usuario) (string, error) {
	jwtSign := os.Getenv("JWTSIGN")
	miClave := []byte(jwtSign)
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.Id.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
