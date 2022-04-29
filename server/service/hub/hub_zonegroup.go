package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hub"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    hubReq "github.com/flipped-aurora/gin-vue-admin/server/model/hub/request"
)

type ZoneGroupService struct {
}

// CreateZoneGroup 创建ZoneGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (zoneGroupService *ZoneGroupService) CreateZoneGroup(zoneGroup hub.ZoneGroup) (err error) {
	err = global.GVA_DB.Create(&zoneGroup).Error
	return err
}

// DeleteZoneGroup 删除ZoneGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (zoneGroupService *ZoneGroupService)DeleteZoneGroup(zoneGroup hub.ZoneGroup) (err error) {
	err = global.GVA_DB.Delete(&zoneGroup).Error
	return err
}

// DeleteZoneGroupByIds 批量删除ZoneGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (zoneGroupService *ZoneGroupService)DeleteZoneGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hub.ZoneGroup{},"id in ?",ids.Ids).Error
	return err
}

// UpdateZoneGroup 更新ZoneGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (zoneGroupService *ZoneGroupService)UpdateZoneGroup(zoneGroup hub.ZoneGroup) (err error) {
	err = global.GVA_DB.Save(&zoneGroup).Error
	return err
}

// GetZoneGroup 根据id获取ZoneGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (zoneGroupService *ZoneGroupService)GetZoneGroup(id uint) (err error, zoneGroup hub.ZoneGroup) {
	err = global.GVA_DB.Where("id = ?", id).First(&zoneGroup).Error
	return
}

// GetZoneGroupInfoList 分页获取ZoneGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (zoneGroupService *ZoneGroupService)GetZoneGroupInfoList(info hubReq.ZoneGroupSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hub.ZoneGroup{})
    var zoneGroups []hub.ZoneGroup
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&zoneGroups).Error
	return err, zoneGroups, total
}
