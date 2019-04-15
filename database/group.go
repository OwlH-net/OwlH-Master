package ndb

import (
    "github.com/astaxie/beego/logs"
	"database/sql"
	"owlhmaster/utils"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

var (
    Gdb *sql.DB
)

func GConn() {
	var err error
	loadDataSQL := map[string]map[string]string{}
	loadDataSQL["groupConn"] = map[string]string{}
	loadDataSQL["groupConn"]["path"] = ""
	loadDataSQL["groupConn"]["cmd"] = "" 
	loadDataSQL, err = utils.GetConf(loadDataSQL)    
    path := loadDataSQL["groupConn"]["path"]
    cmd := loadDataSQL["groupConn"]["cmd"]
	if err != nil {
		logs.Error("GConn Error getting data from main.conf at master: "+err.Error())
	}
	_, err = os.Stat(path) 
	if err != nil {
		panic("database/Group DB -- DB Open Failed: "+err.Error())
	}	
    Gdb, err = sql.Open(cmd,path)
    if err != nil {
		logs.Error("database/Group -- SQL openning Error: "+err.Error()) 
    }
    logs.Info("database/Group -- DB -> sql.Open, DB Ready") 
}