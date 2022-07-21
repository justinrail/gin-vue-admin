package stub

import (
	"math/rand"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/bus"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/timer"
)

func Ready() {
	rand.Seed(time.Now().UnixNano())
	LoadMockCOGs()
	LoadMockCODs()
	LoadMockCOVs()
	LoadMockCOAs()
	LoadMockCOSs()
}

func Start() {
	// timer.NewTimerTask().AddTaskByFunc("func.mock.cog", "@every 5s", func() {
	// 	bus.COGBus <- randomCOG()
	// })

	// timer.NewTimerTask().AddTaskByFunc("func.mock.cod", "@every 3s", func() {
	// 	bus.CODBus <- randomCOD()
	// })

	timer.NewTimerTask().AddTaskByFunc("func.mock.cov", "@every 10s", func() {
		bus.COVBus <- randomCOVs(100)
		//fmt.Printf("covbus count: %d\n", bus.COVBus)
	})

	// timer.NewTimerTask().AddTaskByFunc("func.mock.coa", "@every 2s", func() {
	// 	bus.COABus <- randomCOA()
	// })

	// dataIn <- RandomCOS()
}
