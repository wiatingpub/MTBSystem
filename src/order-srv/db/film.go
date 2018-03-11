package db

import (
	"database/sql"
	"order-srv/entity"
)

// 获取影片详情
func SelectFilmDetail(movieId int64) (*entity.Film,error) {
	film := entity.Film{}
	err := db.Get(&film,"SELECT `img`,`title_cn`,`r_year`,`r_month`,`r_day`,`film_director` FROM `film` WHERE `movie_id` = ?",movieId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &film, err
}

// 更新影片评分
func UpdateFilmScore(movieId int64,score int64) error{

	_,err := db.Exec("UPDATE `film` SET `rating_final` = (`rating_final` + ?)/2 WHERE `movie_id` = ?",score,movieId)
	return err
}


func SelectFilmMessage(movieId int64) (*entity.Film,error) {
	film := entity.Film{}
	err := db.Get(&film,"SELECT `img`,`title_cn` FROM `film` WHERE `movie_id` = ?",movieId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &film, err
}