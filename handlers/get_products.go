package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func GetProducts(ctx *gin.Context) models.ApiResponse[[]*models.GetProducts] {

	var r models.ApiResponse[[]*models.GetProducts]
	r.Status = 400

	products, success := bd.GetProducts()
	if !success {
		r.Message = "Error getting products"
		return r
	}

	r.Status = 200
	r.Data = products

	return r

}
