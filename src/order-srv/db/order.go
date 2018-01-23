package db

import "time"

func InsertOrder(orderPrice float32,mhId int64,userId int64,movieId int64) error {
	today := time.Now().Format("2006-01-02")
	orderNum := time.Now().Unix()
	_,err := db.Exec("INSERT INTO `order`(`order_num`,`order_price`,`create_at`,`mh_id`,`user_id`,`movie_id`) VALUES(?,?,?,?,?,?) ",orderNum,orderPrice,today,mhId,userId,movieId)
	return err
}

func UpdateOrderStatus(orderNum int64) error {

	_,err := db.Exec("UPDATE `order` SET `order_status` = 1 WHERE `order_num` = ?",orderNum)
	return err;
}
