package handlers

import (
	"database/sql"
	"net/http"

	"github.com/ION-Smart/api_template_go/internal/db/repositories"
)

type MainHandler struct {
	mainRepository *repositories.MainRepository
}

func NewMainHandler(db *sql.DB) *MainHandler {
	return &MainHandler{
		repositories.NewMainRepository(db),
	}
}

func (b *MainHandler) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"App working\"}"))
}
