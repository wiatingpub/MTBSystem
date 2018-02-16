package db

import (
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
