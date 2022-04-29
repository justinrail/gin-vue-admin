package shadow

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
)

type LocalDomainCache struct {

	//Gateways 全局网关缓存
	Gateways map[int]*domain.Gateway
	//Devices 全局设备缓存
	Devices map[int]*domain.Device
	//Points 全局信号缓存
	Points map[string]*domain.Point
	//Alarms 全局告警缓存
	Alarms map[string]*domain.Alarm
	//DeviceAdapters 全局监控单元缓存
	DeviceAdapters map[string]*domain.DeviceAdapter
	//ActiveAlarms global activeAlarm cache
	ActiveAlarms map[string]*domain.ActiveAlarm
}

//Build Domain Objects From Entites
func (cache *LocalDomainCache) Load() {

	//load all gateways
	cache.Gateways = map[int]*domain.Gateway{}
	for muIndex := range entityCache.MonitorUnits {
		mu := entityCache.MonitorUnits[muIndex]
		gateway := domain.NewGateway()
		gateway.From(mu)
		cache.Gateways[mu.MonitorUnitId] = gateway
	}

	//global.GVA_LOG.Debug("Load all gateways")

	//load all devices
	cache.Devices = map[int]*domain.Device{}
	for equipmentIndex := range entityCache.Equipments {
		equipment := entityCache.Equipments[equipmentIndex]
		device := domain.NewDevice()
		device.From(equipment)
		cache.Devices[equipment.EquipmentId] = device

		//加载gateway级别的map的device集合
		gateway, gatewayExist := cache.Gateways[device.GatewayID]
		if gatewayExist {
			gateway.Devices[device.DeviceID] = device
		}
	}

	//load all deviceAdapters
	cache.DeviceAdapters = map[string]*domain.DeviceAdapter{}
	for samplerUnitIndex := range entityCache.SamplerUnits {
		samplerUnit := entityCache.SamplerUnits[samplerUnitIndex]
		deviceAdapter := &domain.DeviceAdapter{
			GatewayID:         samplerUnit.MonitorUnitId,
			SamplerUnitID:     samplerUnit.SamplerUnitId,
			ParentID:          samplerUnit.ParentSamplerUnitId,
			DeviceAdapterName: samplerUnit.SamplerUnitName,
			ConnectState:      samplerUnit.ConnectState,
		}

		cache.DeviceAdapters[samplerUnit.GetKey()] = deviceAdapter
	}

	//load all Points & alarms (从Entity的signal中实例化出来) 可能比较吃内存
	cache.Points = map[string]*domain.Point{}
	cache.Alarms = map[string]*domain.Alarm{}

	for deviceIndex := range cache.Devices {
		device := cache.Devices[deviceIndex]
		equipmentTemplateId := device.EquipmentTemplateID

		et, etExist := entityCache.EquipmentTemplates[equipmentTemplateId]
		if etExist {
			//load instance of signal as point
			for sigIndex := range et.Signals {
				sig := et.Signals[sigIndex]

				point := domain.NewPoint()
				point.From(device.DeviceID, sig)

				device.Points[point.PointID] = point
				cache.Points[point.GetKey()] = point
			}

			//load alarms from event & eventConditions
			for eventIndex := range et.Events {
				event := et.Events[eventIndex]

				for eventConditionIndex := range event.EventConditions {
					eventCondition := event.EventConditions[eventConditionIndex]

					alarm := domain.NewAlarm()
					alarm.From(device.DeviceID, event, eventCondition)

					device.Alarms[alarm.GetKey()] = alarm
					cache.Alarms[alarm.GetKey()] = alarm
				}

			}
		}

	}

}
