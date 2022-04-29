package shadow

//业务运行数据的缓存接口
type IHubDomainCache interface {
	Load()
}
