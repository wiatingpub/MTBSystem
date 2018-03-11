package entity

type MovieHall struct {
	MhId     int64   `json:"mh_id" db:"mh_id"`
	MhName    string  `json:"mh_name" db:"mh_name"`
	MhAddress int64   `json:"mh_address" db:"mh_address"`
	CinemaId  int64   `json:"cinema_id" db:"cinema_id"`
}
