package dataManager

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type mysqldb struct {
	id     string
	db     *sql.DB
	source string
}

func CreateMysqlDb(sid string, source string) *mysqldb {
	o := new(mysqldb)
	o.id = sid
	o.source = source
	msdb, err := sql.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if msdb.Ping() != nil {
		return nil
	}
	o.db = msdb
	return o
}

func (msdb *mysqldb) GetErrorCode(err error) string {
	ret := err.Error()
	if n := strings.Index(ret, ":"); n > 0 {
		return strings.TrimSpace(ret[:n])
	} else {
		fmt.Errorf("msdb error information is not mysql return info")
		return ""
	}
}

func (msdb *mysqldb) GetErrorMsg(err error) string {
	ret := err.Error()
	if n := strings.Index(ret, ":"); n > 0 {
		return strings.TrimSpace(ret[n+1:])
	} else {
		fmt.Errorf("msdb error information is not mysql return info")
		return ""
	}
}

func (msdb *mysqldb) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if msdb.db.Ping() != nil {
	}
	return msdb.db.Query(sql, args...)
}

func (msdb *mysqldb) Exec(sql string, args ...interface{}) error {
	if msdb.db.Ping() != nil {
	}
	_, err := msdb.db.Exec(sql, args...)
	return err
}

func (msdb *mysqldb) Begin() (*sql.Tx, error) {
	return msdb.db.Begin()
}

func (msdb *mysqldb) Prepare(query string) (*sql.Stmt, error) {
	return msdb.db.Prepare(query)
}

func (msdb *mysqldb) QueryRow(sql string, args ...interface{}) *sql.Row {
	if msdb.db.Ping() != nil {
	}
	return msdb.db.QueryRow(sql, args...)
}
