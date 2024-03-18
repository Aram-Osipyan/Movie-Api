package user

import (
	"encoding/json"
	"net/http"

	"github.com/Movie-Api/models"
	"github.com/Movie-Api/repositories"
)

type createRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type createResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:role`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	var err error

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}

	user_repository := new(repositories.UserRepository)

	var user *models.User
	if user, err = user_repository.Create(req.Username, req.Password, "user"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := createResponse{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
