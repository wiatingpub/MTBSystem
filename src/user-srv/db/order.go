package db

import (
	"database/sql"
)

func UpdateOrderScore(movieId int64,score int64) error {

	_,err := db.Exec("UPDATE `order` SET `score` = ? WHERE `movie_id` = ?",score,movieId)
	return err
}

func SelectOrderByUidMid(movieId int64,userId int64) (int64,error) {

	var orderNUM int64 = 0;
	err := db.Get(orderNUM,"SELECT order_num FROM film_order WHERE user_id = ? AND movie_id = ?",userId,movieId)
	if err != sql.ErrNoRows {
		return 0,nil
	}
	return orderNUM,err
}