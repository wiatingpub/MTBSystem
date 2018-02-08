package db

import (
	"time"
	"order-srv/entity"
)

func InsertComment(comment *entity.Comment) error {

	today := time.Now().Format("2006-01-02")
	_,err := db.Exec("INSERT INTO `comment`(`film_id`,`title`,`content`,`head_img`,`nick_name`,`create_at`) ",comment.FilmId,comment.Title,comment.Content,comment.HeadImg,comment.NickName,today)
	return err
}