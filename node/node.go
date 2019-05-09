package node

import (
    "github.com/astaxie/beego/logs"
    "strings"
    "owlhmaster/database"
    "errors"
    "owlhmaster/nodeclient"
    "owlhmaster/ruleset"
    "owlhmaster/stap"
    "owlhmaster/utils"
    "regexp"
    "io/ioutil"
)

func findNode(s string) (id string, err error) {
    if ndb.Db == nil {
        logs.Error("Find Node -> no access to database")
        return "", errors.New("Find Node -> no access to database")
    }
    ip, _ := regexp.Compile(`\d+\.\d+\.\d+\.\d+`)
    uuid, _ := regexp.Compile(`\w{8}-\w{4}-\w{4}-\w{4}-\w{12}`)
    sql := "SELECT node_uniqueid FROM nodes where node_param = 'name' and node_value='"+s+"';"
    if ip.MatchString(s) {
        sql = "SELECT node_uniqueid FROM nodes where node_param = 'ip' and node_value='"+s+"';"
    } else if uuid.MatchString(s) {
        sql = "SELECT node_uniqueid FROM nodes where node_param = 'UUID' and node_value='"+s+"';"
    } 
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return "", err
    }
    defer rows.Close()
    if rows.Next() {
        if err = rows.Scan(&id); err == nil {
            return id, nil
        }
    }
    return "", err
}

func DeleteNode(nodeid string)(err error) {
    logs.Info("NODE Delete -> IN")
    if ndb.Db == nil {
        logs.Error("No access to database")
        return errors.New("No access to database")
    }
    stmt, err := ndb.Db.Prepare("delete from nodes where node_uniqueid = ?")
    if err != nil {
        logs.Error("Prepare -> %s", err.Error())
        return err
    }
    _, err = stmt.Exec(&nodeid)
    if err != nil {
        logs.Error("Execute -> %s", err.Error())
        return err
    }
    return nil
}

func GetNodeIPbyUUID(nk string) (ip string, err error) {
    if ndb.Db == nil {
        logs.Error("GetNodeIPbyUUID -> No access to database")
        return "", errors.New("GetNodeIPbyUUID -> No access to database")
    }
    sql := "SELECT node_value FROM nodes where node_param = 'ip' and node_uniqueid='"+nk+"';"
    logs.Info("GetNodeIP -> SQL -> %s", sql)
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return "", err
    }
    defer rows.Close()
    if rows.Next() {
        if err = rows.Scan(&ip); err == nil {
            return ip, nil
        }
    }
    return "", err
}

func getNodeConf(nodeKey string)(conf map[string]string, err error) {
    var param string
    var value string

    if ndb.Db == nil {
        logs.Error("getNodeConf -> No access to database")
        return nil, errors.New("getNodeConf -> No access to database")
    }
    
    sql := "SELECT node_param, node_value FROM nodes where node_uniqueid='"+nodeKey+"';"
    logs.Info("GetNodeConf -> SQL -> %s", sql)
    
    rows, err := ndb.Db.Query(sql)
    
    if err != nil {
        logs.Error(err.Error())
        return nil, err
    }
    
    defer rows.Close()
    for rows.Next() {
        if err = rows.Scan(&param, &value); err != nil {
            logs.Info (" Error rows.Scan -> %s",err.Error())
            continue
        }
        conf[param]=value
    }
    return conf, nil
}

func GetNodePortbyUUID(nk string) (port string, err error) {
    if ndb.Db == nil {
        logs.Error("GetNodePortbyUUID -> no access to database")
        return "", errors.New("GetNodePortbyUUID -> no access to database")
    }
    sql := "SELECT node_value FROM nodes where node_param = 'port' and node_uniqueid='"+nk+"';"
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return "", err
    }
    defer rows.Close()
    if rows.Next() {
        if err = rows.Scan(&port); err == nil {
            return port, nil
        }
    }
    return "", err
}

func getAllNodesIp() (ips map[string]string, err error) {
    var uid string
    var ip string
    if ndb.Db == nil {
        logs.Error("getAllNodesIp -> no access to database")
        return ips, errors.New("getAllNodesIp -> no access to database")
    }
    sql := "SELECT node_uniqueid, node_value FROM nodes where node_param = 'ip';"
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error("Error ndb.Db.Query %s -> %s", sql, err.Error())
        return ips, err
    }
    defer rows.Close()
    for rows.Next() {
        if err = rows.Scan(&uid, &ip); err != nil {
            logs.Info (" Error -> rows.Scan -> %s",err.Error())
        }
        ips[uid]=ip
    }
    return ips, nil
}

