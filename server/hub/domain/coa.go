package domain

import "strconv"

//change of alarm
type COA struct {
	GatewayID        int
	DeviceID         int
	EventID          int
	EventConditionID int
	AlarmKey         string
	SequenceId       string //底端流水号
	NumericValue     float32
	StringValue      string
	StartTime        int64
	EndTime          int64
}

func (c *COA) GetUniqueKey() string {
	return strconv.Itoa(c.DeviceID) + "-" + strconv.Itoa(c.EventID) + "-" + strconv.Itoa(c.EventConditionID) + "-" + strconv.FormatInt(c.StartTime, 10)
}

func (c *COA) Clone() *COA {
	coa := &COA{
		GatewayID:        c.GatewayID,
		DeviceID:         c.DeviceID,
		EventID:          c.EventID,
		EventConditionID: c.EventConditionID,
		AlarmKey:         c.AlarmKey,
		SequenceId:       c.SequenceId,
		NumericValue:     c.NumericValue,
		StringValue:      c.StringValue,
		StartTime:        c.StartTime,
		EndTime:          c.EndTime,
	}

	return coa
}

func (c *COA) ToString() string {
	return c.GetUniqueKey() + "." + strconv.FormatFloat(float64(c.NumericValue), 'f', 10, 32) + "-" + strconv.FormatInt(c.EndTime, 10)
}
