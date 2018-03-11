package db

import (
	"cms-srv/entity"
	"database/sql"
)

func SelectAllPlace(page int64, num int64) ([]*entity.Place, error) {

	places := []*entity.Place{}
	err := db.Select(&places, "SELECT * FROM `place`  ORDER BY `id` DESC LIMIT ?,?", (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return places, nil
}

func SelectPlaceTotal() (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `place`")
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, nil
}

func AddPlace(place *entity.Place) error {

	_, err := db.Exec("INSERT INTO  `place`(`name`,`pinyin_full`,`pinyin_short`) VALUES(?,?,?)",
		place.Name, place.PinyinFull, place.PinyinShort)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdatePlace(place *entity.Place) error {
	_, err := db.Exec("UPDATE `place` SET `name` = ?,`pinyin_full` = ?,`pinyin_short` = ? WHERE `id` = ?", place.Name, place.PinyinFull, place.PinyinShort, place.Id)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func DeletePlace(placeId int64) error {
	_, err := db.Exec("DELETE FROM `place` WHERE `id` = ?", placeId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
