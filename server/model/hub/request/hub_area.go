package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/hub"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AreaSearch struct{
    hub.Area
    request.PageInfo
}
