SET NAMES utf8mb4;

DELETE FROM `place`;

INSERT INTO place(
    `id`,
    `count`,
    `name`,
    `pinyin_full`,
    `pinyin_short`
)VALUES (
    292,
    287,
    '上海',
    'Shanghai',
    'sh'
),(
    222,
    366,
    '深圳',
    'Shenzhen',
    'sz'
),(
    291,
    217,
    '重庆',
    'Chongqing',
    'cq'
)