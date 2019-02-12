package ndb

import (
    "github.com/astaxie/beego/logs"
    "database/sql"
//    "fmt"
//   "time"
    _ "github.com/mattn/go-sqlite3"
    //"errors"
)

var (
    Rdb *sql.DB
)

func RConn() {
    var err error
	Rdb, err = sql.Open("sqlite3", "database/ruleset.db")
    if err != nil {
        panic("rdb/ruleset -- DB Open Failed")
    }
    logs.Info("rdb/ruleset -- DB -> sql.Open, DB Ready") 
}