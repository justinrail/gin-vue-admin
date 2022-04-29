package process

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/base"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/bus"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/air/protocol"
	flow "github.com/trustmaster/goflow"
)

//RegisterRequestAdapter 注册请求包处理
type RegisterRequestAdapter struct {
	flow.Component
	In <-chan *protocol.PipelineObjectPair
}

//OnIn Request handler
func (adapter *RegisterRequestAdapter) Process() {
	//metrics.AppMetrics.HubCollectorCMBRouterCounter.Inc(1)

	for pair := range adapter.In {
		request := &protocol.RegisterRequest{}
		request.FromByteArray(pair.Data, protocol.HeadLength)
		replayRegister(pair)

		//生成gateway注册状态的COG发送到管道
		cog := &domain.COG{}
		cog.GatewayID = int(pair.SourceHostId)
		cog.Address = pair.Url
		cog.Flag = base.GatewayFlagRegister
		cog.Timestamp = time.Now().Unix()

		bus.COGBus <- cog
	}

	//uuid := request.GetHostUUID()

	//考虑本软件是小项目（FSU小于20），所以默认加载所有gateway（根据entity，初始化domain对象）
	//domain对象的上次值也考虑加载
	//将所有domain对象设置为离线
	//从shandow获取gateway的domain对象（初始化时加载到内存），确认此网关（FSU）是否已经接入（joined参数？）
	//
	//如果没接入，生成cog，让topo去更新gateway。
	//如果接入则什么都不做, 也发cog，给后台更新gateway的活动时间或日志
	//返回注册回复包

	//先同步配置
	// if protocol.SiteWebGateways[uuid] == nil {
	// 	//确定唯一运行配置同步功能
	// 	jobID := "synconfig." + uuid
	// 	state := protocol.JobStates[jobID]
	// 	if state == false {
	// 		protocol.JobStates[jobID] = true
	// 		registerRequestAdapter.synConfig(jobID, request)
	// 		protocol.JobStates[jobID] = false
	// 		addSiteWebGateway(request)

	// 	}
	// }
	// fmt.Println("send register response")

	//fmt.Println(request)
}

func replayRegister(pipelineObjectPair *protocol.PipelineObjectPair) {
	data := createRegisterResponsePacket(pipelineObjectPair)
	protocol.SendUDPData(pipelineObjectPair.Url, data)
}

//SaveOrUpdateMonitorUnit 更新数据库配置
// func (registerRequestAdapter *RegisterRequestAdapter) SaveOrUpdateMonitorUnit(cfg *protocol.MainCfg, monitorUnitID int32) {

// }

// func (registerRequestAdapter *RegisterRequestAdapter) gatherConfig(ftpAddress string, monitorUnitID int32) error {

// 	file := registerRequestAdapter.getMonitorUnitXMLFile(ftpAddress, monitorUnitID)
// 	if len(file) == 0 {
// 		return fmt.Errorf("no found available xml file")
// 	}

// 	return nil
// }

func createRegisterResponsePacket(pipelineObjectPair *protocol.PipelineObjectPair) []byte {
	registerResponse := createRegisterResponse(pipelineObjectPair)
	pack := &protocol.AirPacket{}
	pack.MessageType = 71
	pack.DestinationHostID = pipelineObjectPair.DestinationHostId
	pack.SourceHostID = HubEngineID
	pack.RawObject = registerResponse
	pack.RequireAcknowledge = true
	pack.PipelineType = 0
	pack.ProtocolType = 1
	pack.SequenceNumber = pipelineObjectPair.SequenceNumber
	pack.RawPacket = registerResponse.ToByteArray()
	data := pack.Pack()
	return data
}

func createRegisterResponse(pipelineObjectPair *protocol.PipelineObjectPair) *protocol.RegisterResponse {
	registerResponse := &protocol.RegisterResponse{}
	registerResponse.MessageType = 71
	registerResponse.ResultCode = 0
	registerResponse.PipeLinePairs = protocol.GetPipeLines()
	return registerResponse
}

// //getMonitorUnitXMLFile RMU不提供此功能，目前RMU无法获取FTP文件。所以，暂时不实现
// func (registerRequestAdapter *RegisterRequestAdapter) getMonitorUnitXMLFile(ftpAddress string, monitorUnitID int32) string {
// 	user := ftpAddress[6:strings.LastIndex(ftpAddress, ":")]
// 	pwd := ftpAddress[strings.LastIndex(ftpAddress, ":")+1 : strings.Index(ftpAddress, "@")]
// 	//ip := ftpAddress[strings.Index(ftpAddress, "@")+1:]

// 	c, err := ftp.Dial("10.169.42.154:21", ftp.DialWithTimeout(5*time.Second))

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	user = "demo"
// 	pwd = "password"
// 	err = c.Login(user, pwd)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	if err := c.Quit(); err != nil {
// 		fmt.Println(err)
// 	}

// 	return "nil"
// }
