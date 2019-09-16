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

func PingFlow()(path map[string]map[string]string, err error){
	var pingData = map[string]map[string]string{}
    var uniqid string
    var param string
	var value string

	sql := "select flow_uniqueid, flow_param, flow_value from dataflow";
	rows, err := Mdb.Query(sql)
	if err != nil {
		logs.Error("PingFlow Mdb.Query Error : %s", err.Error())
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
            logs.Error("PingFlow -- Query return error: %s", err.Error())
            return nil, err
		}
        if pingData[uniqid] == nil { pingData[uniqid] = map[string]string{}}
        pingData[uniqid][param]=value
	} 
	return pingData,nil
}

func ChangePluginStatus(anode map[string]string) (err error) {	
	updatePluginMaster, err := Mdb.Prepare("update plugins set plugin_value = ? where plugin_uniqueid = ? and plugin_param = ?;")
	if (err != nil){
		logs.Error("ChangePluginStatus UPDATE prepare error : "+err.Error())
		return err
	}
	_, err = updatePluginMaster.Exec(anode["value"], anode["uuid"], anode["param"])
	defer updatePluginMaster.Close()
	if (err != nil){
		logs.Error("ChangePluginStatus UPDATE error: "+err.Error())
		return err
	}
	return nil
}

func ChangeDataflowStatus(anode map[string]string) (err error) {
	updateDataflowMaster, err := Mdb.Prepare("update dataflow set flow_value = ? where flow_uniqueid = ? and flow_param = ?;")
	if (err != nil){
		logs.Error("ChangeDataflowStatus UPDATE prepare error: "+err.Error())
		return err
	}
	_, err = updateDataflowMaster.Exec(anode["value"], anode["uuid"], anode["param"])
	defer updateDataflowMaster.Close()
	if (err != nil){
		logs.Error("ChangeDataflowStatus UPDATE error: "+err.Error())
		return err
	}
	return nil
}

func UpdateMasterNetworkInterface(anode map[string]string) (err error) {
	updateDataflowMaster, err := Mdb.Prepare("update masterconfig set config_value = ? where config_uniqueid = ? and config_param = ?;")
	if (err != nil){
		logs.Error("UpdateMasterNetworkInterface UPDATE prepare error: "+err.Error())
		return err
	}
	_, err = updateDataflowMaster.Exec(anode["value"], anode["uuid"], anode["param"])
	defer updateDataflowMaster.Close()
	if (err != nil){
		logs.Error("UpdateMasterNetworkInterface UPDATE error: "+err.Error())
		return err
	}
	return nil
}

func LoadMasterNetworkValuesSelected()(path map[string]map[string]string, err error){
	var pingData = map[string]map[string]string{}
    var uniqid string
    var param string
	var value string

	sql := "select config_uniqueid, config_param, config_value from masterconfig";
	rows, err := Mdb.Query(sql)
	if err != nil {
		logs.Error("LoadMasterNetworkValuesSelected Mdb.Query Error : %s", err.Error())
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
            logs.Error("LoadMasterNetworkValuesSelected -- Query return error: %s", err.Error())
            return nil, err
		}
        if pingData[uniqid] == nil { pingData[uniqid] = map[string]string{}}
        pingData[uniqid][param]=value
	} 
	return pingData,nil
}

func InsertPluginService(uuid string, param string, value string)(err error){
	updateAnalyzerNode, err := Mdb.Prepare("insert into plugins(plugin_uniqueid, plugin_param, plugin_value) values (?,?,?);")
	if (err != nil){ logs.Error("InsertPluginService INSERT prepare error: "+err.Error()); return err}

	_, err = updateAnalyzerNode.Exec(&uuid, &param, &value)
	if (err != nil){ logs.Error("InsertPluginService INSERT exec error: "+err.Error()); return err}

	defer updateAnalyzerNode.Close()
	
	return nil
}

func DeleteServiceMaster(uuid string)(err error){
	DeleteService, err := Mdb.Prepare("delete from plugins where plugin_uniqueid = ?;")
	if (err != nil){ logs.Error("DeleteServiceMaster UPDATE prepare error: "+err.Error()); return err}

	_, err = DeleteService.Exec(&uuid)
	if (err != nil){ logs.Error("DeleteServiceMaster exec error: "+err.Error()); return err}

	defer DeleteService.Close()
	
	return nil
}

func UpdatePluginValueMaster(uuid string, param string, value string)(err error){
	UpdatePluginValue, err := Mdb.Prepare("update plugins set plugin_value = ? where plugin_uniqueid = ? and plugin_param = ?;")
	if (err != nil){ logs.Error("UpdatePluginValueMaster UPDATE prepare error: "+err.Error()); return err}

	_, err = UpdatePluginValue.Exec(&value, &uuid, &param)
	if (err != nil){ logs.Error("UpdatePluginValueMaster UPDATE exec error: "+err.Error()); return err}

	defer UpdatePluginValue.Close()
	
	return nil
}

func GetPlugins()(path map[string]map[string]string, err error){
	var serviceValues = map[string]map[string]string{}
    var uniqid string
    var param string
	var value string
	rowsQuery, err := Mdb.Query("select plugin_uniqueid, plugin_param, plugin_value from plugins;")
	if err != nil {
		logs.Error("GetPlugins Mdb.Query Error : %s", err.Error())
		return nil, err
	}
	defer rowsQuery.Close()
	for rowsQuery.Next() {
		if err = rowsQuery.Scan(&uniqid, &param, &value); err != nil { logs.Error("GetPlugins -- Query return error: %s", err.Error()); return nil, err}

		if serviceValues[uniqid] == nil { serviceValues[uniqid] = map[string]string{}}
		serviceValues[uniqid][param]=value
	} 
	return serviceValues,nil
}