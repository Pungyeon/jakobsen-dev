package model

import "time"

var (
	// EmptyArticle represents a null return of an article
	EmptyArticle = Article{}
)

// Article is a pointer and description of a blog article
type Article struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	ArticleLink string    `json:"article_link" db:"article_link"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
