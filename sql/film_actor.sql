CREATE TABLE film_actor (
  film_id INT(11) NOT NULL COMMENT '影片编号',
  film_name CHAR(50) NOT NULL COMMENT '影片名称',
  actor_id INT(11) NOT NULL COMMENT '演员编号',
  actor_name CHAR(50) NOT NULL COMMENT '影片名称'
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
