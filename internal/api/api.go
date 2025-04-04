package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ION-Smart/api_template_go/internal/api/handlers"
	"github.com/ION-Smart/api_template_go/internal/api/routes"
	"github.com/rs/cors"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	// Handlers
	mainHandler := handlers.NewMainHandler(s.db)

	mux.Handle("/api/v1/", routes.MainMux(mainHandler))

	// TODO: Middleware

	corsHandler := cors.AllowAll().Handler(mux)

	log.Printf("API Running on PORT %v\n", s.addr)

	err := http.ListenAndServe(s.addr, corsHandler)
	return err
}
