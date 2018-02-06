package entity

type CinemaFilm struct {
	CfId    int64  `json:"cf_id" db:"cf_id"`
	CinemaId  int64 `json:"cinema_id" db:"cinema_id"`
	FilmId   int64  `json:"film_id" db:"film_id"`
	FilmName string `json:"film_name" db:"film_name"`
	HallId int64 `json:"hall_id" db:"hall_id"`
	ReleaseTimeYear    int64  `json:"release_time_year" db:"release_time_year"`
	ReleaseTimeMonth  int64 `json:"release_time_month" db:"release_time_month"`
	ReleaseTimeDay   int64  `json:"release_time_day" db:"release_time_day"`
	ReleaseTime string `json:"release_time" db:"release_time"`
	ReleaseType    string  `json:"release_type" db:"release_type"`
	ReleaseAdd  string `json:"release_add" db:"release_add"`
	Length   int64  `json:"length" db:"length"`
	ReleaseDiscount float32 `json:"release_discount" db:"release_discount"`
}
