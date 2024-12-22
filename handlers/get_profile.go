package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func GetProfile(ctx *gin.Context) models.ApiResponse[models.Usuario] {
	var r models.ApiResponse[models.Usuario]
	r.Status = 400

	id := ctx.Query("id")
	if len(id) == 0 {
		r.Message = "El id es requerido"
		return r
	}

	perfil, err := bd.GetProfile(id)
	if err != nil {
		r.Message = "Error al buscar el perfil " + err.Error()
		return r
	}

	fmt.Println(perfil)
	r.Status = 200
	r.Data = perfil

	return r
}
