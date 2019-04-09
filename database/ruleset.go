package ndb

import (
    "github.com/astaxie/beego/logs"
	"database/sql"
	"owlhmaster/utils"
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
	loadDataSQL := map[string]map[string]string{}
	loadDataSQL["rulesetConn"] = map[string]string{}
	loadDataSQL["rulesetConn"]["path"] = ""
	loadDataSQL["rulesetConn"]["cmd"] = "" 
	loadDataSQL, err = utils.GetConf(loadDataSQL)    
    path := loadDataSQL["rulesetConn"]["path"]
    cmd := loadDataSQL["rulesetConn"]["cmd"]
	if err != nil {
		logs.Error("RConn Error getting data from main.conf at master: "+err.Error())
	}
	
    Rdb, err = sql.Open(cmd,path)
    if err != nil {
        panic("ruleset/servers -- DB Open Failed")
    }
    logs.Info("ruleset/servers -- DB -> sql.Open, DB Ready") 

}