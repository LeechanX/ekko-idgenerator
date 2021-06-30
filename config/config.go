package config

import (
	"time"

	errors "github.com/leechanx/ekko-idgenerator/ekko_errors"

	"github.com/BurntSushi/toml"
	"github.com/smallnest/rpcx/log"
)

type EkkoConfig struct {
	ReleaseDate         string `toml:"release_date"`
	MaxProductionNumber int32  `toml:"max_production_number"`
	FixedWorkerID       int64  `toml:"fixed_worker_id"`
	CoordinatorType     string `toml:"coordinator_type"`
	OpenTimeCheckSwitch bool   `toml:"open_time_check_switch"`
	ReleaseTimestamp    uint64

	Etcd  EtcdConfig  `toml:"etcd"`
	Mysql MySQLConfig `toml:"mysql"`
	Redis RedisConfig `toml:"redis"`
}

var (
	config   EkkoConfig
	checkers = map[string]CoordinatorConfig{}
)

func Parse(path string) error {
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return err
	}
	checkers[CoordinatorTypeMysql] = &config.Mysql
	checkers[CoordinatorTypeEtcd] = &config.Etcd
	checkers[CoordinatorTypeRedis] = &config.Redis
	return check()
}

func check() error {
	if config.ReleaseDate == "" {
		log.Error("module=configure coordinator type = 'none' should set 'base.release_date'")
		return errors.WithMsg("'release_date' is empty")
	}

	if config.MaxProductionNumber < 0 {
		log.Error("module=configure max_production_number should >= 0")
		return errors.WithMsg("'max_production_number' is invalid")
	}

	t, err := time.ParseInLocation(layout, config.ReleaseDate, time.Local)
	if err != nil {
		return err
	}

	config.ReleaseTimestamp = uint64(t.UnixNano() / 1000000)
	log.Infof("module=configure selected coordinator type is %s\n", config.CoordinatorType)
	if config.CoordinatorType == CoordinatorTypeNone {
		if config.FixedWorkerID < 1 {
			log.Error("module=configure coordinator type = 'none' should set 'base.fixed_worker_id' >= 1")
			return errors.WithMsg("'fixed_worker_id' error")
		}
		return nil
	} else if checker, ok := checkers[config.CoordinatorType]; ok {
		return checker.Check()
	}
	log.Error("module=configure coordinator type can only be 'none' or 'mysql' or 'etcd' or 'redis'")
	return errors.WithMsg("coordinator type error")
}

func GetMaxProdNumber() int32 {
	return config.MaxProductionNumber
}

func GetReleaseTimestamp() uint64 {
	return config.ReleaseTimestamp
}

func GetCoordinatorType() string {
	return config.CoordinatorType
}

func GetFixedWorkerID() int64 {
	return config.FixedWorkerID
}

func IsOpenTimeCheckSwitch() bool {
	return config.OpenTimeCheckSwitch
}

func GetClientAddrs() []string {
	return config.Etcd.ClientAddrs
}

func GetDSN() string {
	return config.Mysql.Dsn
}

func GetRedisAddr() string {
	return config.Redis.ServerAddr
}
