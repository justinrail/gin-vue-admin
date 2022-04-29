package bus

import "github.com/flipped-aurora/gin-vue-admin/server/hub/domain"

const (
	//QueueSize 队列大小
	QueueSize = 1024
)

//COGBus bus channel for COG
var COGBus chan *domain.COG

//CORBus bus channel for COR
var CODBus chan *domain.COD

//COVBus bus channel for COV
var COVBus chan []*domain.COV

//COABus bus channel for COA
var COABus chan *domain.COA

//COSBus bus channel for COS
var COSBus chan *domain.COS

func init() {
	COGBus = make(chan *domain.COG, QueueSize)
	CODBus = make(chan *domain.COD, QueueSize)
	COVBus = make(chan []*domain.COV, QueueSize)
	COABus = make(chan *domain.COA, QueueSize)
	COSBus = make(chan *domain.COS, QueueSize)
}
