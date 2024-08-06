package security

import (
	"gin_realworld/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func generateJWT(username, email string) (string, error) {

	key := []byte(config.GetSecret())
	tokenDuration := 24 * time.Hour
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": map[string]string{
			"email":    email,
			"username": username,
		},
		"iat": now.Unix(),
		"exp": now.Add(tokenDuration).Unix(),
	})
	return t.SignedString(key)
}
