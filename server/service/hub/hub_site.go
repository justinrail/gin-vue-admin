package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hub"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    hubReq "github.com/flipped-aurora/gin-vue-admin/server/model/hub/request"
)

type SiteService struct {
}

// CreateSite 创建Site记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteService *SiteService) CreateSite(site hub.Site) (err error) {
	err = global.GVA_DB.Create(&site).Error
	return err
}

// DeleteSite 删除Site记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteService *SiteService)DeleteSite(site hub.Site) (err error) {
	err = global.GVA_DB.Delete(&site).Error
	return err
}

// DeleteSiteByIds 批量删除Site记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteService *SiteService)DeleteSiteByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hub.Site{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSite 更新Site记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteService *SiteService)UpdateSite(site hub.Site) (err error) {
	err = global.GVA_DB.Save(&site).Error
	return err
}

// GetSite 根据id获取Site记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteService *SiteService)GetSite(id uint) (err error, site hub.Site) {
	err = global.GVA_DB.Where("id = ?", id).First(&site).Error
	return
}

// GetSiteInfoList 分页获取Site记录
// Author [piexlmax](https://github.com/piexlmax)
func (siteService *SiteService)GetSiteInfoList(info hubReq.SiteSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hub.Site{})
    var sites []hub.Site
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&sites).Error
	return err, sites, total
}
