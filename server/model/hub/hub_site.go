// 自动生成模板Site
package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Site 结构体
// 如果含有time.Time 请自行import time包
type Site struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:站点名称;"`
      SiteType  *int `json:"siteType" form:"siteType" gorm:"column:site_type;comment:站点类型;"`
      Longitude  *float64 `json:"longitude" form:"longitude" gorm:"column:longitude;comment:经度;"`
      Latitude  *float64 `json:"latitude" form:"latitude" gorm:"column:latitude;comment:纬度;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:说明;"`
}


// TableName Site 表名
func (Site) TableName() string {
  return "hub_site"
}

