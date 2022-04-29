// 自动生成模板Zone
package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Zone 结构体
// 如果含有time.Time 请自行import time包
type Zone struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:区域名称;"`
	ParentId    *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:区域树的父区域关联ID;"`
	ZoneGroupId *int   `json:"zoneGroupId" form:"zoneGroupId" gorm:"column:zone_group_id;comment:系统支持多个区域组，这用来过滤区域;"`
	ObjectType  *int   `json:"objectType" form:"objectType" gorm:"column:object_type;comment:区域节点所关联对象的类型;"`
	ObjectId    *int   `json:"objectId" form:"objectId" gorm:"column:object_id;comment:关联对象ID;"`
}

// TableName Zone 表名
func (Zone) TableName() string {
	return "hub_zone"
}
