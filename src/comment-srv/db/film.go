package db

import (
	"comment-srv/entity"
	"database/sql"
)

func SelectFilmImageAndName(filmId int64) (string, string, error) {

	film := entity.Film{}
	err := db.Get(&film, "SELECT `img`,`title_cn` FROM `film` WHERE `movie_id` = ?", filmId)
	if err == sql.ErrNoRows {
		return "", "", nil
	}
	return film.Img, film.TitleCn, err
}
