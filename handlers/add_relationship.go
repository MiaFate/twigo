package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func AddRelationship(ctx *gin.Context, claim *models.Claim) models.ApiResponse[string] {
	var r models.ApiResponse[string]
	r.Status = 400

	id := ctx.Query("id")
	if len(id) < 1 {
		r.Message = "id is required"
		return r
	}

	var t models.Relationship
	t.UserId = claim.Id.Hex()
	t.UserRelationshipId = id

	status, err := bd.AddRelationship(t)
	if err != nil {
		r.Message = "have been an error adding friend: " + err.Error()
		return r
	}
	if !status {
		r.Message = "have been an error adding friend :C"
		return r
	}

	r.Status = 200
	r.Message = "Friend added"
	return r
}
