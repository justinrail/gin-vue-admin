package stub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/base"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
)

var covCache []*domain.COV
var cogCache []*domain.COG
var codCache []*domain.COD
var cosCache []*domain.COS
var coaCache []*domain.COA

func LoadMockCOGs() {
	gateways := shadow.GetGateways()

	cogCache = make([]*domain.COG, len(gateways))

	for gatewayIndex := range gateways {
		gateway := gateways[gatewayIndex]
		cogCache = append(cogCache, &domain.COG{
			GatewayID: gateway.GatewayID,
			Flag:      base.GatewayFlagUnkown,
			Address:   gateway.IP,
			Timestamp: 0,
		})
	}
}

func LoadMockCODs() {
	devices := shadow.GetDevices()

	codCache = make([]*domain.COD, len(devices))

	for deviceIndex := range devices {
		device := devices[deviceIndex]
		codCache = append(codCache, &domain.COD{
			DeviceID: device.DeviceID,
		})
	}
}

func LoadMockCOAs() {
	alarms := shadow.GetAlarms()

	coaCache = make([]*domain.COA, len(alarms))

	for alarmIndex := range alarms {
		alarm := alarms[alarmIndex]
		coaCache = append(coaCache, &domain.COA{
			GatewayID: alarm.GatewayID,
			AlarmKey:  alarm.GetKey(),
		})
	}
}

func LoadMockCOSs() {

}

func LoadMockCOVs() {
	points := shadow.GetPoints()

	covCache = make([]*domain.COV, len(points))

	for pointIndex := range points {
		point := points[pointIndex]

		covCache = append(covCache, &domain.COV{
			DeviceID: point.DeviceID,
			PointID:  point.PointID,
			IsValid:  true, //默认为true
			PointKey: point.GetKey(),
		})
	}
}