func nodeKeyExists(nodekey string, key string) (id int, err error) {
    if ndb.Db == nil {
        logs.Error("no access to database")
        return 0, errors.New("no access to database")
    }
    sql := "SELECT node_id FROM nodes where node_uniqueid = '"+nodekey+"' and node_param = '"+key+"';"
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return 0, err
    }
    defer rows.Close()
    if rows.Next() {
        if err = rows.Scan(&id); err == nil {
            return id, err
        }
    }
    return 0, nil
}

func nodeExists(nodeid string) (err error) {
    if ndb.Db == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
    }
    sql := "SELECT * FROM nodes where node_uniqueid = '"+nodeid+"';"
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return err
    }
    defer rows.Close()
    if rows.Next() {
        return errors.New("Node Exists " + nodeid)
    } else {
        return nil
    }
}

func nodeKeyUpdate(id int, nkey string, key string, value string) (err error) {
    logs.Info("NODE Key Insert -> IN")
    if ndb.Db == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
    }
    logs.Info("nkey: %s, key: %s, value: %s", nkey, key, value)
    stmt, err := ndb.Db.Prepare("update nodes set node_param = ?, node_value = ? where node_id = ? and node_uniqueid = ?")
    if err != nil {
        logs.Error("Prepare -> %s", err.Error())
        return err
    }
    _, err = stmt.Exec(&key, &value, &id, &nkey)
    if err != nil {
        logs.Error("Execute -> %s", err.Error())
        return err
    }
    return nil
}

func nodeKeyInsert(nkey string, key string, value string) (err error) {
    if ndb.Db == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
    }
    stmt, err := ndb.Db.Prepare("insert into nodes (node_uniqueid, node_param, node_value) values(?,?,?)")
    if err != nil {
        logs.Error("Prepare -> %s", err.Error())
        return err
    }
    _, err = stmt.Exec(&nkey, &key, &value)
    if err != nil {
        logs.Error("Execute -> %s", err.Error())
        return err
    }
    logs.Info("nkey from node.go to stap.go-->"+nkey)
    _,err = stap.Stap(nkey)
    if err != nil {
        logs.Error("Error creating node stap status from nodeKeyInsert at node.go -> %s", err.Error())
        return err
    }
    return nil
}

func AddNode(n map[string]string) (err error) {
    logs.Info("ADD NODE -> IN")
	nodeKey := utils.Generate()
    if _, ok := n["name"]; !ok {
		logs.Error("name empty: "+err.Error())
        return errors.New("name empty")
    }
    if _, ok := n["ip"]; !ok {
		logs.Error("ip empty: "+err.Error())
        return errors.New("ip empty")
    }

    if err := nodeExists(nodeKey); err != nil {
		logs.Error("node exist: "+err.Error())
        return errors.New("name empty")
    }
    
    for key, value := range n {
        err = nodeKeyInsert(nodeKey, key, value)
    }
    if err != nil {
        return err
    }
    return nil
}

func UpdateNode(n map[string]string) (err error) {
    logs.Info("UPDATE NODE -> IN name es -  %s", n["name"])
    var nodeKey string

    if _, ok := n["name"]; !ok {
        return errors.New("name is empty")
    }
    if _, ok := n["ip"]; !ok {
        return errors.New("ip is empty")
    }
    if _, ok := n["id"]; !ok {
        nodeKey = strings.Replace(n["name"], " ", "-",0)+"-"+strings.Replace(n["ip"], ".", "-",0)
    } else {
        nodeKey = n["id"]
    }
    if err := nodeExists(nodeKey); err == nil {
        return errors.New("Node desn't exist, must be created")
    }
    for key, value := range n {
        if id, _ := nodeKeyExists(nodeKey, key); id != 0 {
            err = nodeKeyUpdate(id, nodeKey, key, value)
        } else {
            err = nodeKeyInsert(nodeKey, key, value)
        }
    }
    if err != nil {
        return err
    }
    return nil
}

func getNodeIpbyName(n string)(ip string, err error) {
    if ndb.Db == nil {
        logs.Error("no access to database")
        return "", errors.New("no access to database")
    }
    sql := "select node_value from nodes where node_uniqueid like '%"+n+"%' and node_param = 'ip';"
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return "", err
    }
    defer rows.Close()
    if rows.Next() {
        if err = rows.Scan(&ip); err == nil {
            return ip, err
        }
    }
    return "", errors.New("There is no IP for given node name")
}

func GetAllNodes() (nodes *map[string]map[string]string, err error) {
    var allnodes = map[string]map[string]string{}
    var uniqid string
    var param string
    var value string
    if ndb.Db == nil {
        logs.Error("no access to database")
        return nil, errors.New("no access to database")
    }
    sql := "select node_uniqueid, node_param, node_value from nodes;"
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error("ndb.Db.Query Error : %s", err.Error())
        return nil, err
    }
    for rows.Next() {
        if err = rows.Scan(&uniqid, &param, &value); err != nil {
            logs.Error("GetAllNodes rows.Scan: %s", err.Error())
            return nil, err
        }
        logs.Info ("uniqid: %s, param: %s, value: %s", uniqid,param,value)
        if allnodes[uniqid] == nil { allnodes[uniqid] = map[string]string{}}
        allnodes[uniqid][param]=value
    } 
    return &allnodes, nil
}

