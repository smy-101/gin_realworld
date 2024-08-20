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
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddArticleHandler(r *gin.Engine) {
	articlesGroup := r.Group("/api/articles")
	articlesGroup.GET("", listArticles)
	articlesGroup.GET("/:slug", getArticle)
	articlesGroup.Use(middlewares.AuthMiddleware)
	articlesGroup.POST("", createArticles)
	articlesGroup.PUT("/:slug", editArticles)
	articlesGroup.DELETE("/:slug", deleteArticles)
}

func createArticles(ctx *gin.Context) {
	var req request.CreateArticleRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	slug := strings.ReplaceAll(req.Article.Title, " ", "-") + "-" + uuid.NewString()
	if err := storage.CreateArticle(ctx, &models.Article{
		AuthorUsername: security.GetCurrentUserName(ctx),
		Title:          req.Article.Title,
		Slug:           slug,
		Body:           req.Article.Body,
		Description:    req.Article.Description,
		TagList:        req.Article.TagList,
	}); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	respArticle := &response.Article{}
	respArticle.FromDB(article)
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"article": respArticle,
	})
}

func editArticles(ctx *gin.Context) {
	oldSlug := ctx.Param("slug")
	var req request.CreateArticleRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	oldArticle, err := storage.GetArticleBySlug(ctx, oldSlug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if oldArticle.AuthorUsername != security.GetCurrentUserName(ctx) {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	slug := strings.ReplaceAll(req.Article.Title, " ", "-") + "-" + uuid.NewString()
	if err := storage.UpdateArticle(ctx, oldSlug, &models.Article{
		AuthorUsername: security.GetCurrentUserName(ctx),
		Title:          req.Article.Title,
		Slug:           slug,
		Body:           req.Article.Body,
		Description:    req.Article.Description,
		TagList:        req.Article.TagList,
	}); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	respArticle := &response.Article{}
	respArticle.FromDB(article)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"article": respArticle,
	})

}

func deleteArticles(ctx *gin.Context) {
	slug := ctx.Param("slug")

	oldArticle, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if oldArticle.AuthorUsername != security.GetCurrentUserName(ctx) {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	err = storage.DeleteArticle(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func getArticle(ctx *gin.Context) {
	slug := ctx.Param("slug")
	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	respArticle := &response.Article{}
	respArticle.FromDB(article)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"article": respArticle,
	})
}

func listArticles(ctx *gin.Context) {
	log := logger.New(ctx)
	// limit, offset := cast.ToInt(ctx.Query("limit")), cast.ToInt(ctx.Query("offset"))
	var req request.ListArticleQuery
	if err := ctx.BindQuery(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.Infof("list articles query: %v", utils.JsonMarshal(req))

	articles, err := storage.ListArticles(ctx, &req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	total, err := storage.CountArticles(ctx, &req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var resp response.ListArticlesResponse
	resp.ArticlesCount = total
	for _, article := range articles {
		respArticle := &response.Article{}
		respArticle.FromDB(&article)
		resp.Articles = append(resp.Articles, respArticle)
	}

	ctx.JSON(http.StatusOK, resp)

}
