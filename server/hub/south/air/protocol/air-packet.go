package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// <remarks>
// 报文格式如下所示：
//  0                   1                   2                   3
//  0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |         SOI = 0x7E6D          |        Sequence Number        |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |       |       |A|A|I|P|   | S |               |               |
// |Ver = 1|   PT  |C|K|C|A|   | E |  Message Type |      TTL      |
// |       |       |K|M|T|D|   | C |               |               |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// | Pipeline Type |                                               |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                         Source HostId                         |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                         Destination HostId                    |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |          Info Length          |            Checksum           |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                           Infomation                          |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// </remarks>
const (
	//HeadLength SWTP首部长度
	HeadLength = 24
	//SOI 16位起始标志位 UInt16 SOI = 0x7E6D
	SOI = 0x7E6D
	//SOIIndex 16位起始标志位 int SOI_INDEX = 0
	SOIIndex = 0
	//SOILen 16位起始标志位长度 int SOI_LEN = 2
	SOILen = 2
	//SequenceNumberIndex 16位序号 int SEQUENCENUMBER_INDEX = 2
	SequenceNumberIndex = 2
	//SequenceNumberLen 16位序号长度 int SEQUENCENUMBER_LEN = 2
	SequenceNumberLen = 2
	//VerPTIndex VerPT标志位  int VERPT_INDEX = 4
	VerPTIndex = 4
	//ACKIndex ACK标志位 int ACK_INDEX = 5
	ACKIndex = 5
	//MessageTypeIndex 消息类型索引 int MESSAGETYPE_INDEX = 6
	MessageTypeIndex = 6
	//TTLIndex TTL字段索引 int TTL_INDEX = 7
	TTLIndex = 7
	//PipeTypeIndex 消息管道类型索引 int PIPETYPE_INDEX = 8
	PipeTypeIndex = 8
	//SourceHostIDIndex Source HostId索引 int SOURCE_HOSTID_INDEX = 12
	SourceHostIDIndex = 12
	//SourceHostIDLen Source HostId长度 int SOURCE_HOSTID_LEN = 4
	SourceHostIDLen = 4
	//DestinationHostIDIndex  Destination HostId索引 int DESTINATION_HOSTID_INDEX = 16
	DestinationHostIDIndex = 16
	//DestinationHostIDLen Destination HostId长度 int DESTINATION_HOSTID_LEN = 4
	DestinationHostIDLen = 4
	//InfoLenIndex 信息域长度字段索引 int INFO_lEN_INDEX = 20
	InfoLenIndex = 20
	//InfoLenLen 信息域长度字段长度 int INFO_lEN_LEN = 2
	InfoLenLen = 2
	//CheckSumIndex 校验和字段索引 int CHECKSUM_INDEX = 22
	CheckSumIndex = 22
	//CheckSumLen 校验和字段长度 int CHECKSUM_LEN = 2
	CheckSumLen = 2
	//StationIDIndex STATIONID_INDEX字段索引 int STATIONID_INDEX = 24
	StationIDIndex = 24
	//HostIDIndex HOSTID字段索引 int HOSTID_INDEX = 28
	HostIDIndex = 28
	//EncodingIndex 编码类型字段索引 int ENCODING_INDEX = 32
	EncodingIndex = 32
	//RDNIndex HOSTID字段索引 int RDN_INDEX = 51
	RDNIndex = 51
	//MaxTTL int MaxTtl = 50
	MaxTTL = 50
)

//AirPacket Air传输协议报文
type AirPacket struct {
	version byte
	//TTL 获取或设置数据包的生存期
	TTL byte
	//SourceHostID 获取或设置发送方主机Id
	SourceHostID int32
	//DestinationHostID 获取或设置接受方主机Id
	DestinationHostID int32
	//RawPacket 获取或设置原始数据包
	RawPacket []byte
	//BakPacket liangyongjun 校验用
	BakPacket [3]int8
	//SequenceNumber 获取序列号  UInt16 SequenceNumber
	SequenceNumber uint16
	//RequireAcknowledge 获取或设置数据包是否需要确认
	RequireAcknowledge bool
	//AutoAcknowledge 获取或设置数据包是否自动确认
	AutoAcknowledge bool
	//Padding 获取数据包是否有填充字段
	Padding bool
	//IsBusy 是否忙
	IsBusy bool
	//Compressed 获取或设置数据包是否压缩
	Compressed bool
	//MessageType 获取或设置数据包的消息类型
	MessageType int
	//EncryptType 获取或设置数据包的加密类型
	EncryptType int
	//PipelineType 获取或设置消息管道类型
	PipelineType int
	//ProtocolType 获取或设置数据包的协议类型
	ProtocolType int
	//DataTransferObject 获取或设置数据传输对象
	RawObject interface{} //public DataTransferObject RawObject { get; set; }
}

