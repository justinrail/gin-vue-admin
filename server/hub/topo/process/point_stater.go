package process

import (
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/shadow"
	flow "github.com/trustmaster/goflow"
)

type PointStater struct {
	flow.Component
	In chan []*domain.COV
}

func (stater *PointStater) Process() {

	for covs := range stater.In {
		for covIndex := range covs {
			cov := covs[covIndex]
			point, existPoint := shadow.GetPointByKey(cov.PointKey)

			if existPoint {
				point.UpdateData(cov)
			}
		}
	}
}
