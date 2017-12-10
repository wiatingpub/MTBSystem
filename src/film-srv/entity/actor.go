package entity

type Actor struct {
	Id              int64    	`json:"id" db:"id"`
	ActorName		string		`json:"actor_name" db:"actor_name"`
	ActorPhoto		string		`json:"actor_photo" db:"actor_photo"`
	ActorCountry	string		`json:"actor_country" db:"actor_country"`
}


