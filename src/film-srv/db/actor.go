package db

import (
	"film-srv/entity"
	"database/sql"
)

func SelectActors (fid int64,actor_type int64) ([]*entity.Actor,error){

	actors := []*entity.Actor{}
	err := db.Select(&actors,"SELECT `id`,`name_cn`,`name_en`,`actor_photo` FROM `actor` WHERE `actor_type` = ? AND `id` IN (SELECT `actor_id` FROM `film_actor` WHERE `film_id` = ? )",actor_type,fid)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return actors,err
}