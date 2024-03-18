package artist

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Movie-Api/models"
	"github.com/Movie-Api/repositories"
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

type errorResponse struct {
	Error string `json:"error"`
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

	repository := new(repositories.ArtistRepository)

	var date time.Time
	if date, err = time.Parse("2006-01-02", req.BirthDate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errorResponse{Error: "Invalid birth date format"})
		return
	}

	var artist *models.Artist
	if artist, err = repository.Create(req.Name, req.Sex, date); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := createResponse{
		Id:        artist.Id,
		Name:      artist.Name,
		Sex:       artist.Sex,
		BirthDate: artist.BirthDate,
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
