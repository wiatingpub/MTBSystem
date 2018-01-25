package entity

type Cinema struct {
	CinemaId        int64    	`json:"cinema_id" db:"cinema_id"`
	CinemaName		string		`json:"cinema_name" db:"cinema_name"`
	CinemaAdd		string		`json:"cinema_add" db:"cinema_add"`
	LocationId		int64		`json:"location_id" db:"location_id"`
	CinemaTypes		string		`json:"cinema_types" db:"cinema_types"`
}


