package handlers

import (
	"fmt"
	"net/http"

	"github.com/Movie-Api/handlers/movie"
)

type MovieHandler struct{}

func (h *MovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println("MOVIES#" + r.Method)
	switch {
	case r.Method == http.MethodGet:
		movie.Get(w, r)
		return
	case r.Method == http.MethodPost:
		movie.Create(w, r)
		return
	case r.Method == http.MethodDelete:
		movie.Delete(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
