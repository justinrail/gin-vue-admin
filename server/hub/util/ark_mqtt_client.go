package util

import (
	"fmt"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ArkMQTTClient struct {
	client        mqtt.Client
	config        *MQTTClientConfig
	clientOptions *mqtt.ClientOptions
	Topics        map[string]byte
}

type MQTTClientConfig struct {
	Host     string
	Port     int
	Action   string
	Topic    string
	Username string
	Password string
	Qos      int
	Tls      bool
	CaCert   string
	ClientID string
}

func (arkMQTTClient *ArkMQTTClient) Connect() error {
	opts := mqtt.NewClientOptions()

	config := arkMQTTClient.config
	brokerString := fmt.Sprintf("tcp://%s:%d", config.Host, config.Port)
	opts.AddBroker(brokerString)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	opts.AutoReconnect = true
	opts.SetOnConnectHandler(connectHandler)
	opts.SetConnectionLostHandler(connectLostHandler)

	arkMQTTClient.clientOptions = opts

	//连接
	arkMQTTClient.client = mqtt.NewClient(opts)
	//客户端连接判断
	if token := arkMQTTClient.client.Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {

		return token.Error()
	}

	return nil
}

func (arkMQTTClient *ArkMQTTClient) Close() {
	arkMQTTClient.client.Disconnect(250)

}

func (arkMQTTClient *ArkMQTTClient) Publish(topic string, qos byte, retained bool, payload interface{}) {
	token := arkMQTTClient.client.Publish(topic, qos, retained, payload)
	token.Wait()
}

func (arkMQTTClient *ArkMQTTClient) Subscibe(topics map[string]byte) {
	arkMQTTClient.Topics = topics
}

func (arkMQTTClient *ArkMQTTClient) Listen(callback mqtt.MessageHandler) {

	arkMQTTClient.client.SubscribeMultiple(arkMQTTClient.Topics, callback)
	for {
		time.Sleep(1 * time.Second)
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {

	fmt.Println("MQTT Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {

	fmt.Printf("MQTT Connect lost: %v", err)
}

func (arkMQTTClient *ArkMQTTClient) New() error {

	uid := strconv.FormatInt(time.Now().Unix(), 10)

	arkMQTTClient.config = &MQTTClientConfig{
		Host:     global.GVA_VP.GetString("mqtt-broker-host"),
		Port:     1883,
		ClientID: global.GVA_VP.GetString("mqtt-client-prefix") + uid,
	}

	return arkMQTTClient.Connect()
}
