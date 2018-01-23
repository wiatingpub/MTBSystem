CREATE TABLE cinema_film (
    cf_id INT(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    cinema_id INT(11) NOT NULL DEFAULT 0,
    film_id INT(11) NOT NULL DEFAULT 0,
    release_time_year INT(11) NOT NULL DEFAULT 0,
    release_time_month INT(11) NOT NULL DEFAULT 0,
    release_time CHAR(20) NOT NULL DEFAULT 0,
    release_type CHAR(20) NOT NULL DEFAULT '',
    release_add CHAR(20) NOT NULL DEFAULT '',
    release_discount FLOAT NOT NULL DEFAULT 0.0,
    PRIMARY KEY(cf_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
