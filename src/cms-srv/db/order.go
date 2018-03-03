package db

import (
	"cms-srv/entity"
	"database/sql"
)

func SelectAllOrder(page int64, num int64) ([]*entity.Order, error) {

	orders := []*entity.Order{}
	err := db.Select(&orders, "SELECT * FROM `film_order`  ORDER BY `order_id` DESC LIMIT ?,?", (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return orders, nil
}

func SelectOrderTotal() (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `film_order`")
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, nil
}

func SelectOrderByFilmId(page int64, num int64, mhId int64) ([]*entity.Order, error) {

	orders := []*entity.Order{}
	err := db.Select(&orders, "SELECT * FROM `film_order`  WHERE `mh_id` = ? ORDER BY `order_id` DESC LIMIT ?,?", mhId, (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return orders, nil
}

func SelectOrderTotalByFilmId(filmId int64) (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `film_order` WHERE `movie_id` = ?", filmId)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, nil
}
