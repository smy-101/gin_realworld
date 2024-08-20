package storage

import (
	"context"
	"gin_realworld/models"
	"gin_realworld/params/request"

	"gorm.io/gorm"
)

func CreateArticle(ctx context.Context, article *models.Article) error {
	return gormDb.WithContext(ctx).Create(article).Error
}

func listArticleTx(ctx context.Context, req *request.ListArticleQuery) *gorm.DB {
	tx := gormDb.WithContext(ctx).Model(models.Article{}).
		Select("article.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article.author_username = user.username").
		Order("created_at desc").Offset(int(req.Offset)).Limit(int(req.Limit))

	if req.Tag != "" {
		tx = tx.Where("article.tag_list like ?", "%\""+req.Tag+"\"%")
	}

	return tx
}

func ListArticles(ctx context.Context, req *request.ListArticleQuery) ([]models.Article, error) {
	var articles []models.Article

	err := listArticleTx(ctx, req).Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func CountArticles(ctx context.Context, req *request.ListArticleQuery) (int64, error) {
	var count int64
	err := listArticleTx(ctx, req).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
