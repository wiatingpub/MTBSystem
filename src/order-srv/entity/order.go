package entity

type Order struct {
	OrderId     int64   `json:"order_id" db:"order_id"`
	OrderNum    string  `json:"order_num" db:"order_num"`
	OrderStatus int64   `json:"order_status" db:"order_status"`
	OrderPrice  float32 `json:"order_price" db:"order_price"`
	CreateAt    string  `json:"create_at" db:"create_at"`
	PayAt       string  `json:"pay_at" db:"pay_at"`
	MhId        int64   `json:"mh_id" db:"mh_id"`
	OrderX      int64   `json:"order_x" db:"order_x"`
	OrderY      int64   `json:"order_y" db:"order_y"`
	UserId      int64   `json:"user_id" db:"user_id"`
	MovieId     int64   `json:"movie_id" db:"movie_id"`
	OrderScore  int64   `json:"order_score" db:"order_score"`
	StartTime   string  `json:"start_time" db:"start_time"`
	EndTime     string  `json:"end_time" db:"end_time"`
}
