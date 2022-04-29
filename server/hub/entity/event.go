package entity

import "strconv"

//数据库Event表对象
type Event struct {
	EquipmentTemplateId int
	EventId             int
	EventName           string
	StartType           int
	EndType             int
	StartExpression     string
	SuppressExpression  string
	EventCategory       int
	SignalId            int
	Enable              bool
	Visible             bool
	Description         string
	DisplayIndex        int
	ModuleNo            int

	EventConditions map[int]*EventCondition `gorm:"-"`
}

func (e Event) GetKey() string {
	eventKey := strconv.Itoa(e.EquipmentTemplateId) + "." + strconv.Itoa(e.EventId)
	return eventKey
}

func (e *Event) TableName() string {
	tableName := "fsu_event"
	return tableName
}
