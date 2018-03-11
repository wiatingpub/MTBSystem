CREATE TABLE want_see_record (

  ws_id INT(11) AUTO_INCREMENT NOT NULL COMMENT 'id',
  movie_id INT(11) NOT NULL DEFAULT 0 COMMENT '影片的id',
  user_id INT(11) NOT NULL DEFAULT 0 COMMENT '影片的id',
  create_at char(50) NOT NULL DEFAULT '' COMMENT '记录生成的时间',
  PRIMARY KEY (`ws_id`)
) ENGINE = InnoDB DEFAULT CHARSET =utf8;