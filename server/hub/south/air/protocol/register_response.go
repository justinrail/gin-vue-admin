package protocol

import "encoding/binary"

//RegisterResponse 注册请求返回
type RegisterResponse struct {
	ResultCode    int
	MessageType   int
	PipeLinePairs []*PipeLinePair
}

func (registerResponse *RegisterResponse) FromByteArray(data []byte, index int) {
	registerResponse.ResultCode = (int)(data[index])
	registerResponse.MessageType = 71
	index += 1
	itemCount := binary.LittleEndian.Uint16(data[index:])
	index += LengthLen
	len := 0

	registerResponse.PipeLinePairs = make([]*PipeLinePair, 0)
	// 服务地址集合
	for i := 0; i < (int)(itemCount); i++ {
		// 管道类型
		pipelineType := (byte)(data[index])
		index++
		// 服务器地址长度
		len = (int)(data[index])
		index++
		// 服务器地址
		url := string(data[index : index+len])
		index += len
		registerResponse.PipeLinePairs = append(registerResponse.PipeLinePairs, &PipeLinePair{PipeLineType: pipelineType, URL: url})
	}
}

func (registerResponse *RegisterResponse) ToByteArray() []byte {
	data := make([]byte, ResultCodeLen+LengthLen)
	index := 0
	data[index] = (byte)(registerResponse.ResultCode)
	index += ResultCodeLen
	binary.LittleEndian.PutUint16(data[index:index+LengthLen], (uint16)(len(registerResponse.PipeLinePairs)))
	index += LengthLen
	for _, pp := range registerResponse.PipeLinePairs {
		url := ([]byte)(pp.URL)
		len := len(url)
		data = append(data, pp.PipeLineType)
		data = append(data, (byte)(len))
		data = append([]byte{}, append(data, url...)...)

	}

	return data

}
