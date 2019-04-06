package database

import (
	"testing"

	"github.com/Pungyeon/jakobsen-dev/model"
)

var (
	put     model.Article
	fetched model.Article
	updated model.Article
	err     error
)

func TestArticleDB(t *testing.T) {
	db := New(Options{
		Name:       "test",
		Create:     true,
		Drop:       true,
		Initialise: true,
		Port:       "26257",
	})

	db.Drop()
	db.Create()
	db.Initialise()

	article := model.Article{
		Title:       "first article",
		Description: "This is my first article",
		ArticleLink: "https://raw.githubusercontent.com/Pungyeon/clean-go/master/README.md",
		ImageLink:   "https://external-preview.redd.it/L5a31wsfcT9TcNcvOF3HTOFkXxnKjA7OopCakXxScDg.png?auto=webp&s=bbb10ca8d08363bc2d94996a77619d8bf60c24e8",
		ViewCount:   0,
		Tags:        "ding,dong,doodle",
	}

	t.Run("create new article", func(t *testing.T) {
		put, err = db.PutArticle(article)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("get all articles from database", func(t *testing.T) {
		result, err := db.GetAllArticles()
		if err != nil {
			t.Fatal(err)
		}
		if len(result) != 1 {
			t.Fatal("article was not retrieved from database")
		}
	})

	t.Run("get article by id", func(t *testing.T) {
		fetched, err = db.GetArticle(put.ID)
		if err != nil {
			t.Fatal(err)
		}

		if fetched.ID != put.ID {
			t.Fatal("What? How did that happen?")
		}
	})

	t.Run("update articel properties", func(t *testing.T) {
		fetched.Title = "my first article"
		fetched.Description = "this is definitely my first article"
		fetched.ArticleLink = "https://raw.githubusercontent.com/Pungyeon/docker-example/master/README.md"

		updated, err = db.UpdateArticle(fetched)
		if err != nil {
			t.Fatal(err)
		}

		if updated.Title != fetched.Title {
			t.Fatal("article title was not updated")
		}

		if updated.Description != fetched.Description {
			t.Fatal("article description was not updated")
		}

		if updated.ArticleLink != fetched.ArticleLink {
			t.Fatal("article link was not updated")
		}
	})
}
