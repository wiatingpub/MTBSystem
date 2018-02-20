package db

import (
	"cms-srv/entity"
	"database/sql"
)

func SelectAllPlace(page int64, num int64) ([]*entity.Place, error) {

	places := []*entity.Place{}
	err := db.Select(&places, "SELECT * FROM `place`  ORDER BY `id` DESC LIMIT ?,?", (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return places, nil
}

func SelectPlaceTotal() (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `place`")
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, nil
}
