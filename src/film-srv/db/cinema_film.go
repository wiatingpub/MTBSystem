package db

import (
	"film-srv/entity"
	"database/sql"
)

func SelectFilmMessageCidADay(cinemaId,filmId,year,month,day int64) ([]*entity.CinemaFilm,error) {

	cinemaFilms := []*entity.CinemaFilm{}
	err := db.Select(&cinemaFilms,"SELECT `film_name`,`film_id`,`release_type`,`release_time`,`length`,`release_type`,`release_discount`,`hall_id`,`cinema_id` FROM `cinema_film` WHERE `cinema_id` = ? AND `film_id` = ? AND `release_time_year` = ? AND `release_time_month` = ? AND `release_time_day` = ?",cinemaId,filmId,year,month,day)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return cinemaFilms,err
}
