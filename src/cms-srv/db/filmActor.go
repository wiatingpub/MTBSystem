package db

import (
	"database/sql"
)

func InsertFilmActor(filmID, actorID int64, filmName, actorName string) error {

	_, err := db.Exec("INSERT INTO  `film_actor`(`film_id`,`film_name`,`actor_id`,`actor_name`) VALUES(?,?,?,?)",
		filmID, filmName, actorID, actorName)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