//NewAirPacket 创建Air协议包
func NewAirPacket(data []byte) *AirPacket {
	packet := &AirPacket{
		version:           1,
		SourceHostID:      0,
		DestinationHostID: 0,
		RawPacket:         data,
		BakPacket:         [3]int8{0, 0, 0}, //BakPacket = new sbyte[3];
	}

	packet.parseHeader()

	return packet
}

//IsAirPacket 判断数据包是否合法
func IsAirPacket(RawPacket []byte) bool {
	pktLen := len(RawPacket)
	if pktLen >= HeadLength {
		// 若数据报的长度大于首部长度

		// TODO: 检验校验和
		checksum := binary.LittleEndian.Uint16(RawPacket[CheckSumIndex:])
		// 原始数据包的校验和置为0
		RawPacket[CheckSumIndex] = 0
		RawPacket[CheckSumIndex+1] = 0

		sum := getCheckSum(RawPacket)

		// 还原原始数据包的校验和
		binary.LittleEndian.PutUint16(RawPacket[CheckSumIndex:CheckSumIndex+CheckSumLen], checksum)
		// Array.Copy(BitConverter.GetBytes(checksum), 0, RawPacket, CheckSumIndex, CheckSumLen)

		if sum != checksum {
			fmt.Println("Packet checkout fail required {0} send {1}", checksum, sum)
			return false
		}
	} else {
		fmt.Println("not siteweb packet")
		return false
	}

	return true
}

func (packet *AirPacket) parseHeader() {
	index := SequenceNumberIndex
	packet.SequenceNumber = binary.LittleEndian.Uint16(packet.RawPacket[index:])
	packet.version = (byte)((packet.RawPacket[index+2] & 0xF0) >> 4)
	packet.ProtocolType = (int)(packet.RawPacket[index+2] & 0xF)

	if (packet.RawPacket[index+3] & 0x80) == 0x80 {
		packet.RequireAcknowledge = true
	} else {
		packet.RequireAcknowledge = false
	}

	if (packet.RawPacket[index+3] & 0x40) == 0x40 {
		packet.AutoAcknowledge = true
	} else {
		packet.AutoAcknowledge = false
	}

	if (packet.RawPacket[index+3] & 0x20) == 0x20 {
		packet.Compressed = true
	} else {
		packet.Compressed = false
	}

	if (packet.RawPacket[index+3] & 0x10) == 0x10 {
		packet.Padding = true
	} else {
		packet.Padding = false
	}

	if (packet.RawPacket[index+3] & 0x0) == 0x0 {
		packet.IsBusy = true
	} else {
		packet.IsBusy = false
	}

	packet.EncryptType = (int)((packet.RawPacket[index+3] & 0x3))
	packet.MessageType = (int)(packet.RawPacket[index+4])
	packet.TTL = packet.RawPacket[TTLIndex]
	packet.PipelineType = (int)(packet.RawPacket[index+6])

	//liangyj 20170911
	packet.BakPacket[0] = (int8)(packet.RawPacket[index+7])
	packet.BakPacket[1] = (int8)(packet.RawPacket[index+8])
	packet.BakPacket[2] = (int8)(packet.RawPacket[index+9])

	bufHost := bytes.NewBuffer(packet.RawPacket[SourceHostIDIndex : SourceHostIDIndex+SourceHostIDLen])
	var sourceHostID int32
	binary.Read(bufHost, binary.LittleEndian, &sourceHostID)
	packet.SourceHostID = sourceHostID

	bufDest := bytes.NewBuffer(packet.RawPacket[DestinationHostIDIndex : DestinationHostIDIndex+DestinationHostIDLen])
	var destHostID int32
	binary.Read(bufDest, binary.LittleEndian, &destHostID)
	packet.DestinationHostID = destHostID
}

//获取ark的接受地址
func GetPipeLines() []*PipeLinePair {
	pp := make([]*PipeLinePair, 0)
	pp = append(pp, &PipeLinePair{PipeLineType: 0, URL: global.GVA_VP.GetString("diagnostics-url")})
	pp = append(pp, &PipeLinePair{PipeLineType: 1, URL: global.GVA_VP.GetString("subscibe-url")})
	pp = append(pp, &PipeLinePair{PipeLineType: 5, URL: global.GVA_VP.GetString("publish-url")})
	pp = append(pp, &PipeLinePair{PipeLineType: 6, URL: global.GVA_VP.GetString("event-publish-url")})
	pp = append(pp, &PipeLinePair{PipeLineType: 7, URL: global.GVA_VP.GetString("realsignal-publish-url")})
	pp = append(pp, &PipeLinePair{PipeLineType: 8, URL: global.GVA_VP.GetString("historysignal-pubish-url")})
	return pp
}

