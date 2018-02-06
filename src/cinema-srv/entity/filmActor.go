package entity

type FilmActor struct {
	FilmId    int64  `json:"film_id" db:"film_id"`
	FilmName  string `json:"film_name" db:"film_name"`
	ActorId   int64  `json:"actor_id" db:"actor_id"`
	ActorName string `json:"actor_name" db:"actor_name"`
}
