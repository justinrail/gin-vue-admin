package protocol

import "strconv"

type SignalResponseItem struct {
	SourceHostId  int
	StationId     int
	EquipmentId   int
	DataId        int32
	SequenceId    string
	UpdateTime    int64
	EventSeverity int
	SignalMeaning string
	Flag          int
	IsMasking     bool
	SampleTime    uint32
	BaseTypeId    int32
	SignalType    int
	Value         *DynamicValue
}

func (signalResponseItem *SignalResponseItem) GetKey() string {
	return strconv.Itoa(signalResponseItem.EquipmentId) + "." + strconv.Itoa(int(signalResponseItem.DataId))
}
