DROP TABLE IF EXISTS `company`;

CREATE TABLE `company` (
  `companyid` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `password` varchar(100) NOT NULL,
  PRIMARY KEY (`companyid`),
  UNIQUE KEY `companyid_UNIQUE` (`companyid`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1;


LOCK TABLES `company` WRITE;
INSERT INTO `company` VALUES (1,'321sc3','ascghdccasd@gmail.com','$2a$10$x0VSzkgKjVuwupFydD66Cu.pbO4.lSgmYMYY1xtu4c56FvaL3NdCu'),(2,'sdf','dsfds',''),(3,'3213','ascghdcasd@gmail.com','$2a$10$aJtwVW81eayZzY65qLMFU.ltSXyjU5DtFYYiuJMZ5gkno95ftwKSS'),(7,'3213','ascghdxccasd@gmail.com','$2a$10$ssTOA7.TB3QAFTtD5YEIsuChVtLOWfb83/3yL8caDxGmQ5HiC0356'),(9,'3213','ascs   ghdxccasd@gmail.com','$2a$10$VgfR9RiNAzhfiI3egjGVLOwMbEowQresD35zXAN.QKtJDmpoubLjS'),(11,'3213','ascs cdghdxccasd@gmail.com','$2a$10$7tA7YtK1qCBvZ0.Hz8u7gevzZKo4g.6XE2No6gZK2SBUEuBka.65i');
UNLOCK TABLES;
