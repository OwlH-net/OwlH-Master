package node

import (
    "github.com/astaxie/beego/logs"
    "strings"
    "owlhmaster/database"
	"errors"
	"os"
	"bufio"
    "owlhmaster/nodeclient"
    "owlhmaster/stap"
    "owlhmaster/utils"
    "regexp"
)

// //take Node ip from specific uuid
// func GetNodeIp(uuid string)(nodeIp string, err error){
// 	var ipData string
// 	sqlIP := "select node_value from nodes where node_param = 'ip' and node_uniqueid = '"+uuid+"';"
// 	ip, err := ndb.Db.Query(sqlIP)
// 	if err != nil {
// 		logs.Error("SetRuleset ndb.Db.Query Error  UUID: %s", err.Error())
// 		return "",err
// 	}
// 	defer ip.Close()
// 	if ip.Next() {
// 		// ip.Scan(&ipData)
// 		if err = ip.Scan(&ipData); err != nil {
// 			return "",err
// 		}
// 	}
// 	return ipData,nil
// }

// //take Node ip from specific uuid
// func GetNodePort(uuid string)(nodePort string, err error){
// 	var portData string
// 	sqlPORT := "select node_value from nodes where node_param = 'port' and node_uniqueid = '"+uuid+"';"
// 	port, err := ndb.Db.Query(sqlPORT)
// 	if err != nil {
// 		logs.Error("SetRuleset ndb.Db.Query Error  UUID: %s", err.Error())
// 		return "",err
// 	}
// 	defer port.Close()
// 	if port.Next() {
// 		if err = port.Scan(&portData); err != nil {
// 			return "",err
// 		}
// 	}
// 	return portData,nil
// }

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
        logs.Error("Prepare nodes -> %s", err.Error())
        return err
    }
    _, err = stmt.Exec(&nodeid)
    if err != nil {
        logs.Error("Execute nodes -> %s", err.Error())
        return err
	}
    deleteRulesetNode, err := ndb.Rdb.Prepare("delete from ruleset_node where node_uniqueid = ?")
    if err != nil {
        logs.Error("Prepare ruleset_node -> %s", err.Error())
        return err
    }
    _, err = deleteRulesetNode.Exec(&nodeid)
    if err != nil {
        logs.Error("Execute ruleset_node -> %s", err.Error())
        return err
    }
    return nil
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
    logs.Info("ADD NODE")
    logs.Info(n)
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
        return errors.New("Node doesn't exist, must be created")
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

func NodePing(uuid string) (err error) {
	ipData,portData,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/NodePing ERROR getting node port/ip : "+err.Error())
        return err
	}	
    err = nodeclient.PingNode(ipData,portData)
    if err != nil {
        return err
    }
    return nil
}

//Get specific file from node files
func GetNodeFile(loadFile map[string]string) (data map[string]string, err error) {
	ipData,portData,err := ndb.ObtainPortIp(loadFile["uuid"])
	if err != nil {
		logs.Error("node/GetNodeFile ERROR getting node port/ip: "+err.Error())
		return data, err
	}

	rData,err := nodeclient.GetNodeFile(ipData,portData,loadFile)
	if err != nil {
		logs.Error("node/GetNodeFile ERROR reading file: "+err.Error())
		return data, err
	}

    return rData,nil
}


//Get specific file from node files
func SetNodeFile(loadFile map[string]string) (err error) {
	ipData,portData,err := ndb.ObtainPortIp(loadFile["uuid"])
	if err != nil {
		logs.Error("node/SetNodeFile ERROR getting node port/ip : "+err.Error())
		return err
	}	

	err = nodeclient.SetNodeFile(ipData,portData,loadFile)
	if err != nil {
		logs.Error("node/SetNodeFile ERROR request HTTP: "+err.Error())
		return err
	}
    return nil
}


func GetAllFiles(uuid string) (data map[string]string, err error) {
    // rData := make(map[string]string)
	ipData,portData,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/GetAllFiles ERROR getting node port/ip : "+err.Error())
        return data, err
	}	

	resp,err := nodeclient.GetAllFiles(ipData,portData,uuid)
	if err != nil {
		logs.Error("node/GetAllFiles ERROR connection through http new Request: "+err.Error())
        return data, err
	}

    return resp,err
}

