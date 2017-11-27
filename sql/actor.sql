CREATE TABLE actor (
  id INT(11) AUTO_INCREMENT COMMENT '演员编号',
  actor_name char(50) COMMENT '演员名称',
  actor_photo char(50) COMMENT '演员头像',
  actor_country char(50) COMMENT '演员所属国家',
  PRIMARY KEY ( `id` )
) ENGINE = InnoDB DEFAULT CHARSET =utf8;