func NodePing(n string) (err error) {
    ip, err := GetNodeIPbyUUID(n)
    if err != nil {
        logs.Info("Ping - IP Error -> %s", err.Error())
        return err
    }
    port, err := GetNodePortbyUUID(n)
    if err != nil {
        logs.Info("Ping - PORT Error -> %s", err.Error())
        return err
    }    
    err = nodeclient.PingNode(ip,port)
    if err != nil {
        return err
    }
    return nil
}

//Set ruleset file from Master to Node
func SetRuleset(uuid string) (err error) {
    logs.Info("SetRuleset node -->"+uuid)
    var portData string
    var ipData string
    
    //Take IP from specific uuid
	sqlIP := "select node_value from nodes where node_param = 'ip' and node_uniqueid = '"+uuid+"';"
	ip, err := ndb.Db.Query(sqlIP)
	if err != nil {
		logs.Error("GetAllFiles ndb.Db.Query Error  UUID: %s", err.Error())
		return err
	}
	defer ip.Close()
	if ip.Next() {
		ip.Scan(&ipData)
	}

	//Take PORT from specific uuid
	sqlPORT := "select node_value from nodes where node_param = 'port' and node_uniqueid = '"+uuid+"';"
	port, err := ndb.Db.Query(sqlPORT)
	if err != nil {
		logs.Error("SetRuleset ndb.Db.Query Error  UUID: %s", err.Error())
		return err
	}
	defer port.Close()
	if port.Next() {
		if err = port.Scan(&portData); err != nil {
			return err
		}
	}
    
    rulesetID, err := ruleset.GetRuleSelected(uuid)
    if err != nil {
        logs.Error("SetRuleset node ERROR GetRuleSelected: "+err.Error())
        return err
	}

    path, err := ruleset.GetRulesetPath(rulesetID)
    if err != nil {
        logs.Error("SetRuleset node ERROR GetRulesetPath: "+err.Error())
        return err
    }

    data, err := ioutil.ReadFile(path)
    if err != nil {
        logs.Error("SetRuleset ioutil.ReadFile ERROR: "+err.Error())
        return err
    }

	err = nodeclient.SetRuleset(ipData,portData,data)
	if err != nil {
		logs.Error("nodeclient.SetRuleset ERROR connection through http new Request: "+err.Error())
		return err
	}
    return nil
}

//Get specific file from node files
func GetNodeFile(loadFile map[string]string) (data map[string]string, err error) {
    rData := make(map[string]string)
    var voidArray map[string]string
    var portData string
	var ipData string
	
    //Take IP from specific uuid
	sqlIP := "select node_value from nodes where node_param = 'ip' and node_uniqueid = '"+loadFile["uuid"]+"';"
	ip, err := ndb.Db.Query(sqlIP)
	if err != nil {
		logs.Error("ndb.Db.Query Error  UUID: %s", err.Error())
		return voidArray, err
	}
	defer ip.Close()
	if ip.Next() {
		ip.Scan(&ipData)
	}

	//Take PORT from specific uuid
	sqlPORT := "select node_value from nodes where node_param = 'port' and node_uniqueid = '"+loadFile["uuid"]+"';"
	port, err := ndb.Db.Query(sqlPORT)
	if err != nil {
		logs.Error("ndb.Db.Query Error  UUID: %s", err.Error())
		return voidArray, err
	}
	defer port.Close()
	if port.Next() {
		if err = port.Scan(&portData); err != nil {
			return voidArray, err
		}
	}

	rData,err = nodeclient.GetNodeFile(ipData,portData,loadFile)
	if err != nil {
		logs.Error("node/GetNodeFile ERROR reading file: "+err.Error())
		return nil, err
	}

    return rData,nil
}


//Get specific file from node files
func SetNodeFile(loadFile map[string]string) (err error) {
    logs.Info("SetNodeFile node "+loadFile["uuid"])
    var portData string
    var ipData string
	sqlIP := "select node_value from nodes where node_param = 'ip' and node_uniqueid = '"+loadFile["uuid"]+"';"
	ip, err := ndb.Db.Query(sqlIP)
	if err != nil {
		logs.Error("SetNodeFile ndb.Db.Query Error  UUID: %s", err.Error())
		return err
	}
	defer ip.Close()
	if ip.Next() {
		ip.Scan(&ipData)
	}

	//Take PORT from specific uuid
	sqlPORT := "select node_value from nodes where node_param = 'port' and node_uniqueid = '"+loadFile["uuid"]+"';"
	port, err := ndb.Db.Query(sqlPORT)
	if err != nil {
		logs.Error("SetNodeFile ndb.Db.Query Error  UUID: %s", err.Error())
		return err
	}
	defer port.Close()
	if port.Next() {
		if err = port.Scan(&portData); err != nil {
			return err
		}
	}

	err = nodeclient.SetNodeFile(ipData,portData,loadFile)
	if err != nil {
		logs.Error("node/SetNodeFile ERROR request HTTP: "+err.Error())
		return err
	}
    return nil
}


