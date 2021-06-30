package etcd

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/leechanx/ekko-idgenerator/config"
	errors "github.com/leechanx/ekko-idgenerator/ekko_errors"
	"github.com/leechanx/ekko-idgenerator/util"

	"github.com/smallnest/rpcx/log"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdCoordinator struct {
	client        *clientv3.Client
	leaseID       clientv3.LeaseID
	myWorkerIDKey string
	myReporterKey string
}

func NewEtcdCoordinator(ttl int64) (*EtcdCoordinator, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   config.GetClientAddrs(),
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Errorf("candidate=etcd build etcd address(%v) connection error: %v\n", config.GetClientAddrs(), err)
		return nil, err
	}
	macAddress := util.GetMacAddress()
	if macAddress == 0 {
		log.Errorf("candidate=etcd get mac address error\n")
		return nil, errors.WithMsg("can't get MAC address")
	}
	lease := clientv3.NewLease(client)
	leaseRsp, err := lease.Grant(util.GetTimeoutContext(), ttl)
	if err != nil {
		log.Errorf("candidate=etcd lease.Grant error: %v", err)
		return nil, err
	}
	leaseID := leaseRsp.ID
	leaseChan, err := lease.KeepAlive(util.GetTimeoutContext(), leaseID)
	if err != nil {
		log.Errorf("candidate=etcd lease.KeepAlive error: %v", err)
		return nil, err
	}
	etcdCoordinator := &EtcdCoordinator{
		client:        client,
		leaseID:       leaseID,
		myWorkerIDKey: fmt.Sprintf(etcdWorkerIdTplKey, macAddress),
		myReporterKey: fmt.Sprintf(etcdTimeReportTplKey, macAddress),
	}
	go etcdCoordinator.listenLeaseRespChan(leaseChan)
	return etcdCoordinator, nil
}

func (e *EtcdCoordinator) GetWorkerId() (uint64, error) {
	ctx := util.GetTimeoutContext()
	// try to put workerID key
	_, err := e.client.Put(ctx, e.myWorkerIDKey, "0")
	if err != nil {
		log.Errorf("candidate=etcd get workerID when put meets error: %v\n", err)
		return 0, err
	}
	// get all worker
	ctx = util.GetTimeoutContext()
	getResp, err := e.client.Get(ctx, etcdPrefixKey, clientv3.WithPrefix())
	if err != nil {
		log.Errorf("candidate=etcd get workerID error: %v\n", err)
		return 0, err
	}
	var workerList []*workerInfo
	for _, worker := range getResp.Kvs {
		workerList = append(workerList, &workerInfo{
			key:            string(worker.Key),
			createRevision: worker.CreateRevision,
		})
	}
	// sort by create revision
	sort.Slice(workerList, func(i, j int) bool {
		return workerList[i].createRevision < workerList[j].createRevision
	})
	for i, worker := range workerList {
		if worker.key == e.myWorkerIDKey {
			log.Infof("candidate=etcd got or applied for the workerID %d\n", i+1)
			return uint64(i + 1), nil
		}
	}
	return 0, errors.WithMsg("impossible")
}

func (e *EtcdCoordinator) GetLeastTime() (int64, error) {
	ctx := util.GetTimeoutContext()
	getResp, err := e.client.Get(ctx, e.myReporterKey)
	if err != nil {
		log.Errorf("candidate=etcd get report time error: %v\n", err)
		return 0, err
	}
	if getResp.Count == 0 {
		ctx = util.GetTimeoutContext()
		if _, err := e.client.Put(ctx, e.myReporterKey, "0", clientv3.WithLease(e.leaseID)); err != nil {
			return 0, err
		} else {
			return 0, nil
		}
	}
	return strconv.ParseInt(string(getResp.Kvs[0].Value), 0, 64)
}

func (e *EtcdCoordinator) ReportTime(timestamp int64) error {
	now := strconv.FormatInt(timestamp, 10)
	compare := clientv3.Compare(clientv3.Value(e.myReporterKey), "<", now)
	_, err := e.client.Txn(util.GetTimeoutContext()).If(
		compare).Then(clientv3.OpPut(e.myReporterKey, now, clientv3.WithLease(e.leaseID))).Commit()
	if err != nil {
		log.Warnf("candidate=etcd report time error: %v\n", err)
	}
	return err
}

func (e *EtcdCoordinator) GetAllTimes() (map[string]int64, error) {
	getResp, err := e.client.Get(util.GetTimeoutContext(), etcdTimeReportPrefixKey, clientv3.WithPrefix())
	if err != nil {
		log.Errorf("candidate=etcd get all time report meet error: %v\n", err)
		return nil, err
	}
	data := map[string]int64{}
	for _, report := range getResp.Kvs {
		key := string(report.Key)
		if key == e.myReporterKey {
			continue
		}
		value := string(report.Value)
		timestamp, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			log.Error("candidate=etcd get all time report parse timestamp meet error: %v\n", err)
			return nil, err
		}
		data[key] = timestamp
	}
	return data, nil
}

func (e *EtcdCoordinator) listenLeaseRespChan(leaseKeepResp <-chan *clientv3.LeaseKeepAliveResponse) {
	for {
		select {
		case rsp := <- leaseKeepResp:
			if rsp == nil {
				log.Error("candidate=etcd close the release function")
				return
			}
		}
	}
}
