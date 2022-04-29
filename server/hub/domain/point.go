//测点领域对象 refs: Signal
package domain

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/entity"
)

const (
	PointLogQueueSize = 10
)

//Point point of domain signal
type Point struct {
	PointID       int
	SignalID      int
	PointName     string
	Accuracy      string
	Unit          string
	Max           string
	Min           string
	DeviceID      int
	DataType      int
	AlarmSeverity int
	DefaultValue  string
	StandardID    int
	Cron          string
	Expression    string
	GatewayID     int
	//当前数值值缓冲
	CurrentNumericValue float32
	//当前字符串值缓冲
	CurrentStringValue string
	//数据是否有效
	IsAvailabe bool
	//数值计算（中心计算后的值状态）比如：数据是否有效,0:ok,1:low limit exceed 2: up limit exceed
	CalcState int
	//缓存全部COV的包
	PacketLogs *list.List
	//更新时间
	UpdateTime int64
	lock       sync.Mutex
}

func NewPoint() *Point {
	point := Point{
		PointID:             0,
		PointName:           "",
		Accuracy:            "",
		Unit:                "",
		Max:                 "",
		Min:                 "",
		DeviceID:            0,
		DataType:            0,
		AlarmSeverity:       0,
		DefaultValue:        "",
		StandardID:          0,
		Cron:                "",
		Expression:          "",
		GatewayID:           0,
		CurrentNumericValue: 0,
		CurrentStringValue:  "",
		IsAvailabe:          false,
		CalcState:           0,
		PacketLogs:          &list.List{},
		UpdateTime:          0,
		lock:                sync.Mutex{},
	}

	return &point
}

func (point *Point) From(deviceId int, signal *entity.Signal) {
	point.DeviceID = deviceId
	point.PointID = signal.SignalId
	point.SignalID = signal.SignalId
	point.PointName = signal.SignalName
	point.Accuracy = signal.ShowPrecision
	point.DataType = signal.DataType
	point.StandardID = signal.BaseTypeId
}

func (point *Point) GetKey() string {
	return strconv.Itoa(point.DeviceID) + "." + strconv.Itoa(point.SignalID)
}

//UpdateLimitState  自动计算模拟量的超限情况
// func (point *CorePoint) UpdateLimitState(cov *COV) {
// 	defer corePoint.lock.Unlock()
// 	corePoint.lock.Lock()

// 	if corePoint.CoreDataTypeID == 2 {
// 		corePoint.LimitState = 0
// 		if len(corePoint.Max) > 0 {
// 			max, err := strconv.ParseFloat(corePoint.Max, 32)
// 			exp.CheckError(err)
// 			if float32(max) < corePoint.CurrentNumericValue {
// 				corePoint.LimitState = 2
// 			}
// 		}
// 		if len(corePoint.Min) > 0 {
// 			min, err := strconv.ParseFloat(corePoint.Min, 32)
// 			exp.CheckError(err)
// 			if float32(min) > corePoint.CurrentNumericValue {
// 				corePoint.LimitState = 1
// 			}
// 		}
// 	}
// }

//updatePacketLogs update packet logs
func (point *Point) updatePacketLogs(cov *COV) {

	point.PacketLogs.PushBack(cov.Clone()) // Enqueue

	if point.PacketLogs.Len() > PointLogQueueSize {
		e := point.PacketLogs.Front()

		if e != nil {
			point.PacketLogs.Remove(e)
		}

	}
}

//GetCurrentValue GetCurrentValue
func (point *Point) GetCurrentValue() (string, error) {
	if !point.DataIsReady() {
		return "", fmt.Errorf("data is not ready")
	}

	if point.DataType == 1 {
		return fmt.Sprintf("%d", int(point.CurrentNumericValue)), nil
	} else if point.DataType == 2 {
		if len(point.Accuracy) == 0 {
			return fmt.Sprintf("%.2f %s", point.CurrentNumericValue, point.Unit), nil
		}
		dotPos := strings.Index(point.Accuracy, ".")
		if dotPos < 0 {
			return fmt.Sprintf("%.2f %s", point.CurrentNumericValue, point.Unit), nil
		}

		if dotPos == 1 {
			return fmt.Sprintf("%.1f %s", point.CurrentNumericValue, point.Unit), nil
		}

		if dotPos == 2 {
			return fmt.Sprintf("%.2f %s", point.CurrentNumericValue, point.Unit), nil
		}

		if dotPos == 3 {
			return fmt.Sprintf("%.3f %s", point.CurrentNumericValue, point.Unit), nil
		}

		if dotPos == 4 {
			return fmt.Sprintf("%.4f %s", point.CurrentNumericValue, point.Unit), nil
		}

		return fmt.Sprintf("%.2f %s", point.CurrentNumericValue, point.Unit), nil

	} else if point.DataType == 3 {
		return point.CurrentStringValue, nil
	} else {
		return "", fmt.Errorf("datatype not supported")
	}
}

//UpdateData update corepoint data
func (point *Point) UpdateData(cov *COV) {
	defer point.lock.Unlock()
	point.lock.Lock()

	point.IsAvailabe = cov.IsValid
	point.CurrentNumericValue = cov.CurrentNumericValue
	point.CurrentStringValue = cov.CurrentStringValue
	point.UpdateTime = cov.Timestamp
	point.updatePacketLogs(cov)
}

//DataIsReady 实时数据是否可用(判断gateway，coresource和isvalid)
func (point *Point) DataIsReady() bool {

	/*gw, ok := Gateways.Get(corePoint.GatewayID)
	if ok {
		gateway := gw.(*Gateway)
		if gateway.ConState == enum.GatewayConStateOnline {
			cr, exist := gateway.CoreSources.Get(corePoint.CoreSourceID)
			if exist {
				coresource := cr.(*CoreSource)
				if coresource.State == enum.CoreSourceConStateOnline {
					if corePoint.IsAvailabe {
						return true
					}
				}
			}
		}
	}*/

	return true
}
