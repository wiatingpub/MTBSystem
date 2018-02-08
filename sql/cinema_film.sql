CREATE TABLE cinema_film (
    cf_id INT(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    cinema_id INT(11) NOT NULL DEFAULT 0,
    film_id INT(11) NOT NULL DEFAULT 0,
    hall_id INT(11) NOT NULL DEFAULT 0 COMMENT '哪个场，即几号厅',
    film_name CHAR(100) NOT NULL DEFAULT '' COMMENT '影片名字',
    cinema_name CHAR(100) NOT NULL DEFAULT '' COMMENT '影院名字',
    release_time_year INT(11) NOT NULL DEFAULT 0,
    release_time_month INT(11) NOT NULL DEFAULT 0,
    release_time_day INT(11) NOT NULL DEFAULT 0,
    release_time CHAR(20) NOT NULL DEFAULT 0,
    release_type CHAR(20) NOT NULL DEFAULT '' COMMENT '2D什么的',
    release_add CHAR(20) NOT NULL DEFAULT '',
    length SMALLINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '影片时长',
    release_discount FLOAT NOT NULL DEFAULT 0.0,
    PRIMARY KEY(cf_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
