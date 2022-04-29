package entity

import (
	"strconv"
	"time"
)

type SamplerUnit struct {
	SamplerUnitId       int
	PortId              int
	MonitorUnitId       int
	SamplerId           int
	ParentSamplerUnitId int
	SamplerType         int
	SamplerUnitName     string
	Address             int
	SpUnitInterval      float32
	DllPath             string
	ConnectState        int
	UpdateTime          time.Time
	PhoneNumber         string
	Description         string
}

func (s *SamplerUnit) TableName() string {
	tableName := "fsu_sampler_unit"
	return tableName
}

func (s *SamplerUnit) GetKey() string {
	signalKey := strconv.Itoa(s.MonitorUnitId) + "." + strconv.Itoa(s.SamplerUnitId)
	return signalKey
}