func ShowPorts(uuid string)(data map[string]map[string]string, err error){
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

func ChangeMode(anode map[string]string)(err error){
	uuid := anode["uuid"]
	mode := anode["mode"]
    if ndb.Db == nil {
        logs.Error("ChangeMode -- Can't acces to database")
        return errors.New("ChangeMode -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/ChangeMode ERROR Obtaining Port and Ip: "+err.Error())
        return err
    }
	err = nodeclient.ChangeMode(ipnid,portnid,mode)
	if err != nil {
		logs.Error("node/ChangeMode ERROR http data request: "+err.Error())
        return err
    }
	return nil
}

func ChangeStatus(anode map[string]string)(err error){
	uuid := anode["uuid"]
	status := anode["status"]
    if ndb.Db == nil {
        logs.Error("ChangeStatus -- Can't acces to database")
        return errors.New("ChangeStatus -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/ChangeStatus ERROR Obtaining Port and Ip: "+err.Error())
        return err
    }
	err = nodeclient.ChangeStatus(ipnid,portnid,status)
	if err != nil {
		logs.Error("node/ChangeStatus ERROR http data request: "+err.Error())
        return err
    }
	return nil
}

func DeletePorts(anode map[string]string, uuid string)(err error){
	if ndb.Db == nil {
        logs.Error("DeletePorts -- Can't acces to database")
        return errors.New("DeletePorts -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/DeletePorts ERROR Obtaining Port and Ip: "+err.Error())
        return err
    }
	err = nodeclient.DeletePorts(ipnid,portnid, anode)
	if err != nil {
		logs.Error("node/DeletePorts ERROR http data request: "+err.Error())
        return err
    }
	return nil
}

func DeleteAllPorts(uuid string)(err error){
	if ndb.Db == nil {
        logs.Error("DeleteAllPorts -- Can't acces to database")
        return errors.New("DeleteAllPorts -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/DeleteAllPorts ERROR Obtaining Port and Ip: "+err.Error())
        return err
    }
	err = nodeclient.DeleteAllPorts(ipnid,portnid)
	if err != nil {
		logs.Error("node/DeleteAllPorts ERROR http data request: "+err.Error())
        return err
    }
	return nil
}

func SyncRulesetToNode(anode map[string]string)(err error){
	uuid := anode["uuid"]
	var rulesetUUID string
	
	ipData,portData,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/GetAllFiles ERROR getting node port/ip : "+err.Error())
        return err
	}	
		
	//get ruleset uuid by node uuid
	sqlIP := "select ruleset_uniqueid from ruleset_node where node_uniqueid = '"+uuid+"';"
	ip, err := ndb.Rdb.Query(sqlIP)
	if err != nil {
		logs.Error("SetRuleset ndb.Db.Query Error  UUID: %s", err.Error())
		return err
	}
	defer ip.Close()
	if ip.Next() {
		if err = ip.Scan(&rulesetUUID); err != nil {
			return err
		}
	}
	//read lines by ruleset uuid
	data, err := CreateNewRuleFile(rulesetUUID)

	//send lines to node
	err = nodeclient.SyncRulesetToNode(ipData,portData,data)
	if err != nil {
		logs.Error("nodeclient.SetRuleset ERROR connection through http new Request: "+err.Error())
		return err
	}

	return nil
}

//create new file with all de ruleset rules
func CreateNewRuleFile(uuid string)(data []byte, err error){
	var uniqueid string
	var rulePath string
	var uuidArray []string
	var validID = regexp.MustCompile(`sid:(\d+);`)

	//read rule uuid
	uuidRules, err := ndb.Rdb.Query("select rule_uniqueid from rule_files where rule_value='"+uuid+"'")
	if err != nil {
		logs.Error("CreateNewRuleFile ndb.Rdb.Query Error checking rule_uniqueid for rule_files: %s", err.Error())
		return nil, err
	}
	defer uuidRules.Close()
	for uuidRules.Next() {
		if err = uuidRules.Scan(&uniqueid); err != nil {
			logs.Error("CreateNewRuleFile rows.Scan: %s", err.Error())
			return nil, err
		}
		uuidArray = append(uuidArray, uniqueid)
	}

	//read files paths and
	for x := range uuidArray{
		rules, err := ndb.Rdb.Query("select rule_value from rule_files where rule_param = 'path' and rule_uniqueid= '"+uuidArray[x]+"'")
		if err != nil {
			logs.Error("CreateNewRuleFile ndb.Rdb.Query Error loading files paths: %s", err.Error())
			return nil, err
		}
		defer rules.Close()
		for rules.Next() {
			if err = rules.Scan(&rulePath); err != nil {
				logs.Error("CreateNewRuleFile rows.Scan: %s", err.Error())
				return nil,err
			}
			file, err := os.Open(rulePath)
			if err != nil {
				logs.Error("File reading error: %s", err.Error())
				return nil, err
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan(){
				if validID.MatchString(scanner.Text()){
					data = append(data, scanner.Bytes()...)
					data = append(data, "\n"...)
				}
			}
		}	
	}
	return data,nil
}

func SyncRulesetToAllNodes(anode map[string]string)(err error){
	uuid := anode["uuid"]

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
		logs.Info(nodeID)
		ipData,portData,err := ndb.ObtainPortIp(nodeID)
		if err != nil {
			logs.Error("node/GetAllFiles ERROR getting node port/ip : "+err.Error())
			return err
		}	
		
		data,err := CreateNewRuleFile(uuid)
		if err != nil {
			logs.Error("SyncRulesetToAllNodes node.CreateNewRuleFile query error %s",err.Error())
			return err
		}

		//send lines to node
		err = nodeclient.SyncRulesetToNode(ipData,portData,data)
		if err != nil {
			logs.Error("nodeclient.SetRuleset ERROR connection through http new Request: "+err.Error())
			return err
		}
	}
	return nil
}

