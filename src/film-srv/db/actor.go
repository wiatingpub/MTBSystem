package db

import (
	"film-srv/entity"
	"database/sql"
)

func SelectActorsByMid (aid int64) (*entity.Actor,error){

	actor := entity.Actor{}
	err := db.Select(&actor,"SELECT `id`,`actor_name`,`actor_photo` FROM `actor` WHERE `id` = ?",aid)
	if err == sql.ErrNoRows {
		return nil,nil
	}

	return &actor,err
}