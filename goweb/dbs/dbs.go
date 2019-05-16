package dbs

import (
	"database/sql"
	"vms-go/goweb/logger"
	"vms-go/goweb/models"
)

type DBService struct {
	mysql *MysqlIface
}

var _dbService *DBService

func GetDBService() *DBService {
	return _dbService
}

func NewDBService() error {
	_dbService = &DBService{}
	var err error
	_dbService.mysql, err = NewMysqlIface()
	if err != nil {
		return err
	}

	return nil
}

func (db *DBService) Destory() {
	db.mysql.Close()
}

func (db *DBService) QueryLoggerStats(channel string, sdate, edate string, replys *[]*models.LogStats) {
	var rows *sql.Rows
	var err error
	if channel == "admin" {
		sqlstr := "select channel, gameid, newly, tow_pr, three_pr, seven_pr, retention, logdate from log_stat where logdate>=? and logdate<=?"
		rows, err = db.mysql.QueryV2(sqlstr, sdate, edate)
	} else {
		sqlstr := "select channel, gameid, newly, tow_pr, three_pr, seven_pr, retention, logdate from log_stat where channel=? and logdate>=? and logdate<=?"
		rows, err = db.mysql.QueryV2(sqlstr, channel, sdate, edate)
	}

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		logger.Error("Query: failed: ", err)
		return
	}

	for rows.Next() {
		stats := &models.LogStats{}
		if err := rows.Scan(&stats.Channel, &stats.Game, &stats.Newly, &stats.TowPr, &stats.ThreePr, &stats.SevenPr, &stats.Retention, &stats.LogDate); err != nil {
			logger.Error("Scan: failed: ", err)
			continue
		}

		*replys = append(*replys, stats)
	}
}

func (db *DBService) ReplaceLoggerStats(channel string, game string, newly, tow_pr, three_pr, seven_pr, retention, logdate string) {
	sqlstr := "replace into log_stat values(?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.mysql.ReplaceV2(sqlstr, channel, game, newly, tow_pr, three_pr, seven_pr, retention, logdate)
	if err != nil {
		logger.Error("ReplaceLoggerStats ", err)
		return
	}
	// exec := db.mysql.ReplaceFuture(sqlstr, channel, game, newly, tow_pr, three_pr, seven_pr, retention, logdate)
	// _, err := exec()
	// if err != nil {
	// 	logger.Error("ReplaceLoggerStats ", err)
	// 	return
	// }
}
