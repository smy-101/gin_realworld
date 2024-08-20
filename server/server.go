package server

import (
	"gin_realworld/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunHttpServer() {
	r := gin.Default()
	handler.AddUserHandler(r)
	handler.AddArticleHandler(r)
	handler.AddTagsHandler(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
