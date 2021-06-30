package redis

import (
	"context"
	"strconv"

	"github.com/leechanx/ekko-idgenerator/config"
	errors "github.com/leechanx/ekko-idgenerator/ekko_errors"
	"github.com/leechanx/ekko-idgenerator/util"

	redisV8 "github.com/go-redis/redis/v8"
	"github.com/smallnest/rpcx/log"
)

type RedisCoordinator struct {
	client          *redisV8.Client
	macAddressStr   string
	leastReportTime int64
}

func NewRedisCoordinator() (*RedisCoordinator, error) {
	redisClient := redisV8.NewClient(&redisV8.Options{
		Addr: config.GetRedisAddr(),
	})
	macAddress := util.GetMacAddress()
	if macAddress == 0 {
		log.Errorf("candidate=redis get mac address meet error\n")
		return nil, errors.WithMsg("can't get MAC address")
	}
	return &RedisCoordinator{
		client:        redisClient,
		macAddressStr: strconv.FormatUint(macAddress, 10),
	}, nil
}

func (r *RedisCoordinator) GetWorkerId() (uint64, error) {
	// try to find workerID in redis
	scmd := r.client.HGet(util.GetTimeoutContext(), redisWorkerIDKey, r.macAddressStr)
	if scmd.Err() != nil && scmd.Err() != redisV8.Nil {
		log.Errorf("candidate=redis hget mac address meet error: %v\n", scmd.Err())
		return 0, scmd.Err()
	}
	if scmd.Err() == nil {
		workerID, err := strconv.ParseUint(scmd.Val(), 0, 64)
		if err != nil {
			log.Errorf("candidate=redis parse workerID meet error: %v\n", err)
		} else {
			log.Infof("candidate=redis got the workerID %d\n", workerID)
		}
		return workerID, err
	}
	// generate new workerID
	icmd := r.client.Incr(util.GetTimeoutContext(), redisWorkerIDGenKey)
	if icmd.Err() != nil {
		log.Errorf("candidate=redis incr meet error: %v\n", icmd.Err())
		return 0, icmd.Err()
	}
	workerID := uint64(icmd.Val())
	wcmd := r.client.HSet(context.TODO(), redisWorkerIDKey, r.macAddressStr, workerID)
	if wcmd.Err() != nil {
		log.Errorf("candidate=redis hset meet error: %v\n", wcmd.Err())
		return 0, wcmd.Err()
	}
	log.Infof("candidate=redis apply for a new workerID %d\n", workerID)
	return workerID, nil
}

func (r *RedisCoordinator) GetLeastTime() (int64, error) {
	scmd := r.client.HGet(util.GetTimeoutContext(), redisTimestampReportKey, r.macAddressStr)
	if scmd.Err() == redisV8.Nil {
		return 0, nil
	}
	if scmd.Err() != nil {
		log.Errorf("candidate=redis hget least time meet error: %v\n", scmd.Err())
		return 0, scmd.Err()
	}
	leastReportTime, err := strconv.ParseInt(scmd.Val(), 0, 64)
	if err != nil {
		log.Errorf("candidate=redis parse least time meet error: %v\n", err)
		return 0, err
	}
	r.leastReportTime = leastReportTime
	return leastReportTime, nil
}

func (r *RedisCoordinator) ReportTime(timestamp int64) error {
	if timestamp <= r.leastReportTime {
		log.Warnf("candidate=redis ignore report time %d, least report time %d\n", timestamp, r.leastReportTime)
		return errors.WithMsg("Report time ignored")
	}
	intcmd := r.client.HSet(util.GetTimeoutContext(), redisTimestampReportKey, r.macAddressStr, timestamp)
	if intcmd.Err() != nil {
		log.Warnf("candidate=redis hset time report meet error: %v\n", intcmd.Err())
		return intcmd.Err()
	}
	return nil
}

func (r *RedisCoordinator) GetAllTimes() (map[string]int64, error) {
	results := r.client.HGetAll(util.GetTimeoutContext(), redisTimestampReportKey)
	if results.Err() != nil {
		log.Error("candidate=redis get all time report error: %v\n", results.Err())
		return nil, results.Err()
	}
	data := map[string]int64{}
	for key, value := range results.Val() {
		if key == r.macAddressStr {
			continue
		}
		timestamp, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			log.Error("candidate=redis get all time report parse timestamp meet error: %v\n", err)
			return nil, err
		}
		data[key] = timestamp
	}
	return data, nil
}
