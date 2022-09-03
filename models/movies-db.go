package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Hero, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date, rating, runtime, mpaa_rating,
			created_at, updated_at from movies where id = $1
	`
	row := m.DB.QueryRowContext(ctx, query, id)

	var hero Hero

	err := row.Scan(
		&hero.ID,
		&hero.Name,
		&hero.Description,
	)
	if err != nil {
		return nil, err
	}

	query = `select 
				mg.id, mg.movie_id, mg.genre_id, g.genre_name
			from
				movies_genres mg
				left join genres g on (g.id = mg.genre_id)
			where
				mg.movie_id = $1`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	var genres = make(map[int]string)
	for rows.Next() {
		var hg HeroGenre
		err := rows.Scan(
			&hg.ID,
			&hg.MovieID,
			&hg.GenreID,
			&hg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}

		genres[hg.ID] = hg.Genre.GenreName
	}

	hero.HeroGenre = genres

	return &hero, nil
}

//GET ALL

func (m *DBModel) All(genre ...int) ([]*Hero, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""
	if len(genre) > 0 {
		where = fmt.Sprintf("where id in (select movie_id from movies_genres where genre_id = %d)", genre[0])
	}

	query := fmt.Sprintf(`SELECT id, title, description, year, release_date, rating, runtime, mpaa_rating,
	created_at, updated_at FROM movies %s`, where)

	rows, _ := m.DB.QueryContext(ctx, query)
	defer rows.Close()

	var heroes []*Hero

	for rows.Next() {
		var hero Hero
		err := rows.Scan(
			&hero.ID,
			&hero.Name,
			&hero.Description,
		)
		if err != nil {
			return nil, err
		}

		genreQuery := `select 
				mg.id, mg.hero_id, mg.genre_id, g.genre_name
			from
				heroes_genres mg
				left join genres g on (g.id = mg.genre_id)
			where
				mg.hero_id = $1`

		genreRows, _ := m.DB.QueryContext(ctx, genreQuery, hero.ID)

		genres := make(map[int]string)
		for genreRows.Next() {
			var hg HeroGenre
			err := genreRows.Scan(
				&hg.ID,
				&hg.MovieID,
				&hg.GenreID,
				&hg.Genre.GenreName,
			)
			if err != nil {
				return nil, err
			}
			genres[hg.ID] = hg.Genre.GenreName
		}
		genreRows.Close()

		hero.HeroGenre = genres
		heroes = append(heroes, &hero)

	}
	return heroes, nil
}
