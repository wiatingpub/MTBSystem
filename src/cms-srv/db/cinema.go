package db

import (
	"cms-srv/entity"
	"database/sql"
)

func SelectCinemaName(cinemaId int64) (string, error) {

	var cinemaName string
	err := db.Get(&cinemaName, "SELECT `cinema_name` FROM `cinema` WHERE `cinema_id` = ? LIMIT 1", cinemaId)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return cinemaName, err
}

func InsertCinema(cinema *entity.Cinema) error {

	_, err := db.Exec("INSERT INTO `cinema`(`cinema_name`,`cinema_add`,`location_id`,`cinema_types`,`cinema_card`,`cinema_min_price`,`cinema_support`,`cinema_discount`,`cinema_phone`) VALUES(?,?,?,?,?,?,?,?,?)",
		cinema.CinemaName, cinema.CinemaAdd, cinema.LocationId, cinema.CinemaTypes, cinema.CinemaCard, cinema.CinemaMinPrice, cinema.CinemaSupport, cinema.CinemaDiscount, cinema.CinemaPhone)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func SelectCinema(cinema *entity.Cinema) (*entity.Cinema, error) {

	cinemaTmp := entity.Cinema{}
	err := db.Get(&cinemaTmp, "SELECT * FROM `cinema` WHERE  `cinema_phone`=?", cinema.CinemaName, cinema.LocationId, cinema.CinemaPhone)
	if err == sql.ErrNoRows {
		return &cinemaTmp, nil
	}
	return &cinemaTmp, err
}
