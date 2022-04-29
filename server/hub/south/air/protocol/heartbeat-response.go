package protocol

import "encoding/binary"

//HeartbeatRequest 心跳相应包
type HeartBeatResponse struct {
	MessageType uint32

	HeartbeatTime uint32
}

//FromByteArray 把二进制值转换为对象
func (heartbeatResponse *HeartBeatResponse) FromByteArray(data []byte, index int) {

	// 心跳时间
	//dto.HeartbeatTime = NTPTimeConvert.FromBinary(BitConverter.ToUInt32(data, index));
	heartbeatResponse.HeartbeatTime = binary.LittleEndian.Uint32(data[index:])

	heartbeatResponse.MessageType = 211

}

//ToByteArray 把传输对象转换为二进制数组
func (heartbeatResponse *HeartBeatResponse) ToByteArray() []byte {
	data := make([]byte, StartUpTimeLen)
	index := 0

	// 心跳时间
	//var ntpTime = NTPTimeConvert.ToBinary(HeartbeatTime);
	//Array.Copy(BitConverter.GetBytes(ntpTime), 0, data, index, HostInfo.STARTUPTIME_LEN);
	binary.LittleEndian.PutUint32(data[index:], heartbeatResponse.HeartbeatTime)
	return data
}
