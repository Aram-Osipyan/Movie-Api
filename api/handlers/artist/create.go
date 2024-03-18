package artist

import (
	"encoding/json"
	"net/http"
	"time"
)

type createRequest struct {
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	BirthDate string `json:"birth_date"`
}

type createResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Sex       string    `json:"sex"`
	BirthDate time.Time `json:"birth_date"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

	role := r.Context().Value("role").(string)
	if role != "admin" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
}
