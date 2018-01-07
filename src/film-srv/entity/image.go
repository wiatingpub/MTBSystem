package entity

type Image struct {
	ImageId  int64  `json:"image_id" db:"image_id"`
	MovieId  int64 `json:"movie_id" db:"movie_id"`
	ImageUrl string  `json:"image_url" db:"image_url"`
}
