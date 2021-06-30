package mysql

const (
	checkTimeReportQuery  = "SELECT count(*) FROM time_report WHERE mac_address = ?"
	insertTimeReportExec  = "INSERT INTO time_report(mac_address, ts) VALUES(?, 0)"
	workerIDQuery         = "SELECT id FROM global_worker_id WHERE mac_address = ?"
	genWorkerIDExec       = "INSERT INTO global_worker_id(mac_address) VALUES(?)"
	getTimeReportQuery    = "SELECT ts FROM time_report WHERE mac_address = ?"
	getAllTimeReportQuery = "SELECT mac_address, ts FROM time_report"
	updateTimeReportExec  = "UPDATE time_report SET ts = ? WHERE mac_address = ? AND ts < ?"
)
