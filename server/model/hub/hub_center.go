// 自动生成模板Center
package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Center 结构体
// 如果含有time.Time 请自行import time包
type Center struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;"`
      Longitude  *float64 `json:"longitude" form:"longitude" gorm:"column:longitude;comment:经度;"`
      Latitude  *float64 `json:"latitude" form:"latitude" gorm:"column:latitude;comment:纬度;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:说明;"`
}


// TableName Center 表名
func (Center) TableName() string {
  return "hub_center"
}

