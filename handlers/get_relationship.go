package handlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func GetRelationship(ctx *gin.Context, claim *models.Claim) models.ApiResponse[string] {
	var r models.ApiResponse[string]

	id := ctx.Query("id")
	if len(id) < 1 {
		r.Message = "id parameter is mandatory"
		return r
	}

	var t models.Relationship
	t.UserId = claim.Id.Hex()
	t.FriendId = id

	var resp models.GetRelationshipResponse

	relationExists := bd.GetRelationship(t)
	if !relationExists {
		resp.Status = false
	} else {
		resp.Status = true
	}

	r.Status = 200
	respJson, err := json.Marshal(resp.Status)
	if err != nil {
		r.Status = 500
		r.Message = "Error converting response to JSON"
		return r
	}
	r.Message = string(respJson)
	return r
}
