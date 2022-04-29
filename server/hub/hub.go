package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/south/stub"
)

func Run() {
	shadow.Init()
	// mock接口，用来做数据测试
	stub.Ready()
	go stub.Start()

	// 标准B接口协议
	//air.Ready()
	//go air.Start()
}
