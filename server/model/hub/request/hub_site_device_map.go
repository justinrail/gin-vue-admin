package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/hub"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SiteDeviceMapSearch struct{
    hub.SiteDeviceMap
    request.PageInfo
}
