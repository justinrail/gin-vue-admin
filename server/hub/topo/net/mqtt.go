package net

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/util"
)

var MQTTClient *util.ArkMQTTClient

var EnableMQTT = false

func InitMQTT() {
	EnableMQTT = global.GVA_VP.GetBool("hub.mqtt-send-enable")
	MQTTClient = &util.ArkMQTTClient{}
	MQTTClient.New()

	//now not need to add topic to mqtt client
	// topics := make(map[string]byte, 5)
	// topics["ark/cov"] = 0
	// topics["ark/coa"] = 1
	// topics["ark/cog"] = 1
	// topics["ark/cod"] = 1
	// topics["ark/cos"] = 1

	// mqttClient.Subscibe(topics)
}
