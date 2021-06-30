package generator

import (
	"fmt"

	"github.com/leechanx/ekko-idgenerator/config"
	errors "github.com/leechanx/ekko-idgenerator/ekko_errors"
	"github.com/leechanx/ekko-idgenerator/system"

	"github.com/smallnest/rpcx/log"
)

var (
	idGenerators []*IDGenerator
)

func InitIdGenerator(idc uint64, workerID uint64) error {
	if idc >= idcUpperBound {
		log.Errorf("module=generator idc is %d, but idc upper bound is %d\n", idc, idcUpperBound)
		return errors.WithMsg("idc number shouldn't exceed 15")
	}
	if config.GetMaxProdNumber() > productUpperBound {
		log.Errorf("module=generator configure max product number is %d, but product upper bound is %d\n", config.GetMaxProdNumber(), productUpperBound)
		return errors.WithMsg("configure's max_product_number shouldn't exceed 16")
	}
	var (
		productBits   uint64
		productNumber int32
	)
	for config.GetMaxProdNumber() > productNumber {
		if productNumber == 0 {
			productNumber = 2
		} else {
			productNumber *= 2
		}
		productBits += 1
	}
	workerIDOffset := workerIDOffsetBase + productBits
	var maxWorkerID uint64 = 1
	for i := uint64(0); i < (idcOffset - workerIDOffset); i++ {
		maxWorkerID *= 2
	}
	maxWorkerID -= 1
	log.Infof("module=generator max worker ID is %d, productBits is %d\n", maxWorkerID, productBits)
	if maxWorkerID < workerID {
		log.Errorf("module=generator workerID is %d, but workerID upper bound is %d\n", workerID, maxWorkerID)
		message := fmt.Sprintf("worker ID shouldn't exceed %d", maxWorkerID)
		return errors.WithMsg(message)
	}

	if config.GetMaxProdNumber() == 0 {
		idGenerator := &IDGenerator{
			releaseTimestamp: config.GetReleaseTimestamp(),
		}
		idGenerator.uidBase |= (workerID << workerIDOffset)
		idGenerator.uidBase |= (idc << idcOffset)
		idGenerators = append(idGenerators, idGenerator)
	} else {
		maxProdNumber := uint64(config.GetMaxProdNumber())
		for prod := uint64(0); prod < maxProdNumber; prod++ {
			idGenerator := &IDGenerator{
				releaseTimestamp: config.GetReleaseTimestamp(),
			}
			idGenerator.uidBase |= (prod << productOffset)
			idGenerator.uidBase |= (workerID << workerIDOffset)
			idGenerator.uidBase |= (idc << idcOffset)
			idGenerators = append(idGenerators, idGenerator)
		}
	}
	// init memory
	for i, _ := range idGenerators {
		mmap, err := system.NewMemory(fmt.Sprintf(mmapFilePathTpl, i))
		if err != nil {
			return err
		}
		idGenerators[i].memory = mmap
	}
	return nil
}
