package db

import (
	"cms-srv/entity"
	"database/sql"
	"time"
)

func SelectAllFilms(page int64, num int64) ([]*entity.Film, error) {

	films := []*entity.Film{}
	err := db.Select(&films, "SELECT * FROM `film`  ORDER BY `movie_id` DESC LIMIT ?,?", (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return films, nil
}

func SelectFilmsTotal() (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `film`")
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, nil
}

func SelectFilmsID(cinemaID int64) ([]int64, error) {

	var filmIDs []int64
	err := db.Get(&filmIDs, "SELECT `film_id` FROM `cinema_film` WHERE `cinema_id` = ?", cinemaID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return filmIDs, nil
}

func InsertFilm(film *entity.Film) (int64, error) {

	today := time.Now().Format("2006-01-02")
	var id int64
	_, err := db.Exec("INSERT INTO `film`(`img`,`length`,`film_price`,`film_screenwriter`,`film_director`,`title_cn`,`title_en`,`create_at`,`type`,`film_drama`,`common_special`,`company_issued`,`country`,`is_3D`,`is_DMAX`,`is_IMAX`,`is_IMAX3D`,`r_day`,`r_month`,`r_year`) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		film.Img, film.Length, film.FilmPrice, film.FilmScreenwriter, film.FilmDirector, film.TitleCn, film.TitleEn, today, film.Type, film.FilmDrama, film.CommonSpecial, film.CompanyIssued, film.Country, film.Is3D, film.IsDMAX, film.IsIMAX, film.IsIMAX3D, film.RDay, film.RMonth, film.RYear)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err == nil {
		err = db.Get(&id, "SELECT `movie_id` FROM `film` WHERE `title_cn`= ? LIMIT 1", film.TitleCn)
		if err == sql.ErrNoRows {
			return 0, nil
		}
	}
	return id, err
}

func UpdateFilm(film *entity.Film) error {
	_, err := db.Exec("UPDATE `film` SET `img` = ?,`length` = ?,`film_price`=?,`film_director`=?,`title_cn`=?,`title_en`=?,`type`=?,`film_drama`=?,`common_special`=?,`company_issued`=?,`country`=?,`r_year`=?,`r_month`=?,`r_day`=?,`is_ticking` = ? where `movie_id`=?",
		film.Img, film.Length, film.FilmPrice, film.FilmDirector, film.TitleCn, film.TitleEn, film.Type, film.FilmDrama, film.CommonSpecial, film.CompanyIssued, film.Country, film.RYear, film.RMonth, film.RDay, film.IsTicking, film.MovieId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err

}

func DeleteFilm(movieId int64) error {
	_, err := db.Exec("DELETE FROM `film` WHERE `movie_id` = ?", movieId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
