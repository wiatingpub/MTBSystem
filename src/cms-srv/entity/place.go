package entity

import "share/pb"

type Place struct {
	Id          int64  `json:"id" db:"id"`
	Count       int64  `json:"count" db:"count"`
	Name        string `json:"name" db:"name"`
	PinyinFull  string `json:"pinyin_full" db:"pinyin_full"`
	PinyinShort string `json:"pinyin_short" db:"pinyin_short"`
}

func (p Place) ToProtoDBHotPlayMovies() *pb.Place {
	return &pb.Place{
		Id:          p.Id,
		Count:       p.Count,
		N:           p.Name,
		PinyinFull:  p.PinyinFull,
		PinyinShort: p.PinyinShort,
	}
}

func (p Place) ToProtoPlaceAll() *pb.PlaceAll {
	return &pb.PlaceAll{
		Id:          p.Id,
		Count:       p.Count,
		Name:        p.Name,
		PinyinFull:  p.PinyinFull,
		PinyinShort: p.PinyinShort,
	}
}
