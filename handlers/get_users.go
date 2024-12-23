package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func GetUsers(ctx *gin.Context, claim *models.Claim) models.ApiResponse[[]*models.Usuario] {
	var r models.ApiResponse[[]*models.Usuario]
	r.Status = 400

	page := ctx.Query("page")
	typeUser := ctx.Query("type")
	search := ctx.Query("search")
	userId := claim.Id.Hex()

	if len(page) == 0 {
		page = "1"
	}

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Page must be an int bigger than zero" + err.Error()
		return r
	}

	users, status := bd.GetUsers(userId, int64(pagTemp), search, typeUser)
	if !status {
		r.Message = "Error getting users"
		return r
	}

	// r.Message = "Listado de usuarios"
	// r.Data = append(r.Data, bd.GetUsers()...)
	r.Status = 200
	r.Data = users

	return r

}
