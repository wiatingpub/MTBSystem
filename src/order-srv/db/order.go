package db

import (
	"time"
	"order-srv/entity"
	"database/sql"
)

func InsertOrder(orderPrice float32,mhId int64,userId int64,movieId int64) error {
	today := time.Now().Format("2006-01-02")
	orderNum := time.Now().Unix()
	_,err := db.Exec("INSERT INTO `film_order`(`order_num`,`order_price`,`create_at`,`mh_id`,`user_id`,`movie_id`) VALUES(?,?,?,?,?,?) ",orderNum,orderPrice,today,mhId,userId,movieId)
	return err
}

func UpdateOrderStatus(orderNum string,userId int64) error {

	_,err := db.Exec("UPDATE `film_order` SET `order_status` = 1 WHERE `order_num` = ? AND `user_id` = ?",orderNum,userId)
	return err;
}

func UpdateOrderScore(orderScore int64,userId int64,orderNum string) error {

	_,err := db.Exec("UPDATE `film_order` SET `order_score` = ? WHERE `userId` = ? AND `order_num` = ?",orderScore,userId,orderNum)
	return err;
}

func SelectOrderNumMovieIdMHIdStartTime(userId int64) ([]*entity.Order,error){

	orders := []*entity.Order{}
	err := db.Select(&orders,"SELECT `start_time`,`order_num`,`movie_id`,`mh_id` FROM `film_order` WHERE `user_id` = ? ORDER BY order_id DESC",userId)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return orders,err
}

func SelectLookAlreadyOrders(userId int64) ([]*entity.Order,error){

	orders := []*entity.Order{}
	err := db.Select(&orders,"SELECT `order_score`,`movie_id`,`order_num` FROM `film_order` WHERE `user_id` = ? ORDER BY order_id DESC",userId)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return orders,err

}

func SelectOrderMessage(orderNum string,userId int64) (*entity.Order,error){

	order := entity.Order{}
	err := db.Get(&order,"SELECT `start_time`,`end_time`,`mh_id`,`movie_id`,`order_x`,`order_x`,`create_at`,`order_price` FROM `film_order` WHERE `order_num` = ? AND `user_id` = ?",orderNum,userId)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return &order,err
}

func SelectOrderScore(orderNum string,userId int64) (*entity.Order,error){

	order := entity.Order{}
	err := db.Get(&order,"SELECT `order_score` FROM `film_order` WHERE `order_num` = ? AND `user_id` = ?",orderNum,userId)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return &order,err
}