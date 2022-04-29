package shadow

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
)

//核心业务缓存 暂用本地缓存（内存）
var domainCache LocalDomainCache

//配置对象缓存 暂用本地缓存（内存）
var entityCache LocalEntityCache

var activeAlarmSinker *ActiveAlarmSinker
var activeAlarmLogSinker *ActiveAlarmLogSinker

// 考虑本软件是小项目（FSU小于20），默认初始化时加载所有配置
// 先加载相关entity，再根据entity初始化相关domain对象
// TODO : domain对象除了静态配置，历史数据（比如上次离线数据）未来考虑增加
func Init() {
	domainCache = LocalDomainCache{}
	entityCache = LocalEntityCache{}
	activeAlarmSinker = &ActiveAlarmSinker{}
	activeAlarmLogSinker = &ActiveAlarmLogSinker{}
	entityCache.Load()
	domainCache.Load()
	activeAlarmSinker.Ready()

}

func GetGateways() []*domain.Gateway {
	gs := make([]*domain.Gateway, 0)
	for gatewayIndex := range domainCache.Gateways {
		gateway := domainCache.Gateways[gatewayIndex]
		gs = append(gs, gateway)
	}

	return gs
}

func GetGatewayByID(gatewayID int) (*domain.Gateway, bool) {
	gateway, ok := domainCache.Gateways[gatewayID]
	return gateway, ok
}

func GetDevices() []*domain.Device {
	ds := make([]*domain.Device, 0)

	for deviceIndex := range domainCache.Devices {
		device := domainCache.Devices[deviceIndex]
		ds = append(ds, device)
	}

	return ds
}

func GetDeviceByID(DeviceID int) (*domain.Device, bool) {
	device, ok := domainCache.Devices[DeviceID]
	return device, ok
}

func GetPoints() []*domain.Point {
	points := make([]*domain.Point, 0)

	for pointIndex := range domainCache.Points {
		point := domainCache.Points[pointIndex]
		points = append(points, point)
	}

	return points
}

func GetPointByKey(key string) (*domain.Point, bool) {
	point, ok := domainCache.Points[key]
	return point, ok
}

func GetAlarms() []*domain.Alarm {
	alarms := make([]*domain.Alarm, 0)

	for alarmIndex := range domainCache.Alarms {
		alarm := domainCache.Alarms[alarmIndex]
		alarms = append(alarms, alarm)
	}

	return alarms
}
func GetAlarmByKey(key string) (*domain.Alarm, bool) {
	alarm, ok := domainCache.Alarms[key]
	return alarm, ok
}

func GetActiveAlarmByKey(key string) (*domain.ActiveAlarm, bool) {
	activeAlarm, ok := domainCache.ActiveAlarms[key]
	return activeAlarm, ok
}

//activelarm has end and confirmed
func FinishActiveAlarm(coa *domain.COA) {
	//remove from activeAlarm cache
	activeAlarm, ok := domainCache.ActiveAlarms[coa.GetUniqueKey()]
	if ok {
		activeAlarmLogSinker.Sink(activeAlarm) //save log
		delete(domainCache.ActiveAlarms, coa.GetUniqueKey())
	}
}

func StartActiveAlarm(coa *domain.COA) {
	alarm, ok := domainCache.Alarms[coa.AlarmKey]
	if ok {
		activeAlarm := &domain.ActiveAlarm{
			GatewayID:           alarm.GatewayID,
			DeviceID:            alarm.DeviceID,
			PointID:             alarm.PointID,
			EventId:             alarm.EventId,
			EventConditionId:    alarm.EventConditionId,
			EventName:           alarm.EventName,
			EquipmentTemplateId: alarm.EquipmentTemplateId,
			EventCategory:       alarm.EventCategory,
			Description:         "",
			Frequency:           0,
			Meanings:            alarm.Meanings,
			StandardID:          alarm.StandardID,
			EventSeverity:       alarm.EventSeverity,
			StandardName:        alarm.StandardName,
			CurrentNumericValue: coa.NumericValue,
			CurrentStringValue:  coa.StringValue,
			StartTime:           coa.StartTime,
			EndTime:             0,
			ConfirmTime:         0,
			ConfimerID:          0,
		}

		domainCache.ActiveAlarms[coa.GetUniqueKey()] = activeAlarm
		activeAlarmSinker.Sink(activeAlarm) //save to activeAlarm table
	}

}

func EndActiveAlarm(coa *domain.COA) {
	activeAlarm, ok := domainCache.ActiveAlarms[coa.GetUniqueKey()]
	if ok {
		activeAlarm.EndTime = coa.EndTime
		activeAlarmSinker.Sink(activeAlarm) //update to activeAlarm table
	}
}

func AppendCOVLog(coa *domain.COA) {

}

//从缓冲获取Alarm的Entity配置（Alarm就是EventCondition和Event的合集，避免多维护Event对象）并组装成默认COA
//这个函数可用来放配置不匹配的诊断代码
//但问题是这样要有一次缓存命中，效率较低，如果接受数据慢，可考虑放到bus后处理
//这里不调用此函数，直接拼好给后台，进行命中测试
func GetPlanCOA(equipmentId int, dataId int, conditionId int) (*domain.COA, bool) {
	coa := &domain.COA{}

	alarmKey := strconv.Itoa(equipmentId) + "." + strconv.Itoa(dataId) + "." + strconv.Itoa(conditionId)

	_, existAlarm := domainCache.Alarms[alarmKey]
	if existAlarm {
		coa.AlarmKey = alarmKey
		return coa, true
	} else {
		return nil, false
	}
}

//从缓冲获取信号的SignalEntity配置并组装成默认COV, 这样做可以确保不扔坏数据给后方
//但问题是这样要有一次缓存命中，效率较低，如果接受数据慢，可考虑放到bus后处理
//这里不调用此函数，直接拼好给后台，进行命中测试
func GetPlanCOV(equipmentId int, dataId int) (*domain.COV, bool) {
	cov := &domain.COV{}

	pointKey := strconv.Itoa(equipmentId) + "." + strconv.Itoa(dataId)
	_, existPoint := domainCache.Points[pointKey]
	if existPoint {
		cov.PointKey = pointKey
		cov.DeviceID = equipmentId
		cov.PointID = dataId
		cov.IsValid = false

		return cov, true
	} else {
		return nil, false
	}
}
