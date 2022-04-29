package protocol

//Packet Protocol pack
type Packet struct {
	MessageType int
	Body        interface{}
}
