DROP TABLE IF EXISTS `usercompanyrelation`;
CREATE TABLE `usercompanyrelation` (
  `iduser` int(11) NOT NULL,
  `idcompany` int(11) NOT NULL,
  PRIMARY KEY (`iduser`),
  KEY `idcompany_idx` (`idcompany`),
  CONSTRAINT `idcompany` FOREIGN KEY (`idcompany`) REFERENCES `company` (`companyid`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `iduser` FOREIGN KEY (`iduser`) REFERENCES `user` (`iduser`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


LOCK TABLES `usercompanyrelation` WRITE;
INSERT INTO `usercompanyrelation` VALUES (20,1),(22,1),(24,1),(25,1),(21,2);
UNLOCK TABLES;