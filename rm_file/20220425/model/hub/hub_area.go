// 自动生成模板Area
package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Area 结构体
// 如果含有time.Time 请自行import time包
type Area struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:片区名称;"`
      Longitude  *float64 `json:"longitude" form:"longitude" gorm:"column:longitude;comment:经度;"`
      Latitude  *float64 `json:"latitude" form:"latitude" gorm:"column:latitude;comment:纬度;"`
      Remark  *float64 `json:"remark" form:"remark" gorm:"column:remark;comment:说明;"`
}


// TableName Area 表名
func (Area) TableName() string {
  return "area"
}

