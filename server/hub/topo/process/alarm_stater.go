package process

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
	flow "github.com/trustmaster/goflow"
	"go.uber.org/zap"
)

//Confirm被认为是业务，不在告警机制上处理，在ActiveAlarm处处理
type AlarmStater struct {
	flow.Component
	In  chan *domain.COA
	Out chan<- *domain.COA //提供过滤后的COA给ActiveAlarm Ponder进行告警缓存维护
}

func (stater *AlarmStater) Process() {

	for coa := range stater.In {
		alarm, existAlarm := shadow.GetAlarmByKey(coa.AlarmKey)
		if existAlarm {
			//更新测点和告警状态(同时维护告警日志的状态，确保时序正常发送给后端)
			stater.Normalize(coa, alarm)

			//更新point的告警状态
			pointKey := strconv.Itoa(alarm.DeviceID) + "." + strconv.Itoa(alarm.PointID)
			point, existPoint := shadow.GetPointByKey(pointKey)
			if existPoint {
				if alarm.StartTime > 0 && alarm.EndTime == 0 {
					point.AlarmSeverity = alarm.EventSeverity
				} else {
					point.AlarmSeverity = 0
				}
			}
		}
	}
}

func (stater *AlarmStater) Normalize(coa *domain.COA, alarm *domain.Alarm) {
	//1 如果告警对象未开始(starttime=0)
	//如果COA开始(StartTime>0,EndTime=0)，则更新告警对象为开始状态
	if alarm.StartTime == 0 && coa.StartTime > 0 && coa.EndTime == 0 {
		alarm.StartTime = coa.StartTime
		alarm.EndTime = 0
		stater.Out <- coa //发送告警开始
	}

	//2 如果告警对象未开始(startTime=0)
	//如果COA结束告警，开始(StartTime>0,EndTime>0)
	//则更新则告警对象为结束状态
	if alarm.StartTime == 0 && coa.StartTime > 0 && coa.EndTime > 0 {
		alarm.StartTime = coa.StartTime
		alarm.EndTime = coa.EndTime

		coaStart := coa.Clone() //send start coa
		coaStart.EndTime = 0
		stater.Out <- coaStart

		coaEnd := coa.Clone() //send end coa
		coaEnd.EndTime = 0
		stater.Out <- coaEnd
	}

	//3 如果告警对象开始(startTime>0 && endTime== 0)
	//如果COA为开始告警，开始(StartTime>0,EndTime=0)
	if alarm.StartTime > 0 && alarm.EndTime == 0 && coa.StartTime > 0 && coa.EndTime == 0 {
		//两个告警开始时间不一致，说明不是一个触发点的事件或校时问题
		//因为告警对象已经告警开始，如果更改告警对象的开始时间，活动告警匹配和历史告警表将出错
		if coa.StartTime != alarm.StartTime {
			global.GVA_LOG.Warn("start coa's start time mismatch with alarm : ",
				zap.String("COA", coa.ToString()), zap.String("Alarm", alarm.ToString()))
			//这时不做任何动作，如果cov的开始时间早的话事实上缩短了上报告警的持续时长，但没办法
			//因为告警已经乱序了，后开始的先到,而且已经开用了，这里更改弥补的代价太大（尽量底层保证不乱序）
		} else {
			//TODO 重复发送一条告警的开始，无确认的话，是否是反转业务？如果是需要Overturn +1
			//因为Alarm已经开始（代表已经发过给后续node最新coa，所以这里不需要重复发
		}
	}

	//4 如果告警对象开始(startTime>0 && endTime== 0)
	//如果COA为结束告警(StartTime>0,EndTime>0 )
	if alarm.StartTime > 0 && alarm.EndTime == 0 && coa.StartTime > 0 && coa.EndTime > 0 {

		//如果告警对象开始时间大于coa结束时间，说明COA是历史的一条告警
		//coa 代表告警又重新开始了，coa不能影响现有告警，策略：（1 抛弃，2 存历史）
		if alarm.StartTime >= coa.EndTime {
			global.GVA_LOG.Warn("get a old end cov : ",
				zap.String("COA", coa.ToString()), zap.String("Alarm", alarm.ToString()))
			//not send to Out chan，事实执行抛弃动作
		} else {
			//两个告警开始时间不一致，说明不是一个触发点的事件或校时问题
			//因为告警对象已经告警开始，如果更改告警对象的开始时间，活动告警匹配和历史告警表将出错
			if coa.StartTime != alarm.StartTime {
				global.GVA_LOG.Warn("end coa's start time mismatch with alarm : ",
					zap.String("COA", coa.ToString()), zap.String("Alarm", alarm.ToString()))
				//这时不做任何动作，事实上缩短了上报告警的持续时长，但没办法
				//因为告警已经乱序了，后开始的先到,而且已经开用了，这里更改弥补的代价太大（尽量底层保证不乱序）
			}

			//如果上面都考虑过，那么cov就代表一个正常的结束
			alarm.EndTime = coa.EndTime

			coaEnd := coa.Clone()
			coaEnd.StartTime = alarm.StartTime
			stater.Out <- coaEnd //send coa
		}
	}

	//5 如果告警对象结束(startTime>0 && endTime>0)
	//如果COA为开始告警，开始(StartTime>0,EndTime=0)
	if alarm.StartTime > 0 && alarm.EndTime > 0 && coa.StartTime > 0 && coa.EndTime == 0 {
		//如果新的cov为旧记录，或在旧告警时间段的新发生cov，则抛弃
		//这样可能导致告警量减少，一个告警同时只有一条在用，但业务清晰
		//coa.EndTime <= alarm.StartTime || coa.StartTime < alarm.EndTime 会被抛弃
		if coa.StartTime >= alarm.EndTime { //如果两条告警彻底分开才认为是新告警
			alarm.EndTime = 0
			alarm.StartTime = coa.StartTime
			stater.Out <- coa
		}
	}

	//6 如果告警对象结束(startTime>0 && endTime>0)
	//如果COA为结束告警，开始(StartTime>0,EndTime=0)
	if alarm.StartTime > 0 && alarm.EndTime > 0 && coa.StartTime > 0 && coa.EndTime > 0 {

		//如果coa代表是未来的则进行更新，否则则抛弃（不管）
		if coa.StartTime > alarm.EndTime {
			coaStart := coa.Clone() //send start coa
			coaStart.EndTime = 0
			stater.Out <- coaStart

			coaEnd := coa.Clone() //send end coa
			coaEnd.EndTime = 0
			stater.Out <- coaEnd
		}
	}
}
