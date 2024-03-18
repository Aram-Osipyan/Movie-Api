package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Movie-Api/models"
)

type MovieRepository struct {
	BaseRepository
}

func (repository *MovieRepository) Create(name, description string, release_date time.Time, rating float32, artist_ids *[]int) (*models.Movie, error) {
	db, err := repository.GetConnection()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("INSERT INTO movies (name, description, release_date, rating) VALUES($1, $2, $3, $4) RETURNING id;", name, description, release_date.Format("2006-01-02"), rating)
	defer db.Close()

	if err != nil {
		fmt.Println("query error: ", err)
		return nil, err
	}

	new_movie := models.Movie{Name: name, Description: description, ReleaseDate: release_date, Rating: rating}

	for rows.Next() {
		err = rows.Scan(&new_movie.Id)
		break
	}

	// add artists to the movie
	var values []string
	for _, artist_id := range *artist_ids {
		values = append(values, fmt.Sprintf("(%d, %d)", new_movie.Id, artist_id))
	}

	db.Query(fmt.Sprintf("INSERT INTO movie_artists (movie_id, artist_id) VALUES %s", strings.Join(values, ",")))
	defer fmt.Println("new movie created")

	return &new_movie, err
}

func (repository *MovieRepository) Delete(id string) error {
	db, err := repository.GetConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	queries := []string{
		"DELETE from movies WHERE id = $1;",
		"DELETE from movie_artists WHERE movie_id = $1;",
	}

	for _, query := range queries {
		if _, err := db.Query(query, id); err != nil {
			fmt.Println("query error: ", err)
			return err
		}
	}

	fmt.Println("movie removed")

	return nil
}

func (repository *MovieRepository) Get(sort_by, sort_method string) ([]models.Movie, error) {
	fmt.Println(sort_by, " ", sort_method)
	db, err := repository.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var rows *sql.Rows

	if sort_method != "desc" && sort_method != "asc" {
		fmt.Println("sort_method must be 'desc' or 'asc'")
		return nil, errors.New("sort_method must be 'desc' or 'asc'")
	}

	sortQuery := fmt.Sprintf("SELECT id, name, description, release_date, rating FROM movies ORDER BY $1 %s", sort_method)

	var stmt *sql.Stmt
	if stmt, err = db.Prepare(sortQuery); err != nil {
		fmt.Println("prepare error: ", err)
		return nil, err
	}

	if rows, err = stmt.Query(sort_by); err != nil {
		fmt.Println("query error: ", err)
		return nil, err
	}

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.Id, &movie.Name, &movie.Description, &movie.ReleaseDate, &movie.Rating); err != nil {
			fmt.Println("scan error: ", err)
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}
