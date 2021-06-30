package coordinator

import (
	"time"
	"github.com/smallnest/rpcx/log"
	errors "github.com/leechanx/ekko-idgenerator/ekko_errors"
)

var clockDriftChecker ClockDriftChecker

type NoClockDriftChecker struct {
}

func (NoClockDriftChecker) CheckOtherNodes(coordinator Coordinator) error {
	return nil
}

type StoreClockDriftChecker struct {
}

func (StoreClockDriftChecker) CheckOtherNodes(coordinator Coordinator) error {
	firstData, err := coordinator.GetAllTimes()
	if err != nil {
		return err
	}
	// wait
	<-time.After(reportIntervals * time.Millisecond)
	secondData, err := coordinator.GetAllTimes()
	if err != nil {
		return err
	}
	now := time.Now().UnixNano() / 1000000
	var otherTimestamps []int64
	for key, firstValue := range firstData {
		secondValue, ok := secondData[key]
		if ok && secondValue > firstValue {
			otherTimestamps = append(otherTimestamps, secondValue)
		}
	}
	if len(otherTimestamps) != 0 {
		var diff int64
		for _, timestamp := range otherTimestamps {
			if timestamp > now {
				diff += timestamp - now
			} else {
				diff += now - timestamp
			}
		}
		diff /= int64(len(otherTimestamps))
		if diff > minOtherNecessaryTimeDiff {
			log.Errorf("remote nodes' times = %v, but now time is %d\n", otherTimestamps, now)
			return &errors.ClockDrift
		}
	}
	return nil
}

type ETCDClockDriftChecker struct {
}

func (ETCDClockDriftChecker) CheckOtherNodes(coordinator Coordinator) error {
	data, err := coordinator.GetAllTimes()
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}
	now := time.Now().UnixNano() / 1000000
	var diff int64
	for _, timestamp := range data {
		if timestamp > now {
			diff += timestamp - now
		} else {
			diff += now - timestamp
		}
	}
	diff /= int64(len(data))
	if diff > minOtherNecessaryTimeDiff {
		log.Errorf("remote nodes' average times = %v, but now time is %d\n", diff, now)
		return &errors.ClockDrift
	}
	return nil
}
