package entity

import "time"

type MonitorUnit struct {
	MonitorUnitId        int `gorm:"type:int(11);not null"`
	MonitorUnitName      string
	MonitorUnitCategory  int `gorm:"type:int(11);not null"`
	MonitorUnitCode      string
	WorkStationId        int `gorm:"type:int(11);null"`
	StationId            int `gorm:"type:int(11);null"`
	IpAddress            string
	RunMode              int `gorm:"type:int(11);null"`
	ConfigFileCode       string
	ConfigUpdateTime     time.Time
	SampleConfigCode     string
	SoftwareVersion      string
	Description          string
	StartTime            time.Time
	HeartbeatTime        time.Time
	ConnectState         int `gorm:"type:int(11);null"`
	UpdateTime           time.Time
	IsSync               bool `gorm:"type:bool"`
	SyncTime             time.Time
	IsConfigOK           bool `gorm:"type:bool"`
	ConfigFileCode_Old   string
	SampleConfigCode_Old string
	AppCongfigId         int  `gorm:"type:int(11);null"`
	CanDistribute        bool `gorm:"type:bool"`
	Enable               bool `gorm:"type:bool"`
	ProjectName          string
	ContractNo           string
	InstallTime          time.Time
}

func (s *MonitorUnit) TableName() string {
	tableName := "fsu_monitor_unit"
	return tableName
}
