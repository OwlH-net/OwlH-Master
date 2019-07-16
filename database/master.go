package ndb

import (
	"github.com/astaxie/beego/logs"
	"database/sql"
	"os"
	"owlhmaster/utils"
	_ "github.com/mattn/go-sqlite3"
)

var (
    Mdb *sql.DB
)

func MConn() {
	var err error
	loadDataSQL := map[string]map[string]string{}
	loadDataSQL["masterConn"] = map[string]string{}
	loadDataSQL["masterConn"]["path"] = ""
	loadDataSQL["masterConn"]["cmd"] = "" 
	loadDataSQL, err = utils.GetConf(loadDataSQL)    
    path := loadDataSQL["masterConn"]["path"]
    cmd := loadDataSQL["masterConn"]["cmd"]
	if err != nil {
		logs.Error("MConn Error getting data from main.conf at master: "+err.Error())
	}
	_, err = os.Stat(path) 
	if err != nil {
		panic("Error: dbs/node DB -- DB Open Failed: "+err.Error())
	}	
    Mdb, err = sql.Open(cmd,path)
    if err != nil {
        logs.Error("dbs/node DB -- DB Open Failed: "+err.Error())
    }else{
		logs.Info("dbs/node DB -- DB -> sql.Open, DB Ready") 
	}
}

func PingPlugins()(path map[string]map[string]string, err error){
	var pingData = map[string]map[string]string{}
    var uniqid string
    var param string
	var value string

	sql := "select plugin_uniqueid, plugin_param, plugin_value from plugins";
	rows, err := Mdb.Query(sql)
	if err != nil {
		logs.Error("PingPlugins Mdb.Query Error : %s", err.Error())
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
            logs.Error("PingPlugins -- Query return error: %s", err.Error())
            return nil, err
		}
        if pingData[uniqid] == nil { pingData[uniqid] = map[string]string{}}
        pingData[uniqid][param]=value
	} 
	return pingData,nil
}

func ChangePluginStatus(anode map[string]string) (err error) {
	updateRulesetNode, err := Mdb.Prepare("update plugins set plugin_value = ? where plugin_uniqueid = ? and plugin_param = ?;")
	if (err != nil){
		logs.Error("ChangePluginStatus UPDATE prepare error for update isDownloaded -- "+err.Error())
		return err
	}
	_, err = updateRulesetNode.Exec(anode["value"], anode["uuid"], anode["param"])
	defer updateRulesetNode.Close()
	if (err != nil){
		logs.Error("ChangePluginStatus UPDATE error for update isDownloaded -- "+err.Error())
		return err
	}
	return nil
}