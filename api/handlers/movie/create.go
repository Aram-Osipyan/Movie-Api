package movie

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Movie-Api/models"
	"github.com/Movie-Api/repositories"
)

type createRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"sex"`
	ReleaseDate string  `json:"release_date"`
	Rating      float32 `json:"rating"`
	ArtistsIds  []int   `json:"artists_ids"`
}

type createResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"release_date"`
	Rating      float32 `json:"rating"`
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

	repository := new(repositories.MovieRepository)

	var date time.Time
	if date, err = time.Parse("2006-01-02", req.ReleaseDate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errorResponse{Error: "Invalid release date format"})
		return
	}

	var movie *models.Movie
	if movie, err = repository.Create(req.Name, req.Description, date, req.Rating, &req.ArtistsIds); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := createResponse{
		Id:          movie.Id,
		Name:        movie.Name,
		Description: movie.Description,
		ReleaseDate: movie.ReleaseDate.Format("2006-01-02"),
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
