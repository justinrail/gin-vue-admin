package air

import "github.com/flipped-aurora/gin-vue-admin/server/hub/south/air/process"

//InitGraph creates network structure
func InitGraph() {
	// Add graph nodes

	//解包，确定是不是标准协议
	Vars.Graph.Add("unpacker", new(process.Unpacker))
	//装配包
	Vars.Graph.Add("message_picker", new(process.MessagePicker))
	//分发给不同的处理器进行包的处理
	Vars.Graph.Add("router", new(process.Router))
	//心跳包处理器
	Vars.Graph.Add("heartbeatrequest_adapter", new(process.HeartbeatRequestAdapter))
	//注册包处理器
	Vars.Graph.Add("registerrequest_adapter", new(process.RegisterRequestAdapter))
	Vars.Graph.Add("realsignalresponse_adapter", new(process.RealSignalResponseAdapter))
	Vars.Graph.Add("eventresponse_adapter", new(process.EventResponseAdapter))
	// 历史数据可以简单的在上层用定时器遍历，如果不实现不影响底端运行的话，可先不实现
	// Vars.Graph.Add(new(process.HistorySignalResponseAdapter), "historysignalresponse_adapter")
	// 关联信号可以不使用，所以不用
	// Vars.Graph.Add(new(process.RelatedSignalResponseAdapter), "relatedsignalresponse_adapter")
	// 配置同步校验因底端是数据库方式和文件都存在设计需更改，目前采用中心同步，网关数量少，如果不影响运行可不实现
	// Vars.Graph.Add(new(process.ConfigFileResponseAdapter), "configfileresponse_adapter")
	// TODO 控制命令相关协议需要实现
	// TODO 时间校时命令相关协议需要实现
	//samplerUnit  的协议理解不够，没完全写完
	Vars.Graph.Add("samplerconnectstate_adapter", new(process.SamplerConnectStateAdapter))
	//net connections
	Vars.Graph.Connect("unpacker", "Out", "message_picker", "In")
	Vars.Graph.Connect("message_picker", "Out", "router", "In")
	Vars.Graph.Connect("router", "OutHeartbeatRequest", "heartbeatrequest_adapter", "In")
	Vars.Graph.Connect("router", "OutRegisterRequest", "registerrequest_adapter", "In")
	Vars.Graph.Connect("router", "OutRealSignalResponse", "realsignalresponse_adapter", "In")
	Vars.Graph.Connect("router", "OutEventResponse", "eventresponse_adapter", "In")
	// Vars.Graph.Connect("router", "OutHistorySignalResponse", "historysignalresponse_adapter", "In")
	// Vars.Graph.Connect("router", "OutRelatedSignalResponse", "relatedsignalresponse_adapter", "In")
	// Vars.Graph.Connect("router", "OutConfigFileResponse", "configfileresponse_adapter", "In")
	Vars.Graph.Connect("router", "OutSamplerConnectState", "samplerconnectstate_adapter", "In")

	// Network ports
	Vars.Graph.MapInPort("In", "unpacker", "In")
}
