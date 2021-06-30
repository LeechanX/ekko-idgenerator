package config

import (
	errors "github.com/leechanx/ekko-idgenerator/ekko_errors"

	"github.com/smallnest/rpcx/log"
)

type CoordinatorConfig interface {
	Check() error
}

type EtcdConfig struct {
	ClientAddrs []string `toml:"client_addrs"`
}

func (e *EtcdConfig) Check() error {
	if len(e.ClientAddrs) == 0 {
		log.Error("module=configure coordinator type = 'etcd' should set 'etcd.client_addrs' with length >= 1")
		return errors.WithMsg("'etcd.client_addrs' is empty")
	}
	return nil
}

type MySQLConfig struct {
	Dsn string `toml:"dsn"`
}

func (m *MySQLConfig) Check() error {
	if m.Dsn == "" {
		log.Error("module=configure coordinator type = 'mysql' should set 'mysql.dsn'")
		return errors.WithMsg("'mysql.dsn' is empty")
	}
	return nil
}

type RedisConfig struct {
	ServerAddr string `toml:"server_addr"`
}

func (r *RedisConfig) Check() error {
	if r.ServerAddr == "" {
		log.Error("module=configure coordinator type = 'redis' should set 'redis.server_addr'")
		return errors.WithMsg("'redis.server_addr' is empty")
	}
	return nil
}
