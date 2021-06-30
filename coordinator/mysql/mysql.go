package mysql

import (
	"database/sql"
	"fmt"

	"github.com/leechanx/ekko-idgenerator/config"
	errors "github.com/leechanx/ekko-idgenerator/ekko_errors"
	"github.com/leechanx/ekko-idgenerator/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/smallnest/rpcx/log"
)

type MySQLCoordinator struct {
	db         *sql.DB
	macAddress uint64
}

func NewMySQLCoordinator() (*MySQLCoordinator, error) {
	db, err := sql.Open("mysql", config.GetDSN())
	if err != nil {
		log.Errorf("candidate=mysql open mysql address(%s) meet error: %v\n", config.GetDSN(), err)
		return nil, errors.WithMsg(err.Error())
	}
	macAddress := util.GetMacAddress()
	if macAddress == 0 {
		log.Errorf("candidate=mysql get mac address error\n")
		return nil, errors.WithMsg("can't get MAC address")
	}
	var count int
	// check if there is record
	err = db.QueryRow(checkTimeReportQuery, macAddress).Scan(&count)
	if err != nil {
		log.Errorf("candidate=mysql query time report error: %v\n", err)
		return nil, err
	}
	if count == 0 {
		// insert new time report
		_, err = db.Exec(insertTimeReportExec, macAddress)
		if err != nil {
			log.Errorf("candidate=mysql insert new time report error: %v\n", err)
			return nil, err
		}
	}
	return &MySQLCoordinator{db: db, macAddress: macAddress}, nil
}

func (my *MySQLCoordinator) GetWorkerId() (uint64, error) {
	var workerID uint64
	err := my.db.QueryRow(workerIDQuery, my.macAddress).Scan(&workerID)
	if err == nil {
		log.Infof("candidate=mysql got the workerID %d\n", workerID)
		return workerID, nil
	}
	if err != sql.ErrNoRows {
		log.Errorf("candidate=mysql get workerID error: %v\n", err)
		return 0, errors.WithMsg(err.Error())
	}
	ret, err := my.db.Exec(genWorkerIDExec, my.macAddress)
	if err != nil {
		log.Errorf("candidate=mysql generate workerID error: %v\n", err)
		return 0, errors.WithMsg(err.Error())
	}
	index, err := ret.LastInsertId()
	if err != nil {
		log.Errorf("candidate=mysql get workerID error: %v\n", err)
		return 0, errors.WithMsg(err.Error())
	}
	log.Infof("candidate=mysql apply for a new workerID %d\n", index)
	return uint64(index), nil
}

func (my *MySQLCoordinator) GetLeastTime() (int64, error) {
	var timestamp int64
	err := my.db.QueryRow(getTimeReportQuery, my.macAddress).Scan(&timestamp)
	if err == nil || err == sql.ErrNoRows {
		return timestamp, nil
	}
	log.Errorf("candidate=mysql get least time meet error: %v\n", err)
	return 0, err
}

func (my *MySQLCoordinator) ReportTime(timestamp int64) error {
	res, err := my.db.Exec(updateTimeReportExec, timestamp, my.macAddress, timestamp)
	if err != nil {
		log.Warnf("candidate=mysql report time error: %v\n", err)
		return errors.WithMsg(err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Warnf("candidate=mysql report time error: %v\n", err)
		return err
	}
	if rows == 0 {
		log.Warnf("candidate=mysql report time %d ignored\n", timestamp)
		return errors.WithMsg("Report time ignored")
	}
	return nil
}

func (my *MySQLCoordinator) GetAllTimes() (map[string]int64, error) {
	res, err := my.db.Query(getAllTimeReportQuery)
	if err != nil {
		log.Error("candidate=mysql get all time report meet error: %v\n", err)
		return nil, err
	}
	data := map[string]int64{}
	for res.Next() {
		var address uint64
		var timestamp int64
		if err := res.Scan(&address, &timestamp); err != nil {
			log.Error("candidate=mysql get all time report meet error: %v\n", err)
			return nil, err
		}
		if address != my.macAddress {
			key := fmt.Sprint(address)
			data[key] = timestamp
		}
	}
	return data, nil
}
