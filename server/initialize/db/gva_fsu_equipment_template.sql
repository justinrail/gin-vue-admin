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
-- Table structure for table `fsu_equipment_template`
--

DROP TABLE IF EXISTS `fsu_equipment_template`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `fsu_equipment_template` (
  `EquipmentTemplateId` int NOT NULL,
  `EquipmentTemplateName` varchar(128) NOT NULL,
  `ParentTemplateId` int NOT NULL,
  `Memo` varchar(255) NOT NULL,
  `ProtocolCode` varchar(255) NOT NULL,
  `EquipmentCategory` int NOT NULL,
  `EquipmentType` int NOT NULL,
  `Property` varchar(255) DEFAULT NULL,
  `Description` varchar(255) DEFAULT NULL,
  `EquipmentStyle` varchar(128) DEFAULT NULL,
  `Unit` varchar(255) DEFAULT NULL,
  `Vendor` varchar(255) DEFAULT NULL,
  `EquipmentBaseType` int DEFAULT NULL,
  `StationCategory` int DEFAULT NULL,
  PRIMARY KEY (`EquipmentTemplateId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fsu_equipment_template`
--

LOCK TABLES `fsu_equipment_template` WRITE;
/*!40000 ALTER TABLE `fsu_equipment_template` DISABLE KEYS */;
INSERT INTO `fsu_equipment_template` VALUES (123084848,'自诊断设备',0,'#','7A4EEF969B0167A9533128601C719AA7',99,1,'1/3','','','','',2001,0),(181068179,'配电-三相2',0,'2020/11/13 11:27:22:首次导入模板','ADL4008-00',12,1,'1/3','','','','',1501,0),(224526521,'UPS-短卡',0,'2020/8/24 16:50:02:首次导入模板','PPCUPS8-00',31,1,'1/3','','','','',501,0),(239226515,'PDU_SLAVE1-CH',0,'2021-01-15 14:41:02:首次导入模板','58D176E166E4517618BFFEA32797D4BA',85,1,'1/3','','','','',1503,0),(257572310,'PDU_SLAVE2-CH',0,'2021-01-15 14:41:02:首次导入模板','58D176E166E4517618BFFEA32797D4BB',85,1,'1/3','','','','',1503,0),(274430612,'PDU_SLAVE3-CH',0,'2021-01-15 14:41:02:首次导入模板','58D176E166E4517618BFFEA32797D4BC',85,1,'1/3','','','','',1503,0),(345348143,'采集器',0,'#','6F67F49B95C300D859610F3E37DDD3AE',51,1,'1/3','','','','',1004,0),(359364069,'UPS-FA卡',0,'2020/7/10 15:38:57:首次导入模板','ModbusMS_ZH8-00',31,1,'1/3','','','','',501,0),(373889048,'空调-分体机',0,'2019/7/27 10:29:47:首次导入模板','ASLHR35A8-01',43,1,'1/3','','','','',702,0),(388297395,'Collector',0,'#','6F67F49B95C300D859610F3E37DDD3A2',51,1,'1/3','','','','',1004,0),(396816357,'PDU-CH',0,'2021-01-15 14:41:02:首次导入模板','58D176E166E4517618BFFEA32797D4B8',85,1,'1/3','','','','',1503,0),(402530686,'温湿度',0,'2018/6/20 18:08:47:首次导入模板','FDD44104C4CCC0B94465363A9F86E959',51,1,'1/3','','','','',1006,0),(431318156,'空调-一体机',0,'2017/11/8 16:23:36:首次导入模板','AEnviCoolDC036-00',43,1,'1/3','','','','',702,0),(523718189,'UPS-Sirius-CH',0,'2020/6/29 17:38:49:首次导入模板','USIRIUS8-00',31,1,'1/3','','','','',501,0),(853701543,'电池',0,'#','HSCM02N8-00',24,1,'1/3','','','','',1101,0),(883130628,'UPS-长卡',0,'2020/12/9 16:14:53:首次导入模板','CMCUPS8-00',31,1,'1/3','','','','',501,0),(899026572,'空调-15K/30K',0,'2021/4/27 15:47:12:首次导入模板','EATONEK8-00',43,1,'1/3','','','','',702,0),(932765648,'配电-单相2',0,'2020/11/13 9:58:24:首次导入模板','AcrelDDS13528-00',12,1,'1/3','','','','',1501,0),(942951547,'门锁',0,'2020-03-17 10:22:08:首次导入模板','OSJModbusV18-00',82,1,'1/3','','','','',1001,0),(957680696,'ATS-CH',0,'#','EATON_ATS8-02',85,1,'1/3','','','','',1504,0);
/*!40000 ALTER TABLE `fsu_equipment_template` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-11 10:42:09
