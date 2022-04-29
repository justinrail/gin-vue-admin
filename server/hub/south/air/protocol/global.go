package protocol

import "github.com/flipped-aurora/gin-vue-admin/server/global"

var HubEngineID int32

func Init() {
	HubEngineID = global.GVA_VP.GetInt32("hub.engine-id")
}
