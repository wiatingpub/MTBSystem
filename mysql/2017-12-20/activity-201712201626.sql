-- MySQL dump 10.13  Distrib 5.7.20, for osx10.12 (x86_64)
--
-- Host: localhost    Database: activity
-- ------------------------------------------------------
-- Server version	5.7.20

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
-- Table structure for table `ly_joiner`
--

DROP TABLE IF EXISTS `ly_joiner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ly_joiner` (
  `lid` int(64) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` char(50) NOT NULL DEFAULT '' COMMENT '参与者名字',
  `vote_num` int(64) NOT NULL DEFAULT '0' COMMENT '投票次数',
  `photo` char(100) NOT NULL DEFAULT '' COMMENT '参与者头像',
  `jid` int(64) NOT NULL DEFAULT '0' COMMENT '参与者id',
  `version` int(64) NOT NULL DEFAULT '1' COMMENT '第几期',
  PRIMARY KEY (`lid`),
  KEY `version` (`version`,`jid`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ly_joiner`
--

LOCK TABLES `ly_joiner` WRITE;
/*!40000 ALTER TABLE `ly_joiner` DISABLE KEYS */;
INSERT INTO `ly_joiner` VALUES (1,'桃子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',39939713,1),(2,'毕游侠',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',2171602,1),(3,'指间',1,'http://img15.3lian.com/2015/h1/280/d/4.jpg',24422,1),(4,'王宝宝',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',40791648,1),(5,'大卫',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',10647017,1),(6,'半个橙子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',37279291,1),(7,'弟弟',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',91111942,1),(8,'妹妹',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',82049832,1),(9,'一帆',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',55578162,1),(10,'桃子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',39939713,2),(11,'毕游侠',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',2171602,2),(12,'指间',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',24422,2),(13,'王宝宝',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',40791648,2),(14,'大卫',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',10647017,2),(15,'半个橙子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',37279291,2),(16,'弟弟',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',91111942,2),(17,'妹妹',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',82049832,2),(18,'一帆',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',55578162,2),(19,'桃子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',39939713,3),(20,'毕游侠',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',2171602,3),(21,'指间',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',24422,3),(22,'王宝宝',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',40791648,3),(23,'大卫',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',10647017,3),(24,'半个橙子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',37279291,3),(25,'弟弟',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',91111942,3),(26,'妹妹',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',82049832,3),(27,'一帆',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',55578162,3),(28,'桃子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',39939713,4),(29,'毕游侠',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',2171602,4),(30,'指间',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',24422,4),(31,'王宝宝',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',40791648,4),(32,'大卫',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',10647017,4),(33,'半个橙子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',37279291,4),(34,'弟弟',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',91111942,4),(35,'妹妹',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',82049832,4),(36,'一帆',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',55578162,4),(37,'桃子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',39939713,5),(38,'毕游侠',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',2171602,5),(39,'指间',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',24422,5),(40,'王宝宝',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',40791648,5),(41,'大卫',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',10647017,5),(42,'半个橙子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',37279291,5),(43,'弟弟',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',91111942,5),(44,'妹妹',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',82049832,5),(45,'一帆',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',55578162,5),(46,'桃子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',39939713,6),(47,'毕游侠',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',2171602,6),(48,'指间',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',24422,6),(49,'王宝宝',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',40791648,6),(50,'大卫',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',10647017,6),(51,'半个橙子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',37279291,6),(52,'弟弟',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',91111942,6),(53,'妹妹',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',82049832,6),(54,'一帆',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',55578162,6),(55,'桃子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',39939713,7),(56,'毕游侠',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',2171602,7),(57,'指间',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',24422,7),(58,'王宝宝',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',40791648,7),(59,'大卫',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',10647017,7),(60,'半个橙子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',37279291,7),(61,'弟弟',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',91111942,7),(62,'妹妹',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',82049832,7),(63,'一帆',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',55578162,7),(64,'桃子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',39939713,8),(65,'毕游侠',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',2171602,8),(66,'指间',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',24422,8),(67,'王宝宝',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',40791648,8),(68,'大卫',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',10647017,8),(69,'半个橙子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',37279291,8),(70,'弟弟',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',91111942,8),(71,'妹妹',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',82049832,8),(72,'一帆',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',55578162,8),(73,'桃子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',39939713,9),(74,'毕游侠',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',2171602,9),(75,'指间',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',24422,9),(76,'王宝宝',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',40791648,9),(77,'大卫',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',10647017,9),(78,'半个橙子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',37279291,9),(79,'弟弟',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',91111942,9),(80,'妹妹',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',82049832,9),(81,'一帆',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',55578162,9),(82,'桃子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',39939713,10),(83,'毕游侠',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',2171602,10),(84,'指间',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',24422,10),(85,'王宝宝',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',40791648,10),(86,'大卫',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',10647017,10),(87,'半个橙子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',37279291,10),(88,'弟弟',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',91111942,10),(89,'妹妹',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',82049832,10),(90,'一帆',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',55578162,10),(91,'桃子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',39939713,11),(92,'毕游侠',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',2171602,11),(93,'指间',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',24422,11),(94,'王宝宝',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',40791648,11),(95,'大卫',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',10647017,11),(96,'半个橙子',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',37279291,11),(97,'弟弟',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',91111942,11),(98,'妹妹',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',82049832,11),(99,'一帆',0,'http://img15.3lian.com/2015/h1/280/d/4.jpg',55578162,11);
/*!40000 ALTER TABLE `ly_joiner` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ly_play_plan`
--

DROP TABLE IF EXISTS `ly_play_plan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ly_play_plan` (
  `lid` int(64) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` char(50) NOT NULL DEFAULT '' COMMENT '用户名字',
  `show_url` char(100) NOT NULL DEFAULT '0' COMMENT '房间链接',
  `club_id` int(64) NOT NULL DEFAULT '0' COMMENT '俱乐部id',
  `img` char(100) NOT NULL DEFAULT '' COMMENT '图片',
  `type` char(50) NOT NULL DEFAULT '' COMMENT '类型：  排位  俱乐部赛',
  `day_week` char(50) NOT NULL DEFAULT '' COMMENT '星期几',
  `version` int(64) NOT NULL DEFAULT '0' COMMENT '第几期',
  PRIMARY KEY (`lid`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ly_play_plan`
--

LOCK TABLES `ly_play_plan` WRITE;
/*!40000 ALTER TABLE `ly_play_plan` DISABLE KEYS */;
INSERT INTO `ly_play_plan` VALUES (1,'name1','1',1,'http://192.168.1.50:8080/images/LyingMan/videoposter.png','俱乐部赛','1',1),(2,'name2','2',2,'http://192.168.1.50:8080/images/LyingMan/videoposter.png','排位','2',1),(3,'name3','3',3,'http://192.168.1.50:8080/images/LyingMan/videoposter.png','俱乐部赛','3',1),(4,'name4','4',4,'http://192.168.1.50:8080/images/LyingMan/videoposter.png','排位','4',1),(5,'name5','5',5,'http://192.168.1.50:8080/images/LyingMan/videoposter.png','俱乐部赛','5',1),(6,'name6','6',6,'http://192.168.1.50:8080/images/LyingMan/videoposter.png','排位','6',1),(7,'name7','7',7,'http://192.168.1.50:8080/images/LyingMan/videoposter.png','俱乐部赛','7',1),(8,'name4','4',4,'http://192.168.1.50:8080/images/LyingMan/videoposter.png','排位','4',2),(9,'name5','5',5,'http://192.168.1.50:8080/images/LyingMan/videoposter.png','俱乐部赛','5',2),(10,'name3','3',3,'http://192.168.1.50:8080/images/LyingMan/videoposter.png','俱乐部赛','3',2),(11,'name4','4',4,'http://192.168.1.50:8080/images/LyingMan/videoposter.png','排位','4',3);
/*!40000 ALTER TABLE `ly_play_plan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ly_version_video`
--

DROP TABLE IF EXISTS `ly_version_video`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ly_version_video` (
  `lid` int(64) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type` char(50) NOT NULL DEFAULT '' COMMENT '类型： 正片  复盘',
  `version` int(64) NOT NULL DEFAULT '1' COMMENT '第几期',
  `title` char(50) NOT NULL DEFAULT '' COMMENT '标题名字',
  `video` char(100) NOT NULL DEFAULT '' COMMENT '视频链接',
  `img` char(100) NOT NULL DEFAULT '' COMMENT '图片',
  PRIMARY KEY (`lid`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ly_version_video`
--

LOCK TABLES `ly_version_video` WRITE;
/*!40000 ALTER TABLE `ly_version_video` DISABLE KEYS */;
INSERT INTO `ly_version_video` VALUES (1,'正片1',1,'标题1','video1','http://192.168.1.50:8080/images/LyingMan/videoposter.png'),(2,'复盘1',1,'标题1','video1','http://192.168.1.50:8080/images/LyingMan/videoposter.png'),(3,'正片2',2,'标题2','video2','http://192.168.1.50:8080/images/LyingMan/videoposter.png'),(4,'复盘3',3,'标题3','video3','http://192.168.1.50:8080/images/LyingMan/videoposter.png'),(5,'正片4',4,'标题4','video4','http://192.168.1.50:8080/images/LyingMan/videoposter.png'),(6,'复盘4',4,'标题4','video4','http://192.168.1.50:8080/images/LyingMan/videoposter.png'),(7,'正片5',5,'标题5','video5','http://192.168.1.50:8080/images/LyingMan/videoposter.png'),(8,'复盘5',5,'标题5','video5','http://192.168.1.50:8080/images/LyingMan/videoposter.png');
/*!40000 ALTER TABLE `ly_version_video` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ly_vote_message`
--

DROP TABLE IF EXISTS `ly_vote_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ly_vote_message` (
  `lid` int(64) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `vid` int(64) NOT NULL DEFAULT '0' COMMENT '投票者id',
  `fid` char(50) NOT NULL DEFAULT '' COMMENT '投票者fakeid',
  `jid` int(64) NOT NULL DEFAULT '0' COMMENT '参与者id',
  `voter_name` char(50) NOT NULL DEFAULT '' COMMENT '投票者的名字',
  `version` int(64) NOT NULL DEFAULT '1' COMMENT '第几期',
  `create_at` char(50) NOT NULL DEFAULT '' COMMENT '投票时间',
  PRIMARY KEY (`lid`),
  KEY `version` (`version`,`jid`,`fid`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ly_vote_message`
--

LOCK TABLES `ly_vote_message` WRITE;
/*!40000 ALTER TABLE `ly_vote_message` DISABLE KEYS */;
INSERT INTO `ly_vote_message` VALUES (1,5,'55',39939713,'voterName',1,'2017-12-12'),(2,5,'55',39939713,'voterName',1,'2017-12-12'),(3,5,'55',39939713,'voterName',1,'2017-12-12'),(4,5,'56',39939713,'voterName',1,'2017-12-12'),(5,5,'57',39939713,'voterName',1,'2017-12-12'),(6,5,'58',39939713,'voterName',1,'2017-12-12'),(7,5,'59',39939713,'voterName',1,'2017-12-12'),(8,5,'60',39939713,'voterName',1,'2017-12-12'),(9,5,'61',39939713,'voterName',1,'2017-12-12'),(10,5,'62',39939713,'voterName',1,'2017-12-12'),(11,5,'63',39939713,'voterName',1,'2017-12-12'),(12,5,'64',39939713,'voterName',1,'2017-12-12'),(13,5,'65',39939713,'voterName',1,'2017-12-12'),(14,5,'66',39939713,'voterName',1,'2017-12-12'),(15,5,'67',39939713,'voterName',2,'2017-12-12'),(16,5,'68',39939713,'voterName',2,'2017-12-12'),(17,5,'69',39939713,'voterName',1,'2017-12-12');
/*!40000 ALTER TABLE `ly_vote_message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ly_vote_record`
--

DROP TABLE IF EXISTS `ly_vote_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ly_vote_record` (
  `lid` int(64) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_at` char(50) NOT NULL DEFAULT '' COMMENT '投票时间',
  `vote_times` int(64) NOT NULL DEFAULT '3' COMMENT '投票剩余次数',
  `fid` char(100) NOT NULL DEFAULT '' COMMENT '投票者fakeid',
  `vid` int(64) NOT NULL DEFAULT '0' COMMENT '投票者id',
  `version` int(64) NOT NULL DEFAULT '1' COMMENT '第几期',
  PRIMARY KEY (`lid`),
  KEY `version` (`version`,`fid`,`create_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ly_vote_record`
--

LOCK TABLES `ly_vote_record` WRITE;
/*!40000 ALTER TABLE `ly_vote_record` DISABLE KEYS */;
INSERT INTO `ly_vote_record` VALUES (1,'2017-12-11',3,'51',1,1),(2,'2017-12-11',3,'52',1,1),(3,'2017-12-20',0,'2B28ABDBD4A72EA40FF6FF3B190A466D',0,1);
/*!40000 ALTER TABLE `ly_vote_record` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-12-20 16:26:48
