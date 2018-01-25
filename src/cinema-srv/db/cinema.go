package db

import (
	"cinema-srv/entity"
	"database/sql"
)

func SelectCinemasByLid (locationId int64) ([]*entity.Cinema,error){

	cinemas := []*entity.Cinema{}
	err := db.Select(&cinemas,"SELECT `cinema_name`,`cinema_add` FROM `cinema` WHERE `location_id` = ?)",locationId)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return cinemas,err
}