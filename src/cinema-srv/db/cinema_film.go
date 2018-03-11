package db

import (
	"database/sql"
)

func SelectMidByCid (cinemaId int64,year int64,month int64,day int64) ([]int64,error){

	movies := []int64{}
	err := db.Select(&movies,"SELECT `film_id` FROM `cinema_film` WHERE `cinema_id` = ? AND `release_time_year` = ? AND `release_time_month` = ? AND `release_time_day` = ?",cinemaId,year,month,day)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return movies,err
}

