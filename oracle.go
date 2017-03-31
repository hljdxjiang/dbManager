package dataManager

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	//_ "github.com/mattn/go-oci8"
)

type oracledb struct {
	id     string
	db     *sql.DB
	source string
}

func CreateOracleDb(sid string, source string) *oracledb {
	o := new(oracledb)
	o.id = sid
	o.source = source
	orcldb, err := sql.Open("oci8", source)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if orcldb.Ping() != nil {
		return nil
	}
	o.db = orcldb
	return o
}
func (this *oracledb) GetErrorCode(err error) string {
	ret := err.Error()
	if n := strings.Index(ret, ":"); n > 0 {
		return strings.TrimSpace(ret[:n])
	} else {
		fmt.Errorf("orcldb error information is not oracle return info")
		return ""
	}

}

func (this *oracledb) GetErrorMsg(err error) string {

	ret := err.Error()
	if n := strings.Index(ret, ":"); n > 0 {
		return strings.TrimSpace(ret[n+1:])
	} else {
		fmt.Errorf("orcldb error information is not oracle return info")
		return ""
	}

}

func (orcldb *oracledb) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if orcldb.db.Ping() != nil {
	}
	return orcldb.db.Query(sql, args...)

}

func (orcldb *oracledb) Exec(sql string, args ...interface{}) error {
	if orcldb.db.Ping() != nil {
	}
	_, err := orcldb.db.Exec(sql, args...)
	return err
}

func (orcldb *oracledb) Begin() (*sql.Tx, error) {
	return orcldb.db.Begin()
}

func (orcldb *oracledb) Prepare(query string) (*sql.Stmt, error) {
	return orcldb.db.Prepare(query)
}

func (orcldb *oracledb) QueryRow(sql string, args ...interface{}) *sql.Row {
	if orcldb.db.Ping() != nil {
	}
	return orcldb.db.QueryRow(sql, args...)
}
