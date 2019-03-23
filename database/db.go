package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Pungyeon/jakobsen-dev/model"
	_ "github.com/lib/pq"
)

type Options struct {
	Name       string
	Create     bool
	Drop       bool
	Initialise bool
	Port       string
}

type ConnectionURI struct {
	uri string
}

func NewURI(options Options) ConnectionURI {
	return ConnectionURI{
		uri: fmt.Sprintf(
			"postgresql://root@localhost:%s?dbname=%s&sslmode=disable",
			options.Port, options.Name,
		),
	}
}

const (
	GET_ARTICLE      = `SELECT id, title, description, article_link, created_at FROM articles WHERE id = %d`
	GET_ALL_ARTICLES = `SELECT id, title, description, article_link, created_at FROM articles`
	PUT_ARTICLE      = `INSERT INTO articles (title, description, article_link) VALUES ('%s', '%s', '%s') RETURNING id`
)

func (uri ConnectionURI) String() string {
	return uri.uri
}

type DB struct {
	conn    *sql.DB
	options Options
}

func New(options Options) *DB {
	_db, err := sql.Open("postgres", NewURI(options).String())
	if err != nil {
		panic(err)
	}
	log.Println("connected to database:", NewURI(options).String())

	return &DB{
		conn:    _db,
		options: options,
	}
}

func (db *DB) Create() {
	if !db.options.Create {
		return
	}
	_, err := db.conn.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", db.options.Name))
	if err != nil {
		panic(err)
	}
}

func (db *DB) Drop() {
	if !db.options.Drop {
		return
	}
	_, err := db.conn.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", db.options.Name))
	if err != nil {
		panic(err)
	}
}

func (db *DB) Initialise() {
	if !db.options.Initialise {
		return
	}
	for _, schema := range schemas {
		_, err := db.conn.Exec(schema.Schema)
		if err != nil {
			panic(fmt.Sprintf("Failed to create %s: %v", schema.Name, err))
		}
	}
}

func (db *DB) GetAllArticles() ([]model.Article, error) {
	rows, err := db.conn.Query(GET_ALL_ARTICLES)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []model.Article
	for rows.Next() {
		var article model.Article
		if err := rows.Scan(
			&article.ID, &article.Title, &article.Description, &article.ArticleLink, &article.CreatedAt,
		); err != nil {
			return nil, err
		}
		fmt.Println(article)
		articles = append(articles, article)
	}
	return articles, nil
}

func (db *DB) GetArticle(id int64) (model.Article, error) {
	var article model.Article
	if err := db.conn.QueryRow(
		fmt.Sprintf(GET_ARTICLE, id),
	).Scan(&article.ID, &article.Title, &article.Description, &article.ArticleLink, &article.CreatedAt); err != nil {
		return model.EmptyArticle, err
	}
	return article, nil
}

func (db *DB) PutArticle(article model.Article) (int64, error) {
	var id int64
	if err := db.conn.QueryRow(
		fmt.Sprintf(PUT_ARTICLE,
			article.Title, article.Description, article.ArticleLink),
	).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (db *DB) UpdateArticle() {}
