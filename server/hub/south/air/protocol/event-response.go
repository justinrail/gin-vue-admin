package protocol

import (
	"encoding/binary"
	"math"
)

const (
	EventReportLen = 1

	EventSequenceLen = 32

	EventLen = 4

	EventConditionLen = 4

	EventBaseTypeLen = 4
)

type EventResponse struct {
	IsSend              bool
	EventReport         int
	EquipmentEventItems []*EquipmentEventItem
}

//解析协议包
func (eventResponse *EventResponse) FromByteArray(data []byte) {

	index := HeadLength
	stationID := int(binary.LittleEndian.Uint32(data[index:]))
	index = index + 4 + 4 + 1
	eventResponse.EventReport = int(data[index])
	index += EventReportLen
	equipmentItemCount := int(binary.LittleEndian.Uint16(data[index:]))
	index += LengthLen
	eventResponse.EquipmentEventItems = make([]*EquipmentEventItem, 0)
	for i := 0; i < equipmentItemCount; i++ {
		equipmentItem := &EquipmentEventItem{}
		equipmentItem.EventResponseItems = make([]*EventResponseItem, 0)
		equipmentItem.EquipmentId = int(binary.LittleEndian.Uint32(data[index:]))
		index += EquipmentLen
		eventItemCount := binary.LittleEndian.Uint16(data[index:])
		index += LengthLen
		for j := 0; j < (int)(eventItemCount); j++ {
			eventResponse := &EventResponseItem{}
			eventResponse.TriggerValue = &DynamicValue{}
			eventResponse.StationId = stationID
			eventResponse.EquipmentId = equipmentItem.EquipmentId
			eventResponse.SequenceId = string(data[index : index+EventSequenceLen])
			index += EventSequenceLen
			eventResponse.DataId = int(binary.LittleEndian.Uint32(data[index:]))
			index += EventLen
			utpTime := binary.LittleEndian.Uint32(data[index:])
			eventResponse.StartTime = utpTime
			index += StartUpTimeLen
			utpTime = binary.LittleEndian.Uint32(data[index:])
			eventResponse.EndTime = utpTime
			index += StartUpTimeLen
			eventResponse.ConditionId = int(binary.LittleEndian.Uint32(data[index:]))
			index += EventConditionLen
			eventResponse.Overturn = int(binary.LittleEndian.Uint16(data[index:]))
			index = index + 2
			len := data[index]
			index++
			if 0 < len {
				eventResponse.Meanings = string(data[index : index+int(len)])
			} else {
				eventResponse.Meanings = ""
			}
			index += int(len)
			eventResponse.TriggerValue.ValueType = int(data[index])
			index++
			valueLength := data[index]
			index++
			if eventResponse.TriggerValue.ValueType == 0 {
				bits := binary.LittleEndian.Uint32(data[index:])
				eventResponse.TriggerValue.FloatValue = math.Float32frombits(bits)
			}
			if eventResponse.TriggerValue.ValueType == 1 {
				eventResponse.TriggerValue.StringValue = string(data[index : index+(int)(valueLength)])
			}
			index += (int)(valueLength)
			eventResponse.BaseTypeId = int64(binary.LittleEndian.Uint32(data[index:]))
			index += EventBaseTypeLen
			equipmentItem.EventResponseItems = append(equipmentItem.EventResponseItems, eventResponse)
		}
		eventResponse.EquipmentEventItems = append(eventResponse.EquipmentEventItems, equipmentItem)
	}
}