func GetAllFiles(uuid string) (data map[string]string, err error) {
    var portData string
    var ipData string
    rData := make(map[string]string)

    //Take IP from specific uuid
	sqlIP := "select node_value from nodes where node_param = 'ip' and node_uniqueid = '"+uuid+"';"
	ip, err := ndb.Db.Query(sqlIP)
	if err != nil {
		logs.Error("GetAllFiles ndb.Db.Query Error  UUID: %s", err.Error())
		return rData, err
	}
	defer ip.Close()
	if ip.Next() {
		ip.Scan(&ipData)
	}

	//Take PORT from specific uuid
	sqlPORT := "select node_value from nodes where node_param = 'port' and node_uniqueid = '"+uuid+"';"
	port, err := ndb.Db.Query(sqlPORT)
	if err != nil {
		logs.Error("GetAllFiles ndb.Db.Query Error  UUID: %s", err.Error())
		return rData,err
	}
	defer port.Close()
	if port.Next() {
		if err = port.Scan(&portData); err != nil {
			return rData,err
		}
	}
	logs.Info("GetAllFiles PORT --> "+portData)

	resp,err := nodeclient.GetAllFiles(ipData,portData,uuid)
	if err != nil {
		logs.Error("node/GetAllFiles ERROR connection through http new Request: "+err.Error())
        return rData, err
	}

    return resp,err
}


func SyncRulesetToAllNodes(uuid string)(err error){
	if ndb.Rdb == nil {
        logs.Error("SyncRulesetToAllNodes -- Can't access to database")
        return errors.New("SyncRulesetToAllNodes -- Can't access to database")
    }
	sqlQuery := "SELECT node_uniqueid FROM ruleset_node WHERE ruleset_uniqueid = \""+uuid+"\" ;"
    rows, err := ndb.Rdb.Query(sqlQuery)
    if err != nil {
        logs.Error("SyncRulesetToAllNodes query error %s",err.Error())
        return err
    }
    defer rows.Close()
    for rows.Next() {
		var nodeID string
		err = rows.Scan(&nodeID)
		if err != nil {
			logs.Error("SyncRulesetToAllNodes FOR query error %s",err.Error())
			return err
		}
		err = SetRuleset(nodeID)
		if err != nil {
			logs.Error("SyncRulesetToAllNodes node.SetRuleset query error %s",err.Error())
			return err
		}
	}
	return nil
}

func ShowPorts(uuid string)(data map[string]string, err error){
    if ndb.Db == nil {
        logs.Error("ShowPorts -- Can't acces to database")
        return data,errors.New("ShowPorts -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/ShowPorts ERROR Obtaining Port and Ip: "+err.Error())
        return data,err
    }
	data, err = nodeclient.ShowPorts(ipnid,portnid)
	if err != nil {
		logs.Error("node/ShowPorts ERROR http data request: "+err.Error())
        return data,err
    }
	return data,nil
}
func PingPorts(uuid string)(data map[string]map[string]string, err error){
    if ndb.Db == nil {
        logs.Error("PingPorts -- Can't acces to database")
        return data,errors.New("PingPorts -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/PingPorts ERROR Obtaining Port and Ip: "+err.Error())
        return data,err
    }
	data, err = nodeclient.PingPorts(ipnid,portnid)
	if err != nil {
		logs.Error("node/PingPorts ERROR http data request: "+err.Error())
        return data,err
    }
	return data,nil
}

func ChangeMode(uuid string)(err error){
    if ndb.Db == nil {
        logs.Error("ChangeMode -- Can't acces to database")
        return errors.New("ChangeMode -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/ChangeMode ERROR Obtaining Port and Ip: "+err.Error())
        return err
    }
	err = nodeclient.ChangeMode(ipnid,portnid)
	if err != nil {
		logs.Error("node/ChangeMode ERROR http data request: "+err.Error())
        return err
    }
	return nil
}

func ChangeStatus(uuid string)(err error){
    if ndb.Db == nil {
        logs.Error("ChangeStatus -- Can't acces to database")
        return errors.New("ChangeStatus -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/ChangeStatus ERROR Obtaining Port and Ip: "+err.Error())
        return err
    }
	err = nodeclient.ChangeStatus(ipnid,portnid)
	if err != nil {
		logs.Error("node/ChangeStatus ERROR http data request: "+err.Error())
        return err
    }
	return nil
}