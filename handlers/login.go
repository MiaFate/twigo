package handlers

import (
	"encoding/json"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/jwt"
	"github.com/miafate/twigo/models"
)

func Login(ctx *gin.Context) models.ApiResponse[models.LoginResponse] {
	var t models.Usuario
	var r models.ApiResponse[models.LoginResponse]
	r.Status = 400

	// body := ctx.Value(models.Key("body")).(string)
	// err := json.Unmarshal([]byte(body), &t)
	body, _ := io.ReadAll(ctx.Request.Body)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Usuario o contraseña incorrectos " + err.Error()
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "El email es requerido"
		return r
	}

	userData, existe := bd.Login(t.Email, t.Password)
	if !existe {
		r.Message = "Usuario o contraseña incorrectos"
		return r
	}

	jwtKey, err := jwt.GenerateJWT(userData)
	if err != nil {
		r.Message = "Error al generar el token " + err.Error()
		return r

	}

	ctx.SetCookie("token", jwtKey, 24*60*60, "/", os.Getenv("DOMAIN"), false, true)
	ctx.Header("Access-Control-Allow-Origin", "*")
	r.Status = 200
	r.Data = models.LoginResponse{
		Token: jwtKey,
	}

	return r
}
