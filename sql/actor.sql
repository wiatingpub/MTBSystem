CREATE TABLE actor (
  id INT(11) AUTO_INCREMENT COMMENT '演员编号',
  name_cn char(50) NOT NULL DEFAULT '' COMMENT '演员名称' ,
  name_en char(50) NOT NULL DEFAULT '' COMMENT '演员名称'  ,
  actor_photo char(255) NOT NULL DEFAULT '' COMMENT '演员头像'  ,
  actor_country char(50) NOT NULL   DEFAULT '' COMMENT '演员所属国家' ,
  actor_type INT(11) NOT NULL DEFAULT 1 COMMENT '演员级别，默认1演员，2是导演',
  PRIMARY KEY ( `id` )
) ENGINE = InnoDB DEFAULT CHARSET =utf8;
