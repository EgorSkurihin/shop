package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

var decoder = schema.NewDecoder()

type Server struct {
	addr         string
	router       *mux.Router
	store        *Store
	sessionStore sessions.CookieStore
}

func main() {
	// Create new DataBase object
	store, err := NewStore()
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	// Create Server
	srv := &Server{
		addr:         "localhost:1323",
		router:       mux.NewRouter(),
		store:        store,
		sessionStore: *sessions.NewCookieStore([]byte("super-secret-session-key")),
	}
	srv.setupRoutes()

	//Start server
	log.Println("http://localhost:1323")
	log.Fatal(http.ListenAndServe(":1323", srv.router))
}
