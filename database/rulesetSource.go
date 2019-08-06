package ndb

import (
    "github.com/astaxie/beego/logs"
	"database/sql"
	"owlhmaster/utils"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

var (
    RSdb *sql.DB
)

func RSConn() {
	var err error
	loadDataSQL := map[string]map[string]string{}
	loadDataSQL["rulesetSourceConn"] = map[string]string{}
	loadDataSQL["rulesetSourceConn"]["path"] = ""
	loadDataSQL["rulesetSourceConn"]["cmd"] = "" 
	loadDataSQL, err = utils.GetConf(loadDataSQL)    
    path := loadDataSQL["rulesetSourceConn"]["path"]
    cmd := loadDataSQL["rulesetSourceConn"]["cmd"]
	if err != nil {
		logs.Error("SConn Error getting data from main.conf at master: "+err.Error())
	}
	_, err = os.Stat(path) 
	if err != nil {
		panic("database/RulesetSource DB -- DB Open Failed: "+err.Error())
	}	
    RSdb, err = sql.Open(cmd,path)
    if err != nil {
		logs.Error("database/RulesetSource -- SQL openning Error: "+err.Error()) 
    }
    logs.Info("database/RulesetSource -- DB -> sql.Open, DB Ready") 
}

