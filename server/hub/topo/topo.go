package topo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/bus"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/topo/net"
	flow "github.com/trustmaster/goflow"
)

func Ready() {
	InitVars()
	net.InitMQTT()
	//link
	Vars.Graph.SetInPort("cog_in", bus.COGBus)
	Vars.Graph.SetInPort("cor_in", bus.CODBus)
	Vars.Graph.SetInPort("cov_in", bus.COVBus)
	Vars.Graph.SetInPort("coa_in", bus.COABus)

}

func Start() {
	wait := flow.Run(Vars.Graph)
	<-wait
}
