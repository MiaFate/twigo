package handlers

import (
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func AddProduct(ctx *gin.Context) models.ApiResponse[models.Product] {
	var product models.Product
	var r models.ApiResponse[models.Product]
	r.Status = 400

	body, _ := io.ReadAll(ctx.Request.Body)
	err := json.Unmarshal([]byte(body), &product)
	if err != nil {
		r.Message = "error trying to read the request body " + err.Error()
		return r
	}

	register := models.Product{
		Product: product.Product,
		Date:    time.Now(),
	}

	_, status, err := bd.AddProduct(register)
	if err != nil {
		r.Message = "error trying to add the post " + err.Error()
		return r
	}

	if !status {
		r.Message = "error trying to add the post"
		return r
	}

	r.Status = 200
	r.Message = "product added correctly"
	r.Data = register

	return r
}
