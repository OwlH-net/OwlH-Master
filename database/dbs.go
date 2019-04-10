package ndb

import (
	"github.com/astaxie/beego/logs"
	"database/sql"
	"os"
	"owlhmaster/utils"
)

var (
    Db *sql.DB
)

func Conn() {
	var err error
	loadDataSQL := map[string]map[string]string{}
	loadDataSQL["dbsConn"] = map[string]string{}
	loadDataSQL["dbsConn"]["path"] = ""
	loadDataSQL["dbsConn"]["cmd"] = "" 
	loadDataSQL, err = utils.GetConf(loadDataSQL)    
    path := loadDataSQL["dbsConn"]["path"]
    cmd := loadDataSQL["dbsConn"]["cmd"]
	if err != nil {
		logs.Error("Conn Error getting data from main.conf at master: "+err.Error())
	}
	_, err = os.Stat(path) 
	if err != nil {
		panic("dbs/node DB -- DB Open Failed")
	}	
    Db, err = sql.Open(cmd,path)
    if err != nil {
        logs.Error("dbs/node DB -- DB Open Failed")
    }else{
		logs.Info("dbs/node DB -- DB -> sql.Open, DB Ready") 
	}
}

func Close() {
	Db.Close()
	Rdb.Close()
}