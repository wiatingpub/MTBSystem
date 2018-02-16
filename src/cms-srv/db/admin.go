package db

import (
	"cms-srv/entity"
	"database/sql"
)

func SelectAdmin(user string, password string) (*entity.Admin, error) {

	admin := entity.Admin{}
	err := db.Get(&admin, "SELECT `admin_cinema_id`,`admin_last_login_time`,`admin_num` FROM `admin_user` WHERE `admin_name` = ? AND `admin_password` = ? LIMIT 1", user, password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &admin, err
}
