package models

import "time"

type Movie struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	ReleaseDate time.Time `json:release_date`
	Artists     []Artist  `json:artists`
}
