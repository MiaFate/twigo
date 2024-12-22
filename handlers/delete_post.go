package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func DeletePost(ctx *gin.Context, claim *models.Claim) models.ApiResponse[string] {
	var r models.ApiResponse[string]
	r.Status = 400

	id := ctx.Query("id")
	if len(id) < 1 {
		r.Message = "id is required"
		return r
	}

	err := bd.DeletePost(id, claim.Id.Hex())
	if err != nil {
		r.Message = "Error trying to delete post: " + err.Error()
		return r
	}

	r.Message = "Post Deleted"
	r.Status = 200
	return r
}
