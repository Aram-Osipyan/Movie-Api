package user

import (
	"encoding/json"
	"net/http"

	"github.com/Movie-Api/repositories"
)

type request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type response struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:role`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

	user_repository := new(repositories.UserRepository)
	user, err := user_repository.Create(req.Username, req.Password, "user")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
