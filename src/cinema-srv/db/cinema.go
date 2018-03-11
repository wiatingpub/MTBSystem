package db

import (
	"cinema-srv/entity"
	"database/sql"
)

func SelectCinemasByLid (locationId int64) ([]*entity.Cinema,error){

	cinemas := []*entity.Cinema{}
	err := db.Select(&cinemas,"SELECT `cinema_id`,`cinema_name`,`cinema_add`,`cinema_support`,`cinema_discount`,`cinema_min_price`,`cinema_card` FROM `cinema` WHERE `location_id` = ?",locationId)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return cinemas,err
}

func SelectCinemaByCid(cinemaId int64) (*entity.Cinema,error){

	cinema := entity.Cinema{}
	err := db.Get(&cinema,"SELECT `cinema_name`,`cinema_add`,`cinema_support`,`cinema_card`,`cinema_min_price`,`cinema_discount`,`cinema_id` FROM `cinema` WHERE `cinema_id` = ?",cinemaId)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return &cinema,err
}