package handlers

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func UpdateUser(ctx *gin.Context, claim *models.Claim) models.ApiResponse[models.Usuario] {
	var r models.ApiResponse[models.Usuario]
	r.Status = 400

	var t models.Usuario

	body, _ := io.ReadAll(ctx.Request.Body)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Error al leer el cuerpo de la petición " + err.Error()
		return r
	}

	perfil, status, err := bd.UpdateUser(t, claim.Id.Hex())
	if err != nil {
		r.Message = "Error al actualizar el usuario " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado actualizar el usuario"
		return r
	}

	r.Status = 200
	r.Message = "Usuario actualizado correctamente"
	r.Data = perfil

	return r
}