func (packet *AirPacket) Pack() []byte {
	data := make([]byte, HeadLength)

	//设置SOI
	index := SOIIndex
	binary.LittleEndian.PutUint16(data[index:index+SOILen], (uint16)(SOI))

	index = SequenceNumberIndex
	//设置 序列号
	binary.LittleEndian.PutUint16(data[index:index+SequenceNumberLen], (uint16)(packet.SequenceNumber))

	//设置协议版本 目前是固定
	var version byte = 1
	var aByte byte = 0
	aByte |= (byte)(version << 4)
	aByte |= (byte)(packet.ProtocolType)
	data[VerPTIndex] = aByte

	aByte = 0
	aByte |= Ternary(0x80, 0x0, packet.RequireAcknowledge)
	aByte |= Ternary(0x40, 0x0, packet.AutoAcknowledge)
	aByte |= Ternary(0x20, 0x0, packet.Compressed)
	aByte |= Ternary(0x10, 0x0, packet.Padding)
	aByte |= (byte)(packet.EncryptType)
	data[ACKIndex] = aByte

	data[MessageTypeIndex] = (byte)(packet.MessageType)
	data[TTLIndex] = packet.TTL
	data[PipeTypeIndex] = (byte)(packet.PipelineType)

	index = SourceHostIDIndex
	binary.LittleEndian.PutUint32(data[index:index+SourceHostIDLen], (uint32)(packet.SourceHostID))
	index = DestinationHostIDIndex
	binary.LittleEndian.PutUint32(data[index:index+DestinationHostIDLen], (uint32)(packet.DestinationHostID))
	data[InfoLenIndex] = (byte)(len(packet.RawPacket))
	data = append(data, packet.RawPacket...)

	index = CheckSumIndex
	binary.LittleEndian.PutUint16(data[index:index+CheckSumLen], (uint16)(getCheckSum(data)))
	return data

}

func Ternary(a, b byte, con bool) byte {
	if con {
		return a
	}
	return b
}

func getCheckSum(data []byte) uint16 {
	var checksum uint16 = 0
	for i := 2; i < len(data); i++ {
		checksum = checksum + uint16(data[i])
	}

	return checksum
}

// 		//  /// <summary>
// 		//  /// 填充TTL字段，并增量计算checksum，更新原始数据
// 		//  /// </summary>
// 		//  /// <param name="ttl">TTL字段</param>
// 		//  public void FillTtl(byte ttl)
// 		//  {
// 		// 	 Debug.Assert(RawPacket != null);

// 		// 	 UInt16 checksum = BitConverter.ToUInt16(RawPacket, CHECKSUM_INDEX);
// 		// 	 checksum = (UInt16)((checksum + ttl - RawPacket[TTL_INDEX]) % 65536);

// 		// 	 RawPacket[TTL_INDEX] = ttl;
// 		// 	 Array.Copy(BitConverter.GetBytes(checksum), 0, RawPacket, CHECKSUM_INDEX, 2);
// 		// 	 TTL = ttl;
// 		//  }

// 		//  /// <summary>
// 		//  /// 填充动态配置更新字段，并增量计算checksum，更新原始数据
// 		//  /// </summary>
// 		//  /// <param name="ttl">动态配置更新字段</param>
// 		// 	public void FillDynamicConfig(int hostId, int stationId, string rdn)
// 		//  {
// 		// 	 Debug.Assert(RawPacket != null);

// 		// 	 byte[] newHostId = BitConverter.GetBytes(hostId);
// 		// 	 int incHostId = newHostId[0] + newHostId[1] + newHostId[2] + newHostId[3]
// 		// 		 - RawPacket[HOSTID_INDEX]
// 		// 		 - RawPacket[HOSTID_INDEX + 1]
// 		// 		 - RawPacket[HOSTID_INDEX + 2]
// 		// 		 - RawPacket[HOSTID_INDEX + 3];

// 		// 	 byte[] newStationId = BitConverter.GetBytes(stationId);
// 		// 	 int incStationId = newStationId[0] + newStationId[1] + newStationId[2] + newStationId[3]
// 		// 		 - RawPacket[STATIONID_INDEX]
// 		// 		 - RawPacket[STATIONID_INDEX + 1]
// 		// 		 - RawPacket[STATIONID_INDEX + 2]
// 		// 		 - RawPacket[STATIONID_INDEX + 3];

