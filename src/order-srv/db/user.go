package db

import (
	"database/sql"
	"order-srv/entity"
)

func SelectUserPhoneByUserId(userId int64) (*entity.User,error) {

	user := entity.User{}
	err := db.Get(&user,"SELECT phone FROM user WHERE `userId` = ?",userId)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return &user,err
}

func SelectUserNameByUserId(userId int64) (*entity.User,error) {

	user := entity.User{}
	err := db.Get(&user,"SELECT user_name FROM user WHERE `userId` = ?",userId)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return &user,err
}