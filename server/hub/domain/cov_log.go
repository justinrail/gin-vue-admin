package domain

//历史数据，存时序数据库tsstore，实现基础功能，暂存1个月，仅kv和label，尽量采用ID管理，7天分段
//未来兼容influxDB实现接口隔离
type COVLog struct {
}
