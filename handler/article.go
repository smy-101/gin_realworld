package handler

import (
	"gin_realworld/logger"
	"gin_realworld/params/request"
	"gin_realworld/params/response"
	"gin_realworld/storage"
	"gin_realworld/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddArticleHandler(r *gin.Engine) {
	articlesGroup := r.Group("/api/articles")
	articlesGroup.GET("", listArticles)
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
		resp.Articles = append(resp.Articles, &response.Article{
			Author: &response.ArticleAuthor{
				Bio:       article.AuthorUserBio,
				Following: false,
				Image:     article.AuthorUserImage,
				Username:  article.AuthorUsername,
			},
			Title:          article.Title,
			Slug:           article.Slug,
			Body:           article.Body,
			Description:    article.Description,
			TagList:        article.TagList,
			Favorited:      false,
			FavoritesCount: 0,
			CreatedAt:      article.CreatedAt,
			UpdatedAt:      article.UpdatedAt,
		})
	}
	ctx.JSON(http.StatusOK, resp)

}
