package cascade

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/bus"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
)

func handleCOV(topic string, data []byte) {
	var covs []*domain.COV
	err := json.Unmarshal(data, &covs)
	if err != nil {
		fmt.Println(err)
	} else {
		bus.COVBus <- covs
	}
}