func PingAnalyzer(uuid string)(data map[string]string, err error){
    if ndb.Db == nil {
        logs.Error("PingAnalyzer -- Can't acces to database")
        return data,errors.New("PingAnalyzer -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/PingAnalyzer ERROR Obtaining Port and Ip: "+err.Error())
        return data,err
    }
	data, err = nodeclient.PingAnalyzer(ipnid,portnid)
	if err != nil {
		logs.Error("node/PingAnalyzer ERROR http data request: "+err.Error())
        return data,err
    }
	return data,nil
}

func ChangeAnalyzerStatus(anode map[string]string)(err error){
    if ndb.Db == nil {
        logs.Error("ChangeAnalyzerStatus -- Can't acces to database")
        return errors.New("ChangeAnalyzerStatus -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
	if err != nil {
		logs.Error("node/ChangeAnalyzerStatus ERROR Obtaining Port and Ip: "+err.Error())
        return err
    }
	err = nodeclient.ChangeAnalyzerStatus(ipnid,portnid,anode)
	if err != nil {
		logs.Error("node/ChangeAnalyzerStatus ERROR http data request: "+err.Error())
        return err
    }
	return nil
}

func DeployNode(anode map[string]string)(err error){
	if ndb.Db == nil {
        logs.Error("Deploy -- Can't acces to database")
        return errors.New("Deploy -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
	if err != nil {
		logs.Error("node/Deploy ERROR Obtaining Port and Ip: "+err.Error())
        return err
    }
	err = nodeclient.DeployNode(ipnid,portnid,anode)
	if err != nil {
		logs.Error("node/Deploy ERROR http data request: "+err.Error())
        return err
    }
	return nil

}

func CheckDeploy(uuid string)(anode map[string]string){
	if ndb.Db == nil {
        logs.Error("CheckDeploy -- Can't acces to database")
        return nil
	}
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/CheckDeploy ERROR Obtaining Port and Ip: "+err.Error())
        return nil
	}
	anode = nodeclient.CheckDeploy(ipnid,portnid)
	if err != nil {
		logs.Error("node/CheckDeploy ERROR http data request: "+err.Error())
        return nil
    }
	return anode
}

func ChangeDataflowValues(anode map[string]string)(err error){
	if ndb.Db == nil {
        logs.Error("ChangeDataflowValues -- Can't acces to database")
        return errors.New("ChangeDataflowValues -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
	if err != nil {
		logs.Error("node/ChangeDataflowValues ERROR Obtaining Port and Ip: "+err.Error())
        return err
    }
	err = nodeclient.ChangeDataflowValues(ipnid,portnid,anode)
	if err != nil {
		logs.Error("node/ChangeDataflowValues ERROR http data request: "+err.Error())
        return err
    }
	return nil
}

func UpdateNetworkInterface(anode map[string]string)(err error){
	if ndb.Db == nil {
        logs.Error("UpdateNetworkInterface -- Can't acces to database")
        return errors.New("UpdateNetworkInterface -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
	if err != nil {
		logs.Error("node/UpdateNetworkInterface ERROR Obtaining Port and Ip: "+err.Error())
        return err
    }
	err = nodeclient.UpdateNetworkInterface(ipnid,portnid,anode)
	if err != nil {
		logs.Error("node/UpdateNetworkInterface ERROR http data request: "+err.Error())
        return err
    }
	return nil
}

func LoadDataflowValues(uuid string)(anode map[string]map[string]string, err error){
	if ndb.Db == nil {
        logs.Error("LoadDataflowValues -- Can't acces to database")
        return nil,err
	}
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil { logs.Error("node/LoadDataflowValues ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}

	anode,err = nodeclient.LoadDataflowValues(ipnid,portnid)
	if err != nil { logs.Error("node/LoadDataflowValues ERROR http data request: "+err.Error()); return nil,err}

	return anode,nil
}

func LoadNetworkValues(uuid string)(anode map[string]string, err error){
	if ndb.Db == nil {
        logs.Error("LoadNetworkValues -- Can't acces to database")
        return nil,err
	}
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil { logs.Error("node/LoadNetworkValues ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}

	anode,err = nodeclient.LoadNetworkValues(ipnid,portnid)
	if err != nil { logs.Error("node/LoadNetworkValues ERROR http data request: "+err.Error()); return nil,err}

	return anode,nil
}

func LoadNetworkValuesSelected(uuid string)(anode map[string]map[string]string, err error){
	if ndb.Db == nil {
        logs.Error("LoadNetworkValuesSelected -- Can't acces to database")
        return nil,err
	}
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil { logs.Error("node/LoadNetworkValuesSelected ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}

	anode,err = nodeclient.LoadNetworkValuesSelected(ipnid,portnid)
	if err != nil { logs.Error("node/LoadNetworkValuesSelected ERROR http data request: "+err.Error()); return nil,err}

	return anode,nil
}