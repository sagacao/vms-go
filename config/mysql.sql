DROP TABLE IF EXISTS `log_stat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_stat` (
  `channel` varchar(64) NOT NULL,
  `gameid` varchar(16) NOT NULL,
  `newly` varchar(8) NOT NULL,
  `tow_pr` varchar(8) NOT NULL,
  `three_pr` varchar(8) NOT NULL,
  `seven_pr` varchar(8) NOT NULL,
  `retention` varchar(8) NOT NULL,
  `logdate` date NOT NULL,
  PRIMARY KEY (`channel`, `gameid`, `logdate`),
  KEY `index_channel` (`channel`) USING BTREE,
  KEY `index_gameid` (`gameid`) USING BTREE,
  KEY `index_logdate` (`logdate`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;