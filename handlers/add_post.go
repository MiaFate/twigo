package handlers

import (
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func AddPost(ctx *gin.Context, claim *models.Claim) models.ApiResponse[models.AddPost] {
	var post models.Post
	var r models.ApiResponse[models.AddPost]
	r.Status = 400
	userId := claim.Id.Hex()

	body, _ := io.ReadAll(ctx.Request.Body)
	err := json.Unmarshal([]byte(body), &post)
	if err != nil {
		r.Message = "error trying to read the request body " + err.Error()
		return r
	}

	register := models.AddPost{
		UserId:  userId,
		Message: post.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.AddPost(register)
	if err != nil {
		r.Message = "error trying to add the post " + err.Error()
		return r
	}

	if !status {
		r.Message = "error trying to add the post"
		return r
	}

	r.Status = 200
	r.Message = "post added correctly"
	r.Data = register

	return r
}
