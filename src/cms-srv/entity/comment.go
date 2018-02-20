package entity

import "share/pb"

type Comment struct {
	CommentId int64  `json:"comment_id" db:"comment_id"`
	FilmId    int64  `json:"film_id" db:"film_id"`
	Title     string `json:"title" db:"title"`
	Content   string `json:"content" db:"content"`
	HeadImg   string `json:"head_img" db:"head_img"`
	NickName  string `json:"nick_name" db:"nick_name"`
	CreateAt  string `json:"create_at" db:"create_at"`
	UpNum     int64  `json:"up_num" db:"up_num"`
}

func (c Comment) ToProtoComment() *pb.Comment {
	return &pb.Comment{
		CommentID: c.CommentId,
		FilmID:    c.FilmId,
		Title:     c.Title,
		Content:   c.Content,
		HeadImg:   c.HeadImg,
		NickName:  c.NickName,
		CreateAt:  c.CreateAt,
		UpNum:     c.UpNum,
	}
}
