package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func DeleteRelationship(ctx *gin.Context, claim *models.Claim) models.ApiResponse[string] {
	var r models.ApiResponse[string]

	id := ctx.Query("id")
	if len(id) < 1 {
		r.Message = "id parameter is mandatory"
		return r
	}

	var t models.Relationship
	t.UserId = claim.Id.Hex()
	t.FriendId = id

	status, err := bd.DeleteRelationship(t)
	if err != nil {
		r.Message = "Error trying to delete friend: " + err.Error()
		return r
	}
	if !status {
		r.Message = "Error trying to delete friend"
		return r
	}

	r.Status = 200
	r.Message = "Friend deleted successfully"
	return r
}
