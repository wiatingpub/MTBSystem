package entity

type Film struct {
	filmId  int64 		`json:"film_id" db:"film_id"`
	logo	string		`json:"logo" db:"logo"`
	filmLong	int64	`json:"film_long" db:"film_long"`
	isSelectSeat	int64	`json:"is_select_seat" db:"is_select_seat"`
	filmPrice	float32		`json:"film_price" db:"film_price"`
	filmScreenwriter	int64	`json:"film_screenwriter" db:"film_screenwriter"`
	commentNum	int64	`json:"comment_num" db:"comment_num"`
	filmName	string	`json:"film_name" db:"film_name"`
	isSupportInlineWatch	int64	`json:"is_support_inline_watch" db:"is_support_inline_watch"`
	createAt	string	`json:"create_at" db:"create_at"`
	bigType	int64	`json:"big_type" db:"big_type"`
	filmDrama	string	`json:"film_drama" db:"film_drama"`
	commonSpecial	string	`json:"common_special" db:"common_special"`
	userAccessTimes	int64	`json:"user_access_times" db:"user_access_times"`
	filmBoxoffice	float32	`json:"film_boxoffice" db:"film_boxoffice"`
	filmDirector	string	`json:"film_director" db:"film_director"`
	userLikeWatchTimes	int64	`json:"user_like_watch_times" db:"user_like_watch_times"`
	userCommentTimes	int64	`json:"user_comment_times" db:"user_comment_times"`
	companyIssued	string	`json:"company_issued" db:"company_issued"`
	country	string	`json:"country" db:"country"`
	commentScore	float32	`comment_score:"logo" db:"comment_score"`
	is3D	int64`	json:"is_3D" db:"is_3D"`
	isDMAX	int64`	json:"is_DMAX" db:"is_DMAX"`
	isFilter	int64	`json:"is_filter" db:"is_filter"`
	isHot	int64	`json:"is_hot" db:"is_hot"`
	isIMAX	int64	`json:"is_IMAX" db:"is_IMAX"`
	isIMAX3D	int64	`json:"is_IMAX3D" db:"is_IMAX3D"`
	isNew	int64	`json:"is_new" db:"is_new"`
	isTicking	int64	`json:"is_ticking" db:"is_ticking"`
	rDay	int64	`json:"r_day" db:"r_day"`
	rMonth	int64	`json:"r_month" db:"r_month"`
	rYear	int64	`json:"r_year" db:"r_year"`
}
