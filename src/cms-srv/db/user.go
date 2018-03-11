package db

import (
	"cms-srv/entity"
	"database/sql"
)

func SelectAllUsers(page int64, num int64) ([]*entity.User, error) {

	users := []*entity.User{}
	err := db.Select(&users, "SELECT * FROM `user`  ORDER BY `user_id` DESC LIMIT ?,?", (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return users, nil
}

func SelectUserTotal() (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `user`")
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, nil
}
