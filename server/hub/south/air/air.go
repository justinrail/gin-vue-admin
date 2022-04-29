package air

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/air/process"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/air/protocol"
	flow "github.com/trustmaster/goflow"
)

func Ready() {
	//init global vars
	InitVars()
	protocol.Init()
	process.Init()

	//entity & domain objs be loaded in shadow init

	diagnoseReceiver, err := protocol.NewUDPReceiver("9000", UDPBufferSize,
		Vars.RawPackets, global.GVA_VP.GetString("hub.dsc-default-url"))
	if err == nil {
		Vars.DefaultReceiver = diagnoseReceiver
	}

	/*RdsReceiver, err := protocol.NewUDPReceiver("7000", UDPBufferSize, rawPackets)
	if err == nil {
		rdsReceiver = RdsReceiver
	}*/

	InitGraph()

	Vars.Graph.SetInPort("In", Vars.RawPackets)
	//time.AfterFunc(5*time.Second, func() { protocol.SendUDPData("192.168.199.140:9000", []byte("Hello world!")) })

}

//Start 启动
func Start() {
	// InitCorePointSiteWebGateWayMap()
	// gateways := util.DeepLoadGatewayByCollector("siteweb")
	// for _, gateway := range gateways {
	// 	domain.Gateways.Insert(gateway.ID, &gateway)
	// }

	//goflow 的新版本接口变动非常大，这里请注意使用方法和接口都有变化，一定要看github文档
	wait := flow.Run(Vars.Graph)
	if Vars.DefaultReceiver != nil {
		go Vars.DefaultReceiver.Serve()
	}
	/*if rdsReceiver != nil {
		go rdsReceiver.Serve()
	}*/
	<-wait
}

// func InitCorePointSiteWebGateWayMap() {
// 	CorePointSiteWebGateWayMaps := repo.GetAllCorePointSiteWebGateWayMap()
// 	for _, pp := range CorePointSiteWebGateWayMaps {
// 		if pp.IsSignal() {
// 			protocol.SignalPointSiteWebGateWayMaps[pp.GetSignalKey()] = pp
// 		}
// 		if pp.IsControl() {
// 			protocol.ControlPointSiteWebGateWayMaps[pp.GetControlKey()] = pp
// 		}
// 		if pp.IsEvent() {
// 			protocol.EventPointSiteWebGateWayMaps[pp.GetEventKey()] = pp
// 		}
// 	}

// }
