// 自动生成模板SiteDeviceMap
package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SiteDeviceMap 结构体
// 如果含有time.Time 请自行import time包
type SiteDeviceMap struct {
      global.GVA_MODEL
      SiteId  *int `json:"siteId" form:"siteId" gorm:"column:site_id;comment:站点ID;"`
      DeviceId  *int `json:"deviceId" form:"deviceId" gorm:"column:device_id;comment:设备ID;"`
}


// TableName SiteDeviceMap 表名
func (SiteDeviceMap) TableName() string {
  return "hub_site_device_map"
}