// 		// 	 byte[] newRdn = Encoding.ASCII.GetBytes(rdn);

// 		// 	 int sum = 0;

// 		// 	 for(int i = 0; i < newRdn.Length; i++)
// 		// 	 {
// 		// 		 sum += newRdn[i] - RawPacket[RDN_INDEX];
// 		// 	 }

// 		// 	 UInt16 checksum = BitConverter.ToUInt16(RawPacket, CHECKSUM_INDEX);
// 		// 	 checksum = (UInt16)((checksum + incHostId + incStationId) % 65536);

// 		// 	 Array.Copy(newHostId, 0, RawPacket, HOSTID_INDEX, 4);
// 		// 	 Array.Copy(newStationId, 0, RawPacket, STATIONID_INDEX, 4);
// 		// 	 Array.Copy(newRdn, 0, RawPacket, RDN_INDEX, newRdn.Length);

// 		// 	 Array.Copy(BitConverter.GetBytes(checksum), 0, RawPacket, CHECKSUM_INDEX, 2);
// 		//  }

// 		//  /// <summary>
// 		//  /// 填充SourceHostId, DesHostId字段，并增量计算checksum，更新原始数据
// 		//  /// 只是改了类型，没有修改算法
// 		//  /// 要修改这个算法
// 		//  /// </summary>
// 		//  /// <param name="sourceHostId">发送方主机编号</param>
// 		//  /// <param name="desHostId">接受方主机编号</param>
// 		//  public void FillSourceDesHostId(Int32 sourceHostId, Int32 desHostId)
// 		//  {
// 		// 	 Debug.Assert(RawPacket != null);

// 		// 	 byte[] newHostId = BitConverter.GetBytes(sourceHostId);
// 		// 	 int inc = newHostId[0] + newHostId[1] + newHostId[2] + newHostId[3]
// 		// 		 - RawPacket[SOURCE_HOSTID_INDEX]
// 		// 		 - RawPacket[SOURCE_HOSTID_INDEX + 1]
// 		// 		 - RawPacket[SOURCE_HOSTID_INDEX + 2]
// 		// 		 - RawPacket[SOURCE_HOSTID_INDEX + 3];

// 		// 	 byte[] newDesHostId = BitConverter.GetBytes(desHostId);
// 		// 	 int incDes = newDesHostId[0] + newDesHostId[1] + newDesHostId[2] + newDesHostId[3]
// 		// 		 - RawPacket[DESTINATION_HOSTID_INDEX]
// 		// 		 - RawPacket[DESTINATION_HOSTID_INDEX + 1]
// 		// 		 - RawPacket[DESTINATION_HOSTID_INDEX + 2]
// 		// 		 - RawPacket[DESTINATION_HOSTID_INDEX + 3];

// 		// 	 UInt16 checksum = BitConverter.ToUInt16(RawPacket, CHECKSUM_INDEX);
// 		// 	 checksum = (UInt16)((checksum + inc + incDes) % 65536);

// 		// 	 Array.Copy(BitConverter.GetBytes(sourceHostId), 0, RawPacket, SOURCE_HOSTID_INDEX, SOURCE_HOSTID_LEN);
// 		// 	 Array.Copy(BitConverter.GetBytes(desHostId), 0, RawPacket, DESTINATION_HOSTID_INDEX, DESTINATION_HOSTID_LEN);
// 		// 	 Array.Copy(BitConverter.GetBytes(checksum), 0, RawPacket, CHECKSUM_INDEX, 2);

// 		// 	 SourceHostId = sourceHostId;
// 		// 	 DestinationHostId = desHostId;
// 		//  }

// 		 ///// <summary>
// 		 ///// 创建SiteWeb协议的数据类数据包
// 		 ///// </summary>
// 		 ///// <param name="dto">协议包实体</param>
// 		 ///// <param name="encoding">协议包实体的编码方式</param>
// 		 ///// <param name="pipelineType">协议首部字段：消息管道类型</param>
// 		 ///// <returns>已创建的数据类数据包</returns>
// 		 //public static SiteWebPacket CreateDataPacket(DataTransferObject dto,
// 		 //                                SiteWebEncoding encoding,
// 		 //                                PipelineType pipelineType)
// 		 //{
// 		 //    var packet = CreateDataPacket(dto, encoding, ProtocolType.Swdp,
// 		 //                                 false, false, false,
// 		 //                                 EncryptType.NonEncrypt, 0,
// 		 //                                 pipelineType);

