package shadow

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/entity"
	"go.uber.org/zap"
)

type LocalEntityCache struct {
	//监控单元全局缓存，以MonitorUnitID为key和唯一标示
	MonitorUnits map[int]*entity.MonitorUnit
	//设备全局缓存，以ID为key和唯一标示
	Equipments map[int]*entity.Equipment
	//设备模板全局缓存，以ID为key和唯一标识
	EquipmentTemplates map[int]*entity.EquipmentTemplate
	//模板信号全局缓存，以'EquipmentTemplateId.SignalId'为Key和唯一标示
	Signals map[string]*entity.Signal
	//模板事件全局缓存，以'EquipmentTemplateId.EventId'为Key和唯一标示
	Events map[string]*entity.Event
	//模板事件条件全局缓存，以'EquipmentTemplateId.EventId.ConditionId'为Key和唯一标示
	EventConditions map[string]*entity.EventCondition
	//监控单元全局缓存， 以'MonitorUnitId.SamplerUnitId'为Key
	SamplerUnits map[string]*entity.SamplerUnit
}

func (cache *LocalEntityCache) FindEquipment(equipmentId int) (*entity.Equipment, bool) {
	equipment, existEquipment := cache.Equipments[equipmentId]
	if existEquipment {
		return equipment, true
	} else {
		return nil, false
	}
}

//加载配置缓存（按表加载，方便管理和更新）
func (cache *LocalEntityCache) Load() {
	//Load ALL MonitorUnits
	var mus []entity.MonitorUnit
	global.GVA_DB.Find(&mus)
	cache.MonitorUnits = map[int]*entity.MonitorUnit{}

	for muIndex := range mus {
		cache.MonitorUnits[mus[muIndex].MonitorUnitId] = &mus[muIndex]
	}

	//Load all Equipments
	var equipments []entity.Equipment
	global.GVA_DB.Find(&equipments)
	cache.Equipments = map[int]*entity.Equipment{}

	for equipmentIndex := range equipments {
		cache.Equipments[equipments[equipmentIndex].EquipmentId] = &equipments[equipmentIndex]
	}

	//Load all EquipmentTemplates
	var equipmentTemplates []entity.EquipmentTemplate
	global.GVA_DB.Find(&equipmentTemplates)
	cache.EquipmentTemplates = map[int]*entity.EquipmentTemplate{}

	for equipmentTemplateIndex := range equipmentTemplates {
		equipmentTemplate := equipmentTemplates[equipmentTemplateIndex]
		equipmentTemplate.Signals = map[int]*entity.Signal{}
		equipmentTemplate.Events = map[int]*entity.Event{}
		cache.EquipmentTemplates[equipmentTemplate.EquipmentTemplateId] = &equipmentTemplate
	}

	//Load all EquipmentTemplates' signals
	var sigs []entity.Signal
	global.GVA_DB.Find(&sigs)

	cache.Signals = map[string]*entity.Signal{}

	for sigIndex := range sigs {
		sig := sigs[sigIndex]
		cache.Signals[sig.GetKey()] = &sig

		equipmentTemplate, eqExist := cache.EquipmentTemplates[sig.EquipmentTemplateId]

		if eqExist {
			equipmentTemplate.Signals[sig.SignalId] = &sig
		} else {
			global.GVA_LOG.Warn("signal's equipmentTemplate missing : ",
				zap.Int("EquipmentTemplateId", sig.EquipmentTemplateId))
		}
	}

	//Load all EquipmentTemplates' events
	var events []entity.Event
	global.GVA_DB.Find(&events)

	cache.Events = map[string]*entity.Event{}
	for eventIndex := range events {
		event := events[eventIndex]
		event.EventConditions = map[int]*entity.EventCondition{}

		cache.Events[event.GetKey()] = &event

		equipmentTemplate, eqExist := cache.EquipmentTemplates[event.EquipmentTemplateId]

		if eqExist {
			equipmentTemplate.Events[event.EventId] = &event
		} else {
			global.GVA_LOG.Warn("event's equipmentTemplate missing : ",
				zap.Int("EquipmentTemplateId", event.EquipmentTemplateId))
		}
	}

	//Load all EquipmentTemplates' eventConditions
	var eventConditions []entity.EventCondition
	global.GVA_DB.Find(&eventConditions)

	cache.EventConditions = map[string]*entity.EventCondition{}
	for eventConditionIndex := range eventConditions {
		eventCondition := eventConditions[eventConditionIndex]
		cache.EventConditions[eventCondition.GetKey()] = &eventCondition
		equipmentTemplate, eqExist := cache.EquipmentTemplates[eventCondition.EquipmentTemplateId]
		if eqExist {
			event, eventExist := equipmentTemplate.Events[eventCondition.EventId]
			if eventExist {
				event.EventConditions[eventCondition.EventConditionId] = &eventCondition
			}
		}

	}

	//Load all Sampler Units
	var samplerUnits []entity.SamplerUnit
	global.GVA_DB.Find(&samplerUnits)

	cache.SamplerUnits = map[string]*entity.SamplerUnit{}
	for samplerUnitIndex := range samplerUnits {
		samplerUnit := samplerUnits[samplerUnitIndex]
		cache.SamplerUnits[samplerUnit.GetKey()] = &samplerUnit
	}

}

//找到单条的配置信号实体
func (cache *LocalEntityCache) FindSignal(equipmentId int, dataId int) (*entity.Signal, bool) {

	equipment, existEquipment := cache.Equipments[equipmentId]
	if existEquipment {
		equipmentTemplate, existTemplate := cache.EquipmentTemplates[equipment.EquipmentTemplateId]

		if existTemplate {
			sigKey := strconv.Itoa(equipmentTemplate.EquipmentTemplateId) + "." + strconv.Itoa(dataId)
			signal, existSignal := cache.Signals[sigKey]
			if existSignal {
				return signal, true
			} else {
				return nil, false
			}
		} else {
			return nil, false
		}
	} else {
		return nil, false
	}

}
