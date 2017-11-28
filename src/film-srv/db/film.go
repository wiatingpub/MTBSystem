package db

import (
	"database/sql"
	"film-srv/entity"
)

// 获取正在上映的电影
func GetTickingFilims() ([]*entity.Film, error) {
	films := []*entity.Film{}
	err := db.Select(&films, "SELECT `film_id`,`logo`,`film_long`,`is_select_seat`,`film_price`,`film_screenwriter`,`comment_num`,`film_name` FROM `film`")
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return films, err
}
