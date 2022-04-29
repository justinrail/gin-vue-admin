package protocol

type SamplerConnectStateResponseItem struct {
	SourceHostId int
	StationId    int
	EquipmentId  int
	DataId       int
	SequenceId   string
	UpdateTime   uint32
}
