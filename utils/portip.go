package utils

import (
    "github.com/astaxie/beego/logs"
	"owlhmaster/database"
	"errors"
)

func ObtainPortIp(uuid string)(ip string, port string, err error)  {
	if ndb.Db == nil {
		logs.Warn("obtainPortIp -> Error conexión DB")
		return "","",errors.New("DB NODE obtainPortIp -> Conexión a DB fallida: " + err.Error())
	}
	
	var ipNode string
	var portNode string
	row1 := ndb.Db.QueryRow("SELECT node_value FROM nodes WHERE node_uniqueid = \""+uuid+"\" and node_param = \"ip\";")
	err = row1.Scan(&ipNode)
	if err != nil {
		logs.Error("DB NODE obtainPortIp ipNode -> La Query no ha funcionado bien: %s", err.Error())
		return "","",errors.New("DB NODE -> La Query no ha funcionado bien: " + err.Error())
	}
	row2 := ndb.Db.QueryRow("SELECT node_value FROM nodes WHERE node_uniqueid = \""+uuid+"\" and node_param = \"port\";")
    err = row2.Scan(&portNode)
	if err != nil {
		logs.Error("DB NODE obtainPortIp portNode -> La Query no ha funcionado bien: %s", err.Error())
		return "","",errors.New("DB NODE obtainPortIp portNod -> La Query no ha funcionado bien: " + err.Error())
	}

	return ipNode, portNode, nil
}