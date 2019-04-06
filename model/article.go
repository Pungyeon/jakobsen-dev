package model

import (
	"encoding/json"
	"strings"
	"time"
)

var (
	// EmptyArticle represents a null return of an article
	EmptyArticle = Article{}
)

// Article is a pointer and description of a blog article
type Article struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	ImageLink   string    `json:"image_link" db:"image_link"`
	ArticleLink string    `json:"article_link" db:"article_link"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	ViewCount   int       `json:"view_count" db:"view_count"`
	Tags        string    `db:"tags"`

	SplitTags []string `json:"tags"`
}

func (article *Article) JSON() []byte {
	article.SetTags()
	data, err := json.Marshal(article)
	if err != nil {
		return []byte(`{}`)
	}
	return data
}

func (article *Article) SetTags() Article {
	article.SplitTags = strings.Split(article.Tags, ",")
	return *article
}

type Articles []Article

func (articles Articles) JSON() []byte {
	for i := range articles {
		articles[i] = articles[i].SetTags()
	}
	data, err := json.Marshal(articles)
	if err != nil {
		return []byte(`[]`)
	}
	return data
}
