package entity

import "share/pb"

type Film struct {
	FilmId  int64 		`json:"film_id" db:"film_id"`
	ActorName []string `json:"actor_name" db:"actor_name"`
	Logo	string		`json:"logo" db:"logo"`
	FilmLong	int64	`json:"film_long" db:"film_long"`
	IsSelectSeat	int64	`json:"is_select_seat" db:"is_select_seat"`
	FilmPrice	float32		`json:"film_price" db:"film_price"`
	FilmScreenwriter	string	`json:"film_screenwriter" db:"film_screenwriter"`
	CommentNum	int64	`json:"comment_num" db:"comment_num"`
	FilmName	string	`json:"film_name" db:"film_name"`
	IsSupportInlineWatch	int64	`json:"is_support_inline_watch" db:"is_support_inline_watch"`
	CreateAt	string	`json:"create_at" db:"create_at"`
	BigType	string	`json:"big_type" db:"big_type"`
	FilmDrama	string	`json:"film_drama" db:"film_drama"`
	CommonSpecial	string	`json:"common_special" db:"common_special"`
	UserAccessTimes	int64	`json:"user_access_times" db:"user_access_times"`
	FilmBoxoffice	float32	`json:"film_boxoffice" db:"film_boxoffice"`
	FilmDirector	string	`json:"film_director" db:"film_director"`
	UserLikeWatchTimes	int64	`json:"user_like_watch_times" db:"user_like_watch_times"`
	UserCommentTimes	int64	`json:"user_comment_times" db:"user_comment_times"`
	CompanyIssued	string	`json:"company_issued" db:"company_issued"`
	Country	string	`json:"country" db:"country"`
	CommentScore	float32	`comment_score:"logo" db:"comment_score"`
	Is3D	int64`	json:"is_3D" db:"is_3D"`
	IsDMAX	int64`	json:"is_DMAX" db:"is_DMAX"`
	IsFilter	int64	`json:"is_filter" db:"is_filter"`
	IsHot	int64	`json:"is_hot" db:"is_hot"`
	IsIMAX	int64	`json:"is_IMAX" db:"is_IMAX"`
	IsIMAX3D	int64	`json:"is_IMAX3D" db:"is_IMAX3Ds"`
	IsNew	int64	`json:"is_new" db:"is_new"`
	IsTicking	int64	`json:"is_ticking" db:"is_ticking"`
	RDay	int64	`json:"r_day" db:"r_day"`
	RMonth	int64	`json:"r_month" db:"r_month"`
	RYear	int64	`json:"r_year" db:"r_year"`
}

func (f Film)ToProtoDBHotPlayMovies() *pb.HotMovie {
	return & pb.HotMovie{
		ActorName:f.ActorName,
		CommonSpecial:f.CommonSpecial,
		DirectorName:f.FilmDirector,
		Img:f.Logo,
		Is3D:f.Is3D,
		IsDMAX:f.IsDMAX,
		IsFilter:f.IsFilter,
		IsHot:f.IsHot,
		IsIMAX:f.IsIMAX,
		IsIMAX3D:f.IsIMAX3D,
		IsNew:f.IsNew,
		Length:f.FilmLong,
		MovieId:f.FilmId,
		RDay:f.RDay,
		RMonth:f.RMonth,
		RYear:f.RYear,
		RatingFinal:f.CommentScore,
		Type:f.BigType,
		WantedCount:f.UserLikeWatchTimes,
	}
}