package protocol

type RelatedSignalResponseItem struct {
	SourceHostId               int
	StationId                  int
	EquipmentId                int
	DataId                     int32
	SequenceId                 string
	UpdateTime                 int64
	EventSeverity              int32
	SignalMeaning              string
	Flag                       int32
	IsMasking                  bool
	SampleTime                 uint32
	BaseTypeId                 int64
	SignalType                 int32
	value                      *DynamicValue
	BusinessId                 int32
	StateExpressionId          int32
	SerialId                   int32
	StateExpressionValue       float32
	StateExpressionTriggerTime uint32
}
