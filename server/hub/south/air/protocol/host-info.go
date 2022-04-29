package protocol

import (
	"encoding/binary"
)

//HostInfo 主机信息
type HostInfo struct {
	//HostID 监控单元编号
	HostID int32
	//HostType 主机类型
	HostType int
	//MonitorUnitCode 监控单元编号
	MonitorUnitCode string
	//ConfigFileCode 主机配置MD5码
	ConfigFileCode string
	//SampleConfigCode 采集结构文件MD5码
	SampleConfigCode string
	//StartupTime 启动时间
	StartupTime int64
	//IsChildOfWorkStation 是否某个工作站下的主机
	IsChildOfWorkStation bool
	//WorkStationID 工作站编号
	WorkStationID int32
	//CommunicationState 主机状态(注意如果是挂在工作站下的主机，这个状态对DS无意义)
	CommunicationState int
	//URLAddress Url地址
	URLAddress string
	//PipeLinePairs 消息管道集合
	PipeLinePairs []*PipeLinePair
}

//FromByteArray 把二进制值转换为传输对象
func (hostInfo *HostInfo) FromByteArray(data []byte, index int) int32 {
	// 监控单元编码
	hostInfo.MonitorUnitCode = string(data[index : index+MonitorUnitCodeLen])
	index += MonitorUnitCodeLen
	// 主机类型
	hostInfo.HostType = (int)(data[index])
	index += HostTypeLen
	// 主机配置MD5码
	hostInfo.ConfigFileCode = string(data[index : index+ConfigFileCodeLen])
	index += ConfigFileCodeLen
	// 采集结构文件MD5码
	hostInfo.SampleConfigCode = string(data[index : index+SampleFileCodeLen])
	index += SampleFileCodeLen
	// 启动时间
	hostInfo.StartupTime = int64(binary.LittleEndian.Uint32(data[index:]))
	// StartupTime = NTPTimeConvert.FromBinary(BitConverter.ToUInt32(data, index));
	index += StartUpTimeLen
	// 服务地址集合长度
	itemCount := binary.LittleEndian.Uint16(data[index:])
	index += LengthLen

	len := 0
	hostInfo.PipeLinePairs = make([]*PipeLinePair, 0)
	// 服务地址集合
	for i := 0; i < (int)(itemCount); i++ {
		// 管道类型
		pipelineType := (byte)(data[index])
		index++
		// 服务器地址长度
		len = (int)(data[index])
		index++
		// 服务器地址
		url := string(data[index : index+len])
		index += len
		hostInfo.PipeLinePairs = append(hostInfo.PipeLinePairs, &PipeLinePair{PipeLineType: pipelineType, URL: url})
	}

	return (int32)(index)
}

//GetCheck  GetCheck
func (hostInfo *HostInfo) GetCheck(data []byte, index int, check []uint8) {

	for i := 0; i < MonitorUnitCodeLen; i++ {
		check[i] = (uint8)(data[index+i])
	}

	index += MonitorUnitCodeLen
	index += HostTypeLen
	index += ConfigFileCodeLen
	index += SampleFileCodeLen

	for j := 0; j < 4; j++ {
		check[MonitorUnitCodeLen+j] = (uint8)(data[index+j])
	}
}

