package db

import (
	"database/sql"
	"film-srv/entity"
)

// 获取正在上映的电影
func SelectTickingFilims() ([]*entity.Film, error) {
	films := []*entity.Film{}
	err := db.Select(&films, "SELECT `title_cn`,`is_3D`,`is_DMAX`,`is_IMAX`,`is_IMAX3D`,`img`,`movie_id`,`film_director`,`film_drama` FROM `film`")
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return films, err
}

// 获取影片详情
func SelectFilmDetail(movieId int64) (*entity.Film,error) {
	film := entity.Film{}
	err := db.Get(&film,"SELECT * FROM `film` WHERE `movie_id` = ?",movieId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &film, err
}
