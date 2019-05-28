package dbs

import (
	"database/sql"
	"fmt"
	"nuvem/engine/asyncdb/dbmysql"
	"vms-go/goweb/models"
	"vms-go/goweb/settings"

	"github.com/sagacao/goworld/engine/gwlog"
)

type DBService struct {
	// mysql *MysqlIface

	dbmysql *dbmysql.DBMysql
}

var _dbService *DBService

func GetDBService() *DBService {
	return _dbService
}

func NewDBService() error {
	_dbService = &DBService{}
	var err error
	// _dbService.mysql, err = NewMysqlIface()
	// if err != nil {
	// 	return err
	// }

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&allowOldPasswords=1",
		settings.SvrConfig.Mysql.USER,
		settings.SvrConfig.Mysql.PASSWD,
		settings.SvrConfig.Mysql.HOST,
		settings.SvrConfig.Mysql.DBNAME)
	_dbService.dbmysql, err = dbmysql.OpenMySQL(dataSource)
	if err != nil {
		gwlog.Error("NewDBService err:", err)
		return err
	}

	return nil
}

func (db *DBService) Destory() {
	// db.mysql.Close()
	db.dbmysql.Close()
}

func (db *DBService) QueryLoggerStats(channel string, sdate, edate string, replys *[]*models.LogStats) error {
	var rows *sql.Rows
	var err error
	if channel == "admin" {
		sqlstr := "select channel, gameid, newly, tow_pr, three_pr, seven_pr, retention, logdate from log_stat where logdate>=? and logdate<=?"
		rows, err = db.dbmysql.AsyncQuery(sqlstr, sdate, edate) //db.mysql.QueryV2(sqlstr, sdate, edate)
	} else {
		sqlstr := "select channel, gameid, newly, tow_pr, three_pr, seven_pr, retention, logdate from log_stat where channel=? and logdate>=? and logdate<=?"
		rows, err = db.dbmysql.AsyncQuery(sqlstr, channel, sdate, edate) //db.mysql.QueryV2(sqlstr, channel, sdate, edate)
	}

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		gwlog.Error("Query: failed: ", err)
		return err
	}

	for rows.Next() {
		stats := &models.LogStats{}
		if err := rows.Scan(&stats.Channel, &stats.Game, &stats.Newly, &stats.TowPr, &stats.ThreePr, &stats.SevenPr, &stats.Retention, &stats.LogDate); err != nil {
			gwlog.Error("Scan: failed: ", err)
			continue
		}

		*replys = append(*replys, stats)
	}
	return nil
}

func (db *DBService) ReplaceLoggerStats(channel string, game string, newly, tow_pr, three_pr, seven_pr, retention, logdate string) error {
	sqlstr := "replace into log_stat values(?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.dbmysql.AsyncExec(sqlstr, channel, game, newly, tow_pr, three_pr, seven_pr, retention, logdate) //db.mysql.ReplaceV2(sqlstr, channel, game, newly, tow_pr, three_pr, seven_pr, retention, logdate)
	if err != nil {
		gwlog.Error("ReplaceLoggerStats ", err)
		return err
	}
	return nil
}

func (db *DBService) RemoveLoggerStats(channel string, game string, logdate string) error {
	sqlstr := "delete from log_stat where channel=? and gameid = ? and logdate = ?"
	_, err := db.dbmysql.AsyncExec(sqlstr, channel, game, logdate)
	if err != nil {
		gwlog.Error("RemoveLoggerStats ", err)
		return err
	}
	return nil
}
