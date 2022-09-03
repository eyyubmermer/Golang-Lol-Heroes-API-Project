package models

import (
	"context"
	"time"
)

func (m *DBModel) GetGenre(id int) (*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, genre_name, created_at, updated_at from genres where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var genre Genre

	err := row.Scan(
		&genre.ID,
		&genre.GenreName,
	)
	if err != nil {
		return nil, err
	}

	return &genre, nil
}

// GET ALL
func (m *DBModel) AllGenres() ([]*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, genre_name, created_at, updated_at 
				FROM genres`

	rows, _ := m.DB.QueryContext(ctx, query)
	defer rows.Close()

	var genres []*Genre

	for rows.Next() {
		var genre Genre
		err := rows.Scan(
			&genre.ID,
			&genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres = append(genres, &genre)
	}
	return genres, nil
}
