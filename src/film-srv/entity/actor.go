package entity

type Actor struct {
	Id              int64    	`json:"id" db:"id"`
	NameCN		string		`json:"name_cn" db:"name_cn"`
	NameEN		string		`json:"name_en" db:"name_en"`
	ActorPhoto		string		`json:"actor_photo" db:"actor_photo"`
	ActorCountry	string		`json:"actor_country" db:"actor_country"`
}


