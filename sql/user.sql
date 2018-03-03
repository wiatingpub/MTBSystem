CREATE TABLE user (
  user_id INT(11) AUTO_INCREMENT NOT NULL COMMENT '类型id',
  user_name char(50) NOT NULL DEFAULT '' COMMENT '用户名称',
  password char(50) NOT NULL DEFAULT '' COMMENT '用户的密码',
  create_at char(50) NOT NULL DEFAULT '' COMMENT '用户的注册时间',
  email char(50) NOT NULL DEFAULT '' COMMENT '用户的email',
  phone char(12) NOT NULL DEFAULT 0 COMMENT '用户联系方式',
  PRIMARY KEY (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET =utf8;