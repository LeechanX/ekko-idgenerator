package system

import (
	"os"
	"syscall"

	"github.com/smallnest/rpcx/log"
)

var flock *os.File

func RequireFlock() error {
	var err error
	flock, err = os.Create(flockPath)
	if err != nil {
		log.Errorf("module=flock create or open file %s failed: %s\n", flockPath, err.Error())
		return err
	}

	if err = syscall.Flock(int(flock.Fd()), flockFlag); err != nil {
		log.Errorf("module=flock require flock %s failed: %s\n", flockPath, err.Error())
		return err
	}
	return nil
}

func ReleaseFlock() error {
	err := syscall.Flock(int(flock.Fd()), flockUnlockFlag)
	if err != nil {
		log.Errorf("module=flock release flock %s failed: %s\n", flockPath, err.Error())
		return err
	}
	return nil
}
