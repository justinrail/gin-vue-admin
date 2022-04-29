package protocol

import (
	"strconv"
	"strings"
)

//RegisterRequest 注册请求
type RegisterRequest struct {
	//HostInfo 主机信息
	HostInfo    *HostInfo
	Check       [48]uint8
	MessageType int
	//SourceHostID 源主机ID，应该是MonitorUnitID
	SourceHostID int32
}

//GetRegisterRequestFromByteArray 把二进制值转换为传输对象
func (registerRequest *RegisterRequest) FromByteArray(data []byte, index int) {
	registerRequest.HostInfo = &HostInfo{}
	registerRequest.HostInfo.FromByteArray(data, index)
	registerRequest.MessageType = 70
	//registerRequest.HostInfo.GetCheck(data, index, registerRequest.Check[:])
}

//GetIP 获取gateway 的 IP
func (registerRequest *RegisterRequest) GetIP() string {
	if len(registerRequest.HostInfo.PipeLinePairs) == 0 {
		return ""
	}

	url := registerRequest.HostInfo.PipeLinePairs[0].URL

	return url[strings.Index(url, "//")+2 : strings.LastIndex(url, ":")]
}

//GetFTPAddress 获取FTP的服务地址
func (registerRequest *RegisterRequest) GetFTPAddress() string {
	if len(registerRequest.HostInfo.PipeLinePairs) == 0 {
		return ""
	}

	for _, pr := range registerRequest.HostInfo.PipeLinePairs {
		if strings.ToLower(pr.URL[0:3]) == "ftp" {
			return pr.URL
		}
	}

	return ""
}

//GetUDPAddress 获取UDP的服务地址
func (registerRequest *RegisterRequest) GetUDPAddress() string {
	if len(registerRequest.HostInfo.PipeLinePairs) == 0 {
		return ""
	}

	for _, pr := range registerRequest.HostInfo.PipeLinePairs {
		if strings.ToLower(pr.URL[0:3]) == "udp" {
			return pr.URL[6:]
		}
	}

	return ""
}

//GetHostUUID 根据IP + MonitorUnitID的方式进行UUID生成
func (registerRequest *RegisterRequest) GetHostUUID() string {
	return registerRequest.GetIP() + " " + strconv.Itoa((int)(registerRequest.SourceHostID))
}

//ToByteArray 转为二进制
func (registerRequest *RegisterRequest) ToByteArray() []byte {
	return registerRequest.HostInfo.ToByteArray()
}
