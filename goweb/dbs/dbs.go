package dbs

import (
	"database/sql"
	"errors"
	"fmt"
	"nuvem/engine/asyncdb/dbmongo"
	"nuvem/engine/asyncdb/dbmysql"
	"nuvem/engine/coder"
	"vms-go/goweb/models"
	"vms-go/goweb/settings"

	"github.com/sagacao/goworld/engine/gwlog"
	"gopkg.in/mgo.v2/bson"
)

type DBService struct {
	// mysql *MysqlIface

	dbmysql *dbmysql.DBMysql
	dbmongo *dbmongo.DBMongo
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

	_dbService.dbmongo, err = dbmongo.OpenMongo(settings.SvrConfig.Mongo.ADDR, settings.SvrConfig.Mongo.DATABASE, "")
	if err != nil {
		gwlog.Fatal("NewDBService err:", err)
	}

	return nil
}

func (db *DBService) Destory() {
	// db.mysql.Close()
	db.dbmysql.Close()
	if db.dbmongo != nil {
		db.dbmongo.Close()
	}
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

func (db *DBService) QueryPlayerInfoByUid(collection string, uid string, replys coder.JSON) error {
	query := bson.M{"_id": uid}
	reply, err := db.dbmongo.AsyncQuery(collection, query)
	defer func() {
		if reply != nil {
			reply.Close()
		}
	}()
	if err != nil {
		gwlog.Errorf("QueryPlayerInfoByUid err : %v", err)
		return err
	}

	rownum := 0
	for reply.Next(replys) {
		rownum++
		break
	}

	if rownum == 0 {
		return errors.New("no player")
	}

	return nil
}

func (db *DBService) SavePlayerInfoByUid(collection string, uid string, data coder.JSON) error {
	reply, err := db.dbmongo.AsyncExec(collection, uid, data)
	if err != nil {
		gwlog.Fatal("SavePlayerInfoByUid: ", err)
		return err
	}
	if reply.Matched != 1 {
		gwlog.Debug("SavePlayerInfoByUid", reply.UpsertedId, reply.Updated, reply.Matched, reply.Removed)
	}
	return nil
}

func (db *DBService) InsertGiftCode(game, atype string, acount int, codes []string) error {
	sqlstr := "insert into (gameId, code, atype, acount) mem_gift values"
	for _, v := range codes {
		sqlstr += fmt.Sprintf(" (%s, %s, %s, %v),", game, v, atype, acount)
	}
	sqlstr = sqlstr[0 : len(sqlstr)-1]
	gwlog.Debugf("InsertGiftCode sql:[%s]", sqlstr)
	_, err := db.dbmysql.AsyncExec(sqlstr)
	if err != nil {
		gwlog.Error("InsertGiftCode ", err)
		return err
	}
	return nil
}

func (db *DBService) QueryGiftCodes(game string, replys *[]*models.GiftCode) error {
	var rows *sql.Rows
	var err error
	sqlstr := "select gameId, code, atype, acount from mem_gift where used = 0 and gameId = ?"
	rows, err = db.dbmysql.AsyncQuery(sqlstr, game)

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
		gifts := &models.GiftCode{}
		if err := rows.Scan(&gifts.Game, &gifts.Code, &gifts.Atype, &gifts.Acount); err != nil {
			gwlog.Error("Scan: failed: ", err)
			continue
		}

		*replys = append(*replys, gifts)
	}
	return nil
}
