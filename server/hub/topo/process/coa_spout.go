package process

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	flow "github.com/trustmaster/goflow"
)

type COASpout struct {
	flow.Component
	In         chan *domain.COA
	COAState   chan<- *domain.COA
	COACascade chan<- *domain.COA
}

func (spout *COASpout) Process() {

	for coa := range spout.In {
		spout.COAState <- coa   //本地处理COA
		spout.COACascade <- coa //上送COA
	}
}
