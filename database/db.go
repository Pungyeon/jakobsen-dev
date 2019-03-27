package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Pungyeon/jakobsen-dev/model"
	_ "github.com/lib/pq"
)

// Options represent database input options
type Options struct {
	Name       string
	Create     bool
	Drop       bool
	Initialise bool
	Port       string
}

// ConnectionURI represents a database connection url
type ConnectionURI struct {
	uri string
}

// NewURI will return a new database url
func NewURI(options Options) ConnectionURI {
	return ConnectionURI{
		uri: fmt.Sprintf(
			"postgresql://root@localhost:%s?dbname=%s&sslmode=disable",
			options.Port, options.Name,
		),
	}
}

const (
	GET_ARTICLE      = `SELECT id, title, description, article_link, image_link, view_count, created_at FROM articles WHERE id = %d`
	GET_ALL_ARTICLES = `SELECT id, title, description, article_link, image_link, view_count, created_at FROM articles`
	PUT_ARTICLE      = `
	INSERT INTO articles
		(title, description, article_link, image_link, view_count)
	VALUES ('%s', '%s', '%s', '%s', %d)
	RETURNING id, title, description, article_link, image_link, view_count, created_at
	`
	UPDATE_ARTICLE = `
	UPDATE articles 
		SET title = '%s', description = '%s', article_link = '%s'
	WHERE id = %d
	RETURNING id, title, description, article_link, image_link, view_count, created_at
	`
)

func (uri ConnectionURI) String() string {
	return uri.uri
}

// DB represents a database connection and the operations associated with it
type DB struct {
	conn    *sql.DB
	options Options
}

// New will return a new database
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

// Create will create a new database
func (db *DB) Create() *DB {
	if !db.options.Create {
		return db
	}
	log.Println("Creating Database:", db.options.Name)
	_, err := db.conn.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", db.options.Name))
	if err != nil {
		panic(err)
	}
	return db
}

// Drop will drop the current database
func (db *DB) Drop() *DB {
	if !db.options.Drop {
		return db
	}
	log.Println("Dropping Database:", db.options.Name)
	_, err := db.conn.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", db.options.Name))
	if err != nil {
		panic(err)
	}
	return db
}

// Initialise will initialise all tables in the database
func (db *DB) Initialise() *DB {
	if !db.options.Initialise {
		return db
	}
	for _, schema := range schemas {
		_, err := db.conn.Exec(schema.Schema)
		if err != nil {
			panic(fmt.Sprintf("Failed to create %s: %v", schema.Name, err))
		}
		log.Println("Initialising Database Table:", schema.Name)
	}
	return db
}

// GetAllArticles will return all articles in the database
func (db *DB) GetAllArticles() ([]model.Article, error) {
	rows, err := db.conn.Query(GET_ALL_ARTICLES)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanArticleRows(rows)
}

// GetArticle will retrieve a single article specified by id
func (db *DB) GetArticle(id int64) (model.Article, error) {
	return scanArticleRow(
		db.conn.QueryRow(fmt.Sprintf(GET_ARTICLE, id)))
}

// PutArticle will insert a new article into the database
func (db *DB) PutArticle(article model.Article) (model.Article, error) {
	return scanArticleRow(
		db.conn.QueryRow(fmt.Sprintf(
			PUT_ARTICLE,
			article.Title, article.Description, article.ArticleLink,
			article.ImageLink, article.ViewCount)),
	)
}

// UpdateArticle will update an already existing article, overwriting all dynamic values
// title, description and article_link
func (db *DB) UpdateArticle(article model.Article) (model.Article, error) {
	return scanArticleRow(
		db.conn.QueryRow(fmt.Sprintf(
			UPDATE_ARTICLE,
			article.Title, article.Description,
			article.ArticleLink, article.ID)))
}

func scanArticleRow(row *sql.Row) (model.Article, error) {
	var article model.Article
	if err := row.Scan(
		&article.ID, &article.Title, &article.Description,
		&article.ArticleLink, &article.ImageLink,
		&article.ViewCount, &article.CreatedAt,
	); err != nil {
		return model.EmptyArticle, err
	}
	return article, nil
}

func scanArticleRows(rows *sql.Rows) ([]model.Article, error) {
	var articles []model.Article
	for rows.Next() {
		article, err := scanCurrentInArticleRows(rows)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func scanCurrentInArticleRows(rows *sql.Rows) (model.Article, error) {
	var article model.Article
	if err := rows.Scan(
		&article.ID, &article.Title,
		&article.Description, &article.ArticleLink,
		&article.ImageLink, &article.ViewCount, &article.CreatedAt,
	); err != nil {
		return model.EmptyArticle, err
	}
	return article, nil
}
