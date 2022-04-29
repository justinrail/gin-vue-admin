package protocol

type PipelineObjectPair struct {
	PipelineType       int
	Data               []byte
	RequireAcknowledge bool
	Ttl                byte
	Url                string
	SequenceNumber     uint16
	SourceHostId       int32
	DestinationHostId  int32
	MessageType        int
}
