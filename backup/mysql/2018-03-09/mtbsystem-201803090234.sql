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
  `name_cn` char(50) NOT NULL DEFAULT '',
  `name_en` char(50) NOT NULL DEFAULT '',
  `actor_photo` char(255) NOT NULL DEFAULT '',
  `actor_country` char(50) NOT NULL DEFAULT '',
  `actor_type` int(11) NOT NULL DEFAULT '1' COMMENT 'æ¼”å‘˜çº§åˆ«ï¼Œé»˜è®¤1æ¼”å‘˜ï¼Œ2æ˜¯å¯¼æ¼”',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actor`
--

LOCK TABLES `actor` WRITE;
/*!40000 ALTER TABLE `actor` DISABLE KEYS */;
INSERT INTO `actor` VALUES (1,'瑞恩·库格勒','','https://upload.wikimedia.org/wikipedia/commons/thumb/b/b3/Ryan_Coogler_by_Gage_Skidmore.jpg/220px-Ryan_Coogler_by_Gage_Skidmore.jpg','',2),(2,'查德维克·博斯曼','','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e8/Chadwick_Boseman_by_Gage_Skidmore_July_2017_%28cropped%29.jpg/220px-Chadwick_Boseman_by_Gage_Skidmore_July_2017_%28cropped%29.jpg','',1),(3,'露皮塔·尼永奥','','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ3cAFPpINdvMRK5G_0BnG7tlpjQgZS4JC9TUgsVE4Yjqr2hZgf_Q','',1),(4,'林超贤','','https://gss0.bdstatic.com/94o3dSag_xI4khGkpoWK1HF6hhy/baike/c0%3Dbaike80%2C5%2C5%2C80%2C26/sign=9df1047fac51f3ded7bfb136f5879b7a/78310a55b319ebc4b7a2ba2c8126cffc1f1716a8.jpg','',2),(5,'张译','','http://images.china.cn/attachement/jpg/site1007/20140905/001ec949fb5915738ea401.jpg','',1),(6,'黄景瑜','','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRrAMfjlwT8uPdWkowe5hD1Zg5K69DKl8yg7ub4rSddIG8NMWSE','',1),(7,'陈思诚','','https://i2.wp.com/www.herworldplus.com/sites/default/files/2805_yangmi_embed.jpg','',2),(8,'王宝强','','https://images-na.ssl-images-amazon.com/images/M/MV5BMTQ4NzU3NTY3N15BMl5BanBnXkFtZTcwMDc2MjM1OQ@@._V1_UX214_CR0,0,214,317_AL_.jpg','',1),(9,'刘昊然','','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTgW0qH9w6Zga3YdUaF_a04XlTzoG-nfLKHEAWB2HuF6hxSxwZa','',1),(10,'卫铁','','https://gss3.bdstatic.com/-Po3dSag_xI4khGkpoWK1HF6hhy/baike/c0%3Dbaike80%2C5%2C5%2C80%2C26/sign=52134e265d82b2b7b392319650c4a08a/dcc451da81cb39db45aab174d4160924ab1830ae.jpg','',2),(11,'卫铁','','https://gss3.bdstatic.com/-Po3dSag_xI4khGkpoWK1HF6hhy/baike/c0%3Dbaike80%2C5%2C5%2C80%2C26/sign=52134e265d82b2b7b392319650c4a08a/dcc451da81cb39db45aab174d4160924ab1830ae.jpg','',1),(12,'许诚毅','','https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/Raman_Hui_2008.jpg/220px-Raman_Hui_2008.jpg','',2),(13,'梁朝伟','','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSLcygj-KIMb814n-EmkG2CNVhMKM_18-tAXrFDiCjDKzLI8Q4b','',1),(14,'白百何','','http://ent.news.cn/2015-08/08/128103465_14389116478221n.jpg','',1),(15,'郭德纲','','http://www.bestchinanews.com/1jrttimg/large/ac80006f052c1f98dde','',2),(16,'岳云鹏','','http://chinachristiandaily.com/data/images/full/0/35/3598.jpg','',1),(17,'吴京','','https://upload.wikimedia.org/wikipedia/commons/2/23/Wu_Jing_%28Wolf_Warrior_2%29.jpg','',1),(18,'丁亮','','http://i.ce.cn/district/newarea/sddy/201601/19/W020160119376801678857.jpg','',2),(19,'丁亮','','http://i.ce.cn/district/newarea/sddy/201601/19/W020160119376801678857.jpg','',1),(20,'林汇达','','https://img1.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1492836254.8.webp','',1),(21,'郑保瑞','','http://www1.pictures.zimbio.com/gi/Accident+Photocall+66th+Venice+Film+Festival+groE2jNSlmFl.jpg','',2),(22,'郭富城','','https://images-na.ssl-images-amazon.com/images/M/MV5BMTkyOTQ0MTMyMV5BMl5BanBnXkFtZTcwMDM2Mzc3NA@@._V1_UX214_CR0,0,214,317_AL_.jpg','',1),(23,'冯绍峰','','http://ww4.sinaimg.cn/large/61e7f4aatw1dj4308roiqj.jpg','',1),(24,'瑞恩·库格勒','','https://upload.wikimedia.org/wikipedia/commons/d/df/Ryan_Coogler_at_Sundance_2013.jpg','',2),(25,'查德维克·博斯曼','','data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxATEBUTEhIVFRUWFhUVFRAQFRUVFRUVFRcWFxUVFRUYHSggGBolHRUVITEhJSkrLi4uFx8zODMtNygtLisBCgoKDg0OGBAQGi8eHSYtKy0tKy0rLSstLS0tKy0tLS0tLS0rLS0tLSstLS0tLS0tLS0tLSstKy0rKy03NystN//AABEIARMAtwMBIgACEQEDEQH/','',1),(26,'瑞恩·库格勒','','https://upload.wikimedia.org/wikipedia/commons/d/df/Ryan_Coogler_at_Sundance_2013.jpg','',2),(27,'查德维克·博斯曼','','data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxATEBUTEhIVFRUWFhUVFRAQFRUVFRUVFRcWFxUVFRUYHSggGBolHRUVITEhJSkrLi4uFx8zODMtNygtLisBCgoKDg0OGBAQGi8eHSYtKy0tKy0rLSstLS0tKy0tLS0tLS0rLS0tLSstLS0tLS0tLS0tLSstKy0rKy03NystN//AABEIARMAtwMBIgACEQEDEQH/','',1),(28,'瑞恩·库格勒','','https://upload.wikimedia.org/wikipedia/commons/d/df/Ryan_Coogler_at_Sundance_2013.jpg','',2),(29,'查德维克·博斯曼','','https://i.ytimg.com/vi/engQfhmeY3w/maxresdefault.jpg','',1),(30,'露皮塔·尼永奥','','data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxITEhUTEhIVFhUXFRUVFxUVFRUVFRcVFRYWFxcVFxUYHSggGBolHRUVITEhJSkrLi4uFx8zODMtNygtLisBCgoKDg0OGhAQGi0lHx0rLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLSstLS0tLS0tKy0tLS0tLS0tLf/AABEIAQgAvwMBIgACEQEDEQH/','',1),(31,'瑞恩·库格勒','','https://upload.wikimedia.org/wikipedia/commons/d/df/Ryan_Coogler_at_Sundance_2013.jpg','',2),(32,'查德维克·博斯曼','','https://i.ytimg.com/vi/engQfhmeY3w/maxresdefault.jpg','',1),(33,'露皮塔·尼永奥','','data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxITEhUTEhIVFhUXFRUVFxUVFRUVFRcVFRYWFxcVFxUYHSggGBolHRUVITEhJSkrLi4uFx8zODMtNygtLisBCgoKDg0OGhAQGi0lHx0rLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLSstLS0tLS0tKy0tLS0tLS0tLf/AABEIAQgAvwMBIgACEQEDEQH/','',1),(34,'瑞恩·库格勒','','https://upload.wikimedia.org/wikipedia/commons/d/df/Ryan_Coogler_at_Sundance_2013.jpg','',2),(35,'查德维克·博斯曼','','https://i.ytimg.com/vi/engQfhmeY3w/maxresdefault.jpg','',1),(36,'露皮塔·尼永奥','','http://image11.m1905.cn/uploadfile/2014/0113/20140113074415612349_watermark.jpg','',1),(37,'黄真真','','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQPTvzCgXf811NxBzr5bTNWTKbIt7jSwAHYrVIIeIGVRtqsqmttrA','',2),(38,'陈意涵','','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSKnY-CVovv5LOHHgAaI8RGlWQ0iBApvPDfxaucMc4fljlX_P1L','',1),(39,'薛凯琪','','http://1.bp.blogspot.com/-LocscOBFi2c/UXMOK22TflI/AAAAAAAAB24/qJ63Bwj0Lro/s1600/1+(129).jpg','',1),(40,'威尔·古勒','','https://gss3.bdstatic.com/-Po3dSag_xI4khGkpoWK1HF6hhy/baike/c0%3Dbaike80%2C5%2C5%2C80%2C26/sign=c6baf281b2fb43160e12722841cd2d46/ca1349540923dd54ad993522d109b3de9d824896.jpg','',2),(41,'詹姆斯·柯登','','https://upload.wikimedia.org/wikipedia/commons/thumb/c/ca/James_Corden_at_2015_PaleyFest.jpg/800px-James_Corden_at_2015_PaleyFest.jpg','',1),(42,'多姆纳尔·格里森','','http://img31.mtime.cn/pi/2016/05/11/162619.91319350_369X369X4.jpg','',1),(43,'保罗·维尔奇','','https://p1.meituan.net/movie/b5d2a27a9a7c0f774b0590adefe961d5159373.jpg','',2),(44,'海伦·米伦','','https://gss3.bdstatic.com/7Po3dSag_xI4khGkpoWK1HF6hhy/baike/s%3D220/sign=634b80f2eef81a4c2232ebcbe72b6029/a2cc7cd98d1001e956792689bd0e7bec55e797ea.jpg','',1),(45,'唐纳德·萨瑟兰','','http://n.sinaimg.cn/ent/transform/20170814/rP_h-fyixias0687655.jpg','',1),(46,'克里斯托弗·兰敦','','http://img5.mtime.cn/ph/2017/07/24/173208.32631800_290X440X4.jpg','',2),(47,'杰西卡·罗瑟','','http://images.movie.xunlei.kankan.com/gallery/1495/23fd6d8a731c7eaf624fcffb172e5580.jpg','',1),(48,'伊瑟尔·布罗萨德','','http://images.miiqu.com/uploads/star/p/p_b/p_1501779768.jpg','',1);
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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_user`
--

