package system

import (
	"bytes"
	"encoding/binary"
	"os"

	"github.com/leechanx/ekko-idgenerator/ekko_errors"

	"github.com/smallnest/rpcx/log"

	"golang.org/x/sys/unix"
)

type MMap struct {
	fileOperator *os.File
	data         []byte
}

func NewMemory(filename string) (*MMap, error) {
	fileOperator, err := os.OpenFile(filename, mmapFlag, 0644)
	if err != nil {
		log.Errorf("module=mmap open mmap file %s error: %v\n", filename, err)
		return nil, err
	}
	state, _ := fileOperator.Stat()
	if state.Size() == 0 {
		_, _ = fileOperator.WriteAt(bytes.Repeat([]byte{0}, memorySize), 0)
		state, _ = fileOperator.Stat()
	}
	data, err := unix.Mmap(int(fileOperator.Fd()), 0, int(state.Size()), unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		log.Errorf("module=mmap execute mmap error: %v\n", err)
		return nil, err
	}
	return &MMap{
		fileOperator: fileOperator,
		data:         data,
	}, nil
}

func (tm *MMap) Write(millis, currency uint64) *ekko_errors.RuntimeError {
	var num uint64
	num |= currency
	num |= (millis << mmapSecondsOffset)

	current := binary.BigEndian.Uint64(tm.data)
	if current >= num {
		return &ekko_errors.ClockDrift
	}
	binary.BigEndian.PutUint64(tm.data, num)
	return nil
}

func (tm *MMap) Close() {
	unix.Munmap(tm.data)
	tm.fileOperator.Close()
}
