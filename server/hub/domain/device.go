package domain

import (
	"container/list"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/base"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/entity"
)

const (
	DevicePacketLogQueueSize = 10
)

//Device 采集设备领域对象
type Device struct {
	GatewayID           int
	DeviceID            int
	EquipmentTemplateID int
	DeviceName          string
	ConnectState        int
	Alarms              map[string]*Alarm //Alarm 实际为EventCondition级对象，所以ID用Event+Condition组合
	Points              map[int]*Point
	PacketLogs          *list.List
	UpdateTime          int64
}

//NewDevice create domain device
func NewDevice() *Device {
	device := Device{}
	device.Points = map[int]*Point{}
	device.Alarms = map[string]*Alarm{}
	device.PacketLogs = list.New()
	return &device
}

func (device *Device) UpdateConnectState(flag int) {
	switch flag {
	case base.DeviceConStateUnkown:
		device.ConnectState = base.DeviceConStateUnkown
	case base.DeviceConStateOffline:
		device.ConnectState = base.DeviceConStateOffline
	case base.DeviceConStateOnline:
		device.ConnectState = base.DeviceConStateOnline
	}
}

func (device *Device) From(equipment *entity.Equipment) {
	device.GatewayID = equipment.MonitorUnitId
	device.DeviceID = equipment.EquipmentId
	device.DeviceName = equipment.EquipmentName
	device.EquipmentTemplateID = equipment.EquipmentTemplateId
}

func (device *Device) AppendPacketLogs(cod *COD) {
	device.PacketLogs.PushBack(cod) // Enqueue

	if device.PacketLogs.Len() > DevicePacketLogQueueSize {
		e := device.PacketLogs.Front()

		if e != nil {
			device.PacketLogs.Remove(e)
		}
	}
}
