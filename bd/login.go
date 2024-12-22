package bd

import (
	"github.com/miafate/twigo/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.Usuario, bool) {
	usr, encontrado, _ := UserExist(email)
	if !encontrado {
		return usr, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usr.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usr, false
	}
	return usr, true
}
