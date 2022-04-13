DROP TABLE IF EXISTS `invoices`;
CREATE TABLE `invoices` (
  `idinvoices` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `selling_company_id` int(11) NOT NULL,
  `buying_company_id` int(11) NOT NULL,
  `grand_total` int(11) NOT NULL,
  `acknowledged` tinyint(4) DEFAULT '0',
  `settled` tinyint(4) DEFAULT '0',
  `products` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`idinvoices`),
  UNIQUE KEY `idinvoices_UNIQUE` (`idinvoices`),
  KEY `company_id_idx` (`selling_company_id`),
  KEY `user_id_idx` (`user_id`),
  KEY `buying_company_id_idx` (`buying_company_id`),
  CONSTRAINT `buying_company_id` FOREIGN KEY (`buying_company_id`) REFERENCES `company` (`companyid`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `selling_company_id` FOREIGN KEY (`selling_company_id`) REFERENCES `company` (`companyid`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`iduser`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=latin1;

LOCK TABLES `invoices` WRITE;
INSERT INTO `invoices` VALUES (2,21,1,3,12212,0,0,NULL),(3,21,1,3,2112,0,0,NULL),(4,21,1,3,1212,0,0,NULL),(6,21,1,3,23232,0,0,NULL),(7,21,1,3,1212,0,0,NULL),(9,21,2,1,13223,1,1,NULL),(10,21,2,1,13223,0,0,NULL),(11,21,2,1,132233232,0,1,NULL),(12,24,1,3,91,0,0,NULL),(13,24,1,3,91,0,0,NULL),(14,24,1,3,91,0,0,NULL),(15,24,1,3,91,0,0,NULL),(16,1,1,3,91,0,0,NULL),(17,1,1,3,91,0,0,NULL),(18,1,1,3,46,0,0,NULL),(19,24,1,3,91,0,0,'A'),(20,24,1,3,91,0,0,'A'),(21,1,1,3,46,0,0,'[A 3213]'),(22,1,1,3,46,0,0,'[A 3213]');
UNLOCK TABLES;