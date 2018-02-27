CREATE TABLE film_actor (
  fa_id INT(11)  AUTO_INCREMENT,
  film_id INT(11) NOT NULL DEFAULT 0 COMMENT '影片编号',
  film_name CHAR(50) NOT NULL DEFAULT '' COMMENT '影片名称',
  actor_id INT(11) NOT NULL  DEFAULT 0 COMMENT '演员编号',
  actor_name CHAR(50) NOT NULL  DEFAULT '' COMMENT '影片名称',
  PRIMARY KEY (fa_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
