package shadow

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/timer"
)

type ActiveAlarmLogSinker struct {
	ActiveAlarmArrayA []domain.ActiveAlarm
	ActiveAlarmArrayB []domain.ActiveAlarm
	FlagTurn          bool
	saved             bool
}

func (s *ActiveAlarmLogSinker) Ready() {
	s.ActiveAlarmArrayA = make([]domain.ActiveAlarm, 0)
	s.ActiveAlarmArrayB = make([]domain.ActiveAlarm, 0)
	s.FlagTurn = true
	s.saved = true
	//register process job
	timer.NewTimerTask().AddTaskByFunc("func.sinker.activealarmlog", "@every 1s", func() {
		s.Process()
	})
}

func (s *ActiveAlarmLogSinker) Sink(activeAlarm *domain.ActiveAlarm) {
	if s.FlagTurn == true {
		s.ActiveAlarmArrayA = append(s.ActiveAlarmArrayA, *activeAlarm)
	} else {
		s.ActiveAlarmArrayB = append(s.ActiveAlarmArrayB, *activeAlarm)
	}
}

//定时批量写入数据库(同时删除ActiveAlarm表记录)
func (s *ActiveAlarmLogSinker) Process() {
	if s.saved {
		s.FlagTurn = !s.FlagTurn
		s.saved = false
		if !s.FlagTurn == true {
			s.Save(s.ActiveAlarmArrayA)
			s.ActiveAlarmArrayA = make([]domain.ActiveAlarm, 0)
		} else {
			s.Save(s.ActiveAlarmArrayB)
			s.ActiveAlarmArrayB = make([]domain.ActiveAlarm, 0)
		}
	}
}

func (s *ActiveAlarmLogSinker) Save(activeAlarms []domain.ActiveAlarm) {
	//first add to log table
	//TODO change table name
	global.GVA_DB.CreateInBatches(activeAlarms, 100)
	//then remove from cache
	//global.GVA_DB.Where("name = ?", "jinzhu").Delete(&email)
	//db.Exec("UPDATE orders SET shipped_at = ? WHERE id IN ?", time.Now(), []int64{1, 2, 3})
	s.saved = true
}
