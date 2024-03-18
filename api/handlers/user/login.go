package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Movie-Api/repositories"
	"github.com/golang-jwt/jwt"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

	user_repository := new(repositories.UserRepository)

	user, err := user_repository.Find(req.Username, req.Password)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// TODO: move to environment
	var sampleSecretKey = []byte("BknHBOeY3j7lxUKYK8TUHDsx5J0KUqnIt81TeHllIrY=")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["expiration"] = time.Now().Add(time.Hour)
	claims["role"] = user.Role
	claims["user_id"] = user.Id

	var tokenString string
	if tokenString, err = token.SignedString(sampleSecretKey); err == nil {
		json.NewEncoder(w).Encode(&loginResponse{AccessToken: tokenString})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(tokenString)
	}
}
