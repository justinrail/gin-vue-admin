package protocol

import (
	"encoding/binary"
)

const (
	Idle = 0
	Busy = 1
	Full = 2
)

func CreateAckPacket(sequenceNumber uint16, status int, hostId int) []byte {
	data := make([]byte, HeadLength)
	var version byte = 1
	var protocolType = 15
	var isBusy = false
	index := SOIIndex
	binary.LittleEndian.PutUint16(data[index:index+SOILen], (uint16)(SOI))
	index = SequenceNumberIndex
	binary.LittleEndian.PutUint16(data[index:index+SequenceNumberLen], sequenceNumber)
	index = SourceHostIDIndex
	binary.LittleEndian.PutUint32(data[index:index+SourceHostIDLen], uint32(HubEngineID))
	index = DestinationHostIDIndex
	binary.LittleEndian.PutUint32(data[index:index+DestinationHostIDLen], uint32(hostId))

	data[4] = (byte)(version<<4) | (byte)(protocolType)

	if status != Idle {
		isBusy = true
	}

	data[5] = byte(Ternary(0x40, 0x0, isBusy))

	//checksum := (uint16)(data[2] + data[3] + data[4] + data[5] + data[12] + data[13] + data[14] + data[15] + data[16] + data[17] + data[18] + data[19]);
	checksum := getCheckSum(data)
	index = CheckSumIndex
	binary.LittleEndian.PutUint16(data[index:index+CheckSumLen], checksum)
	return data
}
