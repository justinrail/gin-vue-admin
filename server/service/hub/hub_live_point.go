package hub

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
	hubRes "github.com/flipped-aurora/gin-vue-admin/server/model/hub/response"
)

type LivePointService struct {
}

// GetLivePointsByDevice 获取GetLivePointsByDevice记录
func (srv *LivePointService) GetLivePointsByDevice(deviceId int) []*hubRes.LivePoint {
	livePoints := make([]*hubRes.LivePoint, 0)
	device, existDevice := shadow.GetDeviceByID(deviceId)
	if existDevice {
		for index := range device.Points {
			p := device.Points[index]
			livePoints = append(livePoints, &hubRes.LivePoint{
				DeviceID:   p.DeviceID,
				DeviceName: device.DeviceName,
				PointID:    p.PointID,
				PointName:  p.PointName,
				LiveValue:  fmt.Sprint(p.CurrentNumericValue),
				UpdateTime: p.GetUpdateTimeString(),
			})
		}
	}
	return livePoints
}
