package entity

import "share/pb"

type CinemaFilm struct {
	CfId             int64   `json:"cf_id" db:"cf_id"`
	CinemaId         int64   `json:"cinema_id" db:"cinema_id"`
	FilmId           int64   `json:"film_id" db:"film_id"`
	HallId           int64   `json:"hall_id" db:"hall_id"`
	FilmName         string  `json:"film_name" db:"film_name"`
	CinemaName       string  `json:"cinema_name" db:"cinema_name"`
	ReleaseTimeYear  int64   `json:"release_time_year" db:"release_time_year"`
	ReleaseTimeMonth int64   `json:"release_time_month" db:"release_time_month"`
	ReleaseTimeDay   int64   `json:"release_time_day" db:"release_time_day"`
	ReleaseTime      string  `json:"release_time" db:"release_time"`
	ReleaseType      string  `json:"release_type" db:"release_type"`
	ReleaseAdd       string  `json:"release_add" db:"release_add"`
	Length           int64   `json:"length" db:"length"`
	ReleaseDiscount  float32 `json:"release_discount" db:"release_discount"`
}

func (c CinemaFilm) ToProtoCinemaFilm() *pb.CinemaFilm {
	return &pb.CinemaFilm{
		CfID:             c.CfId,
		CinemaID:         c.CinemaId,
		FilmID:           c.FilmId,
		HallID:           c.HallId,
		FilmName:         c.FilmName,
		CinemaName:       c.CinemaName,
		ReleaseTimeYear:  c.ReleaseTimeYear,
		ReleaseTimeMonth: c.ReleaseTimeMonth,
		ReleaseTimeDay:   c.ReleaseTimeDay,
		ReleaseTime:      c.ReleaseTime,
		ReleaseType:      c.ReleaseType,
		ReleaseAdd:       c.ReleaseAdd,
		Length:           c.Length,
		ReleaseDiscount:  c.ReleaseDiscount,
	}
}
