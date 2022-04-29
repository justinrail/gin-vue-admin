package protocol

type EquipmentSignalItem struct {
	EquipmentId         int
	SignalResponseItems []*SignalResponseItem
}

type EquipmentEventItem struct {
	EquipmentId        int
	EventResponseItems []*EventResponseItem
}

type RelatedSignalItem struct {
	EquipmentId                int
	RelatedSignalResponseItems []*RelatedSignalResponseItem
}

type SamplerConnectStateItem struct {
	EquipmentId                     int
	SamplerConnectStateResponseItem []*SamplerConnectStateResponseItem
}
