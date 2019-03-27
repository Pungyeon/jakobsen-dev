package database

// ArticlesSchema is the schema for the article database table
const ArticlesSchema = `
CREATE TABLE IF NOT EXISTS articles (
	id 				SERIAL PRIMARY KEY,
	title 			STRING,
	description 	STRING,
	image_link 		STRING,
	view_count 		INT,
	article_link 	STRING,
	created_at 		TIMESTAMPTZ DEFAULT NOW()
)
`

// Schema represents a database table schema
type Schema struct {
	Name   string
	Schema string
}

var schemas = []Schema{
	{"articles", ArticlesSchema},
}
