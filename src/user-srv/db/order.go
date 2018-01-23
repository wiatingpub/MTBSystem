package db

func UpdateOrderScore(order_num string,score int64) error {

	_,err := db.Exec("UPDATE `order` SET `score` = ? WHERE `order_num` = ?",score,order_num)
	return err
}