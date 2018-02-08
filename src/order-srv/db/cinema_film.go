package db

import (
	"database/sql"
	"film-srv/entity"
)

func SelectFilmPrice(movieId int64) (float32,error) {

	var price float32;
	err := db.Get(&price,"SELECT release_discount FROM cinema_film WHERE `film_id` = ?",movieId)
	if err == sql.ErrNoRows {
		return 0,nil
	}
	return price,err
}

func SelectFilmNameCinemaName(hallId int64,filmId int64) (*entity.CinemaFilm,error){

	film := entity.CinemaFilm{}
	err:= db.Get(&film,"SELECT `film_name`,`cinema_name` FROM `cinema_film` WHERE `film_id` = ? AND `hall_id` = ?",filmId,hallId)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return &film,err
}