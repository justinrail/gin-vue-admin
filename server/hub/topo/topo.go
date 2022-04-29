package topo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/bus"
	flow "github.com/trustmaster/goflow"
)

func Ready() {
	InitVars()

	//link
	Vars.Graph.SetInPort("cog_in", bus.COGBus)
	Vars.Graph.SetInPort("cor_in", bus.CODBus)
	Vars.Graph.SetInPort("cov_in", bus.COVBus)

}

func Start() {
	wait := flow.Run(Vars.Graph)
	<-wait
}
