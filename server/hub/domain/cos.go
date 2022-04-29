package domain

//Change of SamplerUnit(DTO for DeviceAdapter)
type COS struct {
	GatewayID     int
	DeviceID      int
	SamplerUnitID int
	State         int
	UpdateTime    int64
}
