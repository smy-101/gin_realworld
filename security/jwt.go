package security

import (
	"gin_realworld/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(username, email string) (string, error) {

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

func VerifyJWT(token string) (*jwt.MapClaims, bool, error) {
	var claim jwt.MapClaims
	claims, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetSecret()), nil
	})
	if err != nil {
		return nil, false, err
	}

	if claims.Valid {
		return &claim, true, nil
	}
	return nil, false, nil
}

func GetCurrentUserName(ctx *gin.Context) string {
	mapClaims := ctx.MustGet("user").(*jwt.MapClaims)
	userName := (*mapClaims)["user"].(map[string]interface{})["username"].(string)
	return userName
}

func GetCurrentUserEmail(ctx *gin.Context) string {
	mapClaims := ctx.MustGet("user").(*jwt.MapClaims)
	userName := (*mapClaims)["user"].(map[string]interface{})["email"].(string)
	return userName
}
