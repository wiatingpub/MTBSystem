package db

import "time"

func InsertWantSeeRecord(movieId int64,userId int64) error {

	today := time.Now().Format("2006-01-02")
	_,err := db.Exec("INSERT INTO `want_see_record`(`movie_id`,`user_id`,`create_at`) VALUES(?,?,?) ",movieId,userId,today)
	return err
}
