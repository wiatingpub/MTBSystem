package db

import (
	"cms-srv/entity"
	"database/sql"
)

func SelectAdmin(user string, password string) (*entity.Admin, error) {

	admin := entity.Admin{}
	err := db.Get(&admin, "SELECT `au_id`,`admin_cinema_id`,`admin_last_login_time`,`admin_num` FROM `admin_user` WHERE `admin_name` = ? AND `admin_password` = ? LIMIT 1", user, password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &admin, err
}

func SelectAdminByAUID(auID int64) (*entity.Admin, error) {

	admin := entity.Admin{}
	err := db.Get(&admin, "SELECT `admin_num`,`admin_cinema_id`,`au_id` FROM `admin_user` WHERE `au_id` = ? LIMIT 1", auID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &admin, err
}

func SelectAllAdmin(page int64, num int64) ([]*entity.Admin, error) {

	admins := []*entity.Admin{}
	err := db.Select(&admins, "SELECT * FROM `admin_user`  ORDER BY `au_id` DESC LIMIT ?,?", (page-1)*num, page*num)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return admins, err
}

func SelectAdminTotal() (int64, error) {

	var total int64
	err := db.Get(&total, "SELECT count(*) FROM `admin_user`")
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return total, err
}

func AddAdminUser(admin *entity.Admin) error {

	_, err := db.Exec("INSERT INTO  `admin_user`(`admin_name`,`admin_password`,`admin_cinema_id`,`admin_num`) VALUES(?,?,?,?)",
		admin.AdminName, admin.AdminPassword, admin.CinemaID, admin.AdminNum)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func DeleteAdminUser(auID int64) error {
	_, err := db.Exec("DELETE FROM `admin_user` WHERE `au_id` = ?", auID)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateAdminUser(auID int64, adminName string, adminCinemaID int64) error {

	_, err := db.Exec("UPDATE `admin_user` SET `admin_name` = ?,`admin_cinema_id` = ? WHERE `au_id` = ?",
		adminName, adminCinemaID, auID)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
