package base

const (
	//PacketCOG Packet type is COG
	PacketCOG = iota
	//PacketCOD Packet type is COD
	PacketCOD
	//PacketCOV Packet type is COV
	PacketCOV
	//PacketCOA Packet type is COA
	PacketCOA
	//PacketCmdAck Packet type is CmdAck
	PacketCmdAck
)

const (
	//GatewayFlagUnkown 未知Flag
	GatewayFlagUnkown = iota
	//GatewayFlagRegister COG register event flag
	GatewayFlagRegister
	//GatewayFlagTimeCheckAck gateway时间同步返回
	GatewayFlagTimeCheckAck
	//GatewayFlagHeartbeat gateway 心跳消息
	GatewayFlagHeartbeat
	//GatewayFlagRebootAck gateway重启返回
	GatewayFlagRebootAck
	//GatewayFlagGetConfigAck gateway获取配置返回
	GatewayFlagGetConfigAck
	//GatewayFlagSetConfigAck gateway设置配置返回
	GatewayFlagSetConfigAck
	//GatewayFlagSendConfig gateway上送配置
	GatewayFlagSendConfig
	//GatewayFlagFTPAck gateway FTP信息同步返回
	GatewayFlagFTPAck
	//GatewayFlagAlarm gateway gateway上送告警
	GatewayFlagAlarm
	//GatewayFlagGetDataAck 对gateway请求实时数据返回
	GatewayFlagGetDataAck
	//GatewayFlagSetPointAck 对gateway设置测点返回
	GatewayFlagSetPointAck
	//GatewayFlagGetThresholdAck 对gateway获取测点门限
	GatewayFlagGetThresholdAck
	//GatewayFlagSetThresholdAck 对gateway设置测点门限
	GatewayFlagSetThresholdAck
	//GatewayFlagSetStorageRuleAck 设置存储规则返回
	GatewayFlagSetStorageRuleAck
	//GatewayFlagGetStorageRuleAck 读取存储规则返回
	GatewayFlagGetStorageRuleAck
	//GatewayFlagGetInfoAck 对gateway获取系统信息返回
	GatewayFlagGetInfoAck
	//GatewayFlagSetInfoUpdateIntervalAck 对gateway进行心跳间隔设置返回
	GatewayFlagSetInfoUpdateIntervalAck

	//GatewayConStateUnkown gateway 连接状态未知
	GatewayConStateUnkown
	//GatewayConStateOnline gateway上线
	GatewayConStateOnline
	//GatewayConStateOffline gateway离线
	GatewayConStateOffline
	//GatewaySynStateUnkown gateway 同步状态未知
	GatewaySynStateUnkown
	//GatewaySynStateNeedConfig gateway待配置同步
	GatewaySynStateNeedConfig
	//GatewaySynStateNeedTime gateway待时间同步
	GatewaySynStateNeedTime
	//GatewaySynStateOK gateway同步完成
	GatewaySynStateOK
	//GatewayDebugStateOn gateway调试状态开，对应的Flag还没定义，是因为B接口暂没有此协议
	GatewayDebugStateOn
	//GatewayDebugStateOff gateway调试状态关
	GatewayDebugStateOff

	//DeviceConStateUnkown Device 未知连接状态
	DeviceConStateUnkown
	//DeviceConStateOnline Device 正常通信
	DeviceConStateOnline
	//DeviceConStateOffline Device 离线，通信中断
	DeviceConStateOffline
	//DeviceDebugStateOn Device 调试开
	DeviceDebugStateOn
	//DeviceDebugStateOff Device 调试关
	DeviceDebugStateOff

	//CorePointFlagUnkown corepoint 未知状态
	CorePointFlagUnkown
	//SignalFlagData signal cov 实时数据
	CorePointFlagStart
	//CorePointFlagConfirm corepoint 事件确认
	CorePointFlagConfirm
	//CorePointFlagEnd corepoint 事件结束
	CorePointFlagEnd
)

