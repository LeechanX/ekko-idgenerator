package main

import (
	"context"
	"flag"

	"github.com/leechanx/ekko-idgenerator/config"
	"github.com/leechanx/ekko-idgenerator/coordinator"
	"github.com/leechanx/ekko-idgenerator/generator"
	"github.com/leechanx/ekko-idgenerator/protobuf"
	"github.com/leechanx/ekko-idgenerator/system"

	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/server"
)

var (
	addr       = flag.String("addr", "localhost:8972", "server address")
	idcNumber  = flag.Int("idc", -1, "idc number")
	configPath = flag.String("conf_path", "conf/ekko_server.toml", "configure path")
)

type Ekko struct {
}

func (e *Ekko) GetUniqueID(ctx context.Context,
	req *protobuf.GetUniqueIDRequest,
	rsp *protobuf.GetUniqueIDResponse) error {
	uid, err := generator.Gen(req.GetProduct())
	if err != nil {
		log.Errorf("module=main GetUniqueID return error: %s\n", err.Error())
		return err
	}
	rsp.Uid = uid
	return nil
}

func (e *Ekko) MGetUniqueID(ctx context.Context,
	req *protobuf.MGetUniqueIDRequest,
	rsp *protobuf.MGetUniqueIDResponse) error {
	lowerUid, upperUid, count, err := generator.MultiGen(req.GetProduct(), req.GetCount())
	if err != nil {
		log.Errorf("module=main MGetUniqueID return error: %s\n", err.Error())
		return err
	}
	rsp.LowerUid = lowerUid
	rsp.UpperUid = upperUid
	rsp.Count = count
	return nil
}

func main() {
	flag.Parse()
	if *idcNumber == -1 {
		panic("idc number should be set")
	}
	// server address, configure path, idc
	if err := config.Parse(*configPath); err != nil {
		panic(err)
	}

	// require flock
	if err := system.RequireFlock(); err != nil {
		panic(err)
	}
	defer system.ReleaseFlock()

	ekkoCoordinator, err := coordinator.NewCoordinator()
	if err != nil {
		panic(err)
	}
	if err := coordinator.CheckClockDrift(ekkoCoordinator); err != nil {
		panic(err)
	}
	workerID, err := ekkoCoordinator.GetWorkerId()
	if err != nil {
		panic(err)
	}
	err = generator.InitIdGenerator(uint64(*idcNumber), workerID)
	if err != nil {
		panic(err)
	}
	// start report
	coordinator.Reporter(ekkoCoordinator)

	s := server.NewServer()
	s.Register(&Ekko{}, "")
	s.Serve("tcp", *addr)
}
