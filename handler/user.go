package handler

import (
	"gin_realworld/logger"
	"gin_realworld/models"
	"gin_realworld/params/request"
	"gin_realworld/params/response"
	"gin_realworld/security"
	"gin_realworld/storage"
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
	defaultImage := "https://api.realworld.io/images/smiley-cyrus.jpeg"

	//insert user to db
	if err := storage.CreateUser(ctx, &models.User{
		UserName: body.User.UserName,
		Password: body.User.Password,
		Email:    body.User.Email,
		Image:    defaultImage,
		Bio:      "",
	}); err != nil {
		log.WithError(err).Errorln("create user failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

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
			Image:    defaultImage,
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
	dbUser, err := storage.GetUserByEmail(ctx, body.User.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if dbUser.Password != body.User.Password {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := security.GenerateJWT(dbUser.UserName, body.User.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    body.User.Email,
			Token:    token,
			UserName: dbUser.UserName,
			Bio:      dbUser.Bio,
			Image:    dbUser.Image,
		}})

}