//GetEnumInt return enum int value
func GetEnumInt(str string) int {
	switch str {
	case "GatewayFlagUnkown":
		return GatewayFlagUnkown
	case "GatewayFlagRegister":
		return GatewayFlagRegister
	case "GatewayFlagTimeCheckAck":
		return GatewayFlagTimeCheckAck
	case "GatewayFlagHeartbeat":
		return GatewayFlagHeartbeat
	case "GatewayFlagRebootAck":
		return GatewayFlagRebootAck
	case "GatewayFlagGetConfigAck":
		return GatewayFlagGetConfigAck
	case "GatewayFlagSetConfigAck":
		return GatewayFlagSetConfigAck
	case "GatewayFlagSendConfig":
		return GatewayFlagSendConfig
	case "GatewayFlagFTPAck":
		return GatewayFlagFTPAck
	case "GatewayFlagAlarm":
		return GatewayFlagAlarm
	case "GatewayFlagGetDataAck":
		return GatewayFlagGetDataAck
	case "GatewayFlagSetPointAck":
		return GatewayFlagSetPointAck
	case "GatewayFlagGetThresholdAck":
		return GatewayFlagGetThresholdAck
	case "GatewayFlagSetThresholdAck":
		return GatewayFlagSetThresholdAck
	case "GatewayFlagSetStorageRuleAck":
		return GatewayFlagSetStorageRuleAck
	case "GatewayFlagGetStorageRuleAck":
		return GatewayFlagGetStorageRuleAck
	case "GatewayFlagGetInfoAck":
		return GatewayFlagGetInfoAck
	case "GatewayFlagSetInfoUpdateIntervalAck":
		return GatewayFlagSetInfoUpdateIntervalAck
	case "GatewayConStateUnkown":
		return GatewayConStateUnkown
	case "GatewayConStateOnline":
		return GatewayConStateOnline
	case "GatewayConStateOffline":
		return GatewayConStateOffline
	case "GatewaySynStateUnkown":
		return GatewaySynStateUnkown
	case "GatewaySynStateNeedConfig":
		return GatewaySynStateNeedConfig
	case "GatewaySynStateNeedTime":
		return GatewaySynStateNeedTime
	case "GatewaySynStateOK":
		return GatewaySynStateOK
	case "GatewayDebugStateOn":
		return GatewayDebugStateOn
	case "GatewayDebugStateOff":
		return GatewayDebugStateOff

	case "DeviceConStateUnkown":
		return DeviceConStateUnkown
	case "DeviceConStateOnline":
		return DeviceConStateOnline
	case "DeviceConStateOffline":
		return DeviceConStateOffline
	case "DeviceDebugStateOn":
		return DeviceDebugStateOn
	case "DeviceDebugStateOff":
		return DeviceDebugStateOff
	case "CorePointFlagUnkown":
		return CorePointFlagUnkown
	case "CorePointFlagStart":
		return CorePointFlagStart
	case "CorePointFlagConfirm":
		return CorePointFlagConfirm
	case "CorePointFlagEnd":
		return CorePointFlagEnd
	default:
		return -1
	}
}

//GetEnumString return enum string value
func GetEnumString(numInt int) string {
	switch numInt {
	case GatewayFlagUnkown:
		return "GatewayFlagUnkown"
	case GatewayFlagRegister:
		return "GatewayFlagRegister"
	case GatewayFlagTimeCheckAck:
		return "GatewayFlagTimeCheckAck"
	case GatewayFlagHeartbeat:
		return "GatewayFlagHeartbeat"
	case GatewayFlagRebootAck:
		return "GatewayFlagRebootAck"
	case GatewayFlagGetConfigAck:
		return "GatewayFlagGetConfigAck"
	case GatewayFlagSetConfigAck:
		return "GatewayFlagSetConfigAck"
	case GatewayFlagSendConfig:
		return "GatewayFlagSendConfig"
	case GatewayFlagFTPAck:
		return "GatewayFlagFTPAck"
	case GatewayFlagAlarm:
		return "GatewayFlagAlarm"
	case GatewayFlagGetDataAck:
		return "GatewayFlagGetDataAck"
	case GatewayFlagSetPointAck:
		return "GatewayFlagSetPointAck"
	case GatewayFlagGetThresholdAck:
		return "GatewayFlagGetThresholdAck"
	case GatewayFlagSetThresholdAck:
		return "GatewayFlagSetThresholdAck"
	case GatewayFlagSetStorageRuleAck:
		return "GatewayFlagSetStorageRuleAck"
	case GatewayFlagGetStorageRuleAck:
		return "GatewayFlagGetStorageRuleAck"
	case GatewayFlagGetInfoAck:
		return "GatewayFlagGetInfoAck"
	case GatewayFlagSetInfoUpdateIntervalAck:
		return "GatewayFlagSetInfoUpdateIntervalAck"
	case GatewayConStateUnkown:
		return "GatewayConStateUnkown"
	case GatewayConStateOnline:
		return "GatewayConStateOnline"
	case GatewayConStateOffline:
		return "GatewayConStateOffline"
	case GatewaySynStateUnkown:
		return "GatewaySynStateUnkown"
	case GatewaySynStateNeedConfig:
		return "GatewaySynStateNeedConfig"
	case GatewaySynStateNeedTime:
		return "GatewaySynStateNeedTime"
	case GatewaySynStateOK:
		return "GatewaySynStateOK"
	case GatewayDebugStateOn:
		return "GatewayDebugStateOn"
	case GatewayDebugStateOff:
		return "GatewayDebugStateOff"
	case DeviceConStateOffline:
		return "DeviceConStateOffline"
	case DeviceDebugStateOn:
		return "DeviceDebugStateOn"
	case DeviceDebugStateOff:
		return "DeviceDebugStateOff"
	case CorePointFlagUnkown:
		return "CorePointFlagUnkown"
	case CorePointFlagStart:
		return "CorePointFlagStart"
	case CorePointFlagConfirm:
		return "CorePointFlagConfirm"
	case CorePointFlagEnd:
		return "CorePointFlagEnd"
	default:
		return "NA"
	}
}
