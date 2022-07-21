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
	Vars.Graph.Add("coa_spout", new(process.COASpout))
	Vars.Graph.Add("cov_spout", new(process.COVSpout))
	Vars.Graph.Add("coa_cascade_sender", new(process.COACascadeSender))
	Vars.Graph.Add("cov_cascade_sender", new(process.COVCascadeSender))

	Vars.Graph.Connect("coa_spout", "COAState", "alarm_stater", "In")
	Vars.Graph.Connect("alarm_stater", "Out", "activealarm_ponder", "In")

	Vars.Graph.Connect("coa_spout", "COACascade", "coa_cascade_sender", "In")

	Vars.Graph.Connect("cov_spout", "COVState", "point_stater", "In")
	Vars.Graph.Connect("cov_spout", "COVCascade", "cov_cascade_sender", "In")

	// net.Connect("cov_data_updater", "StateOut", "cov_state_normalizer", "In")
	// net.Connect("cov_state_normalizer", "StateOut", "cov_state_updater", "In")
	// net.Connect("cov_state_updater", "CoreLiteEventOut2", "phoenix_cov_hooker", "In")

	// Network ports
	Vars.Graph.MapInPort("cog_in", "gateway_stater", "In")
	Vars.Graph.MapInPort("cod_in", "device_stater", "In")
	// Vars.Graph.MapInPort("cov_in", "point_stater", "In")
	// Vars.Graph.MapInPort("coa_in", "alarm_stater", "In")
	Vars.Graph.MapInPort("cov_in", "cov_spout", "In")
	Vars.Graph.MapInPort("coa_in", "coa_spout", "In")
}
