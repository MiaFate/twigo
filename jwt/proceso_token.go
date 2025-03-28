package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

var Email string
var IdUser string

func ProcesoToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	token, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.UserExist(claims.Email)
		if encontrado {
			Email = claims.Email
			IdUser = string(claims.Id.Hex())
		}
		return &claims, encontrado, IdUser, nil
	}

	if !token.Valid {
		return &claims, false, string(""), errors.New("token invalido")
	}

	return &claims, false, string(""), err
}
