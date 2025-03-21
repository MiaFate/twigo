package handlers

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/models"
)

func UploadImage(ctx *gin.Context, uploadType string, claim *models.Claim) models.ApiResponse[models.Image] {
	var r models.ApiResponse[models.Image]
	r.Status = 400
	file, err := ctx.FormFile(uploadType)

	if file == nil {
		r.Message = "No file uploaded " + err.Error()
		return r
	}

	userid := claim.Id.Hex()
	fileExt := filepath.Ext(file.Filename)
	var filename string
	var user models.Usuario
	switch uploadType {
	case "A":
		filename = "public/images/avatars/" + userid + fileExt
		user.Avatar = os.Getenv("BASEPATH") + "/images/avatars/" + userid + fileExt
	case "B":
		filename = "public/images/banners/" + userid + fileExt
		user.Banner = os.Getenv("BASEPATH") + "/images/banners/" + userid + fileExt
	}

	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		r.Message = "upload file err: " + err.Error()
		return r
	}

	_, status, err := bd.UpdateUser(user, userid)
	if err != nil {
		r.Message = "Error al actualizar el usuario " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado actualizar el usuario"
		return r
	}

	r.Status = 200
	switch uploadType {
	case "A":
		r.Message = "avatar uploaded"
		r.Data = models.Image{Avatar: user.Avatar}
	case "B":
		r.Message = "banner uploaded"
		r.Data = models.Image{Banner: user.Banner}
	}

	return r
}
