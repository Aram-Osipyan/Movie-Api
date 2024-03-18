package handlers

import (
	"fmt"
	"net/http"

	"github.com/Movie-Api/handlers/user"
)

type UserHandler struct{}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println("USERS#" + r.Method)
	switch {
	case r.Method == http.MethodPost:
		user.Create(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
