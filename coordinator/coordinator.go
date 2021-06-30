package coordinator

import (
	"time"

	"github.com/leechanx/ekko-idgenerator/config"
	"github.com/leechanx/ekko-idgenerator/coordinator/etcd"
	"github.com/leechanx/ekko-idgenerator/coordinator/mysql"
	"github.com/leechanx/ekko-idgenerator/coordinator/none"
	"github.com/leechanx/ekko-idgenerator/coordinator/redis"
	errors "github.com/leechanx/ekko-idgenerator/ekko_errors"

	"github.com/smallnest/rpcx/log"
)

const (
	reportIntervals           = 3000
	minLocalNecessaryTimeDiff = reportIntervals * 2
	minOtherNecessaryTimeDiff = reportIntervals
	leaseTTL                  = (reportIntervals * 3) / 1000
)

func NewCoordinator() (Coordinator, error) {
	if config.GetCoordinatorType() == config.CoordinatorTypeNone {
		clockDriftChecker = &NoClockDriftChecker{}
		return none.NewNoneCoordinator()
	} else if config.GetCoordinatorType() == config.CoordinatorTypeMysql {
		clockDriftChecker = &StoreClockDriftChecker{}
		return mysql.NewMySQLCoordinator()
	} else if config.GetCoordinatorType() == config.CoordinatorTypeEtcd {
		clockDriftChecker = &ETCDClockDriftChecker{}
		return etcd.NewEtcdCoordinator(leaseTTL)
	} else if config.GetCoordinatorType() == config.CoordinatorTypeRedis {
		clockDriftChecker = &StoreClockDriftChecker{}
		return redis.NewRedisCoordinator()
	}
	panic("impossible")
}

// reporter
func Reporter(coordinator Coordinator) {
	go func() {
		ticker := time.NewTicker(reportIntervals * time.Millisecond)
		for {
			now := <-ticker.C
			coordinator.ReportTime(now.UnixNano() / 1000000)
		}
	}()
}

// check time
func CheckClockDrift(coordinator Coordinator) error {
	err := checkMyTimestamp(coordinator)
	if err == nil && config.IsOpenTimeCheckSwitch() {
		log.Infof("check remote nodes now...")
		err = clockDriftChecker.CheckOtherNodes(coordinator)
	}
	return err
}

// check my time
func checkMyTimestamp(coordinator Coordinator) error {
	now := time.Now().UnixNano() / 1000000
	leastTimestamp, err := coordinator.GetLeastTime()
	if err != nil {
		return err
	}
	if leastTimestamp == 0 {
		return nil
	}
	if leastTimestamp+minLocalNecessaryTimeDiff > now {
		log.Errorf("get least time %d, but now time is %d\n", leastTimestamp, now)
		return &errors.ClockDrift
	}
	return nil
}
