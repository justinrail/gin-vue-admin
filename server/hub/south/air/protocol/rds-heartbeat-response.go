package protocol

import "encoding/binary"

type RDSHeartBeatResponse struct {
	MessageType uint32

	HeartbeatTime uint32

	RealtimePort int

	RealtimeInterval int
}

//ToByteArray 把传输对象转换为二进制数组
func (rdsHeartBeatResponse *RDSHeartBeatResponse) ToByteArray() []byte {
	data := make([]byte, StartUpTimeLen+PortLen+IntervalLen)
	index := 0

	// 心跳时间
	//var ntpTime = NTPTimeConvert.ToBinary(HeartbeatTime);
	//Array.Copy(BitConverter.GetBytes(ntpTime), 0, data, index, HostInfo.STARTUPTIME_LEN);
	binary.LittleEndian.PutUint32(data[index:index+StartUpTimeLen], rdsHeartBeatResponse.HeartbeatTime)
	index += StartUpTimeLen
	binary.LittleEndian.PutUint32(data[index:index+PortLen], (uint32)(rdsHeartBeatResponse.RealtimePort))
	index += PortLen
	binary.LittleEndian.PutUint32(data[index:index+IntervalLen], (uint32)(rdsHeartBeatResponse.RealtimeInterval))
	return data
}
