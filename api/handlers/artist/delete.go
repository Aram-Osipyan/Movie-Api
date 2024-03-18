package artist

import (
	"net/http"

	"github.com/Movie-Api/repositories"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("role").(string)
	if role != "admin" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	repository := new(repositories.ArtistRepository)

	var id string = r.PathValue("id")
	if err := repository.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
