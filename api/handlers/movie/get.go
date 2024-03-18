package movie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Movie-Api/models"
	"github.com/Movie-Api/repositories"
)

type getResponse struct {
	Data []models.Movie `json:"data"`
}

func Get(w http.ResponseWriter, r *http.Request) {
	repository := new(repositories.MovieRepository)
	q_sort := r.URL.Query().Get("q_sort")
	fmt.Println("q_sort: ", q_sort)
	if q_sort == "" {
		q_sort = "rating,desc"
	}

	params := strings.Split(q_sort, ",")

	if len(params) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid sort query"))
		return
	}

	var movies []models.Movie
	var err error
	if movies, err = repository.Get(params[0], params[1]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	response := getResponse{
		Data: movies,
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
