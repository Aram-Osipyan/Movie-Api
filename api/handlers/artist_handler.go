package handlers

import (
	"fmt"
	"net/http"

	"github.com/Movie-Api/handlers/user"
)

type ArtistHandler struct{}

func (h *ArtistHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println("ARTISTS#" + r.Method)
	switch {
	case r.Method == http.MethodPost:
		user.Create(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
