package db

import (
	"cms-srv/entity"
	"database/sql"
)

func SelectAllMovieHall(page int64, num int64) ([]*entity.MovieHall, error) {

	movieHalls := []*entity.MovieHall{}
	err := db.Select(&movieHalls, "SELECT * FROM `movie_hall`  ORDER BY `mh_id` DESC LIMIT ?,?", (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return movieHalls, err
}

func SelectAllMovieHallBycinemaID(page int64, num int64, cinemaID int64) ([]*entity.MovieHall, error) {

	movieHalls := []*entity.MovieHall{}
	err := db.Select(&movieHalls, "SELECT * FROM `movie_hall` WHERE `cinema_id` = ? ORDER BY `mh_id` DESC LIMIT ?,?", cinemaID, (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return movieHalls, err
}

func SelectAllMovieHallByMHID(mhId int64) (*entity.MovieHall, error) {

	movieHall := entity.MovieHall{}
	err := db.Get(&movieHall, "SELECT * FROM `movie_hall` WHERE `mh_id` = ? LIMIT 1", mhId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &movieHall, err
}

func SelectMovieHallTotalByCinemaID(cinemaID int64) (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `movie_hall` WHERE `cinema_id` = ? ", cinemaID)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, err
}

func SelectMovieHallTotal() (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `movie_hall`")
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, err
}

func AddMovieHall(movieHall *entity.MovieHall) error {

	_, err := db.Exec("INSERT INTO  `movie_hall`(`mh_name`,`mh_address`,`cinema_id`) VALUES(?,?,?)",
		movieHall.MhName, movieHall.MhAddress, movieHall.CinemaId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func DeleteMovieHall(mhID int64) error {
	_, err := db.Exec("DELETE FROM `movie_hall` WHERE `mh_id` = ?", mhID)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateMovieHall(movieHall *entity.MovieHall) error {
	_, err := db.Exec("UPDATE `movie_hall` SET `mh_name` = ?,`mh_address` = ?,`cinema_id` = ? WHERE `mh_id` = ?",
		movieHall.MhName, movieHall.MhAddress, movieHall.CinemaId, movieHall.MhID)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func SelectAllMovieHallsBycinemaID(cinemaID int64) ([]*entity.MovieHall, error) {

	movieHalls := []*entity.MovieHall{}
	err := db.Select(&movieHalls, "SELECT * FROM `movie_hall` WHERE `cinema_id` = ? ORDER BY `mh_id` DESC", cinemaID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return movieHalls, err
}

func SelectMovieHallHallName(mhID int64) (string, error) {

	var hallName string
	err := db.Get(&hallName, "SELECT `mh_name` FROM `movie_hall` WHERE `mh_id` = ? ", mhID)
	if err == sql.ErrNoRows {
		return hallName, nil
	}
	return hallName, err
}
