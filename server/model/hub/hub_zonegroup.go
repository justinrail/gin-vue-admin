// 自动生成模板ZoneGroup
package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ZoneGroup 结构体
// 如果含有time.Time 请自行import time包
type ZoneGroup struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:区域分组名称;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:分组说明;"`
}


// TableName ZoneGroup 表名
func (ZoneGroup) TableName() string {
  return "hub_zone_group"
}

