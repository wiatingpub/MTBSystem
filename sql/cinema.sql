CREATE TABLE cinema (
    cinema_id INT(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    cinema_name CHAR(100) NOT NULL DEFAULT '',
    cinema_add CHAR(150) NOT NULL DEFAULT '',
    location_id INT(11) NOT NULL DEFAULT 0,
    cinema_types CHAR(100) NOT NULL DEFAULT '',
    PRIMARY KEY(cinema_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
