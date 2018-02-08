package entity

type Cinema struct {
	CinemaId        int64    	`json:"cinema_id" db:"cinema_id"`
	CinemaName		string		`json:"cinema_name" db:"cinema_name"`
	CinemaAdd		string		`json:"cinema_add" db:"cinema_add"`
	LocationId		int64		`json:"location_id" db:"location_id"`
	CinemaTypes		string		`json:"cinema_types" db:"cinema_types"`
	CinemaCard		int64		`json:"cinema_card" db:"cinema_card"`
	CinemaMinPrice	int64		`json:"cinema_min_price" db:"cinema_min_price"`
	CinemaSupport	string		`json:"cinema_support" db:"cinema_support"`
	CinemaAddId	int64		`json:"cinema_add_id" db:"cinema_add_id"`
	CinemaDiscount	int64		`json:"cinema_discount" db:"cinema_discount"`
	CinemaPhone	int64		`json:"cinema_phone" db:"cinema_phone"`
}


