package handlers

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/models"
)

func AddProductsBulk(ctx *gin.Context) *models.Responses {
	var messages models.Messages
	var res = &models.Responses{}

	body, _ := io.ReadAll(ctx.Request.Body)
	err := json.Unmarshal([]byte(body), &messages)
	if err != nil {
		return res
	}
	for _, message := range messages.Messages {
		if message.Msg != nil && message.Msg.ProductId != nil {
			productRes := AddProduct(ctx, message.Msg.ProductId)
			response := models.Response{
				Id:   *message.Msg.ProductId,
				Code: productRes.Status,
			}
			res.Responses = append(res.Responses, &response)
		}
	}
	return res
}
