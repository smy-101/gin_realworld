package handler

import (
	"gin_realworld/logger"
	"gin_realworld/middlewares"
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
	r.GET("/api/profiles/:username", UserProfiles)
	r.Group("/api/user").Use(middlewares.AuthMiddleware).PUT("", EditUser)
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

	hashPassword, err := security.HashPassword(body.User.Password)
	if err != nil {
		log.WithError(err).Errorln("hash password failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//insert user to db
	if err := storage.CreateUser(ctx, &models.User{
		UserName: body.User.UserName,
		Password: hashPassword,
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

	//get userName from db
	dbUser, err := storage.GetUserByEmail(ctx, body.User.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if !security.CheckPassword(body.User.Password, dbUser.Password) {
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

func UserProfiles(ctx *gin.Context) {
	log := logger.New(ctx)
	userName := ctx.Param("username")
	log = log.WithField("userName", userName)
	log.Infof("get user profile by userName: %v\n", userName)

	user, err := storage.GetUserByUsername(ctx, userName)
	if err != nil {
		log.WithError(err).Errorln("get user by userName failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.UserProfileResponse{
		UserProfile: response.UserProfile{
			Username:  user.UserName,
			Bio:       user.Bio,
			Image:     user.Image,
			Following: false,
		},
	})
}

func EditUser(ctx *gin.Context) {
	log := logger.New(ctx)
	log.Infof("edit user:%v", security.GetCurrentUserName(ctx))
	var body request.EditUserRequest
	if err := ctx.BindJSON(&body); err != nil {
		log.WithError(err).Errorln("bind json failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if body.EditUserBody.Password != "" {
		var err error
		body.EditUserBody.Password, err = security.HashPassword(body.EditUserBody.Password)
		if err != nil {
			log.WithError(err).Errorln("hash password failed")
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	dbUser := &models.User{
		UserName: body.EditUserBody.Username,
		Password: body.EditUserBody.Password,
		Email:    body.EditUserBody.Email,
		Image:    body.EditUserBody.Image,
		Bio:      body.EditUserBody.Bio,
	}

	if err := storage.UpdateUserByUserName(ctx, security.GetCurrentUserName(ctx), dbUser); err != nil {
		log.WithError(err).Errorln("update user failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := security.GenerateJWT(dbUser.UserName, dbUser.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    dbUser.Email,
			Token:    token,
			UserName: dbUser.UserName,
			Bio:      dbUser.Bio,
			Image:    dbUser.Image,
		}})
}
