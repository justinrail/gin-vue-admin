package hub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hub"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    hubReq "github.com/flipped-aurora/gin-vue-admin/server/model/hub/request"
)

type CenterService struct {
}

// CreateCenter 创建Center记录
// Author [piexlmax](https://github.com/piexlmax)
func (centerService *CenterService) CreateCenter(center hub.Center) (err error) {
	err = global.GVA_DB.Create(&center).Error
	return err
}

// DeleteCenter 删除Center记录
// Author [piexlmax](https://github.com/piexlmax)
func (centerService *CenterService)DeleteCenter(center hub.Center) (err error) {
	err = global.GVA_DB.Delete(&center).Error
	return err
}

// DeleteCenterByIds 批量删除Center记录
// Author [piexlmax](https://github.com/piexlmax)
func (centerService *CenterService)DeleteCenterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]hub.Center{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCenter 更新Center记录
// Author [piexlmax](https://github.com/piexlmax)
func (centerService *CenterService)UpdateCenter(center hub.Center) (err error) {
	err = global.GVA_DB.Save(&center).Error
	return err
}

// GetCenter 根据id获取Center记录
// Author [piexlmax](https://github.com/piexlmax)
func (centerService *CenterService)GetCenter(id uint) (err error, center hub.Center) {
	err = global.GVA_DB.Where("id = ?", id).First(&center).Error
	return
}

// GetCenterInfoList 分页获取Center记录
// Author [piexlmax](https://github.com/piexlmax)
func (centerService *CenterService)GetCenterInfoList(info hubReq.CenterSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hub.Center{})
    var centers []hub.Center
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&centers).Error
	return err, centers, total
}
