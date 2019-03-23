package main

import (
	"flag"

	"github.com/Pungyeon/jakobsen-dev/app"
	"github.com/Pungyeon/jakobsen-dev/database"
)

func main() {
	port := *flag.String("port", "8000", "port in which to listen to incoming requests for the web server")
	dbcreate := *flag.Bool("db-create", true, "specify if the database should be recreated if it doesn't exist")
	dbdrop := *flag.Bool("db-drop", false, "specify if the database should be dropped completely")
	dbinit := *flag.Bool("db-init", true, "specify whether to initialise the tables of the database, if they don't already exist")
	dbname := *flag.String("db-name", "core", "specify the main database name")
	dbport := *flag.String("db-port", "26257", "specify ithe port to use for connecting to the database")

	db := database.New(database.Options{
		Name:       dbname,
		Create:     dbcreate,
		Drop:       dbdrop,
		Initialise: dbinit,
		Port:       dbport,
	})

	if err := app.New(db, port).Run(); err != nil {
		panic(err)
	}
}
