package db

import (
	"film-srv/entity"
	"database/sql"
)

func SelectFilmActorByMid(mid int64) ([]*entity.FilmActor,error){

	filmActors := []*entity.FilmActor{}
	err := db.Select(&filmActors,"SELECT `actor_name` FROM `film_actor` WHERE `film_id` = ?",mid)
	if err == sql.ErrNoRows {
		return nil,nil
	}

	return filmActors,err
}