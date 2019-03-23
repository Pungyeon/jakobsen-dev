package database

import (
	"fmt"
	"testing"

	"github.com/Pungyeon/jakobsen-dev/model"
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
	}

	id, err := db.PutArticle(article)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(id)

	result, err := db.GetAllArticles()
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 1 {
		t.Fatal("article was not retrieved from database")
	}

	dbarticle, err := db.GetArticle(id)
	if err != nil {
		t.Fatal(err)
	}

	if dbarticle.ID != id {
		t.Fatal("What? How did that happen?")
	}
}
