package handlers

import (
	"fmt"
	"net/http"

	"github.com/Movie-Api/handlers/user"
)

type LoginHandler struct{}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println("Login#" + r.Method)
	switch {
	case r.Method == http.MethodPost:
		user.Login(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
