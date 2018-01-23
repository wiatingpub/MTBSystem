package db

import (
	"database/sql"
)

func SelectFilmPrice(movieId int64) (float32,error) {

	var price float32;
	err := db.Get(&price,"SELECT release_discount FROM cinema_film WHERE `film_id` = ?",movieId)
	if err == sql.ErrNoRows {
		return 0,nil
	}
	return price,err
}
