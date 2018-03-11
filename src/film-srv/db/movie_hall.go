package db

import "database/sql"

func SelectMHName(mhId int64) (string,error) {

	var mhName string
	err := db.Get(&mhName,"SELECT `mh_name` FROM `movie_hall` WHERE `mh_id` = ?",mhId)
	if err == sql.ErrNoRows {
		return "",nil
	}
	return mhName,err
}
