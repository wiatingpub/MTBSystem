package entity

type User struct {
	UserId int64 `json:"user_id" db:"user_id"`
	UserName string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
	CreateAt string `json:"create_at" db:"create_at"`
	Email	 string `json:"email" db:"email"`
	Phone int64 `json:"phone" db:"phone"`
}