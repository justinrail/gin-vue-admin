package entity

import "strconv"

//EventCondition 表对象
type EventCondition struct {
	EventConditionId    int
	EquipmentTemplateId int
	EventId             int
	StartOperation      string
	StartCompareValue   float32
	StartDelay          int
	EndOperation        string
	EndCompareValue     float32
	EndDelay            int
	Frequency           int
	FrequencyThreshold  int
	Meanings            string
	EquipmentState      int
	BaseTypeId          int
	EventSeverity       int
	StandardName        int
}

func (ec *EventCondition) GetKey() string {
	eventConditionKey := strconv.Itoa(ec.EquipmentTemplateId) +
		"." + strconv.Itoa(ec.EventId) + "." + strconv.Itoa(ec.EventConditionId)

	return eventConditionKey
}

func (ec *EventCondition) TableName() string {
	tableName := "fsu_event_Condition"
	return tableName
}
