package db

import "time"

func InsertUser(userName string,password string) error{

	createAt := time.Now().Format("2016-01-01")
	_,err := db.Exec("INSERT INTO `user`(`user_name`,`password`,`create_at`) VALUES(?,?,?)",userName,password,createAt)
	return err
}
