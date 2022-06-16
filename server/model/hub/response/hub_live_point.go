package response

type LivePoint struct {
	DeviceID   int    `json:"deviceID"`
	DeviceName string `json:"deviceName"`
	PointID    int    `json:"pointID"`
	PointName  string `json:"pointName"`
	LiveValue  string `json:"liveValue"`
	UpdateTime string `json:"updateTime"`
}
