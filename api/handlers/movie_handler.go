package handlers

import (
	"fmt"
	"net/http"

	"github.com/Movie-Api/handlers/artist"
)

type MovieHandler struct{}

func (h *MovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println("ARTISTS#" + r.Method)
	switch {
	case r.Method == http.MethodPost:
		artist.Create(w, r)
		return
	case r.Method == http.MethodPatch:
		artist.Update(w, r)
		return
	case r.Method == http.MethodDelete:
		artist.Delete(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
