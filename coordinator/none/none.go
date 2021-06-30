package none

import (
	"github.com/leechanx/ekko-idgenerator/config"
)

type NoCoordinator struct {
}

func NewNoneCoordinator() (*NoCoordinator, error) {
	return &NoCoordinator{}, nil
}

func (*NoCoordinator) GetWorkerId() (uint64, error) {
	return uint64(config.GetFixedWorkerID()), nil
}

func (*NoCoordinator) GetLeastTime() (int64, error) {
	return 0, nil
}

func (*NoCoordinator) ReportTime(timestamp int64) error {
	return nil
}

func (*NoCoordinator) GetAllTimes() (map[string]int64, error) {
	return nil, nil
}
