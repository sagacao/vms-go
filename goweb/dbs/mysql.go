package dbs

import (
	"database/sql"
	"fmt"
	"sync"
	"vms-go/goweb/logger"
	"vms-go/goweb/settings"
)

type MysqlIface struct {
	dbname     string
	dataSource string
	db         *sql.DB

	signalChan chan string
	wg         sync.WaitGroup
}

func NewMysqlIface() (*MysqlIface, error) {
	_mface := &MysqlIface{}

	_mface.dbname = settings.SvrConfig.Mysql.DBNAME
	_mface.dataSource = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&allowOldPasswords=1",
		settings.SvrConfig.Mysql.USER, settings.SvrConfig.Mysql.PASSWD, settings.SvrConfig.Mysql.HOST, settings.SvrConfig.Mysql.DBNAME)

	var err error
	_mface.db, err = sql.Open("mysql", _mface.dataSource)
	if err != nil {
		logger.Error("Open db '%s' Error : (%v)", settings.SvrConfig.Mysql.DBNAME, err)
		return nil, err
	}
	_mface.db.SetMaxOpenConns(10)
	_mface.db.SetMaxIdleConns(2)

	err = _mface.db.Ping()
	if err != nil {
		logger.Error("db '%s' Ping Error : (%v)", _mface.dbname, err)
		return nil, err
	}

	logger.Info("Mysql Has Open at ", settings.SvrConfig.Mysql.HOST)
	return _mface, nil
}

func (ms *MysqlIface) Close() {
	ms.db.Close()
}

func (ms *MysqlIface) QueryFuture(query string, args ...interface{}) func() (*sql.Rows, error) {
	var rows *sql.Rows
	var err error
	c := make(chan struct{}, 1)
	go func() {
		defer close(c)
		rows, err = ms.db.Query(query, args...)
	}()

	return func() (*sql.Rows, error) {
		<-c
		return rows, err
	}
}

func (ms *MysqlIface) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return ms.db.Query(query, args...)
}

func (ms *MysqlIface) ReplaceFuture(query string, args ...interface{}) func() (sql.Result, error) {
	var res sql.Result
	var err error
	c := make(chan struct{}, 1)
	go func() {
		defer close(c)
		res, err = ms.db.Exec(query, args...)
	}()

	return func() (sql.Result, error) {
		<-c
		return res, err
	}
}
