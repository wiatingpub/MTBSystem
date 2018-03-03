package db

import (
	"order-srv/entity"
	"time"
)

func InsertComment(comment *entity.Comment) error {

	today := time.Now().Format("2006-01-02")
	_, err := db.Exec("INSERT INTO `comment`(`film_id`,`content`,`nick_name`,`create_at`) VALUES(?,?,?,?)", comment.FilmId, comment.Content, comment.NickName, today)
	return err
}
