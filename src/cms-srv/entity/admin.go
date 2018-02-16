package entity

type Admin struct {
	AuID               int64  `json:"au_id" db:"au_id"`
	AdminName          int64  `json:"admin_name" db:"admin_name"`
	AdminPassword      string `json:"admin_password" db:"admin_password"`
	CinemaID           int64  `json:"admin_cinema_id" db:"admin_cinema_id"`
	AdminLastLoginTime string `json:"admin_last_login_time" db:"admin_last_login_time"`
	AdminNum           int64  `json:"admin_num" db:"admin_num"`
}
