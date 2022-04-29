package entity

type Sampler struct {
	SamplerId        int
	SamplerName      string
	SamplerType      int
	ProtocolCode     string
	DLLCode          string
	DLLVersion       string
	ProtocolFilePath string
	DLLFilePath      string
	DllPath          string
	Setting          string
	Description      string
	SoCode           string
	SoPath           string
}

func (s *Sampler) TableName() string {
	tableName := "fsu_sampler"
	return tableName
}
