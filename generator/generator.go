package generator

import (
	"sync"
	"time"

	errors "github.com/leechanx/ekko-idgenerator/ekko_errors"
	"github.com/leechanx/ekko-idgenerator/system"
)

type IDGenerator struct {
	timestamp        uint64
	currency         uint64
	uidBase          uint64
	releaseTimestamp uint64
	lock             sync.Mutex
	memory           *system.MMap
}

func (g *IDGenerator) genTimestamp() (uint64, *errors.RuntimeError) {
	millis, ok := g.noClockDrift()
	if !ok {
		return 0, &errors.ClockDrift
	}

	var currency uint64
	if millis == g.timestamp {
		if g.currency+1 >= CurrencyBound {
			return 0, &errors.CurrencyExceed
		}
		currency = g.currency + 1
	} else {
		currency = 0
	}
	if err := g.memory.Write(millis, currency); err != nil {
		return 0, err
	}
	// generate uid
	uid := g.gen(millis, currency)
	g.timestamp = millis
	g.currency = currency
	return uid, nil
}

func (g *IDGenerator) noClockDrift() (uint64, bool) {
	now := uint64(time.Now().UnixNano() / 1000000)
	if now < g.releaseTimestamp {
		return 0, false
	}
	millis := now - g.releaseTimestamp
	if millis < g.timestamp {
		return 0, false
	}
	return millis, true
}

func (g *IDGenerator) mGenTimestamp(count uint64) (lowerUid uint64, upperUid uint64, err *errors.RuntimeError) {
	millis, ok := g.noClockDrift()
	if !ok {
		err = &errors.ClockDrift
		return
	}

	var rCurrency, rMillis uint64
	// time is equal with least time
	if millis == g.timestamp {
		if g.currency+count < CurrencyBound {
			// can get results from this millis
			// generate lower ID and upper ID
			lowerUid = g.gen(millis, g.currency+1)
			upperUid = g.gen(millis, g.currency+count)
			// to record result
			rCurrency = g.currency + count
			rMillis = millis
		} else {
			currentMillisCount := CurrencyBound - g.currency - 1
			nextMillisCount := count - currentMillisCount
			// generate lower ID
			if currentMillisCount != 0 {
				lowerUid = g.gen(millis, g.currency+1)
			} else {
				// use next millis
				lowerUid = g.gen(millis+1, 0)
			}
			// generate upper ID use next millis
			upperUid = g.gen(millis+1, nextMillisCount-1)
			// to record result
			rCurrency = nextMillisCount - 1
			rMillis = millis + 1
		}
	} else {
		// generate lower ID and upper ID
		lowerUid = g.gen(millis, 0)
		upperUid = g.gen(millis, count-1)
		// to record result
		rCurrency = count - 1
		rMillis = millis
	}

	err = g.memory.Write(rMillis, rCurrency)
	if err != nil {
		return
	}
	g.timestamp = rMillis
	g.currency = rCurrency
	return
}

func (g *IDGenerator) gen(millis, currency uint64) uint64 {
	uid := g.uidBase
	uid |= millis << TimestampOffset
	uid |= currency
	return uid
}
