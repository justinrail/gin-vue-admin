package protocol

import (
	"encoding/binary"
	"math"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
)

//RealSignalResponse 实时数据返回包
type RealSignalResponse struct {
	MessageType          int
	StationId            int
	HostId               int
	EquipmentSignalItems []*EquipmentSignalItem
	engineID             int
}

func (realSignalResponse *RealSignalResponse) GetCOVs() []*domain.COV {
	covs := make([]*domain.COV, 0)
	for _, pp := range realSignalResponse.EquipmentSignalItems {
		for _, cc := range pp.SignalResponseItems {
			//从缓冲获取信号的SignalEntity配置并组装成默认对象（实在不装配也可能行）
			cov, existCOV := shadow.GetPlanCOV(cc.EquipmentId, int(cc.DataId))

			//如果组装默认对象成功，则更新数据部分
			if existCOV {
				cov.IsValid = true
				cov.CurrentNumericValue = cc.Value.FloatValue
				cov.CurrentStringValue = cc.Value.StringValue
				cov.Timestamp = int64(cc.SampleTime)

				covs = append(covs, cov)
			}
		}
	}

	return covs
}

func (realSignalResponse *RealSignalResponse) FromByteArray(data []byte) {
	index := HeadLength
	stationID := int(binary.LittleEndian.Uint32(data[index:]))
	index = EncodingIndex
	index++
	equipmentItemCount := binary.LittleEndian.Uint16(data[index:])
	index += LengthLen
	realSignalResponse.EquipmentSignalItems = make([]*EquipmentSignalItem, 0)
	for i := 0; i < int(equipmentItemCount); i++ {
		equipmentItem := &EquipmentSignalItem{}
		equipmentItem.SignalResponseItems = make([]*SignalResponseItem, 0)
		equipmentItem.EquipmentId = int(binary.LittleEndian.Uint32(data[index:]))
		index += EquipmentLen

		signalResponseItemCount := binary.LittleEndian.Uint16(data[index:])
		index += LengthLen
		for j := 0; j < (int)(signalResponseItemCount); j++ {
			signalResponseItem := &SignalResponseItem{}
			signalResponseItem.Value = &DynamicValue{}
			signalResponseItem.StationId = stationID
			signalResponseItem.EquipmentId = equipmentItem.EquipmentId
			signalResponseItem.DataId = int32(binary.LittleEndian.Uint32(data[index:]))
			index += SignalLen
			signalResponseItem.Flag = int(data[index])
			index += SignalTypeLen
			signalResponseItem.EventSeverity = int(data[index])
			index += EventSeverityLen
			utpTime := binary.LittleEndian.Uint32(data[index:])
			signalResponseItem.SampleTime = utpTime
			index += StartUpTimeLen
			signalResponseItem.Value.ValueType = int(data[index])
			index += SignalTypeLen
			valueLength := data[index]
			index += ValueTypeLen
			if signalResponseItem.Value.ValueType == 0 {
				bits := binary.LittleEndian.Uint32(data[index:])
				signalResponseItem.Value.FloatValue = math.Float32frombits(bits)
			}
			if signalResponseItem.Value.ValueType == 1 {
				signalResponseItem.Value.StringValue = string(data[index : index+(int)(valueLength)])
			}
			index += (int)(valueLength)
			//signalResponseItem.BaseTypeId = (int32)(binary.LittleEndian.Uint32(data[index:]))
			index += SignalBaseTypeLen
			equipmentItem.SignalResponseItems = append(equipmentItem.SignalResponseItems, signalResponseItem)
		}
		realSignalResponse.EquipmentSignalItems = append(realSignalResponse.EquipmentSignalItems, equipmentItem)
	}
}
