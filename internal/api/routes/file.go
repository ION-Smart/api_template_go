package routes

import (
	"net/http"

	"github.com/ION-Smart/api_template_go/internal/api/handlers"
)

func MainMux(handler *handlers.MainHandler) http.Handler {
	mainMux := http.NewServeMux()

	mainMux.HandleFunc("GET /", handler.Info)

	return http.StripPrefix("/api/v1/", mainMux)
}
