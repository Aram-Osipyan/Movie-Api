package repositories

import (
	"fmt"
	"time"

	"github.com/Movie-Api/models"
)

type ArtistRepository struct {
	BaseRepository
}

func (repository *ArtistRepository) Create(name, sex string, birth_date time.Time) (*models.Artist, error) {
	db, err := repository.GetConnection()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("INSERT INTO artists (name, sex, birth_date) VALUES($1, $2, $3) RETURNING id;", name, sex, birth_date.Format("2006-01-02"))
	defer db.Close()

	if err != nil {
		fmt.Println("query error: ", err)
		return nil, err
	}

	fmt.Println("new artist created")

	new_artist := models.Artist{Name: name, Sex: sex, BirthDate: birth_date}

	for rows.Next() {
		err = rows.Scan(&new_artist.Id)
		break
	}

	return &new_artist, err
}

func (repository *ArtistRepository) Update(id string, name, sex string, birth_date time.Time) (*models.Artist, error) {
	db, err := repository.GetConnection()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("UPDATE artists SET name = $1, sex = $2, birth_date = $3 WHERE id = $4 RETURNING id, name, sex, birth_date;", name, sex, birth_date.Format("2006-01-02"), id)
	defer db.Close()

	if err != nil {
		fmt.Println("query error: ", err)
		return nil, err
	}

	fmt.Println("new artist created")

	artist := models.Artist{Name: name, Sex: sex, BirthDate: birth_date}

	for rows.Next() {
		err = rows.Scan(&artist.Id, &artist.Name, &artist.Sex, &artist.BirthDate)
		break
	}

	return &artist, err
}

func (repository *ArtistRepository) Delete(id string) error {
	db, err := repository.GetConnection()
	if err != nil {
		return err
	}

	_, err = db.Query("DELETE from artists WHERE id = $1;", id)
	defer db.Close()

	if err != nil {
		fmt.Println("query error: ", err)
		return err
	}

	fmt.Println("new artist removed")

	return nil
}
