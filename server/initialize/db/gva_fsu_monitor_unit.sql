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
-- Table structure for table `fsu_monitor_unit`
--

DROP TABLE IF EXISTS `fsu_monitor_unit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `fsu_monitor_unit` (
  `MonitorUnitId` int NOT NULL,
  `MonitorUnitName` varchar(128) NOT NULL,
  `MonitorUnitCategory` int NOT NULL,
  `MonitorUnitCode` varchar(128) NOT NULL,
  `WorkStationId` int DEFAULT NULL,
  `StationId` int DEFAULT NULL,
  `IpAddress` varchar(128) DEFAULT NULL,
  `RunMode` int DEFAULT NULL,
  `ConfigFileCode` char(32) DEFAULT NULL,
  `ConfigUpdateTime` datetime DEFAULT NULL,
  `SampleConfigCode` char(32) DEFAULT NULL,
  `SoftwareVersion` varchar(64) DEFAULT NULL,
  `Description` varchar(255) DEFAULT NULL,
  `StartTime` datetime DEFAULT NULL,
  `HeartbeatTime` datetime DEFAULT NULL,
  `ConnectState` int NOT NULL DEFAULT '2',
  `UpdateTime` datetime NOT NULL,
  `IsSync` tinyint(1) NOT NULL DEFAULT '1',
  `SyncTime` datetime DEFAULT NULL,
  `IsConfigOK` tinyint(1) NOT NULL DEFAULT '1',
  `ConfigFileCode_Old` char(32) DEFAULT NULL,
  `SampleConfigCode_Old` char(32) DEFAULT NULL,
  `AppCongfigId` int DEFAULT NULL,
  `CanDistribute` tinyint(1) NOT NULL,
  `Enable` tinyint(1) NOT NULL,
  `ProjectName` varchar(255) DEFAULT NULL,
  `ContractNo` varchar(255) DEFAULT NULL,
  `InstallTime` datetime DEFAULT NULL,
  PRIMARY KEY (`MonitorUnitId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fsu_monitor_unit`
--

LOCK TABLES `fsu_monitor_unit` WRITE;
/*!40000 ALTER TABLE `fsu_monitor_unit` DISABLE KEYS */;
INSERT INTO `fsu_monitor_unit` VALUES (731955649,'Nurse监控单元',2,'148562009',NULL,262776526,'192.168.2.141',1,NULL,NULL,NULL,NULL,'','2022-03-10 15:36:42','2022-03-10 15:52:45',0,'2021-03-19 10:09:22',0,NULL,0,'00000000000000000000000000000000','00000000000000000000000000000000',1,1,1,NULL,NULL,NULL),(755010001,'Stand工作站自诊断虚拟监控单元',0,'755010001',NULL,-755,'',1,NULL,NULL,NULL,'','',NULL,NULL,2,'2022-02-09 18:11:49',0,NULL,1,NULL,NULL,-1,1,1,NULL,NULL,NULL);
/*!40000 ALTER TABLE `fsu_monitor_unit` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-11 10:42:22
