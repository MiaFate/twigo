package secretmanager

import (
	"fmt"

	"github.com/miafate/twigo/models"
)

func GetSecret(secretName string) (models.Secret, error{
	var datosSecret := models.Secret
	fmt.Println("> Pido secreto " + secretName)

	svc := secretmanager.NewFromConfig(configuration.Cfg)
	clave, err := scv.GetSecretValue(configuration.Ctx, &secretsmanager.GetSecretValueInput{})
}