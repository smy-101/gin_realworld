package security

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

const salt = "sfasfjasfj"

func HashPassword(password string) (string, error) {
	password += salt
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bcryptPassword), nil
}

func CheckPassword(plain, hash string) bool {
	plain += salt
	bcryptPassword, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(bcryptPassword, []byte(plain))
	return err == nil
}
