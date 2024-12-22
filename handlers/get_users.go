package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func GetUsers(ctx *gin.Context) models.ApiResponse[[]models.Usuario] {
	var r models.ApiResponse[[]models.Usuario]
	r.Status = 200
	// r.Message = "Listado de usuarios"
	r.Data = append(r.Data, bd.GetUsers()...)

	return r

}
