package shadow

import "github.com/flipped-aurora/gin-vue-admin/server/hub/domain"

type ActiveAlarmSinker struct {
}

func (s *ActiveAlarmSinker) Ready() {

}

//暂时直接更新数据库，以后要改成批量避免堵塞
func (s *ActiveAlarmSinker) Sink(activeAlarm *domain.ActiveAlarm) {

}
