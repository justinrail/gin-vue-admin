package domain

import "time"

type DeviceAdapter struct {
	GatewayID         int
	SamplerUnitID     int
	ParentID          int
	DeviceID          int
	DeviceAdapterName string
	ConnectState      int
	UpdateTime        time.Time
}
