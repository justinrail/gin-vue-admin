package cascade

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/util"
)

func Ready() {
	mqttClient = &util.ArkMQTTClient{}
	mqttClient.New()

	//add topic to mqtt client
	topics := make(map[string]byte, 5)
	topics["ark/cov"] = 0
	topics["ark/coa"] = 1
	topics["ark/cog"] = 1
	topics["ark/cod"] = 1
	topics["ark/cos"] = 1

	mqttClient.Subscibe(topics)
}

func Start() {
	// mqttClient.Publish("ark/cov/1", 0, false, "23")
	// time.Sleep(time.Duration(3) * time.Second)
	// mqttClient.Publish("ark/cov/1", 1, false, "TEST")

	//listen
	if !global.GVA_VP.GetBool("hub.mqtt-recv-enable") {
		return
	}

	go mqttClient.Listen(func(c mqtt.Client, m mqtt.Message) {
		handleCOV(m.Topic(), m.Payload())
		//fmt.Printf("sub [%s] %s\n", m.Topic(), string(m.Payload()))
	})
}
