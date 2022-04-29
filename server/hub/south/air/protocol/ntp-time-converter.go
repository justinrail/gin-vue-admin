package protocol

import "time"

//设置基准时间

func GetNTPTime() uint32 {
	t := time.Now()
	_, offset := t.Zone()
	timestamp := t.Unix()
	currenttime := time.Unix(timestamp+int64(offset), 0)
	return (uint32)(currenttime.Unix())
}
