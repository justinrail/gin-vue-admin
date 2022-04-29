package entity

type EquipmentTemplate struct {
	EquipmentTemplateId   int
	EquipmentTemplateName string
	ParentTemplateId      int
	Memo                  string
	ProtocolCode          string
	EquipmentCategory     int
	EquipmentType         int
	Property              string
	Description           string
	EquipmentStyle        string
	Unit                  string
	Vendor                string
	EquipmentBaseType     int
	StationCategory       int
	Signals               map[int]*Signal `gorm:"-"`
	Events                map[int]*Event  `gorm:"-"`
}

func (s *EquipmentTemplate) TableName() string {
	tableName := "fsu_equipment_template"
	return tableName
}
