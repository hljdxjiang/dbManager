package dataManager

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"

	simplejson "github.com/bitly/go-simplejson"
)

var (
	dbLock = new(sync.RWMutex)
	dbMap  = make(map[string]db)
)

type db interface {
	// Query database
	Query(sql string, args ...interface{}) (*sql.Rows, error)
	// Query one row
	QueryRow(sql string, args ...interface{}) *sql.Row
	// Execute
	Exec(sql string, args ...interface{}) error
	// Begin transaction
	Begin() (*sql.Tx, error)
	// Prepare
	Prepare(query string) (*sql.Stmt, error)
	// Get Error Code
	GetErrorCode(err error) string
	// Get Message info
	GetErrorMsg(err error) string
}

func register(dsn string, d db) {
	dbLock.Lock()
	defer dbLock.Unlock()
	if d == nil {
		fmt.Errorf("sql: Register driver is nil")
	}
	if _, dup := dbMap[dsn]; dup {
		fmt.Println("reregister diver. dsn is :", dsn)
	}
	dbMap[dsn] = d
}

func DbControl(dbname string) (db, error) {
	fmt.Println(len(dbMap))
	if val, ok := dbMap[dbname]; ok {
		if val == nil {
			return nil, errors.New("db[" + dbname + "] is nil")
		}
		return val, nil
	} else {
		return nil, errors.New("db[" + dbname + "]has not registered")
	}
}

func Init(fid string) error {
	conf, err := InitConfig(fid)
	if err == nil {
		cont, err := ioutil.ReadFile(conf.GetFile())
		if err != nil {
			return err
		}
		js, err := simplejson.NewJson(cont)
		if err != nil {
			return err
		}
		arr := js.MustArray()
		fmt.Println("arrlen", len(arr))
		for i := 0; i < len(arr); i++ {
			dbtype := js.GetIndex(i).Get("dbtype").MustString()
			id := js.GetIndex(i).Get("id").MustString()
			source := js.GetIndex(i).Get("source").MustString()
			if dbtype == "mysql" {
				register(id, CreateMysqlDb(id, source))
			} else if dbtype == "oracle" {
				//register(id, CreateOracleDb(id, source))
			}
		}

	}
	return err
}