// 		 //    return packet;
// 		 //}

// 		 /// <summary>
// 		 /// 创建SiteWeb协议的数据类数据包
// 		 /// </summary>
// 		 /// <param name="dto">协议包实体</param>
// 		 /// <param name="encoding">协议包实体的编码方式</param>
// 		 /// <param name="requireAcknowledge">协议首部字段：是否需要确认</param>
// 		 /// <param name="autoAcknowledge">协议首部字段：是否自动确认</param>
// 		 /// <param name="compressed">协议首部字段：是否压缩</param>
// 		 /// <param name="encryptType">协议首部字段：加密类型</param>
// 		 /// <param name="ttl">协议首部字段：TTL</param>
// 		 /// <param name="pipelineType">协议首部字段：消息管道类型</param>
// 		 /// <returns>已创建的数据类数据包</returns>
// 		 public static SiteWebPacket CreateDataPacket(DataTransferObject dto,
// 										 SiteWebEncoding encoding,
// 										 bool requireAcknowledge,
// 										 byte ttl,
// 										 PipelineType pipelineType)
// 		 {
// 			 var packet = CreateDataPacket(dto, encoding, ProtocolType.Swdp,
// 										  requireAcknowledge, false, false,
// 										  EncryptType.NonEncrypt, ttl,
// 										  pipelineType);

// 			 return packet;
// 		 }

// 		 /// <summary>
// 		 /// 创建SiteWeb协议数据包
// 		 /// </summary>
// 		 /// <param name="dto">协议包实体</param>
// 		 /// <param name="encoding">协议包实体的编码方式</param>
// 		 /// <param name="protocolType">协议首部字段：协议类型</param>
// 		 /// <param name="requireAcknowledge">协议首部字段：是否需要确认</param>
// 		 /// <param name="autoAcknowledge">协议首部字段：是否自动确认</param>
// 		 /// <param name="compressed">协议首部字段：是否压缩</param>
// 		 /// <param name="encryptType">协议首部字段：加密类型</param>
// 		 /// <param name="ttl">协议首部字段：TTL</param>
// 		 /// <param name="pipelineType">协议首部字段：消息管道类型</param>
// 		 /// <returns>已创建的数据包</returns>
// 		 public static SiteWebPacket CreateDataPacket(DataTransferObject dto,
// 										 SiteWebEncoding encoding,
// 										 ProtocolType protocolType,
// 										 bool requireAcknowledge,
// 										 bool autoAcknowledge,
// 										 bool compressed,
// 										 EncryptType encryptType,
// 										 byte ttl,
// 										 PipelineType pipelineType)
// 		 {
// 			 Debug.Assert(dto != null);

// 			 var packet = new SiteWebPacket();

// 			 // 设置协议首部字段
// 			 packet.ProtocolType = protocolType;
// 			 packet.RequireAcknowledge = requireAcknowledge;
// 			 packet.AutoAcknowledge = autoAcknowledge;
// 			 packet.Compressed = compressed;
// 			 packet.EncryptType = encryptType;
// 			 packet.TTL = ttl;
// 			 packet.PipelineType = pipelineType;

// 			 // 装载消息实体
// 			 DataMessage message = new DataMessage(encoding);
// 			 message.LoadDataTransferObject(dto);

// 			 packet.Pack(message);

// 			 return packet;
// 		 }

// 		 /// <summary>
// 		 /// 创建注册响应包
// 		 /// </summary>
// 		 /// <param name="hostId">主机ID</param>
// 		 /// <param name="resultCode">注册结果代码</param>
// 		 /// <param name="token">通信令牌</param>
// 		 /// <param name="license">license字段内容</param>
// 		 /// <returns>>注册响应数据包</returns>
// 		 public static SiteWebPacket CreateRegisterResponsePacket(ushort sequenceNubmer,  HostType hostType,
// 			 ResultCode resultCode, AddressFamily family)
// 		 {
// 			 var response = new RegisterResponse();
// 			 response.ResultCode = resultCode;
// 			 response.Pipelines = family == AddressFamily.InterNetwork
// 				 ? Communication.Communicator.Instance.Pipelines
// 				 : Communication.Communicator.Instance.PipelinesV6;

// 			 return CreateCommunicationPacket(sequenceNubmer, response);
// 		 }

