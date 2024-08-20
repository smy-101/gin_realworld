package storage

import (
	"context"
	"encoding/json"
	"gin_realworld/models"
	"gin_realworld/utils"
	"testing"
	"time"
)

func TestCreateArticle(t *testing.T) {
	ctx := context.TODO()
	err := CreateArticle(ctx, &models.Article{
		AuthorUsername: "xx",
		Title:          "xxx",
		Slug:           "xxx",
		Body:           "111",
		Description:    "111",
		TagList:        []string{"111"},
	})
	if err != nil {
		t.Fatal(err)
	}
}

var data = `[]`

func TestDataImport(t *testing.T) {
	ctx := context.TODO()
	var articles []map[string]interface{}
	err := json.Unmarshal([]byte(data), &articles)
	if err != nil {
		t.Fatal(err)
	}

	for _, article := range articles {
		var tagList []string
		for _, tag := range article["tagList"].([]interface{}) {
			tagList = append(tagList, tag.(string))
		}
		createdAt, err := time.Parse(time.RFC3339Nano, article["createdAt"].(string))
		if err != nil {
			t.Logf("parse time failed")
			continue
		}
		updatedAt, err := time.Parse(time.RFC3339Nano, article["updatedAt"].(string))
		if err != nil {
			t.Logf("parse time failed")
			continue
		}

		err = CreateArticle(ctx, &models.Article{
			AuthorUsername: article["author"].(map[string]interface{})["username"].(string),
			Title:          article["title"].(string),
			Slug:           article["slug"].(string),
			Body:           article["body"].(string),
			Description:    article["description"].(string),
			TagList:        tagList,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
		})
		if err != nil {
			t.Errorf("crate article failed")
		}
	}
}

func TestListArticles(t *testing.T) {
	ctx := context.TODO()
	articles, err := ListArticles(ctx, 10, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("articles:%v\n", utils.JsonMarshal(articles))
}
