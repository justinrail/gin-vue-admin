package process

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
	flow "github.com/trustmaster/goflow"
	"go.uber.org/zap"
)

type ActiveAlarmPonder struct {
	flow.Component
	In        chan *domain.COA
	COALogOut chan *domain.COA
}

func (ponder *ActiveAlarmPonder) Process() {

	for coa := range ponder.In {
		activeAlarm, existAcitveAlarm := shadow.GetActiveAlarmByKey(coa.GetUniqueKey())
		if existAcitveAlarm {
			if coa.EndTime == 0 { //if activeAlarm is exist, meaning alarm is occurring, so coa only end is acceptable
				global.GVA_LOG.Warn("start coa duplicate send to activeAlarm Ponder : ",
					zap.String("COA", coa.ToString()))
				return
			} else if activeAlarm.ConfirmTime == 0 { //如果告警not confirmed
				shadow.EndActiveAlarm(coa)
			} else { //if activeAlarm is confirmed
				shadow.FinishActiveAlarm(coa) //save and remove alarm to log table
			}
		} else { //if activeAlarm not exist
			if coa.EndTime > 0 { //if coa is ended
				global.GVA_LOG.Warn("end coa not exist in activeAlarm Ponder : ",
					zap.String("COA", coa.ToString()))
				return
			} else { //if coa == 0
				shadow.StartActiveAlarm(coa)
			}
		}

		shadow.AppendCOVLog(coa)
		ponder.COALogOut <- coa

	}

}
