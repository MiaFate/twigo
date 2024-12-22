package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func GetProfile(ctx *gin.Context) models.ApiResponse[any] {
	var r models.ApiResponse[any]
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

	// respJson, err := json.Marshal(perfil)
	// if err != nil {
	// 	r.Message = "Error al convertir el perfil a json " + err.Error()
	// 	return r
	// }

	r.Status = 200
	r.Data = perfil

	return r
}