// 		 /// <summary>
// 		 /// 创建心跳请求包
// 		 /// </summary>
// 		 /// <param name="monitorUnitCode">监控单元编码</param>
// 		 /// <param name="hostType">主机类型</param>
// 		 /// <param name="hostVersion">主机版本</param>
// 		 /// <param name="heartbeatTime">心跳（发送）时间</param>
// 		 /// <param name="payload">综合负载值</param>
// 		 /// <param name="report">主机信息报告</param>
// 		 /// <returns>已创建的心跳请求包</returns>
// 		 public static SiteWebPacket CreateHeartbeatRequestPacket(HostType hostType, HostVersion hostVersion,
// 			 DateTime heartbeatTime, float payload, string report)
// 		 {
// 			 var request = new HeartbeatRequest();

// 			 request.HostType = hostType;
// 			 request.HeartbeatTime = heartbeatTime;
// 			 request.Report = report;

// 			 return CreateCommunicationPacket(request);
// 		 }

// 		 /// <summary>
// 		 /// 创建心跳响应包
// 		 /// </summary>
// 		 /// <param name="hostId">主机ID</param>
// 		 /// <param name="hostType">主机类型</param>
// 		 /// <param name="heartbeatTime">心跳（发送）时间</param>
// 		 /// <returns>已创建的心跳响应包</returns>
// 		 public static SiteWebPacket CreateHeartbeatResponsePacket(DateTime heartbeatTime)
// 		 {
// 			 var response = new HeartbeatResponse();

// 			 response.HeartbeatTime = heartbeatTime;

// 			 return CreateCommunicationPacket(response);
// 		 }

// 		 /// <summary>
// 		 /// 创建接入允许通知包
// 		 /// </summary>
// 		 /// <param name="resultCode">结果(1:允许接入非自产设备 0:不允许)</param>
// 		 /// <param name="description">描述</param>
// 		 /// <returns>已创建的接入允许通知包</returns>
// 		 public static SiteWebPacket CreateAccessConfigNotifyResponsePacket(Boolean resultCode, String description)
// 		 {
// 			 var response = new AccessConfigNotifyResponse();

// 			 response.ResultCode = resultCode;
// 			 response.ResultDescription = description;

// 			 return CreateCommunicationPacket(response);
// 		 }

// 		 /// <summary>
// 		 /// 创建SiteWeb协议的通信控制类数据包
// 		 /// </summary>
// 		 /// <param name="dto">协议包实体</param>
// 		 /// <param name="requireAcknowledge">数据包是否需要确认</param>
// 		 /// <param name="ttl">数据包的生存时间</param>
// 		 /// <returns>已创建的通信控制数据包</returns>
// 		 private static SiteWebPacket CreateCommunicationPacket(DataTransferObject dto,
// 										 bool requireAcknowledge,
// 										 int ttl)
// 		 {
// 			 Debug.Assert(dto != null);

// 			 var packet = new SiteWebPacket();

// 			 // 设置协议首部字段
// 			 packet.RequireAcknowledge = requireAcknowledge;
// 			 packet.TTL = (byte)ttl;
// 			 packet.ProtocolType = ProtocolType.Swcp;
// 			 packet.PipelineType = PipelineType.Diagnostics;

// 			 // 装载消息实体
// 			 var message = new CommunicationMessage();
// 			 message.LoadDataTransferObject(dto);

// 			 packet.Pack(message);

// 			 return packet;
// 		 }

// 		 /// <summary>
// 		 /// 创建SiteWeb协议的通信控制类数据包
// 		 /// </summary>
// 		 /// <param name="dto">协议包实体</param>
// 		 /// <returns>已创建的通信控制数据包</returns>
// 		 private static SiteWebPacket CreateCommunicationPacket(UInt16 sequenceNumber, DataTransferObject dto)
// 		 {
// 			 Debug.Assert(dto != null);

// 			 var packet = new SiteWebPacket();

// 			 // 设置协议首部字段
// 			 // 要求带上包序号，所有RequireAcknowledge设置为true
// 			 packet.RequireAcknowledge = true;
// 			 packet.ProtocolType = ProtocolType.Swcp;
// 			 packet.PipelineType = PipelineType.Diagnostics;

// 			 // 装载消息实体
// 			 CommunicationMessage message = new CommunicationMessage();
// 			 message.LoadDataTransferObject(dto);

// 			 packet.Pack(message, sequenceNumber);

// 			 return packet;
// 		 }

// 		 /// <summary>
// 		 /// 创建SiteWeb协议的通信控制类数据包
// 		 /// </summary>
// 		 /// <param name="dto">协议包实体</param>
// 		 /// <returns>已创建的通信控制数据包</returns>
// 		 private static SiteWebPacket CreateCommunicationPacket(DataTransferObject dto)
// 		 {
// 			 Debug.Assert(dto != null);

