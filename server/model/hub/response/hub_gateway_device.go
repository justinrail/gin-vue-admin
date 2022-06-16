package response

type GatewayDevice struct {
	GatewayID      int    `json:"gatewayID"`
	GatewayName    string `json:"gatewayName"`
	GatewayAddress string `json:"gatewayAddress"`
	Devices        []*LiteDevice
}
