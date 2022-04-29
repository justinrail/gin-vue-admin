package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hub"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    hubReq "github.com/flipped-aurora/gin-vue-admin/server/model/hub/request"
)

type ZoneService struct {
}

// CreateZone 创建Zone记录
// Author [piexlmax](https://github.com/piexlmax)
func (zoneService *ZoneService) CreateZone(zone hub.Zone) (err error) {
	err = global.GVA_DB.Create(&zone).Error
	return err
}

// DeleteZone 删除Zone记录
// Author [piexlmax](https://github.com/piexlmax)
func (zoneService *ZoneService)DeleteZone(zone hub.Zone) (err error) {
	err = global.GVA_DB.Delete(&zone).Error
	return err
}

// DeleteZoneByIds 批量删除Zone记录
// Author [piexlmax](https://github.com/piexlmax)
func (zoneService *ZoneService)DeleteZoneByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hub.Zone{},"id in ?",ids.Ids).Error
	return err
}

// UpdateZone 更新Zone记录
// Author [piexlmax](https://github.com/piexlmax)
func (zoneService *ZoneService)UpdateZone(zone hub.Zone) (err error) {
	err = global.GVA_DB.Save(&zone).Error
	return err
}

// GetZone 根据id获取Zone记录
// Author [piexlmax](https://github.com/piexlmax)
func (zoneService *ZoneService)GetZone(id uint) (err error, zone hub.Zone) {
	err = global.GVA_DB.Where("id = ?", id).First(&zone).Error
	return
}

// GetZoneInfoList 分页获取Zone记录
// Author [piexlmax](https://github.com/piexlmax)
func (zoneService *ZoneService)GetZoneInfoList(info hubReq.ZoneSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hub.Zone{})
    var zones []hub.Zone
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&zones).Error
	return err, zones, total
}
