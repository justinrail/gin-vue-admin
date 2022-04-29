package protocol

import (
	"encoding/binary"
	"fmt"
)

//HeartbeatRequest 心跳请求包
type HeartBeatRequest struct {
	//HostType 主机类型
	HostType byte

	MessageType uint32

	//HeartbeatTime 心跳时间,utc 最大2038年
	HeartbeatTime uint32
	//Report 主机报告
	Report string
	//SourceHostID 源主机ID，应该是MonitorUnitID
	SourceHostID int32
}

//FromByteArray 把二进制值转换为对象
func (heartbeatRequest *HeartBeatRequest) FromByteArray(data []byte, index int) {

	// 主机类型
	heartbeatRequest.HostType = data[index]
	index += HostTypeLen

	// 心跳时间
	//dto.HeartbeatTime = NTPTimeConvert.FromBinary(BitConverter.ToUInt32(data, index));
	heartbeatRequest.HeartbeatTime = binary.LittleEndian.Uint32(data[index:])
	index += StartUpTimeLen

	// 报告字符串长度
	var length = (int)(data[index])
	index++
	// 报告
	heartbeatRequest.Report = (string)(data[index : index+length])
	heartbeatRequest.MessageType = 210

}

//ToByteArray 把传输对象转换为二进制数组
func (heartbeatRequest *HeartBeatRequest) ToByteArray() []byte {
	data := make([]byte, HostTypeLen+StartUpTimeLen)
	index := 0

	// 主机类型
	data[index] = heartbeatRequest.HostType
	index += HostTypeLen

	// 心跳时间
	//var ntpTime = NTPTimeConvert.ToBinary(HeartbeatTime);
	//Array.Copy(BitConverter.GetBytes(ntpTime), 0, data, index, HostInfo.STARTUPTIME_LEN);
	binary.LittleEndian.PutUint32(data[index:index+StartUpTimeLen], heartbeatRequest.HeartbeatTime)
	index += StartUpTimeLen

	// 报告长度
	reportBytes := ([]byte)(heartbeatRequest.Report)

	data = append(data, reportBytes...)

	data[index] = (byte)(len(reportBytes))
	index++

	return data
}

//ToString 把传输对象内的信息以字符串返回
func (heartbeatRequest *HeartBeatRequest) ToString() string {

	return fmt.Sprintf("HeartbeatRequest: HostType:%d HeartbeatTime:%d Report:%s",
		heartbeatRequest.HostType, heartbeatRequest.HeartbeatTime, heartbeatRequest.Report)
}