//ToByteArray 把传输对象转换为二进制
func (hostInfo *HostInfo) ToByteArray() []byte {
	data := make([]byte, MonitorUnitCodeLen+HostTypeLen+ConfigFileCodeLen+SampleFileCodeLen+StartUpTimeLen+LengthLen)
	index := 0

	// 监控单元编码
	copy(data[index:index+MonitorUnitCodeLen], ([]byte)(hostInfo.MonitorUnitCode))
	index += MonitorUnitCodeLen
	// 主机类型
	data[index] = (byte)(hostInfo.HostType)
	index += HostTypeLen
	// 主机配置MD5码
	copy(data[index:index+ConfigFileCodeLen], ([]byte)(hostInfo.ConfigFileCode))
	index += ConfigFileCodeLen
	// 采集结构文件MD5码
	copy(data[index:index+SampleFileCodeLen], ([]byte)(hostInfo.SampleConfigCode))
	index += SampleFileCodeLen
	// 启动时间
	//var ntpTime = NTPTimeConvert.ToBinary(StartupTime);
	binary.LittleEndian.PutUint32(data[index:index+StartUpTimeLen], (uint32)(hostInfo.StartupTime))
	index += StartUpTimeLen
	// 服务地址集合
	binary.LittleEndian.PutUint16(data[index:index+LengthLen], (uint16)(len(hostInfo.PipeLinePairs)))
	index += LengthLen

	for _, pp := range hostInfo.PipeLinePairs {
		url := ([]byte)(pp.URL)
		len := len(url)
		data = append(data, pp.PipeLineType)
		data = append(data, (byte)(len))
		data = append([]byte{}, append(data, url...)...)
	}

	//Array.Copy(BitConverter.GetBytes((ushort)Pipelines.Count), 0, data, index, LENGTH_LEN);

	// var len = 0;
	// foreach (KeyValuePair<PipelineType, String> item in Pipelines)
	// {
	// 	var url = Encoding.ASCII.GetBytes(item.Value);
	// 	len = url.Length;
	// 	Array.Resize(ref data, index + 2 + len);
	// 	// 消息管道
	// 	data[index] = (byte)item.Key;
	// 	index += 1;
	// 	// 服务地址
	// 	data[index] = (byte)len;
	// 	index += 1;
	// 	url.CopyTo(data, index);
	// 	index += len;
	// }

	return data
}

//         /// <summary>
//         /// 复制
//         /// </summary>
//         public void Copy(HostInfo other)
//         {
//             this.HostType = other.HostType;
//             this.HostId = other.HostId;
//             this.ConfigFileCode = other.ConfigFileCode;
//             this.SampleConfigCode = other.SampleConfigCode;
//             this.StartupTime = other.StartupTime;
//             this.State = other.State;
//             ClonePipeline(other.Pipelines);
//         }

//         /// <summary>
//         /// 字符串形式返回对象信息
//         /// </summary>
//         /// <returns>字符串</returns>
// 		public override string ToString()
// 		{
//             var rtn = new StringBuilder("\r\n");

//             rtn.AppendFormat("{0, -10}{1, -12}{2, -33}{3, -33}{4, -21}{5}\r\n",
//                     "HostType", "HostId", "ConfigFileCode", "SampleConfigCode", "StartupTime", "HostIP");
//             rtn.AppendFormat("{0, -10}{1, -12}{2, -33}{3, -33}{4, -21}{5}\r\n",
//                 HostType, HostId, ConfigFileCode, SampleConfigCode, StartupTime, HostIP);
//             rtn.AppendFormat("{0, -30}{1}\r\n", "Key", "Url");
//             foreach (var pipe in Pipelines)
//             {
//                 rtn.AppendFormat("{0, -30}{1}\r\n", pipe.Key, pipe.Value);
//             }
//             return rtn.ToString();
//         }

//         /// <summary>
//         /// 判断2个传输对象是否相同
//         /// </summary>
//         /// <param name="obj">要判断的对象</param>
//         /// <returns>
//         /// true：相同
//         /// false：不同
//         /// </returns>
//         public override bool Equals(Object obj)
//         {
//             var other = (HostInfo)obj;

//             if (other == null)
//             {
//                 return false;
//             }

//             var rtn = true;

//             rtn = rtn && HostType == other.HostType;
//             rtn = rtn && HostId == other.HostId;
//             rtn = rtn && ConfigFileCode == other.ConfigFileCode;
//             rtn = rtn && SampleConfigCode == other.SampleConfigCode;
//             rtn = rtn && StartupTime.ToString() == other.StartupTime.ToString();
//             rtn = rtn && State == other.State;
//             rtn = rtn && HostIP == other.HostIP;

//             return rtn;
//         }
