package db

import (
	"database/sql"
	"film-srv/entity"
)

// 获取正在上映的电影
func SelectTickingFilims() ([]*entity.Film, error) {
	films := []*entity.Film{}
	err := db.Select(&films, "SELECT `title_cn`,`img`,`movie_id` FROM `film`")
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return films, err
}
