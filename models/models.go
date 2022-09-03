package models

import (
	"database/sql"
)

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Hero struct {
	ID          int            `json:"id"`
	Name        string         `json:"title"`
	Description string         `json:"description"`
	HeroGenre   map[int]string `json:"genres"`
}

type Genre struct {
	ID        int    `json:"id"`
	GenreName string `json:"genre_name"`
}

type HeroGenre struct {
	ID      int   `json:"-"`
	MovieID int   `json:"-"`
	GenreID int   `json:"-"`
	Genre   Genre `json:"genre"`
}