// 			 var packet = new SiteWebPacket();

// 			 // 设置协议首部字段
// 			 packet.ProtocolType = ProtocolType.Swcp;
// 			 packet.PipelineType = PipelineType.Diagnostics;

// 			 // 装载消息实体
// 			 CommunicationMessage message = new CommunicationMessage();
// 			 message.LoadDataTransferObject(dto);

// 			 packet.Pack(message);

// 			 return packet;
// 		 }
// 		 #endregion

// 		 /// <summary>
// 		 /// 确定指定的 Object 是否等于当前的 Object
// 		 /// </summary>
// 		 /// <param name="obj">与当前的 Object 进行比较的 Object</param>
// 		 /// <returns>如果指定的 Object 等于当前的 Object，则为 true；否则为 false</returns>
// 		 public override bool Equals(Object obj)
// 		 {
// 			 // Check for null and compare run-time types.
// 			 if (obj == null || GetType() != obj.GetType()) return false;

// 			 SiteWebPacket other = (SiteWebPacket)obj;

// 			 return (SequenceNumber == other.SequenceNumber)
// 						 && (m_version == other.m_version)
// 						 && (RequireAcknowledge == other.RequireAcknowledge)
// 						 && (AutoAcknowledge == other.AutoAcknowledge)
// 						 && (Compressed == other.Compressed)
// 						 && (Padding == other.Padding)
// 						 && (EncryptType == other.EncryptType)
// 						 && (ProtocolType == other.ProtocolType)
// 						 && (m_messageType == other.m_messageType)
// 						 && (TTL == other.TTL)
// 						 && (SourceHostId == other.SourceHostId)
// 						 && (DestinationHostId == other.DestinationHostId)
// 						 && (PipelineType == other.PipelineType);
// 		 }

// 		 /// <summary>
// 		 /// 获取数据包的哈希码
// 		 /// </summary>
// 		 /// <returns>哈希码</returns>
// 		 public override int GetHashCode()
// 		 {
// 			 return SequenceNumber;
// 		 }

// 		 /// <summary>
// 		 /// 把协议包内的信息以字符串返回
// 		 /// </summary>
// 		 /// <returns>String，表示当前的协议报文</returns>
// 		 public override string ToString()
// 		 {
// 			 var rtn = new StringBuilder("\r\n");

// 			 if (ProtocolType == ProtocolType.Acknowledge)
// 			 {
// 				 rtn.AppendFormat("{0, -10}{1, -16}{2, -12}{3, -12}{4, -6}{5}\r\n",
// 					 "SN", "PT", "SrcHostId", "DesHostId", "ACK", "IsBusy");
// 				 rtn.AppendFormat("{0, -10}{1, -16}{2, -12}{3, -12}{4, -6}{5}\r\n",
// 					 SequenceNumber, ProtocolType, SourceHostId,
// 					 DestinationHostId, RequireAcknowledge, IsBusy);

// 			 }
// 			 else
// 			 {
// 				 rtn.AppendFormat("{0, -10}{1, -6}{2, -12}{3, -12}{4, -6}{5, -23}{6, -6}{7}\r\n",
// 					 "SN", "PT", "SrcHostId", "DesHostId", "ACK", "MessageType", "TTL", "PipelineType");
// 				 rtn.AppendFormat("{0, -10}{1, -6}{2, -12}{3, -12}{4, -6}{5, -23}{6, -6}{7}\r\n",
// 					 SequenceNumber, ProtocolType, SourceHostId,
// 					 DestinationHostId,RequireAcknowledge, m_messageType,
// 					 TTL, PipelineType);
// 			 }

// 			 if (RawObject != null)
// 				 rtn.Append(RawObject.ToString());

// 			 return rtn.ToString();
// 		 }

// 		 /// <summary>
// 		 /// 计算校验和
// 		 /// </summary>
// 		 /// <returns>校验和</returns>
// 		 private UInt16 Checksum()
// 		 {
// 			 UInt16 checksum = 0;
// 			 for (int i = 2; i < RawPacket.Length; i++)
// 			 {
// 				 checksum += RawPacket[i];
// 			 }

// 			 return checksum;
// 		 }

// 		 static public bool CalcuDataCheck(sbyte[] szBak, sbyte[] szData)
// 		 {
// 			 sbyte[] szCheckRuslt = new sbyte[4];

// 			 CalcuDataCheck(ref szCheckRuslt, szData, 36);

