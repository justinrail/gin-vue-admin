package protocol

type EventResponseItem struct {
	IsSecondary         bool
	BaseTypeId          int64
	IsNotify            bool
	ConditionId         int
	TriggerValue        *DynamicValue
	Flag                int
	Meanings            string
	InstructionStatus   string
	InstructionStatusId string
	InstructionId       string
	StartTime           uint32
	EndTime             uint32
	EventSeverity       int
	Overturn            int
	IsConfirmFist       bool
	IsEndFist           bool
	ConfirmUserName     string
	ConfirmUser         int
	ConfirmTime         uint32
	Memo                string
	SeverityName        string
	EventStatus         int
	Value               string
	SourceHostId        int
	StationId           int
	EquipmentId         int
	DataId              int
	SequenceId          string
	UpdateTime          uint64
}
