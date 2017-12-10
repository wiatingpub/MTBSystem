package db

import (
	"database/sql"
	"film-srv/entity"
)

// 获取正在上映的电影
func SelectTickingFilims() ([]*entity.Film, error) {
	films := []*entity.Film{}
	err := db.Select(&films, "SELECT `type`,`title_cn`,`title_en`,`common_special`,`director_name`,`img`,`is_3D`,`is_DMAX`,`is_filter`,`is_hot`,`is_IMAX`,`is_IMAX3D`,`is_new`,`length`,`movie_id`,`r_day`,`r_month`,`r_year`,`rating_final`,`wanted_count` FROM `film`")
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return films, err
}
