package db

import (
	"cms-srv/entity"
	"database/sql"
)

func SelectAllCinemaFilm(page int64, num int64) ([]*entity.CinemaFilm, error) {

	cinemaFilms := []*entity.CinemaFilm{}
	err := db.Select(&cinemaFilms, "SELECT * FROM `cinema_film`  ORDER BY `cf_id` DESC LIMIT ?,?", (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return cinemaFilms, err
}

func SelectAllCinemaFilmByCinemaID(page int64, num int64, cinemaID int64) ([]*entity.CinemaFilm, error) {

	cinemaFilms := []*entity.CinemaFilm{}
	err := db.Select(&cinemaFilms, "SELECT * FROM `cinema_film` WHERE `cinema_id` = ? ORDER BY `cf_id` DESC LIMIT ?,?", cinemaID, (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return cinemaFilms, err
}

func SelectAllCinemaFilmByMHID(mhId int64) (*entity.CinemaFilm, error) {

	cinemaFilms := entity.CinemaFilm{}
	err := db.Select(&cinemaFilms, "SELECT * FROM `cinema_film` WHERE `hall_id` = ? ", mhId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &cinemaFilms, err
}

func SelectAllCinemaFilmTotalByCinemaID(cinemaID int64) (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `cinema_film` WHERE `cinema_id` = ? ", cinemaID)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, err
}

func SelecttAllCinemaFilmTotal() (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `cinema_film`")
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, err
}

func AddAllCinemaFilm(cinemaFilm *entity.CinemaFilm) error {

	_, err := db.Exec("INSERT INTO  `cinema_film`(`cinema_id`,`film_id`,`hall_id`,`film_name`,`cinema_name`,`release_time_year`,`release_time_month`,`release_time_day`,`release_time`,`release_type`,`release_add`,`length`,`release_discount`) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)", cinemaFilm.CinemaId, cinemaFilm.FilmId, cinemaFilm.HallId, cinemaFilm.FilmName, cinemaFilm.CinemaName, cinemaFilm.ReleaseTimeYear, cinemaFilm.ReleaseTimeMonth, cinemaFilm.ReleaseTimeDay, cinemaFilm.ReleaseTime, cinemaFilm.ReleaseType, cinemaFilm.ReleaseAdd, cinemaFilm.Length, cinemaFilm.ReleaseDiscount)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func DeleteCinemaFilm(cfID int64) error {
	_, err := db.Exec("DELETE FROM `cinema_film` WHERE `cf_id` = ?", cfID)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateCinemaFilm(cinemaFilm *entity.CinemaFilm) error {
	_, err := db.Exec("UPDATE `cinema_film` SET `cinema_id` = ?,`film_id` = ?,`hall_id` = ?,`film_name` = ?,`cinema_name` = ?,`release_time_year` = ?,`release_time_month` = ?,`release_time_day` = ?,`release_time` = ?,`release_type` = ?,`release_add` = ?,`length` = ?,`release_discount` = ? WHERE `cf_id` = ?",
		cinemaFilm.CinemaId, cinemaFilm.FilmId, cinemaFilm.HallId, cinemaFilm.FilmName, cinemaFilm.CinemaName, cinemaFilm.ReleaseTimeYear, cinemaFilm.ReleaseTimeMonth, cinemaFilm.ReleaseTimeDay, cinemaFilm.ReleaseTime, cinemaFilm.ReleaseType, cinemaFilm.ReleaseAdd, cinemaFilm.Length, cinemaFilm.ReleaseDiscount, cinemaFilm.CfId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
