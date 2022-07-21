package domain

import (
	"strconv"
	"time"
)

//COV change of device signal's value (refs: activeSignal)
type COV struct {
	//GateWayID int  因为EquipmentId或DeviceID全局唯一所以不需要Gateway的ID做DTO
	DeviceID int
	PointID  int
	IsValid  bool
	PointKey string `json:"-"`
	//当前数值值
	CurrentNumericValue float32 `json:"NumValue"`
	//当前字符串值
	CurrentStringValue string `json:",omitempty"`

	Timestamp int64
}

//Clone clone cov
func (cov *COV) Clone() *COV {
	pcov := &COV{}
	pcov.DeviceID = cov.DeviceID
	pcov.CurrentNumericValue = cov.CurrentNumericValue
	pcov.PointID = cov.PointID
	pcov.CurrentStringValue = cov.CurrentStringValue
	pcov.IsValid = cov.IsValid
	pcov.Timestamp = cov.Timestamp

	return pcov
}

func (cov *COV) toString() string {
	stamp := time.Unix(cov.Timestamp, 0)
	return strconv.Itoa(cov.PointID) + "~" + strconv.Itoa(cov.DeviceID) + "~" +
		"~1~" + cov.CurrentStringValue + "~" + stamp.Format("2006-01-02 15:04:05")
}
