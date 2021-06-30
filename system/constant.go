package system

import (
	"os"
	"syscall"
)

const (
	mmapSecondsOffset = 10
	memorySize        = 64
	mmapFlag          = os.O_RDWR | os.O_CREATE
	flockPath         = "/opt/ekko_flock"
	flockFlag         = syscall.LOCK_EX | syscall.LOCK_NB
	flockUnlockFlag   = syscall.LOCK_UN
)
