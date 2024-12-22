package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func GetPosts(ctx *gin.Context) models.ApiResponse[[]*models.GetPosts] {

	var r models.ApiResponse[[]*models.GetPosts]
	r.Status = 400

	id := ctx.Query("id")
	page := ctx.Query("page")

	if len(id) < 1 {
		r.Message = "ID parameter is mandatory"
		return r
	}

	if len(page) < 1 {
		page = "1"
	}

	pag, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "page must be a number bigger than zero"
		return r
	}

	posts, success := bd.GetPosts(id, int64(pag))
	if !success {
		r.Message = "Error getting posts"
		return r
	}

	r.Status = 200
	r.Data = posts

	return r

}