LOCK TABLES `admin_user` WRITE;
/*!40000 ALTER TABLE `admin_user` DISABLE KEYS */;
INSERT INTO `admin_user` VALUES (1,'admin','123456',-1,'',1),(7,'新光影城','xgyc',6,'',0),(12,'哈艺时代影城','hysdyc',11,'',0);
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
  `cinema_phone` char(15) NOT NULL DEFAULT '',
  PRIMARY KEY (`cinema_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cinema`
--

LOCK TABLES `cinema` WRITE;
/*!40000 ALTER TABLE `cinema` DISABLE KEYS */;
INSERT INTO `cinema` VALUES (6,'新光影城','广州从化广州大学华软软件学院',1,'',1,8,'退签|饮料',8,'0'),(11,'哈艺时代影城','从化太平大道12号',1,'',0,8,'退签|饮料',8,'18814128405');
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
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cinema_film`
--

LOCK TABLES `cinema_film` WRITE;
/*!40000 ALTER TABLE `cinema_film` DISABLE KEYS */;
INSERT INTO `cinema_film` VALUES (7,6,30,2,'爱在记忆消逝前','一号厅',2018,2,28,'23:54','冒险 / 喜剧 / 剧情','',112,35),(8,6,29,2,'比得兔','一号厅',2018,2,28,'23:54','动画 / 冒险 / 喜剧','',93,35),(9,6,28,2,'闺蜜2','一号厅',2018,2,28,'23:54','喜剧 / 剧情','',109,35),(10,6,27,2,'黑豹','一号厅',2018,2,28,'23:55','动作 / 冒险 / 科幻','',134,35),(11,6,17,2,'唐人街探案2','一号厅',2018,2,28,'23:55','喜剧 / 动作 / 悬疑','',120,35),(12,6,17,2,'唐人街探案2','一号厅',2018,2,28,'23:55','喜剧 / 动作 / 悬疑','',120,35),(13,6,31,2,'忌日快乐','一号厅',2018,2,28,'00:39','恐怖 / 悬疑 / 惊悚','',96,35),(14,6,31,2,'忌日快乐','新光影城',2018,3,1,'00:58','恐怖 / 悬疑 / 惊悚','',96,35),(15,6,31,2,'忌日快乐','新光影城',2018,3,2,'00:58','恐怖 / 悬疑 / 惊悚','',96,35),(16,6,30,2,'爱在记忆消逝前','新光影城',2018,3,1,'00:58','冒险 / 喜剧 / 剧情','',112,35),(17,6,30,2,'爱在记忆消逝前','新光影城',2018,3,2,'00:58','冒险 / 喜剧 / 剧情','',112,35),(18,6,29,2,'比得兔','新光影城',2018,2,28,'00:59','动画 / 冒险 / 喜剧','',93,35),(19,6,29,2,'比得兔','新光影城',2018,3,1,'01:00','动画 / 冒险 / 喜剧','',93,35),(20,6,27,2,'黑豹','新光影城',2018,3,1,'01:00','动作 / 冒险 / 科幻','',134,35),(21,6,27,2,'黑豹','新光影城',2018,2,28,'01:00','动作 / 冒险 / 科幻','',134,35),(22,6,26,2,'黑豹','新光影城',2018,2,28,'01:00','动作 / 冒险 / 科幻','',134,35),(23,6,26,2,'黑豹','新光影城',2018,3,1,'01:00','动作 / 冒险 / 科幻','',134,35),(24,6,25,2,'黑豹','新光影城',2018,3,1,'01:00','动作 / 冒险 / 科幻','',134,35),(25,6,24,2,'黑豹','新光影城',2018,3,1,'01:00','动作 / 冒险 / 科幻','',134,35),(26,6,23,2,'黑豹','新光影城',2018,3,1,'01:00','动作 / 冒险 / 科幻','',134,35),(27,6,22,2,'西游记女儿国','新光影城',2018,3,1,'01:00','喜剧 / 爱情 / 动作','',116,35),(28,6,21,2,'熊出没·变形记','新光影城',2018,3,1,'01:00','动画 / 喜剧 / 冒险','',90,35),(29,6,20,2,'祖宗十九代','二号厅',2018,3,9,'01:00','奇幻 / 喜剧','',95,35),(30,6,19,2,'捉妖记2','二号厅',2018,3,9,'01:00','喜剧 / 奇幻 / 动作','',100,35),(31,6,17,2,'唐人街探案2','新光影城',2018,3,1,'01:00','喜剧 / 动作 / 悬疑','',120,35),(32,6,16,2,'红海行动','新光影城',2018,3,1,'01:00','动作 / 剧情','',138,35),(33,6,16,2,'红海行动','新光影城',2018,2,28,'01:01','动作 / 剧情','',138,35),(34,6,17,2,'唐人街探案2','新光影城',2018,2,28,'01:01','喜剧 / 动作 / 悬疑','',120,35),(38,6,23,4,'黑豹','新光影城',2018,3,3,'22:36','动作 / 冒险 / 科幻','',134,30),(43,6,28,4,'闺蜜2','新光影城',2018,3,3,'22:36','喜剧 / 剧情','',109,30),(44,6,29,4,'比得兔','新光影城',2018,3,3,'22:36','动画 / 冒险 / 喜剧','',93,30),(45,6,30,4,'爱在记忆消逝前','新光影城',2018,3,3,'22:36','冒险 / 喜剧 / 剧情','',112,30),(46,6,31,4,'忌日快乐','新光影城',2018,3,3,'22:36','恐怖 / 悬疑 / 惊悚','',96,30),(47,11,16,5,'红海行动','哈艺时代影城',2018,3,3,'23:22','动作 / 剧情','',138,50),(48,11,17,5,'唐人街探案2','哈艺时代影城',2018,3,3,'23:22','喜剧 / 动作 / 悬疑','',120,50),(49,11,18,5,'厉害了，我的国','哈艺时代影城',2018,3,3,'23:22','纪录片','',90,50),(50,11,19,6,'捉妖记2','哈艺时代影城',2018,3,3,'23:22','喜剧 / 奇幻 / 动作','',90,40),(51,11,21,6,'熊出没·变形记','哈艺时代影城',2018,3,3,'23:22','动画 / 喜剧 / 冒险','',90,40),(52,11,22,6,'西游记女儿国','哈艺时代影城',2018,3,3,'23:22','喜剧 / 爱情 / 动作','',116,40),(53,11,28,6,'闺蜜2','哈艺时代影城',2018,3,3,'23:22','喜剧 / 剧情','',109,40),(54,11,29,6,'比得兔','哈艺时代影城',2018,3,3,'23:22','动画 / 冒险 / 喜剧','',93,40),(55,11,30,6,'爱在记忆消逝前','哈艺时代影城',2018,3,3,'23:22','冒险 / 喜剧 / 剧情','',112,40),(56,11,31,6,'忌日快乐','哈艺时代影城',2018,3,3,'23:22','恐怖 / 悬疑 / 惊悚','',96,40);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (4,20,'','不错不错','','18814128406','2018-03-03',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `film`
--

LOCK TABLES `film` WRITE;
/*!40000 ALTER TABLE `film` DISABLE KEYS */;
INSERT INTO `film` VALUES (16,'http://img5.mtime.cn/mt/2018/02/17/150049.57218452_1280X720X2.jpg',138,0,40,'','林超贤',0,'红海行动','Operation Red Sea',0,'2018-02-27','动作 / 剧情','\\蛟龙突击队\\演绎神兵天降','《红海行动》是一部2018年上映的軍事动作电影，根據2015年也門撤僑行動的真實事件改編。由博纳影业出品，林超贤执导，冯骥编剧，张译、海清、黄景瑜、杜江主演；张涵予、王彦霖、任达华客串演出。此片作為2018年春節賀歲電影於2018年2月16日在中國內地各大電影院上映，票房甚佳；次月1日（3月1日）於香港上映。',0,0,0,0,'博纳影业','中国',0,0,0,0,0,0,0,0,1,27,2,2018),(17,'http://img5.mtime.cn/mt/2018/02/05/093619.43082530_1280X720X2.jpg',120,0,40,'','陈思诚',0,'唐人街探案2','Detective Chinatown 2',0,'2018-02-27','喜剧 / 动作 / 悬疑','王宝强刘昊然肖央“大闹”美利坚','《唐人街探案2》是万达影视传媒有限公司、上海骋亚影视文化传媒有限公司、北京红色果实文化传播有限公司联合出品，陈思诚执导，王宝强、刘昊然领衔主演，肖央、刘承羽、尚语贤、王迅、杨金赐、王成思等主演的悬疑动作喜剧冒险系列电影，是电影《唐人街探案》的续集。',0,0,0,0,'万达影视传媒有限公司','中国',0,0,0,0,0,0,0,0,1,27,2,2018),(18,'http://img5.mtime.cn/mt/2018/02/22/142010.67656990_1280X720X2.jpg',90,0,90,'','卫铁',0,'厉害了，我的国','Amazing China',0,'2018-02-27','纪录片','厉害了，我的国','2018“厉害了我的国?改革开放40年”大型主题系列活动春节又有大动作！ 对于讲究团圆的中国人来说，全家福不仅定格家庭的美好时光，也见证着祖国的历史变迁。 春节来临，快打开手机，一起来“幸福照相馆”定制独一无二的全家福吧！ 众筹视频展播活动也已正式开启，想知道自己的作品上电视了没？',0,0,0,0,'国家影院','中国',0,0,0,0,0,0,0,0,1,27,2,2018),(19,'http://img5.mtime.cn/mt/2018/01/05/175658.82351501_1280X720X2.jpg',90,0,40,'','许诚毅',0,'捉妖记2','Monster Hunt 2',0,'2018-02-27','喜剧 / 奇幻 / 动作','梁朝伟胡巴笨笨组成“作妖三宝”','《捉妖记2》由无锡影都传媒有限公司、安乐（北京）电影发行有限公司、蓝色星空影业出品的喜剧奇幻片，作为《捉妖记》续集，由梁朝伟、白百何、井柏然、李宇春、杨祐宁等主演。该片讲述了天荫和小岚找回胡巴团聚一堂的合家欢故事，该片于2018年2月16日（大年初一）在中国大陆上映。',0,0,0,0,'无锡影都传媒有限公司','中国',2,0,0,0,0,0,0,0,1,27,2,2018),(20,'http://img5.mtime.cn/mt/2018/01/24/192236.63399013_1280X720X2.jpg',95,0,40,'','郭德纲',0,'祖宗十九代','The Face Of My Gene',0,'2018-02-27','奇幻 / 喜剧','岳云鹏无限穿越调配基因','《祖宗十九代》是北京德云社文化传播有限公司出品的喜剧片，由郭德纲执导、岳云鹏、吴京、吴秀波、井柏然、林志玲、王宝强等主演。该片讲述了立志成为作家却一事无成的青年小贝，在经历了一趟奇幻之旅后重拾自信，勇敢面对人生的故事。影片于2018年2月16日（大年初一）在中国内地上映',0,0,0,0,'北京德云社文化传播有限公司','中国',3.3125,0,0,0,0,0,0,0,1,27,2,2018),(21,'http://img5.mtime.cn/mt/2018/02/01/101421.62196189_1280X720X2.jpg',90,0,40,'','丁亮',0,'熊出没·变形记','Boonie Bears: The Big Shrink',0,'2018-02-27','动画 / 喜剧 / 冒险','开启“微观世界”冒险之旅','《熊出没·变形记》是《熊出没》系列第五部大电影，由丁亮、林汇达执导，于2018年2月16日在中国内地上映。影片讲述了光头强、熊大、熊二遭遇变形缩小后所展开的一段微观世界的冒险故事。',0,0,0,0,'万达国际影院','中国',0,0,0,0,0,0,0,0,2,27,2,2018),(22,'http://img5.mtime.cn/mt/2018/01/23/094118.33045114_1280X720X2.jpg',116,0,40,'','郑保瑞',0,'西游记女儿国','The Monkey King 3',0,'2018-02-27','喜剧 / 爱情 / 动作','西游最特别一难全新演绎','《西游记女儿国》是由星皓影业有限公司出品的喜剧魔幻片，由郑保瑞执导，郭富城、冯绍峰、赵丽颖、小沈阳、罗仲谦、林志玲、梁咏琪、刘涛、苑琼丹、潘斌龙、施诗领衔主演。该片讲述了唐僧师徒四人在取经路上与各路妖魔鬼怪斗智斗勇的故事。',0,0,0,0,'星皓影业有限公司','中国',0,0,0,0,0,0,0,0,2,27,2,2018),(27,'http://img5.mtime.cn/mt/2018/01/30/102123.92074166_1280X720X2.jpg',134,0,40,'','瑞恩·库格勒',0,'黑豹','Black Panther',0,'2018-02-27','动作 / 冒险 / 科幻','劇情改編自漫威漫畫旗下同名漫畫人物故事','在漫威影业的影片《黑豹》中，特查拉在其父亲——前瓦坎达国王去世之后，回到了这个科技先进但与世隔绝的非洲国家，继任成为新一任“黑豹”及国王。当旧敌重现时，作为“黑豹”及国王的特查拉身陷两难境地，眼看着瓦坎达及全世界陷于危难之中。面对背叛与危险，这位年轻的国王必须联合同盟，释放“黑豹”全部力量，奋力捍卫他的人民。',0,0,0,0,'漫威漫畫公司','美國',0,0,0,0,0,0,0,0,1,27,2,2018),(28,'http://img5.mtime.cn/mt/2018/01/10/092346.98748326_1280X720X2.jpg',109,0,40,'','黄真真',0,'闺蜜2','GIRLS 2',0,'2018-02-27','喜剧 / 剧情','该片讲述希汶、Kimmy和嘉岚在越南度过的一段疯狂的婚前冒险之旅','《閨蜜2》（英语：Girls Ⅱ）（又稱《閨蜜2之單挑越南黑幫》），是一部於2018年上映的愛情喜劇電影。是電影《閨蜜》的第二部曲，由陳意涵、薛凱琪、張鈞甯領銜主演，台灣於2018年3月2日上映。 演員列表[编辑]. 演員姓名, 角色名稱, 介紹. 陳意涵, 希汶. 薛凱琪, Kimmy. 張鈞甯, 嘉嵐. 參考資料[编辑]. 跳转^ 飛車躲黑道大鬧街頭。',0,0,0,0,'万达影业公司','中国',0,0,0,0,0,0,0,0,1,27,2,2018),(29,'http://img5.mtime.cn/mt/2018/02/27/151252.12272322_1280X720X2.jpg',93,0,40,'','威尔·古勒',0,'比得兔','Peter Rabbit',0,'2018-02-27','动画 / 冒险 / 喜剧','比得兔，又譯彼得兔，（英语：Peter Rabbit）是一個虛構的圖畫小說擬人角色。','比得兔，又譯彼得兔，（英语：Peter Rabbit）是一個虛構的圖畫小說擬人角色，他的作者是英國女性作家暨插畫家碧雅翠絲·波特（Helen Beatrix Potter）。比得兔最早出現在1902年所出版的童書：《比得兔的故事》（The Tale of Peter Rabbit），之後，碧雅翠絲·波特又陸續出版五本和比得兔有關的童書。',0,0,0,0,'古卫儿影业','美国',0,0,0,0,0,0,0,0,1,27,2,2018),(30,'http://img5.mtime.cn/mt/2018/02/12/152914.18554032_1280X720X2.jpg',112,0,40,'','保罗·维尔奇',0,'爱在记忆消逝前','The Leisure Seeker',0,'2018-02-27','冒险 / 喜剧 / 剧情','老少女牵手健忘老暖男爱情大逃亡','该片改编自Michael Zadoorian的同名小说，讲述了一对患重病的老夫妻开着老旧的房车从马塞诸塞州出发前往海明威的故居，在旅途中回顾人生的故事。  　　约翰是一位文学教师和作家，海明威的死忠，尽管阿兹海默症已经让他的记忆越来越差，甚至会忘记妻子和儿女的名字，却依旧可以随口说出海明威的所有作品，《老人和海》更是他的钟爱。因为妻子艾拉患重病需要住院， 约翰就只能进入养老院，携手走过一辈子的二人面临永远分开的未来。于是，他们有了一个疯狂的决定，决心驾驶名叫“leisure seeker 求闲者”的房车，',0,0,0,0,'意大利国家影视公司','意大利',0,0,0,0,0,0,0,0,1,27,2,2018),(31,'http://img5.mtime.cn/mt/2018/01/08/104524.96607652_1280X720X2.jpg',96,0,40,'','克里斯托弗·兰敦',0,'忌日快乐','Happy Death Day',0,'2018-02-28','恐怖 / 悬疑 / 惊悚','一部2017年美國恐怖喜劇片','《忌日快樂》（英语：Happy Death Day）是一部2017年美國恐怖喜劇片，由克里斯多福·B·藍登執導，史考特·洛沃德爾（英语：Scott Lobdell）撰寫劇本。電影主演包括潔西卡·羅瑟和伊瑟瑞爾·布薩德（英语：Israel Broussard），其劇情描述一名在生日被謀殺的女大學生，竟遭遇到時間循環而不斷地重新回到那一天，使她決定藉此找出兇手。',0,0,0,0,'美国国际娱乐','美国',0,1,1,0,0,0,0,0,1,1,3,2018);
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
  `film_id` int(11) NOT NULL DEFAULT '0',
  `film_name` char(50) NOT NULL DEFAULT '',
  `actor_id` int(11) NOT NULL DEFAULT '0',
  `actor_name` char(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`fa_id`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `film_actor`
--

LOCK TABLES `film_actor` WRITE;
/*!40000 ALTER TABLE `film_actor` DISABLE KEYS */;
INSERT INTO `film_actor` VALUES (5,15,'黑豹',3,'露皮塔·尼永奥'),(6,16,'红海行动',4,'林超贤'),(7,16,'红海行动',5,'张译'),(8,16,'红海行动',6,'黄景瑜'),(9,17,'唐人街探案2',7,'陈思诚'),(10,17,'唐人街探案2',8,'王宝强'),(11,17,'唐人街探案2',9,'刘昊然'),(12,18,'厉害了，我的国',10,'卫铁'),(13,18,'厉害了，我的国',10,'卫铁'),(14,19,'捉妖记2',12,'许诚毅'),(15,19,'捉妖记2',13,'梁朝伟'),(16,19,'捉妖记2',14,'白百何'),(17,20,'祖宗十九代',15,'郭德纲'),(18,20,'祖宗十九代',16,'岳云鹏'),(19,20,'祖宗十九代',17,'吴京'),(20,21,'熊出没·变形记',18,'丁亮'),(21,21,'熊出没·变形记',18,'丁亮'),(22,21,'熊出没·变形记',20,'林汇达'),(23,22,'西游记女儿国',21,'郑保瑞'),(24,22,'西游记女儿国',22,'郭富城'),(25,22,'西游记女儿国',23,'冯绍峰'),(26,23,'黑豹',24,'瑞恩·库格勒'),(27,23,'黑豹',24,'瑞恩·库格勒'),(28,23,'黑豹',24,'瑞恩·库格勒'),(29,23,'黑豹',29,'查德维克·博斯曼'),(30,23,'黑豹',24,'瑞恩·库格勒'),(31,23,'黑豹',29,'查德维克·博斯曼'),(32,23,'黑豹',24,'瑞恩·库格勒'),(33,23,'黑豹',29,'查德维克·博斯曼'),(34,23,'黑豹',36,'露皮塔·尼永奥'),(35,28,'闺蜜2',37,'黄真真'),(36,28,'闺蜜2',38,'陈意涵'),(37,28,'闺蜜2',39,'薛凯琪'),(38,29,'比得兔',40,'威尔·古勒'),(39,29,'比得兔',41,'詹姆斯·柯登'),(40,29,'比得兔',42,'多姆纳尔·格里森'),(41,30,'爱在记忆消逝前',43,'保罗·维尔奇'),(42,30,'爱在记忆消逝前',44,'海伦·米伦'),(43,30,'爱在记忆消逝前',45,'唐纳德·萨瑟兰'),(44,31,'忌日快乐',46,'克里斯托弗·兰敦'),(45,31,'忌日快乐',47,'杰西卡·罗瑟'),(46,31,'忌日快乐',48,'伊瑟尔·布罗萨德');
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
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `film_order`
--

LOCK TABLES `film_order` WRITE;
/*!40000 ALTER TABLE `film_order` DISABLE KEYS */;
INSERT INTO `film_order` VALUES (29,'1520072884',1,35,'2018-03-03','0',2,6,2,1,20,-1,'',''),(30,'1520072884',1,35,'2018-03-03','0',2,6,3,1,20,-1,'',''),(31,'1520074108',0,35,'2018-03-03','0',2,7,2,1,20,-1,'',''),(32,'1520074108',0,35,'2018-03-03','0',2,6,4,1,20,-1,'',''),(33,'1520074628',0,35,'2018-03-03','0',2,6,4,1,18,-1,'',''),(34,'1520074628',0,35,'2018-03-03','0',2,6,3,1,18,-1,'',''),(35,'1520074801',0,35,'2018-03-03','0',2,6,3,1,19,-1,'',''),(36,'1520074801',0,35,'2018-03-03','0',2,6,5,1,19,-1,'',''),(37,'1520075199',1,35,'2018-03-03','0',2,6,4,1,19,-1,'01:01',''),(38,'1520075199',1,35,'2018-03-03','0',2,6,5,1,19,-1,'01:01',''),(39,'1520075695',1,35,'2018-03-03','0',2,6,3,1,20,3,'01:01',''),(40,'1520075695',1,35,'2018-03-03','0',2,5,4,1,20,3,'01:01',''),(41,'1520075771',1,35,'2018-03-03','0',2,5,5,1,20,4,'01:01',''),(42,'1520075771',1,35,'2018-03-03','0',2,7,3,1,20,4,'01:01',''),(43,'1520076113',1,35,'2018-03-03','0',2,5,4,1,20,3,'01:01',''),(44,'1520076356',1,35,'2018-03-03','0',2,7,3,1,19,4,'01:01',''),(45,'1520082772',1,35,'2018-03-03','0',2,6,5,1,20,3,'01:01',''),(46,'1520562773',1,35,'2018-03-09','0',2,4,3,6,19,-1,'01:00','');
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movie_hall`
--

LOCK TABLES `movie_hall` WRITE;
/*!40000 ALTER TABLE `movie_hall` DISABLE KEYS */;
INSERT INTO `movie_hall` VALUES (1,'一号厅','{\"x\":12,\"y\":12,\"no\":\"[\\\"1no1\\\",\\\"2no1\\\"]\"}',3),(2,'二号厅','{\"x\":11,\"y\":11,\"no\":\"[\\\"1no2\\\",\\\"2no3\\\"]\"}',6),(4,'一号厅','{\"x\":10,\"y\":10,\"no\":\"[\\\"1no1\\\",\\\"2no2\\\"]\"}',6),(5,'一号厅','{\"x\":10,\"y\":10,\"no\":\"[\\\"1no1\\\",\\\"5no1\\\",\\\"4no1\\\",\\\"3no1\\\",\\\"2no1\\\"]\"}',11),(6,'二号厅','{\"x\":10,\"y\":10,\"no\":\"[\\\"1no1\\\",\\\"3no1\\\",\\\"2no1\\\",\\\"4no1\\\",\\\"\\\"]\"}',11);
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
  `phone` char(12) NOT NULL DEFAULT '',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'李锡钒','3144c55546345df13cf8a36d2be16ee7','2018-03-03','111','18814128406'),(2,'ricoder','3144c55546345df13cf8a36d2be16ee7','2018-03-03','111','18814128405'),(3,'李锡钒','e10adc3949ba59abbe56e057f20f883e','2018-03-08','111',''),(4,'李锡钒','e10adc3949ba59abbe56e057f20f883e','2018-03-08','111',''),(5,'李锡钒','e10adc3949ba59abbe56e057f20f883e','2018-03-08','111',''),(6,'李锡钒','e10adc3949ba59abbe56e057f20f883e','2018-03-08','1359704355@qq.com','18814128406');
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

-- Dump completed on 2018-03-09  2:34:37
