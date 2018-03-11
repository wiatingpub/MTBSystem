-- MySQL dump 10.13  Distrib 5.5.57, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: mtbsystem
-- ------------------------------------------------------
-- Server version	5.5.57-0ubuntu0.14.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `actor`
--

DROP TABLE IF EXISTS `actor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `actor` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'æ¼”å‘˜ç¼–å·',
  `name_cn` char(50) DEFAULT NULL COMMENT 'æ¼”å‘˜åç§°',
  `name_en` char(50) DEFAULT NULL COMMENT 'æ¼”å‘˜åç§°',
  `actor_photo` char(255) DEFAULT NULL COMMENT 'æ¼”å‘˜å¤´åƒ',
  `actor_country` char(50) DEFAULT NULL COMMENT 'æ¼”å‘˜æ‰€å±žå›½å®¶',
  `actor_type` int(11) NOT NULL DEFAULT '1' COMMENT 'æ¼”å‘˜çº§åˆ«ï¼Œé»˜è®¤1æ¼”å‘˜ï¼Œ2æ˜¯å¯¼æ¼”',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actor`
--

LOCK TABLES `actor` WRITE;
/*!40000 ALTER TABLE `actor` DISABLE KEYS */;
INSERT INTO `actor` VALUES (1,'瑞恩·库格勒','','https://upload.wikimedia.org/wikipedia/commons/thumb/b/b3/Ryan_Coogler_by_Gage_Skidmore.jpg/220px-Ryan_Coogler_by_Gage_Skidmore.jpg','',2),(2,'查德维克·博斯曼','','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e8/Chadwick_Boseman_by_Gage_Skidmore_July_2017_%28cropped%29.jpg/220px-Chadwick_Boseman_by_Gage_Skidmore_July_2017_%28cropped%29.jpg','',1),(3,'露皮塔·尼永奥','','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ3cAFPpINdvMRK5G_0BnG7tlpjQgZS4JC9TUgsVE4Yjqr2hZgf_Q','',1);
/*!40000 ALTER TABLE `actor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_user`
--

DROP TABLE IF EXISTS `admin_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin_user` (
  `au_id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_name` char(30) NOT NULL DEFAULT '' COMMENT 'ç™»é™†è€…çš„åå­—',
  `admin_password` char(20) NOT NULL DEFAULT '' COMMENT 'ç™»å½•å¯†ç ',
  `admin_cinema_id` int(11) NOT NULL DEFAULT '0',
  `admin_last_login_time` char(20) NOT NULL DEFAULT '' COMMENT 'ä¸Šæ¬¡ç™»å½•æ—¶é—´',
  `admin_num` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT 'æƒé™,0ä¸ºé»˜è®¤å€¼ï¼Œ1ä¸ºæ€»ç®¡',
  PRIMARY KEY (`au_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_user`
--

LOCK TABLES `admin_user` WRITE;
/*!40000 ALTER TABLE `admin_user` DISABLE KEYS */;
INSERT INTO `admin_user` VALUES (1,'admin','123456',-1,'',1),(7,'新光影城','xgyc',6,'',0);
/*!40000 ALTER TABLE `admin_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cinema`
--

DROP TABLE IF EXISTS `cinema`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cinema` (
  `cinema_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'è‡ªå¢žid',
  `cinema_name` char(100) NOT NULL DEFAULT '',
  `cinema_add` char(150) NOT NULL DEFAULT '',
  `location_id` int(11) NOT NULL DEFAULT '0' COMMENT 'å½±é™¢åŸŽå¸‚å¯¹åº”çš„ä½ç½®',
  `cinema_types` char(100) NOT NULL DEFAULT '',
  `cinema_card` int(11) NOT NULL DEFAULT '0' COMMENT 'å½±åŸŽå¡',
  `cinema_min_price` int(11) NOT NULL DEFAULT '0' COMMENT 'å‡ å…ƒèµ·',
  `cinema_support` char(200) NOT NULL DEFAULT '' COMMENT 'å½±é™¢æä¾›çš„æ”¯æŒï¼ŒåŒ…æ‹¬é€€ç­¾ç­‰,ç”¨|éš”å¼€',
  `cinema_discount` int(11) NOT NULL DEFAULT '0' COMMENT 'å½±åŸŽå¡æœ€ä½Žå‡ä»·å¤šå°‘å…ƒ',
  `cinema_phone` int(11) NOT NULL DEFAULT '0' COMMENT 'å½±é™¢ç”µè¯',
  PRIMARY KEY (`cinema_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cinema`
--

LOCK TABLES `cinema` WRITE;
/*!40000 ALTER TABLE `cinema` DISABLE KEYS */;
INSERT INTO `cinema` VALUES (6,'新光影城','广州从化广州大学华软软件学院',1,'',1,8,'退签|饮料',8,0);
/*!40000 ALTER TABLE `cinema` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cinema_film`
--

DROP TABLE IF EXISTS `cinema_film`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cinema_film` (
  `cf_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'è‡ªå¢žid',
  `cinema_id` int(11) NOT NULL DEFAULT '0',
  `film_id` int(11) NOT NULL DEFAULT '0',
  `hall_id` int(11) NOT NULL DEFAULT '0' COMMENT 'å“ªä¸ªåœºï¼Œå³å‡ å·åŽ…',
  `film_name` char(100) NOT NULL DEFAULT '' COMMENT 'å½±ç‰‡åå­—',
  `cinema_name` char(100) NOT NULL DEFAULT '' COMMENT 'å½±é™¢åå­—',
  `release_time_year` int(11) NOT NULL DEFAULT '0',
  `release_time_month` int(11) NOT NULL DEFAULT '0',
  `release_time_day` int(11) NOT NULL DEFAULT '0',
  `release_time` char(20) NOT NULL DEFAULT '0',
  `release_type` char(20) NOT NULL DEFAULT '' COMMENT '2Dä»€ä¹ˆçš„',
  `release_add` char(20) NOT NULL DEFAULT '',
  `length` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT 'å½±ç‰‡æ—¶é•¿',
  `release_discount` float NOT NULL DEFAULT '0',
  PRIMARY KEY (`cf_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cinema_film`
--

LOCK TABLES `cinema_film` WRITE;
/*!40000 ALTER TABLE `cinema_film` DISABLE KEYS */;
INSERT INTO `cinema_film` VALUES (3,3,1,1,'西游伏妖篇','新光影城',2018,2,25,'11:37','奇幻 / 动作 / 喜剧','',108,55),(4,3,1,1,'西游伏妖篇','新光影城',2018,2,25,'11:45','奇幻 / 动作 / 喜剧','',108,55),(5,6,15,2,'黑豹','新光影城',2018,2,25,'21:37','动作 / 冒险 / 科幻','',134,34.5);
/*!40000 ALTER TABLE `cinema_film` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comment` (
  `comment_id` int(11) NOT NULL AUTO_INCREMENT,
  `film_id` int(11) NOT NULL DEFAULT '0',
  `title` varchar(150) NOT NULL DEFAULT '' COMMENT 'æ ‡é¢˜',
  `content` text NOT NULL COMMENT 'å†…å®¹',
  `head_img` varchar(100) NOT NULL DEFAULT '',
  `nick_name` char(50) NOT NULL DEFAULT '',
  `create_at` char(20) NOT NULL DEFAULT '',
  `up_num` smallint(5) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `film`
--

DROP TABLE IF EXISTS `film`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `film` (
  `movie_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'å½±ç‰‡ç¼–å·',
  `img` char(255) DEFAULT '' COMMENT 'å½±ç‰‡logo',
  `length` int(11) DEFAULT '0' COMMENT 'å½±ç‰‡æ—¶é•¿',
  `is_select_seat` int(11) DEFAULT '0' COMMENT 'æ˜¯å¦æ”¯æŒé€‰åº§',
  `film_price` float DEFAULT '0' COMMENT 'å½±ç‰‡ä»·æ ¼',
  `film_screenwriter` char(255) DEFAULT '' COMMENT 'å½±ç‰‡ç¼–å‰§',
  `film_director` char(255) DEFAULT '' COMMENT 'å½±ç‰‡å¯¼æ¼”',
  `comment_num` int(11) DEFAULT '0' COMMENT 'è¯„è®ºäººæ•°',
  `title_cn` char(100) DEFAULT NULL COMMENT 'å½±ç‰‡åå­—',
  `title_en` char(100) DEFAULT NULL COMMENT 'å½±ç‰‡åå­—',
  `is_support_inline_watch` int(11) DEFAULT '0' COMMENT 'æ˜¯å¦æ”¯æŒçº¿ä¸Šè§‚çœ‹',
  `create_at` char(20) DEFAULT '' COMMENT 'è®°å½•åˆ›å»ºæ—¶é—´',
  `type` char(100) DEFAULT '' COMMENT 'å½±ç‰‡ç§ç±»',
  `film_drama` text COMMENT 'å½±ç‰‡å‰§æƒ…',
  `common_special` char(255) DEFAULT '',
  `user_access_times` int(11) DEFAULT '0' COMMENT 'ç”¨æˆ·è®¿é—®æ¬¡æ•°',
  `film_boxoffice` float DEFAULT '0' COMMENT 'å½±ç‰‡ç¥¨æˆ¿',
  `wanted_count` int(11) DEFAULT '0' COMMENT 'ç”¨æˆ·æƒ³çœ‹æ¬¡æ•°',
  `user_comment_times` int(11) DEFAULT '0' COMMENT 'ç”¨æˆ·è¯„åˆ†æ¬¡æ•°',
  `company_issued` char(255) DEFAULT '' COMMENT 'å‘è¡Œå…¬å¸',
  `country` char(50) DEFAULT '' COMMENT 'å‘è¡Œå›½å®¶',
  `rating_final` float DEFAULT '0' COMMENT 'è¯„åˆ†',
  `is_3D` int(11) DEFAULT '0' COMMENT 'æ˜¯å¦3d',
  `is_DMAX` int(11) DEFAULT '0' COMMENT 'æ˜¯å¦is_DMAX',
  `is_filter` int(11) DEFAULT '0' COMMENT 'æ˜¯å¦is_filter',
  `is_hot` int(11) DEFAULT '0' COMMENT 'æ˜¯å¦çƒ­æ’­',
  `is_IMAX` int(11) DEFAULT '0' COMMENT 'æ˜¯å¦is_IMAX',
  `is_IMAX3D` int(11) DEFAULT '0' COMMENT 'æ˜¯å¦is_IMAX3D',
  `is_new` int(11) DEFAULT '0' COMMENT 'æ˜¯å¦æ–°ç‰‡',
  `is_ticking` int(11) DEFAULT '0' COMMENT 'æ˜¯å¦æ˜¯æ­£åœ¨é”€å”®',
  `r_day` int(11) DEFAULT '0' COMMENT 'ä¸Šæ˜ æ—¶é—´-æ—¥',
  `r_month` int(11) DEFAULT '0' COMMENT 'ä¸Šæ˜ æ—¶é—´-æœˆ',
  `r_year` int(11) DEFAULT '0' COMMENT 'ä¸Šæ˜ æ—¶é—´-å¹´',
  PRIMARY KEY (`movie_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `film`
--

LOCK TABLES `film` WRITE;
/*!40000 ALTER TABLE `film` DISABLE KEYS */;
INSERT INTO `film` VALUES (1,'http://img5.mtime.cn/mt/2017/10/23/101938.17733324_1280X720X2.jpg',108,0,35,'','徐克',0,'西游伏妖篇','Journey to the West: Demon Chapter',0,'2018-02-23','奇幻 / 动作 / 喜剧','唐僧\\重色轻友\\与悟空反目','唐僧\\重色轻友\\与悟空反目',0,0,0,0,'好莱坞影片公司','美国',0,0,0,0,0,0,0,0,2,23,2,2018),(2,'http://img5.mtime.cn/mt/2018/02/17/150049.57218452_1280X720X2.jpg',138,0,40,'','林超贤',0,'红海行动','Operation Red Sea',0,'2018-02-23','动作 / 剧情','蛟龙突击队\\演绎神兵天降','《红海行动》是一部2018年上映的軍事动作电影，根據2015年也門撤僑行動的真實事件改編。由博纳影业出品，林超贤执导，冯骥编剧，张译、海清、黄景瑜、杜江主演；张涵予、白冰、王彦霖、任达华客串演出。此片作為2018年春節賀歲電影於2018年2月16日在中國內地各大電影院上映，票房甚佳；次月1日（3月1日）於香港上映。',0,0,0,0,'未知公司','中国',0,0,0,0,0,0,0,0,1,23,2,2018),(3,'http://img5.mtime.cn/mt/2018/02/05/093619.43082530_1280X720X2.jpg',120,0,38,'','陈思诚',0,'唐人街探案2','Detective Chinatown 2',0,'2018-02-23','喜剧 / 动作 / 悬疑','王宝强刘昊然肖央“大闹”美利坚','王宝强刘昊然肖央“大闹”美利坚',0,0,0,0,'未知公司','中国',0,0,0,0,0,0,0,0,1,23,2,2018),(4,'http://img5.mtime.cn/mt/2017/01/12/181512.62044353_1280X720X2.jpg',102,0,35,'','韩寒',0,'乘风破浪','Duckweed',0,'2018-02-25','喜剧','邓超彭于晏一起街头热血','邓超彭于晏一起街头热血',0,0,0,0,'天谕娱乐','中国',0,0,0,0,0,0,0,0,1,25,2,2018),(5,'http://img5.mtime.cn/mt/2018/01/05/175658.82351501_1280X720X2.jpg',110,0,35,'','许诚毅',0,'捉妖记2','Monster Hunt 2',0,'2018-02-25','喜剧 / 奇幻 / 动作','梁朝伟胡巴笨笨组成“作妖三宝”','梁朝伟胡巴笨笨组成“作妖三宝”',0,0,0,0,'天朝娱乐','中国',0,0,0,0,0,0,0,0,1,28,2,2018),(6,'http://img5.mtime.cn/mt/2018/01/24/192236.63399013_1280X720X2.jpg',95,0,40,'','郭德纲',0,'祖宗十九代','The Face Of My Gene',0,'2018-02-25','奇幻 / 喜剧','岳云鹏无限穿越调配基因','岳云鹏无限穿越调配基因',0,0,0,0,'天朝娱乐','中国',0,0,0,0,0,0,0,0,2,27,2,2018),(7,'http://img5.mtime.cn/mt/2018/02/01/101421.62196189_1280X720X2.jpg',90,0,40,'','丁亮',0,'熊出没·变形记','Boonie Bears: The Big Shrink',0,'2018-02-25','动画 / 喜剧 / 冒险','开启“微观世界”冒险之旅','开启“微观世界”冒险之旅',0,0,0,0,'微光科技','中国',0,0,0,0,0,0,0,0,2,27,2,2018),(8,'http://img5.mtime.cn/mt/2018/01/23/094118.33045114_1280X720X2.jpg',116,0,40,'','郑保瑞',0,'西游记女儿国','The Monkey King 3',0,'2018-02-25','喜剧 / 爱情 / 动作','西游最特别一难全新演绎','《西游记·女儿国》是星皓影业出品、郑保瑞执导的第三部西游记题材电影，2016年11月开机拍摄。[1]2018年2月15日于台湾、香港及马来西亚上映，而中国大陆则于2月16日上映。',0,0,0,0,'星皓影业','中国',0,0,0,0,0,0,0,0,1,25,2,2018),(15,'http://img5.mtime.cn/mt/2018/01/30/102123.92074166_1280X720X2.jpg',134,0,40,'','瑞恩·库格勒',0,'黑豹','Black Panther',0,'2018-02-25','动作 / 冒险 / 科幻','漫威电影宇宙系列','《黑豹》（英语：Black Panther）是一部於2018年上映的美国超级英雄电影，劇情改編自漫威漫畫旗下同名漫畫人物故事。驚奇工作室製片，華特迪士尼工作室電影負責發行，本片为漫威电影宇宙系列的第十八部电影。',0,0,0,0,'漫威电影宇宙','美国',0,0,0,0,0,0,0,0,1,25,2,2018);
/*!40000 ALTER TABLE `film` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `film_actor`
--

DROP TABLE IF EXISTS `film_actor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `film_actor` (
  `fa_id` int(11) NOT NULL AUTO_INCREMENT,
  `film_id` int(11) NOT NULL COMMENT 'å½±ç‰‡ç¼–å·',
  `film_name` char(50) NOT NULL COMMENT 'å½±ç‰‡åç§°',
  `actor_id` int(11) NOT NULL COMMENT 'æ¼”å‘˜ç¼–å·',
  `actor_name` char(50) NOT NULL COMMENT 'å½±ç‰‡åç§°',
  PRIMARY KEY (`fa_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `film_actor`
--

LOCK TABLES `film_actor` WRITE;
/*!40000 ALTER TABLE `film_actor` DISABLE KEYS */;
INSERT INTO `film_actor` VALUES (1,15,'黑豹',0,'查德维克·博斯曼'),(2,15,'黑豹',0,'露皮塔·尼永奥'),(3,15,'黑豹',1,'瑞恩·库格勒'),(4,15,'黑豹',2,'查德维克·博斯曼'),(5,15,'黑豹',3,'露皮塔·尼永奥');
/*!40000 ALTER TABLE `film_actor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `film_order`
--

DROP TABLE IF EXISTS `film_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `film_order` (
  `order_id` int(11) NOT NULL AUTO_INCREMENT,
  `order_num` varchar(50) NOT NULL DEFAULT '' COMMENT 'è®¢å•ç¼–å·',
  `order_status` int(11) NOT NULL DEFAULT '0' COMMENT '0:ä¸‹å•æœªæ”¯ä»˜ï¼Œ1ï¼šä¸‹å•æ”¯ä»˜ï¼Œ2ï¼šé€€å•',
  `order_price` float NOT NULL DEFAULT '0',
  `create_at` char(20) NOT NULL DEFAULT '0',
  `pay_at` char(20) NOT NULL DEFAULT '0',
  `mh_id` int(11) NOT NULL DEFAULT '0',
  `order_x` int(11) NOT NULL DEFAULT '0' COMMENT 'ç¬¬å‡ åˆ—',
  `order_y` int(11) NOT NULL DEFAULT '0' COMMENT 'ç¬¬å‡ è¡Œ',
  `user_id` int(11) NOT NULL DEFAULT '0',
  `movie_id` int(11) NOT NULL DEFAULT '0',
  `order_score` int(11) NOT NULL DEFAULT '-1',
  `start_time` char(20) NOT NULL DEFAULT '' COMMENT 'æ ¼å¼å¦‚ï¼š2017-07-15 20:05',
  `end_time` char(20) NOT NULL DEFAULT '' COMMENT 'æ ¼å¼å¦‚ï¼š2017-07-15 20:05',
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `film_order`
--

LOCK TABLES `film_order` WRITE;
/*!40000 ALTER TABLE `film_order` DISABLE KEYS */;
/*!40000 ALTER TABLE `film_order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `image`
--

DROP TABLE IF EXISTS `image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `image` (
  `image_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'è‡ªå¢žid',
  `movie_id` int(11) NOT NULL COMMENT 'å½±ç‰‡ç¼–å·',
  `image_url` char(100) NOT NULL COMMENT 'å½±ç‰‡å›¾ç‰‡',
  PRIMARY KEY (`image_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `image`
--

LOCK TABLES `image` WRITE;
/*!40000 ALTER TABLE `image` DISABLE KEYS */;
/*!40000 ALTER TABLE `image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `movie_hall`
--

DROP TABLE IF EXISTS `movie_hall`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `movie_hall` (
  `mh_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'è‡ªå¢žid',
  `mh_name` char(20) NOT NULL DEFAULT '' COMMENT 'å³å‡ å·åŽ…',
  `mh_address` text NOT NULL COMMENT 'åº§ä½è¡¨ï¼Œjsonè¡¨ç¤ºï¼Œ{"x":5,"y":6,"no":["xnoy"]}ï¼Œå…¶ä¸­xè¡¨ç¤ºåˆ—ï¼Œyè¡¨ç¤ºè¡Œ,è¡¨ç¤ºä¸å…è®¸åçš„ä½ç½®',
  `cinema_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`mh_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movie_hall`
--

LOCK TABLES `movie_hall` WRITE;
/*!40000 ALTER TABLE `movie_hall` DISABLE KEYS */;
INSERT INTO `movie_hall` VALUES (1,'一号厅','{\"x\":12,\"y\":12,\"no\":\"[\\\"1no1\\\",\\\"2no1\\\"]\"}',3),(2,'一号厅','{\"x\":11,\"y\":11,\"no\":\"[\\\"1no2\\\",\\\"2no3\\\"]\"}',6);
/*!40000 ALTER TABLE `movie_hall` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `place`
--

DROP TABLE IF EXISTS `place`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `place` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'åœ°ç‚¹ç¼–å·',
  `count` int(11) DEFAULT '0' COMMENT 'å½±ç‰‡ä¸ªæ•°',
  `name` char(50) DEFAULT '' COMMENT 'åœ°ç‚¹åç§°',
  `pinyin_full` char(50) DEFAULT '' COMMENT 'å…¨æ‹¼',
  `pinyin_short` char(10) DEFAULT '' COMMENT 'ç¼©å†™',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `place`
--

LOCK TABLES `place` WRITE;
/*!40000 ALTER TABLE `place` DISABLE KEYS */;
INSERT INTO `place` VALUES (1,0,'从化','conghua','ch'),(2,0,'广州','guangzhou','gz');
/*!40000 ALTER TABLE `place` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `type`
--

DROP TABLE IF EXISTS `type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `type` (
  `t_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ç±»åž‹id',
  `t_name` char(50) DEFAULT '' COMMENT 'ç±»åž‹åç§°',
  PRIMARY KEY (`t_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `type`
--

LOCK TABLES `type` WRITE;
/*!40000 ALTER TABLE `type` DISABLE KEYS */;
/*!40000 ALTER TABLE `type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ç±»åž‹id',
  `user_name` char(50) NOT NULL DEFAULT '' COMMENT 'ç”¨æˆ·åç§°',
  `password` char(50) NOT NULL DEFAULT '' COMMENT 'ç”¨æˆ·çš„å¯†ç ',
  `create_at` char(50) NOT NULL DEFAULT '' COMMENT 'ç”¨æˆ·çš„æ³¨å†Œæ—¶é—´',
  `email` char(50) NOT NULL DEFAULT '' COMMENT 'ç”¨æˆ·çš„email',
  `phone` int(11) NOT NULL DEFAULT '0' COMMENT 'ç”¨æˆ·è”ç³»æ–¹å¼',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `want_see_record`
--

DROP TABLE IF EXISTS `want_see_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `want_see_record` (
  `ws_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `movie_id` int(11) NOT NULL DEFAULT '0' COMMENT 'å½±ç‰‡çš„id',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT 'å½±ç‰‡çš„id',
  `create_at` char(50) NOT NULL DEFAULT '' COMMENT 'è®°å½•ç”Ÿæˆçš„æ—¶é—´',
  PRIMARY KEY (`ws_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `want_see_record`
--

LOCK TABLES `want_see_record` WRITE;
/*!40000 ALTER TABLE `want_see_record` DISABLE KEYS */;
/*!40000 ALTER TABLE `want_see_record` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-02-26 12:59:52
