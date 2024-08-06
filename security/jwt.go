package security

import (
	"crypto/rsa"
	"gin_realworld/config"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {
	var err error
	var bytes []byte
	bytes, err = os.ReadFile(config.GetPrivateKeyLocation()) //还可以将私钥和公钥放在环境变量中
	if err != nil {
		panic(err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}

	bytes, err = os.ReadFile(config.GetPublicKeyLocation())
	if err != nil {
		panic(err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}
}

func GenerateJWT(username, email string) (string, error) {
	//key := []byte(config.GetSecret())
	tokenDuration := 24 * time.Hour
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodRS512,
		jwt.MapClaims{
			"user": map[string]string{
				"email":    email,
				"username": username,
			},
			"iat": now.Unix(),
			"exp": now.Add(tokenDuration).Unix(),
		})
	return t.SignedString(privateKey)
}
