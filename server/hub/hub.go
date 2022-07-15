package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/cascade"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/stub"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/topo"
)

func Run() {
	shadow.Init()
	// mock接口，用来做数据测试
	stub.Ready()
	go stub.Start()

	// 标准B接口协议
	//air.Ready()
	//go air.Start()

	//数据处理层启动
	topo.Ready()
	go topo.Start()

	//级联接口
	cascade.Ready()
	go cascade.Start()
}
