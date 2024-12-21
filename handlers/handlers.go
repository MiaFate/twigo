package handlers

import (
	"context"
	"fmt"

	"github.com/miafate/twigo/jwt"
	"github.com/miafate/twigo/models"
	"github.com/miafate/twigo/routers"
)

func Manejadores(ctx context.Context, request ) models.RespApi {

	var r models.RespApi
	r.Status = 400

	isOk, statusCode, msg, claim := validoAuthorization(ctx, request)
	if !isOk {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Register(ctx)
		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		}
		//
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		}
		//
	}

}

func validoAuthorization(ctx context.Context, request) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "register" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "", models.Claim{}
	}
	token := request.headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}

	claim, todoOK, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !todoOK {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token " + msg)
			return false, 401, msg, models.Claim{}
		}
	}
	fmt.Println("Token OK")
	return true, 200, msg, *claim

}
