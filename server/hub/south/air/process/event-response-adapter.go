package process

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/bus"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/air/protocol"
	flow "github.com/trustmaster/goflow"
)

type EventResponseAdapter struct {
	flow.Component
	In <-chan *protocol.PipelineObjectPair
}

func (adapter *EventResponseAdapter) Process() {
	for pair := range adapter.In {
		eventresponse := &protocol.EventResponse{}
		eventresponse.FromByteArray(pair.Data)
		// 处理活动事件
		for eventItemIndex := range eventresponse.EquipmentEventItems {
			eventItem := eventresponse.EquipmentEventItems[eventItemIndex]

			for eventReponseItemIndex := range eventItem.EventResponseItems {
				eventReponseItem := eventItem.EventResponseItems[eventReponseItemIndex]

				//这里COV装配测试了，如果不装配后期装配也是方案
				coa, existCOA := shadow.GetPlanCOA(eventReponseItem.EquipmentId,
					eventReponseItem.DataId, eventReponseItem.ConditionId)

				if existCOA {
					coa.StartTime = int64(eventReponseItem.StartTime)
					coa.EndTime = int64(eventReponseItem.EndTime)
					coa.DeviceID = eventReponseItem.EquipmentId
					coa.EventID = eventReponseItem.DataId
					coa.EventConditionID = eventReponseItem.ConditionId
					coa.NumericValue = eventReponseItem.TriggerValue.FloatValue
					coa.StringValue = eventReponseItem.TriggerValue.StringValue

					bus.COABus <- coa
				}
			}
		}
		ack := protocol.CreateAckPacket(pair.SequenceNumber, 0, int(HubEngineID))
		protocol.SendUDPData(pair.Url, ack)
	}

}
