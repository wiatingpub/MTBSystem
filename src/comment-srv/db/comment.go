package db

import (
	"comment-srv/entity"
	"database/sql"
	"time"
)

func SelectHotComment(movieId int64) ([]*entity.Comment, error) {

	comments := []*entity.Comment{}
	err := db.Select(&comments, "SELECT * FROM `comment`")
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return comments, err
}

func InsertHotComment(comment *entity.Comment) error {

	today := time.Now().Format("2006-01-02")
	_, err := db.Exec("INSERT INTO `comment`(`film_id`,`title`,`content`,`head_img`,`nick_name`,`create_at`) ", comment.FilmId, comment.Title, comment.Content, comment.HeadImg, comment.NickName, today)
	return err
}

func UpdateHotComment(commentID int64) error {
	_, err := db.Exec("UPDATE `comment` SET `up_num` = `up_num`+1 WHERE `comment_id` = ?", commentID)
	return err
}

func SelectUpNum(commentID int64) (int64, error) {

	var upNUm int64
	err := db.Get(&upNUm, "SELECT `up_num` FROM `comment` WHERE `comment_id` = ?", commentID)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return upNUm, err
}

func SelectMyComment(userId int64) ([]*entity.Comment, error) {

	comments := []*entity.Comment{}
	err := db.Select(&comments, "SELECT * FROM `comment` WHERE `user_id` = ?", userId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return comments, err
}

func DeleteComment(commentID int64) error {
	_, err := db.Exec("DELETE FROM `comment` WHERE `comment_id` = ?", commentID)
	return err
}
