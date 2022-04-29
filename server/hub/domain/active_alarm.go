package domain

//活动告警，设计模式，在内存建立缓冲，在信号和告警领域对象关联活动状态
//告警先在内存组织好，活动变化的日志change和缓冲表在内存，采用sink模式异步写日志到mysql，历史直接写mysql
//即使活动告警重启，也仅仅丢失1秒内（sink周期，如果全写完的情况下）
type ActiveAlarm struct {
	GatewayID           int
	DeviceID            int
	PointID             int
	EventId             int
	EventConditionId    int
	EventName           string
	EquipmentTemplateId int
	StartType           int
	EndType             int
	EventCategory       int
	Enable              bool
	Visible             bool
	Description         string
	DisplayIndex        int
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

	//更新时间
	UpdateTime int64
	//开始时间
	StartTime int64
	//结束时间
	EndTime int64

	ConfirmTime int64
	ConfimerID  int
}
