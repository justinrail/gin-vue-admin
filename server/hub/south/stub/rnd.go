package stub

import (
	"math/rand"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/hub/base"
	"github.com/flipped-aurora/gin-vue-admin/server/hub/domain"
)

//RandomCOG generate random cog
func randomCOG() *domain.COG {
	cog := cogCache[rand.Intn(len(cogCache))]

	//暂时放用到的，没有测试用途的暂不放入
	flags := []int{base.GatewayFlagUnkown, base.GatewayFlagRegister, base.GatewayFlagTimeCheckAck,
		base.GatewayFlagHeartbeat, base.GatewayFlagHeartbeat, base.GatewayFlagRebootAck}

	rndFlag := flags[rand.Intn(len(flags))]

	cog.Flag = rndFlag
	cog.Timestamp = time.Now().Unix()

	return cog
}

func randomCOD() *domain.COD {
	cod := codCache[rand.Intn(len(codCache))]

	//暂时放用到的，没有测试用途的暂不放入
	flags := []int{base.DeviceConStateUnkown, base.DeviceConStateOnline, base.DeviceConStateOffline,
		base.DeviceDebugStateOn, base.GatewayDebugStateOff}

	rndFlag := flags[rand.Intn(len(flags))]

	cod.Flag = rndFlag
	cod.UpdateTime = time.Now().Unix()

	return cod
}

//randomCOV generate random cov
func randomCOV() *domain.COV {
	indexRnd := rand.Intn(len(covCache))
	cov := covCache[indexRnd]
	if cov != nil {
		cov.CurrentNumericValue = rand.Float32() * 20
		//StringValue not mock
		cov.Timestamp = time.Now().Unix()
	}

	return cov

}

func randomCOVs(count int) []*domain.COV {
	covs := make([]*domain.COV, count)

	for i := 0; i < count; i++ {
		cov := randomCOV()
		if cov != nil {
			covs = append(covs, cov)
		}
	}

	return covs
}

//RandomCOA random coa packet
func randomCOA() *domain.COA {
	coa := coaCache[rand.Intn(len(coaCache))]

	coa.NumericValue = rand.Float32() * 20
	//StringValue not mock

	//根据COA的上次状态，如果上次没有发送过则发送开始
	if coa.StartTime == 0 {
		coa.StartTime = time.Now().Unix() - int64(rand.Intn(20000))
		coa.EndTime = 0
	} else if coa.EndTime == 0 { //如果上次发送过开始，则发送结束，并清除StartTime和EndTime
		coa.EndTime = time.Now().Unix()
	} else { //如果开始和结束都生成过，重置, 生成新的告警
		coa.StartTime = time.Now().Unix() - int64(rand.Intn(20000))
		coa.EndTime = 0
	}

	return coa
}
