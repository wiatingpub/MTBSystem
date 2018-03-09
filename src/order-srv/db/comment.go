package db

import (
	"order-srv/entity"
	"time"
)

func InsertComment(comment *entity.Comment) error {

	today := time.Now().Format("2006-01-02")
	_, err := db.Exec("INSERT INTO `comment`(`user_id`,`film_id`,`content`,`nick_name`,`create_at`,`title`) VALUES(?,?,?,?,?,?)", comment.UserId, comment.FilmId, comment.Content, comment.NickName, today, comment.Title)
	return err
}
