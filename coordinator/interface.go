package coordinator

type Coordinator interface {
	GetWorkerId() (uint64, error)
	GetLeastTime() (int64, error)
	ReportTime(int64) error
	GetAllTimes() (map[string]int64, error)
}

type ClockDriftChecker interface {
	CheckOtherNodes(Coordinator) error
}
