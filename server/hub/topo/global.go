package topo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/topo/process"
	"github.com/trustmaster/goflow"
)

var Vars *TopoVars

//系统标准B接口的全局对象体
type TopoVars struct {
	Graph *goflow.Graph
}

func InitVars() {
	Vars = &TopoVars{}
	Vars.Graph = goflow.NewGraph()
	initGraph()
}

func initGraph() {

	// 网关状态管理器
	Vars.Graph.Add("gateway_stater", new(process.GatewayStater))
	Vars.Graph.Add("device_stater", new(process.DeviceStater))
	Vars.Graph.Add("point_stater", new(process.PointStater))
	Vars.Graph.Add("alarm_stater", new(process.AlarmStater))
	Vars.Graph.Add("activealarm_ponder", new(process.ActiveAlarmPonder))

	Vars.Graph.Connect("alarm_stater", "Out", "activealarm_ponder", "In")
	// net.Connect("cov_spout", "COVOut1", "cov_data_updater", "In")
	// net.Connect("cov_data_updater", "StateOut", "cov_state_normalizer", "In")
	// net.Connect("cov_state_normalizer", "StateOut", "cov_state_updater", "In")
	// net.Connect("cov_state_updater", "CoreLiteEventOut2", "phoenix_cov_hooker", "In")
	// net.Connect("cov_spout", "COVOut2", "cov_ashbin", "In")

	// Network ports
	Vars.Graph.MapInPort("cog_in", "gateway_stater", "In")
	Vars.Graph.MapInPort("cod_in", "device_stater", "In")
	Vars.Graph.MapInPort("cov_in", "point_stater", "In")
	Vars.Graph.MapInPort("coa_in", "alarm_stater", "In")

}
