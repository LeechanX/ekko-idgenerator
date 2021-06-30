package generator

import (
	"github.com/leechanx/ekko-idgenerator/config"
	errors "github.com/leechanx/ekko-idgenerator/ekko_errors"
)

func Gen(prod int32) (uint64, *errors.RuntimeError) {
	if prod < 0 || prod > config.GetMaxProdNumber() {
		return 0, &errors.InvalidRequest
	}

	g := idGenerators[prod]
	g.lock.Lock()
	defer g.lock.Unlock()
	return g.genTimestamp()
}

func MultiGen(prod int32, count uint32) (lowerUid, upperUid uint64,
	realCount uint32, err *errors.RuntimeError) {
	if prod < 0 || prod > config.GetMaxProdNumber() {
		err = &errors.InvalidRequest
		return
	}
	if count > maxMultiGetCount || count == 0 {
		realCount = maxMultiGetCount
	} else {
		realCount = count
	}
	g := idGenerators[prod]
	g.lock.Lock()
	defer g.lock.Unlock()

	if count == 1 {
		lowerUid, err = g.genTimestamp()
		upperUid = lowerUid
	} else {
		lowerUid, upperUid, err = g.mGenTimestamp(uint64(realCount))
	}
	return
}
