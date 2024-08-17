package middlewares

import (
	"gin_realworld/logger"
	"gin_realworld/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	log := logger.New(ctx)
	token := ctx.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	claims, ok, err := security.VerifyJWT(token)
	if err != nil || !ok {
		log.WithError(err).Error("verify jwt failed")
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	ctx.Set("user", claims)
	ctx.Next()
}
