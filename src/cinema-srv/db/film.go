package db

import (
	"cinema-srv/entity"
	"database/sql"
)

// 获取影片详情
func SelectFilmMesage(movieId int64) (*entity.Film, error) {
	film := entity.Film{}
	err := db.Get(&film, "SELECT `title_cn`,`rating_final`,`length`,`img`,`type`,`movie_id` FROM `film` WHERE `movie_id` = ?", movieId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &film, err
}

// 获取演员信息
func SelectActorNameByMid(mid int64) ([]*entity.FilmActor, error) {

	filmActors := []*entity.FilmActor{}
	err := db.Select(&filmActors, "SELECT DISTINCT(`actor_name`) FROM `film_actor` WHERE `film_id` = ?", mid)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return filmActors, err
}
