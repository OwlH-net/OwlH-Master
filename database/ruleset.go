package ndb

import (
    "github.com/astaxie/beego/logs"
	"database/sql"
	"owlhmaster/utils"
	"os"
	_ "github.com/mattn/go-sqlite3"
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
	_, err = os.Stat(path) 
	if err != nil {
		panic("ruleset/Ruleset DB -- DB Open Failed: "+err.Error())
	}	
    Rdb, err = sql.Open(cmd,path)
    if err != nil {
		logs.Error("ruleset/Ruleset DB -- SQL openning Error: "+err.Error()) 
    }
    logs.Info("ruleset/Ruleset DB -- DB -> sql.Open, DB Ready") 
}
