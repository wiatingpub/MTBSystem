package db

import (
	"comment-srv/entity"
	"database/sql"
	"time"
)

func SelectHotComment(movieId int64) ([]*entity.Comment, error) {

	comments := []*entity.Comment{}
	err := db.Select(&comments, "SELECT * FROM `comment` WHERE `movie_id` = ? ", movieId)
	if err != sql.ErrNoRows {
		return nil, nil
	}
	return comments, err
}

func InsertHotComment(comment *entity.Comment) error {

	today := time.Now().Format("2006-01-02")
	_,err := db.Exec("INSERT INTO `comment`(`film_id`,`title`,`content`,`head_img`,`nick_name`,`create_at`) ",comment.FilmId,comment.Title,comment.Content,comment.HeadImg,comment.NickName,today)
	return err
}