package ndb

import (
    "github.com/astaxie/beego/logs"
)

func ObtainPortIp(uuid string)(ip string, port string, err error)  {
	if Db == nil {
		logs.Error("obtainPortIp -> Error conexión DB")
		return "","",err
	}
	
	var ipNode string
	var portNode string
	row1 := Db.QueryRow("SELECT node_value FROM nodes WHERE node_uniqueid = \""+uuid+"\" and node_param = \"ip\";")
	err = row1.Scan(&ipNode)
	if err != nil {
		logs.Error("DB NODE obtainPortIp ipNode -> ndb.Db.QueryRow error: %s", err.Error())
		return "","",err
	}
	row2 := Db.QueryRow("SELECT node_value FROM nodes WHERE node_uniqueid = \""+uuid+"\" and node_param = \"port\";")
    err = row2.Scan(&portNode)
	if err != nil {
		logs.Error("DB NODE obtainPortIp portNode -> row2.Scan error: %s", err.Error())
		return "","",err
	}

	return ipNode, portNode, nil
}

func ObtainNodeName(uuid string)(name string, err error)  {
	if Db == nil {logs.Error("obtainPortIp -> Error conexión DB"); return "",err}
	
	row2 := Db.QueryRow("SELECT node_value FROM nodes WHERE node_uniqueid = \""+uuid+"\" and node_param = \"name\";")
    err = row2.Scan(&name)
	if err != nil {logs.Error("DB NODE obtainPortIp portNode -> row2.Scan error: %s", err.Error()); return "",err}

	return name, nil
}