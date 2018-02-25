package db

import (
	"cms-srv/entity"
	"database/sql"
)

func InsertActor(actor *entity.Actor, actorType int64) (error, int64) {

	var id int64
	_, err := db.Exec("INSERT INTO  `actor`(`name_cn`,`actor_photo`,`actor_type`) VALUES(?,?,?)",
		actor.NameCN, actor.ActorPhoto, actorType)
	if err == sql.ErrNoRows {
		return nil, 0
	}
	if err == nil {
		err = db.Get(&id, "SELECT `id` FROM `actor` WHERE `name_cn`= ? AND `actor_photo` = ? LIMIT 1", actor.NameCN, actor.ActorPhoto)
		if err == sql.ErrNoRows {
			return nil, 0
		}
	}
	return err, id
}
