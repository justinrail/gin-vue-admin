package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hub"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    hubReq "github.com/flipped-aurora/gin-vue-admin/server/model/hub/request"
)

type SiteDeviceMapService struct {
}

// CreateSiteDeviceMap 创建SiteDeviceMap记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteDeviceMapService *SiteDeviceMapService) CreateSiteDeviceMap(siteDeviceMap hub.SiteDeviceMap) (err error) {
	err = global.GVA_DB.Create(&siteDeviceMap).Error
	return err
}

// DeleteSiteDeviceMap 删除SiteDeviceMap记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteDeviceMapService *SiteDeviceMapService)DeleteSiteDeviceMap(siteDeviceMap hub.SiteDeviceMap) (err error) {
	err = global.GVA_DB.Delete(&siteDeviceMap).Error
	return err
}

// DeleteSiteDeviceMapByIds 批量删除SiteDeviceMap记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteDeviceMapService *SiteDeviceMapService)DeleteSiteDeviceMapByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hub.SiteDeviceMap{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSiteDeviceMap 更新SiteDeviceMap记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteDeviceMapService *SiteDeviceMapService)UpdateSiteDeviceMap(siteDeviceMap hub.SiteDeviceMap) (err error) {
	err = global.GVA_DB.Save(&siteDeviceMap).Error
	return err
}

// GetSiteDeviceMap 根据id获取SiteDeviceMap记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteDeviceMapService *SiteDeviceMapService)GetSiteDeviceMap(id uint) (err error, siteDeviceMap hub.SiteDeviceMap) {
	err = global.GVA_DB.Where("id = ?", id).First(&siteDeviceMap).Error
	return
}

// GetSiteDeviceMapInfoList 分页获取SiteDeviceMap记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteDeviceMapService *SiteDeviceMapService)GetSiteDeviceMapInfoList(info hubReq.SiteDeviceMapSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hub.SiteDeviceMap{})
    var siteDeviceMaps []hub.SiteDeviceMap
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&siteDeviceMaps).Error
	return err, siteDeviceMaps, total
}
