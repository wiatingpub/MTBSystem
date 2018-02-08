CREATE TABLE film_order (
    order_id INT(11) AUTO_INCREMENT NOT NULL,
    order_num VARCHAR(50) NOT NULL DEFAULT '' COMMENT '订单编号',
    order_status INT(11) NOT NULL DEFAULT 0 COMMENT '0:下单未支付，1：下单支付，2：退单',
    order_price FLOAT NOT NULL DEFAULT 0.0,
    create_at CHAR(20) NOT NULL DEFAULT '0',
    pay_at CHAR(20) NOT NULL DEFAULT '0',
    mh_id INT(11) NOT NULL DEFAULT 0 ,
    order_x INT(11) NOT NULL DEFAULT 0 COMMENT '第几列',
    order_y INT(11) NOT NULL DEFAULT 0 COMMENT '第几行',
    user_id INT(11) NOT NULL DEFAULT 0 ,
    movie_id INT(11) NOT NULL DEFAULT 0 ,
    order_score INT(11) NOT NULL DEFAULT -1,
    start_time CHAR(20) NOT NULL DEFAULT '' COMMENT '格式如：2017-07-15 20:05',
    end_time CHAR(20) NOT NULL DEFAULT '' COMMENT '格式如：2017-07-15 20:05',
    PRIMARY KEY (order_id)
) ENGINE = InnoDB DEFAULT CHARSET =utf8;