DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `iduser` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `password` varchar(1000) NOT NULL,
  PRIMARY KEY (`iduser`),
  UNIQUE KEY `iduser_UNIQUE` (`iduser`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=latin1;

LOCK TABLES `user` WRITE;
INSERT INTO `user` VALUES (1,'sdsd','dssd',''),(2,'3213','awdasd@gmail.com','$2a$10$y5qGjUn5v8G09PRfgCkb2e7iJhsR0lwCacY3WNKu37VD5DhA.t3i2'),(4,'3213','awdsasd@gmail.com','$2a$10$IW066GM2XcW/M11PQ6RTVeEVa18cMNgYlRP3twOd4wv65ii3F5YNe'),(6,'3213','awdsdsasd@gmail.com','$2a$10$c6R4ZOmUjPtC3q245Rbs6.19kgqaiiA29IyaMAiMKEl/8l8ENyFfi'),(7,'3213','awdsdsdcasd@gmail.com','$2a$10$fcCuX5jpyja6SW.GFdmY0OAc3.zYnB4ugfMgcSx5zKOljKJ1j12Hy'),(11,'3213','awddcasd@gmail.com','$2a$10$FoMn5VRY9fTSF55eJBD27uHDLzQwoqW57lLPvl2KgGzlvHscDhMsi'),(13,'3213','ascwddcasd@gmail.com','$2a$10$c.EAFot1tNgLO2svPRCNZe.VP56yEnlNQMf9bxusbbMkeYRp1ykSi'),(15,'3213','ascscwddcasd@gmail.com','$2a$10$y2TfQcxvUAsMuOYBYHC5cOauAgQFfUlWHnZ/fuMh1lUvnSS0Bst5q'),(17,'3213','ascsccswddcasd@gmail.com','$2a$10$3kZ4XnWexE2JO.jhVPPFuuu8EC1B1d7pTF0V30dgkBqJDFoOB7sIK'),(20,'3213','ascsccswdsdcasd@gmail.com','$2a$10$xSGfQzYvONA06nlBFiX6eOSIRL3nFWVyG/XDdqM4lPnj0K7G4Sa/W'),(21,'321sc3','ascghdccasd@gmail.com','$2a$10$PC7occjqP5wetOyKUtD41e2SQTItHQbJuRZZnuOnALg3Gj9XPRBdW'),(22,'3213','ascghdcasd@gmail.com','$2a$10$tYHaNZdRf/vjMzJumD.j/uQ8d70V5dICPtZmKPs6.F1UXeTbEpn6O'),(24,'3213','asccasd@gmail.com','$2a$10$J9D/Ton.968GOcTJdnZScO/SYih2.zmwizrrGG5g7zTz2sQhE6nM.'),(25,'3213','asccasfd@gmail.com','$2a$10$AdUqvmuWOhwueA/SpdizIuKinL8dJo.bLhSblFg3GnPAehd6n/GuO');
UNLOCK TABLES;