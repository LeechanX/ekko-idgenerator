package client

import (
	"github.com/leechanx/ekko-idgenerator/generator"
)

func fetchTimestampInfo(uid uint64) (uint64, uint64) {
	millis := uid >> generator.TimestampOffset
	currency := uid & generator.CurrencyMask
	return millis, currency
}

func spreadOutUidList(lowerUid, upperUid uint64) []uint64 {
	lowerMillis, lowerCurrency := fetchTimestampInfo(lowerUid)
	upperMillis, upperCurrency := fetchTimestampInfo(upperUid)

	var uidList []uint64
	if lowerMillis == upperMillis {
		for currency := lowerCurrency; currency <= upperCurrency; currency++ {
			uidList = append(uidList, lowerUid|currency)
		}
	} else {
		for currency := lowerCurrency; currency < generator.CurrencyBound; currency++ {
			uidList = append(uidList, lowerUid|currency)
		}
		for currency := uint64(0); currency <= upperCurrency; currency++ {
			uidList = append(uidList, upperUid|currency)
		}
	}
	return uidList
}
