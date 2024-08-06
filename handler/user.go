package handler

import (
	"gin_realworld/logger"
	"gin_realworld/params/request"
	"gin_realworld/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUserHandler(r *gin.Engine) {
	usersGroup := r.Group("/api/users")
	usersGroup.POST("", userRegistration)
	usersGroup.POST("/login", userLogin)
}

func userRegistration(ctx *gin.Context) {
	log := logger.New(ctx)
	var body request.UserRegistrationRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Error("error while binding json")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.WithField("user", utils.JsonMarshal(body)).Infof("user registration")
}

func userLogin(ctx *gin.Context) {
	log := logger.New(ctx)
	var body request.UserLoginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Error("error while binding json")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.WithField("user", utils.JsonMarshal(body)).Infof("user login")
}
