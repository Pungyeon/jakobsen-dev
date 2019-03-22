package database

const ArticlesSchema = `
CREATE TABLE IF NOT EXISTS articles (
	id SERIAL 		PRIMARY KEY,
	title 			STRING,
	descriptiong 	STRING,
	article_link 	STRING,
	created_at 		TIMESTAMPTZ DEFAULT NOW()
)
`

type Schema struct {
	Name   string
	Schema string
}

var schemas = []Schema{
	{"articles", ArticlesSchema},
}
