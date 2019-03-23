package app

import (
	"fmt"
	"net/http"

	"github.com/Pungyeon/jakobsen-dev/database"
	"github.com/Pungyeon/jakobsen-dev/handler"
)

// App is the class containing all operations of the website application
type App struct {
	port        string
	httpHandler *handler.HTTPHandler
}

// New returns a new app
func New(db *database.DB, port string) *App {
	return &App{
		httpHandler: handler.New(db, port),
		port:        port,
	}
}

// Run will initialise the database connection and start the http listener
// returning an error if anything goes astray
func (app *App) Run() error {
	server := &http.Server{
		Handler: app.httpHandler.InitialiseRouter(),
		Addr:    "0.0.0.0:" + app.port,
	}
	fmt.Println("Server running on:", server.Addr)
	return server.ListenAndServe()
}
