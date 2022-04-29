package entity

import "strconv"

//测点配置对象 refs: CfgSignal
type Signal struct {
	EquipmentTemplateId int
	SignalId            int
	Enable              bool
	Visible             bool
	Description         string
	SignalName          string
	SignalCategory      int
	SignalType          int
	ChannelNo           int
	ChannelType         int
	Expression          string
	DataType            int
	ShowPrecision       string
	Unit                string
	StoreInterval       float32
	AbsValueThreshold   float32
	PercentThreshold    float32
	StaticsPeriod       int
	BaseTypeId          int
	ChargeStoreInterVal float32
	ChargeAbsValue      float32
	DisplayIndex        int
	MDBSignalId         int
	ModuleNo            int

	//   ` tinyint(1) NOT NULL,
	//   `Visible` tinyint(1) NOT NULL,
	//   `Description` varchar(255) DEFAULT NULL,
	//   `SignalName` varchar(128) NOT NULL,
	//   `SignalCategory` int(11) NOT NULL,
	//   `SignalType` int(11) NOT NULL,
	//   `ChannelNo` int(11) NOT NULL,
	//   `ChannelType` int(11) NOT NULL,
	//   `Expression` varchar(1024) DEFAULT NULL,
	//   `DataType` int(11) DEFAULT NULL,
	//   `ShowPrecision` varchar(20) DEFAULT NULL,
	//   `Unit` varchar(64) DEFAULT NULL,
	//   `StoreInterval` float DEFAULT NULL,
	//   `AbsValueThreshold` float DEFAULT NULL,
	//   `PercentThreshold` float DEFAULT NULL,
	//   `StaticsPeriod` int(11) DEFAULT NULL,
	//   `BaseTypeId` decimal(10,0) DEFAULT NULL,
	//   `ChargeStoreInterVal` float DEFAULT NULL,
	//   `ChargeAbsValue` float DEFAULT NULL,
	//   `DisplayIndex` int(11) NOT NULL,
	//   `MDBSignalId` int(11) DEFAULT NULL,
	//   `ModuleNo` int(11) NOT NULL DEFAULT '0',
}

func (s *Signal) GetKey() string {
	signalKey := strconv.Itoa(s.EquipmentTemplateId) + "." + strconv.Itoa(s.SignalId)
	return signalKey
}
func (s *Signal) TableName() string {
	tableName := "fsu_signal"
	return tableName
}
