CREATE TABLE movie_hall (
  mh_id INT(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  mh_name CHAR(20) NOT NULL DEFAULT '' COMMENT '即几号厅',
  mh_address TEXT NOT NULL DEFAULT '' COMMENT '座位表，json表示，{"x":5,"y":6,"no":["xnoy"]}，其中x表示列，y表示行,表示不允许坐的位置',
  cinema_id INT(11) NOT NULL DEFAULT 0,
  PRIMARY KEY(mh_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
