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

func UpdateArticle(ctx context.Context, oldSlug string, article *models.Article) error {
	return gormDb.WithContext(ctx).Where("slug = ?", oldSlug).Updates(article).Error
}

func DeleteArticle(ctx context.Context, slug string) error {
	return gormDb.WithContext(ctx).Model(models.Article{}).Where("slug = ?", slug).Delete(models.Article{}).Error
}

func listArticleTx(ctx context.Context, req *request.ListArticleQuery) *gorm.DB {
	tx := gormDb.WithContext(ctx).Model(models.Article{}).
		Select("article.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article.author_username = user.username").
		Order("created_at desc")

	if req.Tag != "" {
		tx = tx.Where("article.tag_list like ?", "%\""+req.Tag+"\"%")
	}

	return tx
}

func ListArticles(ctx context.Context, req *request.ListArticleQuery) ([]models.Article, error) {
	var articles []models.Article

	err := listArticleTx(ctx, req).Offset(int(req.Offset)).Limit(int(req.Limit)).Find(&articles).Error
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

func GetArticleBySlug(ctx context.Context, slug string) (*models.Article, error) {
	var article models.Article
	if err := gormDb.WithContext(ctx).Model(models.Article{}).
		Select("article.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article.author_username = user.username").
		Where("slug = ?", slug).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}
