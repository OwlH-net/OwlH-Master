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

// func ChangePluginStatus(anode map[string]string) (err error) {	
// 	updatePluginMaster, err := Mdb.Prepare("update plugins set plugin_value = ? where plugin_uniqueid = ? and plugin_param = ?;")
// 	if (err != nil){
// 		logs.Error("ChangePluginStatus UPDATE prepare error : "+err.Error())
// 		return err
// 	}
// 	_, err = updatePluginMaster.Exec(anode["value"], anode["uuid"], anode["param"])
// 	defer updatePluginMaster.Close()
// 	if (err != nil){
// 		logs.Error("ChangePluginStatus UPDATE error: "+err.Error())
// 		return err
// 	}
// 	return nil
// }

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

func GetChangeControl()(path map[string]map[string]string, err error){
	var serviceValues = map[string]map[string]string{}
    var uniqid string
    var param string
	var value string
	rowsQuery, err := Mdb.Query("select control_uniqueid, control_param, control_value from changerecord;")
	if err != nil {
		logs.Error("GetChangeControl Mdb.Query Error : %s", err.Error())
		return nil, err
	}
	defer rowsQuery.Close()
	for rowsQuery.Next() {
		if err = rowsQuery.Scan(&uniqid, &param, &value); err != nil { logs.Error("GetChangeControl -- Query return error: %s", err.Error()); return nil, err}

		if serviceValues[uniqid] == nil { serviceValues[uniqid] = map[string]string{}}
		serviceValues[uniqid][param]=value
	} 
	return serviceValues,nil
}

func InsertChangeControl(uuid string, param string, value string)(err error){
	insertChangeControlValues, err := Mdb.Prepare("insert into changerecord(control_uniqueid, control_param, control_value) values (?,?,?);")
	if (err != nil){ logs.Error("InsertChangeControl prepare error: "+err.Error()); return err}

	_, err = insertChangeControlValues.Exec(&uuid, &param, &value)
	if (err != nil){ logs.Error("InsertChangeControl exec error: "+err.Error()); return err}

	defer insertChangeControlValues.Close()
	
	return nil
}

func GetIncidents()(path map[string]map[string]string, err error){
	var serviceValues = map[string]map[string]string{}
    var uniqid string
    var param string
	var value string
	rowsQuery, err := Mdb.Query("select incidents_uniqueid, incidents_param, incidents_value from incidents;")
	if err != nil {
		logs.Error("GetIncidents Mdb.Query Error : %s", err.Error())
		return nil, err
	}
	defer rowsQuery.Close()
	for rowsQuery.Next() {
		if err = rowsQuery.Scan(&uniqid, &param, &value); err != nil { logs.Error("GetIncidents -- Query return error: %s", err.Error()); return nil, err}

		if serviceValues[uniqid] == nil { serviceValues[uniqid] = map[string]string{}}
		serviceValues[uniqid][param]=value
	} 
	return serviceValues,nil
}

func PutIncident(uuid string, param string, value string)(err error){
	PutIncidentValues, err := Mdb.Prepare("insert into incidents(incidents_uniqueid, incidents_param, incidents_value) values (?,?,?);")
	if (err != nil){ logs.Error("PutIncident prepare error: "+err.Error()); return err}

	_, err = PutIncidentValues.Exec(&uuid, &param, &value)
	if (err != nil){ logs.Error("PutIncident exec error: "+err.Error()); return err}

	defer PutIncidentValues.Close()
	
	return nil
}

func GetAllGroups()(groups map[string]map[string]string, err error){
	var allgroups = map[string]map[string]string{}
    var uniqid string
    var param string
    var value string
	if Mdb == nil { logs.Error("no access to database"); return nil, err}
	
    sql := "select group_uniqueid, group_param, group_value from groups;"
    rows, err := Mdb.Query(sql)
	if err != nil { logs.Error("Mdb.Query Error : %s", err.Error()); return nil, err}
	
    for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil { logs.Error("GetAllGroups rows.Scan: %s", err.Error()); return nil, err}
		
        if allgroups[uniqid] == nil { allgroups[uniqid] = map[string]string{}}
        allgroups[uniqid][param]=value
	} 
    return allgroups, nil
}

func GroupExists(uuid string) (err error) {
	if Mdb == nil {logs.Error("no access to database"); return err}
	
    sql := "SELECT * FROM groups where group_uniqueid = '"+uuid+"';"
    rows, err := Mdb.Query(sql)
	if err != nil {logs.Error("Error on query groupExist at group.go "+err.Error()); return err}
	
    defer rows.Close()
    if rows.Next() {
        return err
    } else {
        return nil
    }
}

func DeleteGroup(uuid string) (err error) {
	if Mdb == nil { logs.Error("DeleteGroup -- Can't acces to database"); return err}

	stmt, err := Mdb.Prepare("delete from groups where group_uniqueid = ?")
	if err != nil {logs.Error("Prepare DeleteGroup -> %s", err.Error()); return err}
	
    _, err = stmt.Exec(&uuid)
    if err != nil {logs.Error("Execute DeleteGroup -> %s", err.Error()); return err}

	return nil
}

func InsertGroup(uuid string, param string, value string)(err error){
	if Mdb == nil {logs.Error("no access to database"); return err}
	insertGroup, err := Mdb.Prepare("insert into groups(group_uniqueid, group_param, group_value) values(?,?,?);")
	if err != nil {logs.Error("Prepare InsertGroup-> %s", err.Error()); return err}

	_, err = insertGroup.Exec(&uuid, &param, &value)
	if err != nil {logs.Error("Execute InsertGroup-> %s", err.Error()); return err}

	return nil
}

func UpdateGroupData(uuid string, param string, value string)(err error){
	if Mdb == nil {logs.Error("no access to database"); return err}
	updateGroup, err := Mdb.Prepare("update groups set group_value = ? where group_param = ? and group_uniqueid = ?")
	if err != nil {logs.Error("Prepare UpdateGroupData-> %s", err.Error()); return err}

	_, err = updateGroup.Exec(&value, &param, &uuid)
	if err != nil {logs.Error("Execute UpdateGroupData-> %s", err.Error()); return err}

	return nil
}

func InsertGroupNodes(uuid string, param string, value string)(err error){
	if Mdb == nil {logs.Error("no access to database"); return err}
	insertGroupnodesValues, err := Mdb.Prepare("insert into groupnodes(gn_uniqueid, gn_param, gn_value) values(?,?,?);")
	if err != nil {logs.Error("Prepare InsertGroupNodes-> %s", err.Error()); return err}

	_, err = insertGroupnodesValues.Exec(&uuid, &param, &value)
	if err != nil {logs.Error("Execute InsertGroupNodes-> %s", err.Error()); return err}

	return nil
}