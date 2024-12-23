package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func GetFriendsPosts(ctx *gin.Context, claim *models.Claim) models.ApiResponse[[]models.GetFriendsPosts] {
	var r models.ApiResponse[[]models.GetFriendsPosts]
	r.Status = 400
	userId := claim.Id.Hex()
	page := ctx.Query("page")

	if len(page) < 1 {
		page = "1"
	}

	pag, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "page must be a number bigger than zero"
		return r
	}

	posts, success := bd.GetFriendsPosts(userId, int64(pag))
	if !success {
		r.Message = "Error getting posts"
		return r
	}

	if len(posts) < 1 {
		r.Message = "No post found"
		return r
	}

	r.Status = 200
	r.Data = posts

	return r
}
