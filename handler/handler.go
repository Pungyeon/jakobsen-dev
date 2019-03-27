package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pungyeon/jakobsen-dev/database"
	"github.com/Pungyeon/jakobsen-dev/model"
	"github.com/gorilla/mux"
)

// HTTPHandler is a handler for all http requests
type HTTPHandler struct {
	router *mux.Router
	db     *database.DB
}

// New will return a new handler
func New(db *database.DB, port string) *HTTPHandler {
	return &HTTPHandler{
		db:     db.Drop().Create().Initialise(),
		router: mux.NewRouter(),
	}
}

// InitialiseRouter will initialise all routes of the HTTP handler
func (handler *HTTPHandler) InitialiseRouter() *mux.Router {
	handler.router.HandleFunc("/", handler.index)
	handler.router.HandleFunc("/api/articles", handler.articlesGet).Methods("GET")
	handler.router.HandleFunc("/api/articles/all", handler.articlesGetAll).Methods("GET")
	// handler.router.HandleFunc("/api/articles", handler.articlesPost).Methods("POST")

	return handler.router
}

func (handler *HTTPHandler) index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (handler *HTTPHandler) articlesGetAll(w http.ResponseWriter, r *http.Request) {
	dbarticles, err := handler.db.GetAllArticles()
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(dbarticles)
}

func (handler *HTTPHandler) articlesGet(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}
	dbarticle, err := handler.db.GetArticle(article.ID)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(dbarticle)
}

func handleError(err error, w http.ResponseWriter, status int) {
	log.Println(err)
	w.WriteHeader(status)
}
