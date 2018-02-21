package entity

import "share/pb"

type MovieHall struct {
	MhID      int64  `json:"mh_id" db:"mh_id"`
	MhName    string `json:"mh_name" db:"mh_name"`
	MhAddress string `json:"mh_address" db:"mh_address"`
	CinemaId  int64  `json:"cinema_id" db:"cinema_id"`
}

func (mh MovieHall) ToProtoMovieHall() *pb.MovieHall {
	return &pb.MovieHall{
		MhId:      mh.MhID,
		MhName:    mh.MhName,
		MhAddress: mh.MhAddress,
		CinemaId:  mh.CinemaId,
	}
}
