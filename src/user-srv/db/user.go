package db

import (
	"time"
	"user-srv/entity"
	"database/sql"
)

func InsertUser(userName string,password string,email string) error{

	today := time.Now().Format("2006-01-02")
	_,err := db.Exec("INSERT INTO `user`(`user_name`,`password`,`create_at`,`email`) VALUES(?,?,?,?)",userName,password,today,email)
	return err
}

func SelectUserByEmail(email string) (*entity.User,error) {

	user := entity.User{}
	err := db.Get(&user,"SELECT * FROM user WHERE `email` = ?",email)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return &user,err
}

func SelectUserByPasswordName(name string,password string) (*entity.User,error) {

	user := entity.User{}
	err := db.Get(&user,"SELECT * FROM user WHERE `user_name` = ? AND `password` = ? ",name,password)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return &user,err
}
