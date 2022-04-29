package protocol

import "encoding/binary"

type SamplerConnectStateResponse struct {
	MonitorUnitId            int32
	SamplerUnitId            int32
	SampleTime               uint32
	ConnectState             int32
	SamplerConnectStateItems []*SamplerConnectStateItem
}

func (samplerConnectStateResponse *SamplerConnectStateResponse) FromByteArray(data []byte, index int) {
	samplerConnectStateResponse.MonitorUnitId = int32(binary.LittleEndian.Uint32(data[index:]))
	index += 4
	equipmentCount := int(data[index])
	index++
	samplerConnectStateResponse.SamplerConnectStateItems = make([]*SamplerConnectStateItem, 0)
	for i := 0; i < equipmentCount; i++ {
		equipmentItem := &SamplerConnectStateItem{}
		equipmentItem.EquipmentId = int(binary.LittleEndian.Uint32(data[index:]))
		index += 4
		samplerConnectStateResponse.SamplerConnectStateItems = append(samplerConnectStateResponse.SamplerConnectStateItems, equipmentItem)
	}
	samplerConnectStateResponse.SamplerUnitId = int32(binary.LittleEndian.Uint32(data[index:]))
	index += 4
	utpTime := binary.LittleEndian.Uint32(data[index:])
	samplerConnectStateResponse.SampleTime = utpTime
	index += 4
	samplerConnectStateResponse.ConnectState = int32(data[index])
}
