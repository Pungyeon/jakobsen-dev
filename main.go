package main

import (
	"flag"

	_ "github.com/lib/pq"
)

func main() {
	port := *flag.String("port", "8000", "port in which to listen to incoming requests for the web server")

	db := database.New(database.Options{
		URI:        "postgresql://root@localhost:26257?dbname=test&sslmode=disable",
		DBName:     "test",
		Create:     true,
		Drop:       true,
		Initialise: true,
	})

	if err := app.New(db, port).Run(); err != nil {
		panic(err)
	}
}
