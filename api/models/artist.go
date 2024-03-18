package models

import "time"

type Artist struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Sex       string    `json:"sex"`
	BirthDate time.Time `json:birth_date`
}
