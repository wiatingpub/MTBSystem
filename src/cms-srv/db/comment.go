package db

import (
	"cms-srv/entity"
	"database/sql"
)

func SelectAllComment(page int64, num int64) ([]*entity.Comment, error) {

	comments := []*entity.Comment{}
	err := db.Select(&comments, "SELECT * FROM `comment`  ORDER BY `comment_id` DESC LIMIT ?,?", (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return comments, nil
}

func SelectCommentTotal() (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `comment`")
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, nil
}

func SelectCommentsByCID(page int64, num int64, filmId int64) ([]*entity.Comment, error) {

	comment := []*entity.Comment{}
	err := db.Select(&comment, "SELECT * FROM `comment`  WHERE `film_id` = ? ORDER BY `comment_id` DESC LIMIT ?,?", filmId, (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return comment, nil
}

func SelectCommentsTotalByCID(filmId int64) (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `comment` WHERE `film_id` = ?", filmId)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, nil
}
