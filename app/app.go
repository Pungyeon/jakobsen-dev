package app

import (
	"log"
	"net/http"

	"github.com/Pungyeon/jakobsen-dev/database"
)

type App struct {
	port string
	db   *database.DB
}

func New(db *database.DB, port string) *App {
	return &App{
		port: port,
		db:   db,
	}
}

func (app *App) Run() error {
	app.db.Drop()
	app.db.Create()
	app.db.Initialise()

	http.HandleFunc("/", app.indexHandler)

	log.Println("running server on port:", app.port)
	return http.ListenAndServe(":"+app.port, nil)
}

func (app *App) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
