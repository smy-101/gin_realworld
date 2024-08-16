package handler

import (
	"gin_realworld/logger"
	"gin_realworld/params/request"
	"gin_realworld/params/response"
	"gin_realworld/security"
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

	//TODO: insert user to db

	token, err := security.GenerateJWT(body.User.UserName, body.User.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    body.User.Email,
			Token:    token,
			UserName: body.User.UserName,
			Bio:      "",
			Image:    "https://api.realworld.io/images/smiley-cyrus.jpeg",
		}})
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

	//TODO: get userName from db
	userName := "jack"

	token, err := security.GenerateJWT(userName, body.User.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    body.User.Email,
			Token:    token,
			UserName: userName,
			Bio:      "",
			Image:    "https://api.realworld.io/images/smiley-cyrus.jpeg",
		}})

}
