package db

import (
	"database/sql"
	"order-srv/entity"
)

func SelectUserPhoneByUserId(userId int64) (*entity.User, error) {

	user := entity.User{}
	err := db.Get(&user, "SELECT phone FROM user WHERE `user_id` = ?", userId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func SelectUserNameByUserId(userId int64) (*entity.User, error) {

	user := entity.User{}
	err := db.Get(&user, "SELECT user_name FROM user WHERE `user_id` = ?", userId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func UpdateUserPhone(userID int64, phone int64) error {

	_, err := db.Exec("UPDATE `user` SET `phone` = ? WHERE `user_id` = ?", phone, userID)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
