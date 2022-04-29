package entity

import "time"

type Equipment struct {
	StationId           int
	EquipmentId         int
	EquipmentName       string
	EquipmentNo         string
	EquipmentModule     string
	EquipmentStyle      string
	AssetState          int
	Price               float32
	UsedLimit           float32
	UsedDate            time.Time
	BuyDate             time.Time
	Vendor              string
	Unit                string
	EquipmentCategory   int
	EquipmentType       int
	EquipmentClass      int
	EquipmentState      int
	EventExpression     string
	StartDelay          float32
	EndDelay            float32
	Property            string
	Description         string
	EquipmentTemplateId int
	HouseId             int
	MonitorUnitId       int
	WorkStationId       int
	SamplerUnitId       int
	DisplayIndex        int
	ConnectState        int
	UpdateTime          time.Time
	ParentEquipmentId   string
	RatedCapacity       string
	InstalledModule     string
	ProjectName         string
	ContractNo          string
	InstallTime         time.Time
	EquipmentSN         string
	SO                  string
}

func (s *Equipment) TableName() string {
	tableName := "fsu_equipment"
	return tableName
}