// 			 if (szBak[0] == szCheckRuslt[0] &&
// 				 szBak[1] == szCheckRuslt[1] &&
// 					 szBak[2] == szCheckRuslt[2])
// 			 {
// 				 return true;
// 			 }
// 			 return false;
// 		 }

// 		 static void CalcuDataCheck(ref sbyte[] szResult, sbyte[] szData, int nLen)
// 		 {
// 			 int i = 0;
// 			 ushort wCheck = 0;
// 			 for (i = 0; i < nLen; i++)
// 			 {
// 				 szResult[0] = (sbyte)((sbyte)szResult[0] ^ (sbyte)szData[i]);
// 				 wCheck += (ushort)((byte)szData[i]);
// 			 }
// 			 szResult[1] = (sbyte)((wCheck / 0x100) & 0xFF);
// 			 szResult[2] = (sbyte)((wCheck % 0x100) & 0xFF);
// 		 }

// 		 #region IPacket Members

// 		 /
// 		 /// <summary>
// 		 /// 把消息封装为数据包
// 		 /// </summary>
// 		 /// <param name="message">待封装的消息</param>
// 		 /// <returns>数据包</returns>
// 		 public byte[] Pack(SiteWebMessage message)
// 		 {
// 			 return Pack(message, SequenceNumberGenerator.Instance.Gennerate());
// 		 }

// 		 /// <summary>
// 		 /// 把消息封装为数据包
// 		 /// </summary>
// 		 /// <param name="message">待封装的消息</param>
// 		 /// <returns>数据包</returns>
// 		 public byte[] Pack(SiteWebMessage message, UInt16 sequenceNumber)
// 		 {
// 			 Debug.Assert(message != null);

// 			 if (message is DataMessage)
// 			 {
// 				 ProtocolType = ProtocolType.Swdp;
// 			 }
// 			 else if (message is CommunicationMessage)
// 			 {
// 				 ProtocolType = ProtocolType.Swcp;
// 			 }

// 			 RawObject = message.RawObject;
// 			 byte[] info = message.Serialize();

// 			 int infoLen = info.Length;
// 			 RawPacket = new byte[HEAD_LENGTH + infoLen];

// 			 // 封装协议首部
// 			 Array.Copy(BitConverter.GetBytes(SOI), 0, RawPacket, SOI_INDEX, SOI_LEN);
// 			 SequenceNumber = sequenceNumber;

// 			 if (RequireAcknowledge)	// 若数据包需要确认，则产生SN
// 			 {
// 				 Array.Copy(BitConverter.GetBytes(SequenceNumber), 0, RawPacket, SEQUENCENUMBER_INDEX, SEQUENCENUMBER_LEN);
// 			 }

// 			 byte aByte = 0;
// 			 aByte |= (byte)(m_version << 4);
// 			 aByte |= (byte)ProtocolType;
// 			 RawPacket[VERPT_INDEX] = aByte;

// 			 aByte = 0;
// 			 aByte |= (byte)(RequireAcknowledge ? 0x80 : 0x0);
// 			 aByte |= (byte)(AutoAcknowledge ? 0x40 : 0x0);
// 			 aByte |= (byte)(Compressed ? 0x20 : 0x0);
// 			 aByte |= (byte)(Padding ? 0x10 : 0x0);
// 			 aByte |= (byte)((byte)EncryptType);
// 			 RawPacket[ACK_INDEX] = aByte;

// 			 m_messageType = message.MessageType;
// 			 RawPacket[MESSAGETYPE_INDEX] = (byte)m_messageType;
// 			 RawPacket[TTL_INDEX] = TTL;
// 			 RawPacket[PIPETYPE_INDEX] = (byte)PipelineType;
// 			 Array.Copy(BitConverter.GetBytes(SourceHostId), 0, RawPacket, SOURCE_HOSTID_INDEX, SOURCE_HOSTID_LEN);
// 			 Array.Copy(BitConverter.GetBytes(DestinationHostId), 0, RawPacket, DESTINATION_HOSTID_INDEX, DESTINATION_HOSTID_LEN);
// 			 Array.Copy(BitConverter.GetBytes(infoLen), 0, RawPacket, INFO_lEN_INDEX, INFO_lEN_LEN);

// 			 info.CopyTo(RawPacket, HEAD_LENGTH);

// 			 // TODO: 压缩和加密信息域

// 			 // TODO: 计算校验和，并填充该字段
// 			 Array.Copy(BitConverter.GetBytes(Checksum()), 0, RawPacket, CHECKSUM_INDEX, CHECKSUM_LEN);

// 			 return RawPacket;
// 		 }

// 	 }
//  }
