package db

import (
	"database/sql"
	"film-srv/entity"
)

func SelectActors(fid int64, actor_type int64) ([]*entity.Actor, error) {

	actors := []*entity.Actor{}
	err := db.Select(&actors, "SELECT * FROM `actor` WHERE `actor_type` = ? AND `id` IN (SELECT `actor_id` FROM `film_actor` WHERE `film_id` = ? )", actor_type, fid)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return actors, err
}
