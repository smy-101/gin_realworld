package storage

import (
	"context"
	"gin_realworld/models"
)

func CreateArticle(ctx context.Context, article *models.Article) error {
	return gormDb.WithContext(ctx).Create(article).Error
}
