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
-- Table structure for table `fsu_sampler`
--

DROP TABLE IF EXISTS `fsu_sampler`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `fsu_sampler` (
  `SamplerId` int NOT NULL,
  `SamplerName` varchar(128) NOT NULL,
  `SamplerType` smallint NOT NULL,
  `ProtocolCode` varchar(255) NOT NULL,
  `DLLCode` varchar(255) NOT NULL,
  `DLLVersion` varchar(32) NOT NULL,
  `ProtocolFilePath` varchar(255) NOT NULL,
  `DLLFilePath` varchar(255) NOT NULL,
  `DllPath` varchar(255) NOT NULL,
  `Setting` varchar(255) DEFAULT NULL,
  `Description` varchar(255) DEFAULT NULL,
  `SoCode` varchar(255) NOT NULL,
  `SoPath` varchar(255) NOT NULL,
  PRIMARY KEY (`SamplerId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fsu_sampler`
--

LOCK TABLES `fsu_sampler` WRITE;
/*!40000 ALTER TABLE `fsu_sampler` DISABLE KEYS */;
INSERT INTO `fsu_sampler` VALUES (198516103,'申菱HR35A',18,'ASLHR35A8-01                    ','                                ','','','','HR35A.DLL','9600,N,8,1','','                                ',''),(620093339,'HSCM02',18,'HSCM02N8-00                     ','                                ','','','','HSCM02N.DLL','9600,N,8,1','','                                ',''),(648647482,'Sirius-UPS',18,'USIRIUS8-00                     ','                                ','','','','USirius.DLL','9600,N,8,1','','                                ',''),(658081705,'ModbusMS_ZH',18,'ModbusMS_ZH8-00                 ','                                ','','','','ModbusMS.DLL','9600,N,8,1','','                                ',''),(669064163,'CMCUPS',18,'CMCUPS8-00                      ','                                ','','','','CMCUPS.DLL','9600,N,8,1','','                                ',''),(706834934,'EATONEK',18,'EATONEK8-00                     ','                                ','','','','EATONEK.DLL','9600,N,8,1','','                                ',''),(734164121,'英维克EC03U7-DC03直流空调',18,'AEnviCoolDC036-00               ','                                ','','','','DC03.DLL','9600,n,8,1','','                                ',''),(755000001,'工作站自诊断采集器',0,'LSC工作站自诊断设备6-00         ','                                ',' ',' ',' ','GTBusinessServer.exe','',' ','                                ',''),(755000002,'GMU-IO设备',18,'E4C01DC788BFBC280A91B32FD57178E7','                                ','','','','GMU-IO.so','comm_io_dev.so','','                                ',''),(761318384,'伊顿-ATS',18,'EATON_ATS8-02                   ','                                ','','','','EATON_ATS-CH.DLL','9600,N,8,1','','                                ',''),(770534919,'伊顿-PDU',18,'58D176E166E4517618BFFEA32797D4B8','                                ','','','','EATON_PDU-CH.DLL','9600,N,8,1','','                                ',''),(780043357,'生久机柜电子门锁',18,'OSJModbusV18-00                 ','                                ','','','','SJModbus.DLL','9600,N,8,1','','                                ',''),(809355814,'AcrelDDS1352',18,'AcrelDDS13528-00                ','                                ','','','','AcrelDDS1352.DLL','9600,N,8,1','','                                ',''),(814401698,'ADL400',18,'ADL4008-00                      ','                                ','','','','ADL400.DLL','9600,N,8,1','','                                ',''),(837384954,'ADL400',18,'ADL4008-00                      ','                                ','','','','ADL400.DLL','9600,N,8,1','','                                ',''),(866668335,'伊顿-PDU',18,'58D176E166E4517618BFFEA32797D4BA','                                ','','','','EATON_PDU_SLAVE1.DLL','9600,N,8,1','','                                ',''),(876089188,'伊顿-PDU',18,'58D176E166E4517618BFFEA32797D4BB','                                ','','','','EATON_PDU_SLAVE2.DLL','9600,N,8,1','','                                ',''),(885568333,'伊顿-PDU',18,'58D176E166E4517618BFFEA32797D4BC','                                ','','','','EATON_PDU_SLAVE3.DLL','9600,N,8,1','','                                ',''),(941801628,'UPS',18,'PPCUPS8-00                      ','                                ','','','','PPCUPS.DLL','9600,n,8,1','','                                ',''),(979105873,'Collector',18,'6F67F49B95C300D859610F3E37DDD3A2','                                ','','','','iView-IO-EN.dll','9600,N,8,1','','                                ','');
/*!40000 ALTER TABLE `fsu_sampler` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-11 10:42:25
