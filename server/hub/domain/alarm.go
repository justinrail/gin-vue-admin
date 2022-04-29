package domain

import (
	"container/list"
	"fmt"
	"strconv"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/entity"
)

//EventCondition的领域对象，内容是Event和Condition的集合，避免内存查询使用多Event一级
type Alarm struct {
	GatewayID           int
	DeviceID            int
	PointID             int
	EventId             int
	EventConditionId    int
	EventName           string
	EquipmentTemplateId int
	StartType           int
	EndType             int
	StartExpression     string
	SuppressExpression  string
	EventCategory       int
	Enable              bool
	Visible             bool
	Description         string
	DisplayIndex        int
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
	StandardID          int
	EventSeverity       int
	StandardName        int

	//当前数值值缓冲
	CurrentNumericValue float32
	//当前字符串值缓冲
	CurrentStringValue string
	//缓存COA包中间的告警
	StateLogs *list.List

	//更新时间
	UpdateTime int64
	//开始时间
	StartTime int64
	//结束时间
	EndTime int64
	lock    sync.Mutex
}

func (alarm *Alarm) GetKey() string {
	return strconv.Itoa(alarm.DeviceID) + "." + strconv.Itoa(alarm.EventId) + "." + strconv.Itoa(alarm.EventConditionId)
}

func (alarm *Alarm) From(deviceId int, event *entity.Event, eventCondition *entity.EventCondition) {
	alarm.DeviceID = deviceId
	alarm.EventId = event.EventId
	alarm.Enable = event.Enable
	alarm.EventCategory = event.EventCategory
	alarm.EventSeverity = eventCondition.EventSeverity
	alarm.EventName = event.EventName
	alarm.Meanings = eventCondition.Meanings
	alarm.EventConditionId = eventCondition.EventConditionId
	alarm.StandardID = eventCondition.BaseTypeId
	alarm.StandardName = eventCondition.StandardName
}

func (alarm *Alarm) ToString() string {
	return fmt.Sprintf("%s %s %s %.2f %d %d ", alarm.GetKey(), alarm.EventName, alarm.Meanings,
		alarm.CurrentNumericValue, alarm.StartTime, alarm.EndTime)
}

func NewAlarm() *Alarm {
	alarm := &Alarm{
		GatewayID:           0,
		DeviceID:            0,
		PointID:             0,
		EventId:             0,
		EventConditionId:    0,
		EventName:           "",
		EquipmentTemplateId: 0,
		StartType:           0,
		EndType:             0,
		StartExpression:     "",
		SuppressExpression:  "",
		EventCategory:       0,
		Enable:              false,
		Visible:             false,
		Description:         "",
		DisplayIndex:        0,
		StartOperation:      "",
		StartCompareValue:   0,
		StartDelay:          0,
		EndOperation:        "",
		EndCompareValue:     0,
		EndDelay:            0,
		Frequency:           0,
		FrequencyThreshold:  0,
		Meanings:            "",
		EquipmentState:      0,
		StandardID:          0,
		EventSeverity:       0,
		StandardName:        0,
		CurrentNumericValue: 0,
		CurrentStringValue:  "",
		StateLogs:           &list.List{},
		UpdateTime:          0,
		StartTime:           0,
		EndTime:             0,
		lock:                sync.Mutex{},
	}

	return alarm
}
