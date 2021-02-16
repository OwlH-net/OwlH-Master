package ndb

import (
	"errors"
	"owlhmaster/utils"

	"github.com/astaxie/beego/logs"
)

func ObtainPortIp(uuid string) (ip string, port string, err error) {
	if Db == nil {
		logs.Error("obtainPortIp -> Error conexión DB")
		return "", "", err
	}

	var ipNode string
	var portNode string
	row1 := Db.QueryRow("SELECT node_value FROM nodes WHERE node_uniqueid = \"" + uuid + "\" and node_param = \"ip\";")
	err = row1.Scan(&ipNode)
	if err != nil {
		logs.Error("DB NODE obtainPortIp ipNode -> ndb.Db.QueryRow error: %s", err.Error())
		return "", "", err
	}

	row2 := Db.QueryRow("SELECT node_value FROM nodes WHERE node_uniqueid = \"" + uuid + "\" and node_param = \"port\";")
	err = row2.Scan(&portNode)
	if err != nil {
		logs.Error("DB NODE obtainPortIp portNode -> row2.Scan error: %s", err.Error())
		return "", "", err
	}

	return ipNode, portNode, nil
}

func ObtainNodeName(uuid string) (name string, err error) {
	if Db == nil {
		logs.Error("ObtainNodeName -> Error conexión DB")
		return "", err
	}

	row2 := Db.QueryRow("SELECT node_value FROM nodes WHERE node_uniqueid = \"" + uuid + "\" and node_param = \"name\";")
	err = row2.Scan(&name)
	if err != nil {
		logs.Error("DB NODE ObtainNodeName portNode -> row2.Scan error: %s", err.Error())
		return "", err
	}

	return name, nil
}

func GetAllNodes() (nodes map[string]map[string]string, err error) {
	var allnodes = map[string]map[string]string{}
	var uniqid string
	var param string
	var value string
	if Db == nil {
		logs.Error("no access to database")
		return nil, errors.New("no access to database")
	}
	sql := "select node_uniqueid, node_param, node_value from nodes;"
	rows, err := Db.Query(sql)
	if err != nil {
		logs.Error("Db.Query Error : %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
			logs.Error("GetAllNodes rows.Scan: %s", err.Error())
			return nil, err
		}
		if allnodes[uniqid] == nil {
			allnodes[uniqid] = map[string]string{}
		}
		allnodes[uniqid][param] = value
	}
	return allnodes, nil
}

func GetNodeById(uuid string) (nodes map[string]map[string]string, err error) {
	var allnodes = map[string]map[string]string{}
	var uniqid string
	var param string
	var value string
	if Db == nil {
		logs.Error("GetNodeById no access to database")
		return nil, errors.New("GetNodeById no access to database")
	}

	sql := "select node_uniqueid, node_param, node_value from nodes where node_uniqueid = '" + uuid + "';"
	rows, err := Db.Query(sql)
	if err != nil {
		logs.Error("GetNodeById Db.Query Error : %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
			logs.Error("GetNodeById rows.Scan: %s", err.Error())
			return nil, err
		}

		if allnodes[uniqid] == nil {
			allnodes[uniqid] = map[string]string{}
		}
		allnodes[uniqid][param] = value
	}

	return allnodes, nil
}

func GetTokenByUuid(uuid string) (err error) {
	if Db == nil {
		logs.Error("GetTokenByUuid obtainPortIp -> Error conexión DB")
		return err
	}

	var token string
	query := "SELECT node_value FROM nodes WHERE node_uniqueid = '" + uuid + "' and node_param = 'token';"
	row1 := Db.QueryRow(query)
	err = row1.Scan(&token)
	if err != nil {
		logs.Warn("GetTokenByUuid error: %s", err.Error())
		return err
	}

	//set token for httpRequest header
	utils.TokenMasterValidated = token

	return nil
}
