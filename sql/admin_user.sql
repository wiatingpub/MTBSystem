CREATE TABLE admin_user(
    au_id INT(11) NOT NULL AUTO_INCREMENT,
    admin_name CHAR(30) NOT NULL DEFAULT '' COMMENT '登陆者的名字',
    admin_password CHAR(20) NOT NULL DEFAULT '' COMMENT '登录密码',
    admin_cinema_id INT(11) NOT NULL DEFAULT 0 COMMENT 'admin的时候为-1，初始值为0',
    admin_last_login_time CHAR(20) NOT NULL DEFAULT '' COMMENT '上次登录时间',
    admin_num TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限,0为默认值，1为总管',
    PRIMARY KEY(au_id)
)ENGINE=InnoDB DEFAULT CHARSET = utf8;