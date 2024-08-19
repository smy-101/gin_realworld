package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type TagList []string

type Article struct {
	Id             int64  `db:"id"`
	AuthorUsername string `gorm:"column:author_username"`
	Title          string
	Slug           string
	Body           string
	Description    string
	TagList        TagList `gorm:"type:string"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (a Article) TableName() string {
	return "article"
}

func (j *TagList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	err := json.Unmarshal(bytes, j)
	return err
}

func (j TagList) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.Marshal(j)
}
