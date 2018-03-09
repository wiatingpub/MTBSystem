package entity

type Comment struct {
	CommentId int64  `json:"comment_id" db:"comment_id"`
	FilmId    int64  `json:"film_id" db:"film_id"`
	Title     string `json:"title" db:"title"`
	Content   string `json:"content" db:"content"`
	HeadImg   string `json:"head_img" db:"head_img"`
	NickName  string `json:"nick_name" db:"nick_name"`
	UpNum     int64  `json:"up_num" db:"up_num"`
	CreateAt  string `json:"create_at" db:"create_at"`
	UserId    int64  `json:"user_id" db:"user_id"`
}
