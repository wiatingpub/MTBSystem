package entity

import (
	"share/pb"
)

type User struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	UserName string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
	CreateAt string `json:"create_at" db:"create_at"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
}

func (u User) ToProtoUser() *pb.User {
	return &pb.User{
		UserId:   u.UserId,
		UserName: u.UserName,
		Password: u.Password,
		CreateAt: u.CreateAt,
		Email:    u.Email,
		Phone:    u.Phone,
	}
}
