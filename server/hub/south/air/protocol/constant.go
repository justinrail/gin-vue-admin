package protocol

const (
	//LengthLen 变长集合长度位  Int32 LENGTH_LEN = 2
	LengthLen = 2
	//HostTypeLen 主机类型长度 Int32 HOSTTYPE_LEN = 1
	HostTypeLen = 1
	//MonitorUnitCodeLen 监控单元编码码长度 Int32 MONITORUNITCODE_LEN = 32
	MonitorUnitCodeLen = 32
	//ConfigFileCodeLen 主机配置MD5码长度  Int32 CONFIGFILECODE_LEN = 32
	ConfigFileCodeLen = 32
	//SampleFileCodeLen 采集结构文件MD5码长度 Int32 SAMPLEFILECODE_LEN = 32
	SampleFileCodeLen = 32
	//StartUpTimeLen 启动时间长度 Int32 STARTUPTIME_LEN = 4
	StartUpTimeLen = 4

	PortLen = 4

	IntervalLen = 4

	ResultCodeLen = 1

	EquipmentLen = 4

	SignalLen = 4

	SignalTypeLen = 1

	EventSeverityLen = 1

	ValueTypeLen = 1

	SignalBaseTypeLen = 4
)
