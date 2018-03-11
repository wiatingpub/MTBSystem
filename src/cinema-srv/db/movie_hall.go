package db

import "database/sql"

func SelectMHAddress(mhId int64) (string,error) {

	var mhAddress string
	err := db.Get(&mhAddress,"SELECT `mh_address` FROM `movie_hall` WHERE `mh_id` = ?",mhId)
	if err == sql.ErrNoRows {
		return "",nil
	}
	return mhAddress,err
}
