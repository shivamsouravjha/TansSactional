DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `idproduct` int(11) NOT NULL AUTO_INCREMENT,
  `productname` varchar(100) DEFAULT NULL,
  `productprice` int(11) DEFAULT NULL,
  `id_company` int(11) DEFAULT NULL,
  PRIMARY KEY (`idproduct`),
  UNIQUE KEY `idproduct_UNIQUE` (`idproduct`),
  KEY `id_company_idx` (`id_company`),
  CONSTRAINT `id_company` FOREIGN KEY (`id_company`) REFERENCES `company` (`companyid`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

LOCK TABLES `product` WRITE;
INSERT INTO `product` VALUES (1,'A',1,1),(2,NULL,NULL,1),(3,'ASD',NULL,1),(4,'3213',45,2),(5,'3213',45,2);
UNLOCK TABLES;