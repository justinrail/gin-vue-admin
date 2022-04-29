-- MySQL dump 10.13  Distrib 8.0.28, for Win64 (x86_64)
--
-- Host: localhost    Database: gva
-- ------------------------------------------------------
-- Server version	8.0.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `fsu_station`
--

DROP TABLE IF EXISTS `fsu_station`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `fsu_station` (
  `StationId` int NOT NULL,
  `StationName` varchar(255) NOT NULL,
  `Latitude` decimal(20,17) DEFAULT NULL,
  `Longitude` decimal(20,17) DEFAULT NULL,
  `SetupTime` datetime DEFAULT NULL,
  `CompanyId` int DEFAULT NULL,
  `ConnectState` int NOT NULL DEFAULT '2',
  `UpdateTime` datetime NOT NULL,
  `StationCategory` int NOT NULL,
  `StationGrade` int NOT NULL,
  `StationState` int NOT NULL,
  `ContactId` int DEFAULT NULL,
  `SupportTime` int DEFAULT NULL,
  `OnWayTime` float DEFAULT NULL,
  `SurplusTime` float DEFAULT NULL,
  `FloorNo` varchar(50) DEFAULT NULL,
  `PropList` varchar(255) DEFAULT NULL,
  `Acreage` float DEFAULT NULL,
  `BuildingType` int DEFAULT NULL,
  `ContainNode` tinyint(1) NOT NULL DEFAULT '0',
  `Description` varchar(255) DEFAULT NULL,
  `BordNumber` int DEFAULT NULL,
  `CenterId` int NOT NULL,
  `Enable` tinyint(1) NOT NULL DEFAULT '1',
  `StartTime` datetime DEFAULT NULL,
  `EndTime` datetime DEFAULT NULL,
  `ProjectName` varchar(255) DEFAULT NULL,
  `ContractNo` varchar(255) DEFAULT NULL,
  `InstallTime` datetime DEFAULT NULL,
  PRIMARY KEY (`StationId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fsu_station`
--

LOCK TABLES `fsu_station` WRITE;
/*!40000 ALTER TABLE `fsu_station` DISABLE KEYS */;
INSERT INTO `fsu_station` VALUES (-755,'eaton-管理站点',NULL,NULL,NULL,NULL,1,'2022-02-09 18:11:48',3,2,1,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,0,'',NULL,755,1,NULL,NULL,NULL,NULL,NULL),(262776526,'Nurse站点',22.55329000000000000,113.88308000000000000,NULL,NULL,0,'2021-03-19 10:09:22',2,1,1,NULL,NULL,NULL,NULL,NULL,'10',NULL,NULL,0,'',NULL,755,1,NULL,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `fsu_station` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-11 10:42:23
