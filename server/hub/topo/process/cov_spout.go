package process

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	flow "github.com/trustmaster/goflow"
)

type COVSpout struct {
	flow.Component
	In         chan []*domain.COV
	COVState   chan<- []*domain.COV
	COVCascade chan<- []*domain.COV
}

func (spout *COVSpout) Process() {

	for covs := range spout.In {
		spout.COVState <- covs   //本地处理COV实时数据
		spout.COVCascade <- covs //上送COV
	}
}
