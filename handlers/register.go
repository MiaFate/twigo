package handlers

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func Register(ctx *gin.Context) models.ApiResponse[models.Usuario] {
	var t models.Usuario
	var r models.ApiResponse[models.Usuario]
	r.Status = 400
	fmt.Println("Entre a Registro")

	// body := ctx.Value(models.Key("body")).(string)
	body, _ := io.ReadAll(ctx.Request.Body)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Email no puede estar vacio"
		fmt.Println(r.Message)
		return r
	}
	if len(t.Password) < 6 {
		r.Message = "Debe especificar una contraseña de al menos 6 caracteres"
		fmt.Println(r.Message)
		return r
	}
	_, encontrado, _ := bd.UserExist(t.Email)
	if encontrado {
		r.Message = "El usuario ya existe con ese email"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.AddRegistry(t)
	if err != nil {
		r.Message = "Ocurrio un error al intentar realizar el registro del usuario " + err.Error()
		fmt.Println(r.Message)
		return r
	}
	if !status {
		r.Message = "Ocurrio un error al intentar realizar el registro del usuario"
		fmt.Println(r.Message)
	}

	r.Status = 200
	r.Message = "Usuario creado correctamente"
	fmt.Println(r.Message)
	return r

}
