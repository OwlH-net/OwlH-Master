package node

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "errors"
    "os"
    "bufio"
    "time"
    "io/ioutil"
    "owlhmaster/nodeclient"
    "owlhmaster/utils"
    "regexp"
)

// //take Node ip from specific uuid
// func GetNodeIp(uuid string)(nodeIp string, err error){
//     var ipData string
//     sqlIP := "select node_value from nodes where node_param = 'ip' and node_uniqueid = '"+uuid+"';"
//     ip, err := ndb.Db.Query(sqlIP)
//     if err != nil {
//         logs.Error("SetRuleset ndb.Db.Query Error  UUID: %s", err.Error())
//         return "",err
//     }
//     defer ip.Close()
//     if ip.Next() {
//         // ip.Scan(&ipData)
//         if err = ip.Scan(&ipData); err != nil {
//             return "",err
//         }
//     }
//     return ipData,nil
// }

// //take Node ip from specific uuid
// func GetNodePort(uuid string)(nodePort string, err error){
//     var portData string
//     sqlPORT := "select node_value from nodes where node_param = 'port' and node_uniqueid = '"+uuid+"';"
//     port, err := ndb.Db.Query(sqlPORT)
//     if err != nil {
//         logs.Error("SetRuleset ndb.Db.Query Error  UUID: %s", err.Error())
//         return "",err
//     }
//     defer port.Close()
//     if port.Next() {
//         if err = port.Scan(&portData); err != nil {
//             return "",err
//         }
//     }
//     return portData,nil
// }

func AddNode(n map[string]string) (err error) {
    //check if exists a node with the same ip and port
    nodes,err:= ndb.GetAllNodes()
    for id := range nodes {
        if nodes[id]["ip"] == n["ip"]{
            if nodes[id]["port"] == n["port"]{
                return errors.New("AddNode - There is already a node with the same IP and Port")
            }
        }
    }
    
    //add node token to db
    uuid := utils.Generate()
    err = ndb.InsertNodeKey(uuid, "nodeuser", n["nodeuser"]); if err != nil {logs.Error("AddNode Insert node user error: "+err.Error()); return err}
    err = ndb.InsertNodeKey(uuid, "nodepass", n["nodepass"]); if err != nil {logs.Error("AddNode Insert node pass error: "+err.Error()); return err}
    err = ndb.InsertNodeKey(uuid, "name", n["name"]); if err != nil {logs.Error("AddNode Insert node name error: "+err.Error()); return err}
    err = ndb.InsertNodeKey(uuid, "port", n["port"]); if err != nil {logs.Error("AddNode Insert node port error: "+err.Error()); return err}
    err = ndb.InsertNodeKey(uuid, "ip", n["ip"]); if err != nil {logs.Error("AddNode Insert node ip error: "+err.Error()); return err}

    //get token
    login := make(map[string]string)
    masterid, err := ndb.LoadMasterID()
    login["user"] = n["nodeuser"]
    login["pass"] = n["nodepass"]
    login["master"] = masterid

    //Get token from node  
    token,err := nodeclient.GetNodeToken(n["ip"],n["port"], login)
    if err != nil {
        err = ndb.InsertNodeKey(uuid, "token", "wait"); if err != nil {logs.Error("AddNode Insert node token error: "+err.Error()); return err}
        return err
    }else{
        err = ndb.InsertNodeKey(uuid, "token", token); if err != nil {logs.Error("AddNode Insert node token error: "+err.Error()); return err}
    }    

    //Load token
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("AddNode Error loading node token: %s",err); return err}
    //Save node values into node db  
    nodeValues, err := ndb.GetNodeById(uuid); if err!=nil{logs.Error("AddNode Error loading node values: %s",err); return err}
    //delete data for node
    delete(nodeValues[uuid], "nodeuser")
    delete(nodeValues[uuid], "nodepass")
    delete(nodeValues[uuid], "token")
    err = nodeclient.SaveNodeInformation(n["ip"],n["port"], nodeValues)
    if err != nil {logs.Error("AddNode Error updating node data"); return err}    
    
    return nil
}

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
    //get node ip and port
    ipData,portData,err := ndb.ObtainPortIp(nodeid)
    if err != nil {logs.Error("node/DeleteNode ERROR getting node port/ip: "+err.Error()); return err}    

    //delete node from database
    err = ndb.DeleteNode(nodeid)
    if err != nil {logs.Error("DeleteNode error for uuid: "+nodeid+": "+ err.Error()); return err}
    
    //delete ruleset for this node
    err = ndb.DeleteRulesetNodeByNode(nodeid)
    if err != nil {logs.Error("DeleteNode DeleteRulesetNodeByNode error for uuid: "+nodeid+": "+ err.Error()); return err}

    //delete node for group
    groupnodes,err := ndb.GetGroupNodesByValue(nodeid)
    if err != nil {logs.Error("DeleteNode Execute ruleset_node -> %s", err.Error()); return err}
    for x := range groupnodes{
        err = ndb.DeleteNodeGroupById(x)
        if err != nil {logs.Error("DeleteNode error for uuid: "+x+": "+ err.Error()); return err}
    } 

    //delete node information at node db
    err = nodeclient.DeleteNode(ipData,portData)
    if err != nil {logs.Error("node/DeleteNode nodeclient ERROR: "+err.Error()); return err}   

    return nil
}

// func getNodeConf(nodeKey string)(conf map[string]string, err error) {
//     var param string
//     var value string

//     if ndb.Db == nil {
//         logs.Error("getNodeConf -> No access to database")
//         return nil, errors.New("getNodeConf -> No access to database")
//     }
    
//     sql := "SELECT node_param, node_value FROM nodes where node_uniqueid='"+nodeKey+"';"
//     logs.Info("GetNodeConf -> SQL -> %s", sql)
    
//     rows, err := ndb.Db.Query(sql)
    
//     if err != nil {
//         logs.Error(err.Error())
//         return nil, err
//     }
    
//     defer rows.Close()
//     for rows.Next() {
//         if err = rows.Scan(&param, &value); err != nil {
//             logs.Info ("Error rows.Scan -> %s",err.Error())
//             continue
//         }
//         conf[param]=value
//     }
//     return conf, nil
// }

func GetAllNodes()(data map[string]map[string]string, err error){
    allNodes,err := ndb.GetAllNodes()
    if err != nil {logs.Error("GetAllNodes error: "+err.Error()); return nil, err}

    for id := range allNodes {
        if allNodes[id]["token"] == "wait"{
            //get token
            login := make(map[string]string)
            masterid, err := ndb.LoadMasterID()
            if err != nil {logs.Error("node/GetAllNodes ERROR getting master id: "+err.Error()); return nil,err}    
            node, err := ndb.GetNodeById(id)
            if err != nil {logs.Error("node/GetAllNodes ERROR getting node id: "+err.Error()); return nil,err}    
            login["user"] = node[id]["nodeuser"]
            login["pass"] = node[id]["nodepass"]
            login["master"] = masterid
    
            //Get token from node 
            ipData,portData,err := ndb.ObtainPortIp(id)
            if err != nil { logs.Error("node/GetAllNodes ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}     
            token,err := nodeclient.GetNodeToken(ipData, portData, login)
            logs.Notice(token)
            if err != nil {
                logs.Emergency("node/GetAllNodes ERROR getting node id. Pending registering...")
            }else{
                err = ndb.UpdateNode(id, "token", token)  
                if err != nil {logs.Error("node/GetAllNodes ERROR updating node token: "+err.Error()); return nil,err}  
                allNodes[id]["token"] = token
                //delete data for node
                delete(allNodes[id], "nodeuser")
                delete(allNodes[id], "nodepass")
                delete(allNodes[id], "token")
                err = ndb.GetTokenByUuid(id); if err!=nil{logs.Error("GetAllNodes Error loading node token: %s",err); return nil,err}
                err = nodeclient.SaveNodeInformation(ipData, portData, allNodes)
                if err != nil {logs.Error("GetAllNodes Error updating node data"); return nil,err}    
            }
        }
    }

    return allNodes,nil
}

// func getAllNodesIp() (ips map[string]string, err error) {
//     var uid string
//     var ip string
//     if ndb.Db == nil {
//         logs.Error("getAllNodesIp -> no access to database")
//         return ips, errors.New("getAllNodesIp -> no access to database")
//     }
//     sql := "SELECT node_uniqueid, node_value FROM nodes where node_param = 'ip';"
//     rows, err := ndb.Db.Query(sql)
//     if err != nil {
//         logs.Error("Error ndb.Db.Query %s -> %s", sql, err.Error())
//         return ips, err
//     }
//     defer rows.Close()
//     for rows.Next() {
//         if err = rows.Scan(&uid, &ip); err != nil {
//             logs.Info (" Error -> rows.Scan -> %s",err.Error())
//         }
//         ips[uid]=ip
//     }
//     return ips, nil
// }

func nodeKeyExists(nodekey string, key string) (id int, err error) {
    nodesExists,err := ndb.NodeKeyExists(nodekey, key)
    if err != nil {logs.Error("Get all nodes error: "+err.Error()); return nodesExists, err}
    return nodesExists,err
}

func nodeExists(nodeid string) (err error) {
    node,err := ndb.GetNodeById(nodeid)
    if err != nil {logs.Error("Get node error: "+err.Error()); return err}
    if len(node) == 0 {logs.Error("Node not exists: "+err.Error()); return errors.New("Node does not exists.")}
    return err
}

// func nodeKeyUpdate(nkey string, key string, value string) (err error) {
//     err = ndb.UpdateNode(nkey, key, value)
//     if err != nil {logs.Error("Get node error: "+err.Error()); return err}
//     return err
// }

// func nodeKeyInsert(nkey string, key string, value string) (err error) {
//     err = ndb.InsertNodeKey(nkey, key, value)
//     if err != nil {logs.Error("Insert node error: "+err.Error()); return err}
//     return err
// }

func UpdateNode(n map[string]string) (err error) {
    //cehck if exists a node with the same ip and port
    nodes,err:= ndb.GetAllNodes()
    for id := range nodes {
        if nodes[id]["ip"] == n["ip"]{
            if nodes[id]["port"] == n["port"]{
                if id != n["id"]{
                    return errors.New("There is already a node with the same IP and Port")
                }
            }
        }
    }

    //update node
    err = ndb.UpdateNode(n["id"], "name", n["name"]);  if err != nil {logs.Error("UpdateNode name error: "+err.Error()); return err}
    err = ndb.UpdateNode(n["id"], "ip", n["ip"]);  if err != nil {logs.Error("UpdateNode ip error: "+err.Error()); return err}
    err = ndb.UpdateNode(n["id"], "port", n["port"]);  if err != nil {logs.Error("UpdateNode port error: "+err.Error()); return err}

    //update node
    nodeValues, err := ndb.GetNodeById(n["id"])
    if err != nil {logs.Error("node/NodePing ERROR getting node data for update : "+err.Error()); return err}

    err = ndb.GetTokenByUuid(n["id"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}

    ipnid,portnid,err := ndb.ObtainPortIp(n["id"])
    if err != nil { logs.Error("node/GetChangeControlNode ERROR Obtaining Port and Ip: "+err.Error()); return err}    
    err = nodeclient.UpdateNodeData(ipnid,portnid, nodeValues)
    if err != nil {logs.Error("Error updating node data")}

    return nil
}

func getNodeIpbyName(n string)(ip string, err error) {
    ip,err = ndb.GetNodeIpbyName(n)
    if err != nil {logs.Error("node/GetNodeIpbyName ERROR getting node port/ip: "+err.Error()); return "",err}    
    return ip,err
}

func NodePing(uuid string) (nodeResp map[string]string, err error) {
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("NodePing Error loading node token: %s",err); return nil, err}    
    ipData,portData,err := ndb.ObtainPortIp(uuid)
    if err != nil {logs.Error("node/NodePing ERROR getting node port/ip: "+err.Error()); return nil,err}    
    
    nodeResp, err = nodeclient.PingNode(ipData,portData)
    if err != nil {return nil,err}
    return nodeResp,err
}

func GetServiceStatus(uuid string) (err error) {
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipData,portData,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/GetServiceStatus ERROR getting node port/ip : "+err.Error())
        return err
    }    
    err = nodeclient.GetServiceStatus(ipData,portData)
    if err != nil {
        return err
    }
    return nil
}

func DeployService(uuid string)(err error){
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipData,portData,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/DeployService ERROR getting node port/ip : "+err.Error())
        return err
    }    
    err = nodeclient.DeployService(ipData,portData)
    if err != nil {
        return err
    }
    return nil
}

//Get specific file from node files
func GetNodeFile(loadFile map[string]string) (values map[string]string, err error) {  
    loadDataValue := map[string]map[string]string{}
    loadDataValue["analyzer"] = map[string]string{}
    loadDataValue["analyzer"]["conf"] = ""
    loadDataValue, err = utils.GetConf(loadDataValue)
    if err != nil {logs.Error("GetNodeFile error getting path from main.conf"); return nil,err}

    rData := make(map[string]string)
    if loadFile["file"] == "group-analyzer"{        
        fileReaded, err := ioutil.ReadFile(loadDataValue["analyzer"]["conf"])
        if err != nil {logs.Error("node/GetNodeFile ERROR getting analyzer from master: "+err.Error()); return nil, err}

        rData["fileContent"] = string(fileReaded)
        rData["fileName"] = loadFile["file"]        
    }else{
        err = ndb.GetTokenByUuid(loadFile["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return nil, err}
        ipData,portData,err := ndb.ObtainPortIp(loadFile["uuid"])
        if err != nil {logs.Error("node/GetNodeFile ERROR getting node port/ip: "+err.Error()); return nil, err}
    
        rData,err = nodeclient.GetNodeFile(ipData,portData,loadFile)
        if err != nil {logs.Error("node/GetNodeFile ERROR reading file: "+err.Error()); return nil, err}
    }

    return rData,nil
}


//Get specific file from node files
func SetNodeFile(saveFile map[string]string) (err error) {
    loadDataValue := map[string]map[string]string{}
    loadDataValue["analyzer"] = map[string]string{}
    loadDataValue["analyzer"]["conf"] = ""
    loadDataValue, err = utils.GetConf(loadDataValue)
    if err != nil {logs.Error("GetNodeFile error getting path from main.conf"); return err}

    if saveFile["uuid"] == "local"{
        bytearray := []byte(saveFile["content"])
        err = utils.WriteNewDataOnFile(loadDataValue["analyzer"]["conf"], bytearray)
    }else{
        err = ndb.GetTokenByUuid(saveFile["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
        ipData,portData,err := ndb.ObtainPortIp(saveFile["uuid"])
        if err != nil {logs.Error("node/SetNodeFile ERROR getting node port/ip : "+err.Error()); return err}    
    
        err = nodeclient.SetNodeFile(ipData,portData,saveFile)
        if err != nil {logs.Error("node/SetNodeFile ERROR request HTTP: "+err.Error()); return err}
    }
    return nil
}


func GetAllFiles(uuid string) (data map[string]string, err error) {
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil, err}
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
    
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil, err}
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

func PingPluginsNode(uuid string)(data map[string]map[string]string, err error){
    if ndb.Db == nil {
        logs.Error("PingPluginsNode -- Can't acces to database")
        return data,errors.New("PingPluginsNode -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil, err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/PingPluginsNode ERROR Obtaining Port and Ip: "+err.Error())
        return data,err
    }
    data, err = nodeclient.PingPluginsNode(ipnid,portnid)
    if err != nil {
        logs.Error("node/PingPluginsNode ERROR http data request: "+err.Error())
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
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
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
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
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
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
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
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
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

func PingPorts(uuid string)(data map[string]map[string]string, err error){
    if ndb.Db == nil {
        logs.Error("PingPorts -- Can't acces to database")
        return data,errors.New("PingPorts -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
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

func SyncRulesetToNode(anode map[string]string)(err error){
    rulesetUUID,err := ndb.GetRulesetUUID(anode["uuid"])
    if err != nil {logs.Error("SyncRulesetToNode/GetRulesetUUID error: "+err.Error()); return err}

    //read lines by ruleset uuid
    data, err := CreateNewRuleFile(rulesetUUID)
    if err != nil {logs.Error("nodeclient.SetRuleset ERROR creating a nunique ruleset file: "+err.Error()); return err}

    //send lines to node
    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipData,portData,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {logs.Error("node/GetAllFiles ERROR getting node port/ip : "+err.Error()); return err}    
    err = nodeclient.SyncRulesetToNode(ipData,portData,data)
    if err != nil {logs.Error("nodeclient.SetRuleset ERROR connection through http new Request: "+err.Error()); return err}

    return nil
}

//create new file with all de ruleset rules
func CreateNewRuleFile(uuid string)(data []byte, err error){
    var uniqueid string
    var rulePath string
    var uuidArray []string
    var validID = regexp.MustCompile(`sid:\s?(\d+);`)

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
            if err = rules.Scan(&rulePath); err != nil {logs.Error("CreateNewRuleFile rows.Scan: %s", err.Error()); return nil,err}
            file, err := os.Open(rulePath)
            if err != nil {
                logs.Error("File reading error: %s .Skipping file.", err.Error())
                continue
                // return nil, err
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
        err = ndb.GetTokenByUuid(nodeID); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
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
        if err != nil {logs.Error("nodeclient.SetRuleset ERROR connection through http new Request: "+err.Error()); return err}
    }
    return nil
}

func PingAnalyzer(uuid string)(data map[string]string, err error){
    if ndb.Db == nil {
        logs.Error("PingAnalyzer -- Can't acces to database")
        return data,errors.New("PingAnalyzer -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
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
    var nodeExists bool = true
    if ndb.Db == nil {logs.Error("ChangeAnalyzerStatus -- Can't acces to database"); return errors.New("ChangeAnalyzerStatus -- Can't acces to database")}
    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { 
        if anode["type"] != "groups" {
            logs.Error("node/ChangeAnalyzerStatus ERROR Obtaining Port and Ip: "+err.Error())
            return err
        }else{
            nodeExists = false
            logs.Error("node/ChangeAnalyzerStatus ERROR Obtaining Port and Ip for groups: "+err.Error())
        }
    }

    if nodeExists{
        err = nodeclient.ChangeAnalyzerStatus(ipnid,portnid,anode)
        if err != nil {logs.Error("node/ChangeAnalyzerStatus ERROR http data request: "+err.Error()); return err}
    }

    return nil
}

func DeployNode(anode map[string]string)(err error){
    if ndb.Db == nil {
        logs.Error("Deploy -- Can't acces to database")
        return errors.New("Deploy -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
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
    err := ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil}
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
    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
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
    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
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
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
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
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/LoadNetworkValues ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}

    anode,err = nodeclient.LoadNetworkValues(ipnid,portnid)
    if err != nil { logs.Error("node/LoadNetworkValues ERROR http data request: "+err.Error()); return nil,err}

    return anode,nil
}

func LoadNetworkValuesSelected(uuid string)(anode map[string]map[string]string, err error){
    if ndb.Db == nil {logs.Error("LoadNetworkValuesSelected -- Can't acces to database");return nil,err}
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/LoadNetworkValuesSelected ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}

    anode,err = nodeclient.LoadNetworkValuesSelected(ipnid,portnid)
    if err != nil { logs.Error("node/LoadNetworkValuesSelected ERROR http data request: "+err.Error()); return nil,err}

    return anode,nil
}

func SaveSocketToNetwork(anode map[string]string)(err error){    
    if ndb.Db == nil {logs.Error("SaveSocketToNetwork -- Can't acces to database");return err}
    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/SaveSocketToNetwork ERROR Obtaining Port and Ip: "+err.Error()); return err}

    err = nodeclient.SaveSocketToNetwork(ipnid,portnid,anode)
    if err != nil { logs.Error("node/SaveSocketToNetwork ERROR http data request: "+err.Error()); return err}

    return err
}

func SaveNewLocal(anode map[string]string)(err error){
    if ndb.Db == nil {logs.Error("SaveNewLocal -- Can't acces to database");return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/SaveNewLocal ERROR Obtaining Port and Ip: "+err.Error()); return err}

    err = nodeclient.SaveNewLocal(ipnid,portnid,anode)
    if err != nil { logs.Error("node/SaveNewLocal ERROR http data request: "+err.Error()); return err}

    return err
}

func SaveVxLAN(anode map[string]string)(err error){
    if ndb.Db == nil {logs.Error("SaveVxLAN -- Can't acces to database");return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/SaveVxLAN ERROR Obtaining Port and Ip: "+err.Error()); return err}

    err = nodeclient.SaveVxLAN(ipnid,portnid,anode)
    if err != nil { logs.Error("node/SaveVxLAN ERROR http data request: "+err.Error()); return err}

    return err
}

func SocketToNetworkList(uuid string)(data map[string]map[string]string, err error){
    if ndb.Db == nil {
        logs.Error("SocketToNetworkList -- Can't acces to database")
        return nil,err
    }
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/SocketToNetworkList ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}

    anode,err := nodeclient.SocketToNetworkList(ipnid,portnid)
    if err != nil { logs.Error("node/SocketToNetworkList ERROR http data request: "+err.Error()); return nil,err}

    return anode,nil
}

func SaveSocketToNetworkSelected(anode map[string]string)(err error){
    if ndb.Db == nil {logs.Error("SaveSocketToNetworkSelected -- Can't acces to database");return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/SaveSocketToNetworkSelected ERROR Obtaining Port and Ip: "+err.Error()); return err}

    err = nodeclient.SaveSocketToNetworkSelected(ipnid,portnid,anode)
    if err != nil { logs.Error("node/SaveSocketToNetworkSelected ERROR http data request: "+err.Error()); return err}

    return err
}

func DeleteDataFlowValueSelected(anode map[string]string)(err error){
    if ndb.Db == nil {logs.Error("DeleteDataFlowValueSelected -- Can't acces to database");return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/DeleteDataFlowValueSelected ERROR Obtaining Port and Ip: "+err.Error()); return err}

    err = nodeclient.DeleteDataFlowValueSelected(ipnid,portnid,anode)
    if err != nil { logs.Error("node/DeleteDataFlowValueSelected ERROR http data request: "+err.Error()); return err}

    return err
}


func GetNodeMonitor(uuid string)(data map[string]interface{}, err error){
    if ndb.Db == nil { logs.Error("GetNodeMonitor -- Can't acces to database"); return data,err}

    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil, err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/GetNodeMonitor ERROR Obtaining Port and Ip: "+err.Error()); return data,err}

    data,err = nodeclient.GetNodeMonitor(ipnid,portnid)
    if err != nil { logs.Error("node/GetNodeMonitor ERROR http data request: "+err.Error()); return data,err}

    return data,nil
}

func GetMainconfData(uuid string)(data map[string]map[string]string, err error){
    if ndb.Db == nil { logs.Error("GetMainconfData -- Can't acces to database"); return data,err}

    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil, err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/GetMainconfData ERROR Obtaining Port and Ip: "+err.Error()); return data,err}

    data,err = nodeclient.GetMainconfData(ipnid,portnid)
    if err != nil { logs.Error("node/GetMainconfData ERROR http data request: "+err.Error()); return data,err}

    return data,nil
}

func ChangeServiceStatus(anode map[string]string)(err error){
    if ndb.Db == nil {logs.Error("ChangeServiceStatus -- Can't acces to database");return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/ChangeServiceStatus ERROR Obtaining Port and Ip: "+err.Error()); return err}

    err = nodeclient.ChangeServiceStatus(ipnid,portnid,anode)
    if err != nil { logs.Error("node/ChangeServiceStatus ERROR http data request: "+err.Error()); return err}

    return err
}

func ChangeMainServiceStatus(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("ChangeMainServiceStatus -- Can't acces to database"); return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/ChangeMainServiceStatus ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.ChangeMainServiceStatus(ipnid,portnid,anode)
    if err != nil { logs.Error("node/ChangeMainServiceStatus ERROR http data request: "+err.Error()); return err}

    return nil
}

func DeleteService(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("DeleteService -- Can't acces to database"); return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/DeleteService ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.DeleteService(ipnid,portnid,anode)
    if err != nil { logs.Error("node/DeleteService ERROR http data request: "+err.Error()); return err}

    return nil
}

func DeployStapService(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("DeployStapService -- Can't acces to database"); return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/DeployStapService ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.DeployStapService(ipnid,portnid,anode)
    if err != nil { logs.Error("node/DeployStapService ERROR http data request: "+err.Error()); return err}

    return nil
}

func StopStapService(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("StopStapService -- Can't acces to database"); return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/StopStapService ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.StopStapService(ipnid,portnid,anode)
    if err != nil { logs.Error("node/StopStapService ERROR http data request: "+err.Error()); return err}

    return nil
}

func ModifyStapValues(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("ModifyStapValues -- Can't acces to database"); return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/ModifyStapValues ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.ModifyStapValues(ipnid,portnid,anode)
    if err != nil { logs.Error("node/ModifyStapValues ERROR http data request: "+err.Error()); return err}

    return nil
}

func ReloadFilesData(uuid string)(data map[string]map[string]string, err error){
    if ndb.Db == nil { logs.Error("ReloadFilesData -- Can't acces to database"); return nil,err}

    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/ReloadFilesData ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}
    
    data,err = nodeclient.ReloadFilesData(ipnid,portnid)
    if err != nil { logs.Error("node/ReloadFilesData ERROR http data request: "+err.Error()); return nil,err}

    return data,nil
}

func AddMonitorFile(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("AddMonitorFile -- Can't acces to database"); return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/AddMonitorFile ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.AddMonitorFile(ipnid,portnid,anode)
    if err != nil { logs.Error("node/AddMonitorFile ERROR http data request: "+err.Error()); return err}

    return nil
}

func PingMonitorFiles(uuid string)(data map[string]map[string]string, err error){
    if ndb.Db == nil { logs.Error("PingMonitorFiles -- Can't acces to database"); return nil,err}

    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil, err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/PingMonitorFiles ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}
    
    data,err = nodeclient.PingMonitorFiles(ipnid,portnid)
    if err != nil { logs.Error("node/PingMonitorFiles ERROR http data request: "+err.Error()); return nil,err}

    return data,nil
}

func DeleteMonitorFile(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("DeleteMonitorFile -- Can't acces to database"); return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/DeleteMonitorFile ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.DeleteMonitorFile(ipnid,portnid,anode)
    if err != nil { logs.Error("node/DeleteMonitorFile ERROR http data request: "+err.Error()); return err}

    return nil
}

func GetChangeControlNode(uuid string)(data map[string]map[string]string, err error){
    if ndb.Db == nil { logs.Error("GetChangeControlNode -- Can't acces to database"); return nil,err}

    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/GetChangeControlNode ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}
    
    data,err = nodeclient.GetChangeControlNode(ipnid,portnid)
    if err != nil { logs.Error("node/GetChangeControlNode ERROR http data request: "+err.Error()); return nil,err}

    return data,nil
}

func GetIncidentsNode(uuid string)(data map[string]map[string]string, err error){
    if ndb.Db == nil { logs.Error("GetIncidentsNode -- Can't acces to database"); return nil,err}

    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/GetIncidentsNode ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}
    
    data,err = nodeclient.GetIncidentsNode(ipnid,portnid)
    if err != nil { logs.Error("node/GetIncidentsNode ERROR http data request: "+err.Error()); return nil,err}

    return data,nil
}

func PutIncidentNode(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("PutIncidentNode -- Can't acces to database"); return err}

    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/PutIncidentNode ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.PutIncidentNode(ipnid,portnid,anode)
    if err != nil { logs.Error("node/PutIncidentNode ERROR http data request: "+err.Error()); return err}

    return nil
}

func SyncRulesetToAllGroupNodes(anode map[string]string)(err error){
    logs.Info(anode);
    nodesID,err := ndb.GetGroupNodesByUUID(anode["uuid"])
    if err != nil {logs.Error("SyncRulesetToAllGroupNodes error getting all nodes for a groups: "+err.Error()); return err}
    
    for x := range nodesID {
        //get node data by uuid
        if ndb.Db == nil { logs.Error("PutIncidentNode -- Can't acces to database"); return err}
        
        err = ndb.GetTokenByUuid(nodesID[x]["nodesid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
        ipnid,portnid,err := ndb.ObtainPortIp(nodesID[x]["nodesid"])
        if err != nil { logs.Error("node/PutIncidentNode ERROR Obtaining Port and Ip: "+err.Error()); return err}

        //get all rulesets for this node
        allGroupsNodes,err := ndb.GetAllGroupNodes()
        if err != nil {logs.Error("SyncRulesetToAllGroupNodes error getting all groupsnodes: "+err.Error()); return err}

        var rulesetsList []string
        for y := range allGroupsNodes {
            //get all ruleset id if this node is in this group
            if allGroupsNodes[y]["nodesid"] == nodesID[x]["nodesid"] {
                //get ruleset assigned to this group
                allGroups,err := ndb.GetAllGroups()
                if err != nil {logs.Error("SyncRulesetToAllGroupNodes error getting all groups for get their rulesetID: "+err.Error()); return err}

                groupsRuleset := allGroups[allGroupsNodes[y]["groupid"]]["rulesetID"]
                rulesetsList = append(rulesetsList, groupsRuleset)
            }
        }
        var rulePaths []string
        for ruleID := range rulesetsList{
            rules,err := ndb.GetAllRuleFiles()
            if err != nil {logs.Error("SyncRulesetToAllGroupNodes/GetAllRuleFiles error getting all rule files: "+err.Error()); return err}
            for r := range rules{
                if rules[r]["sourceUUID"] == rulesetsList[ruleID]{
                    rulePaths = append(rulePaths, rules[r]["path"])
                }
            }
        } 
        
        AllEnabledLines,err := utils.MergeAllFiles(rulePaths)
        if AllEnabledLines == nil { return errors.New("There are no rules for synchronize. Please, select a valid ruleset.")}

        //send lines to node
        err = nodeclient.SyncRulesetToNode(ipnid,portnid,AllEnabledLines)
        if err != nil {logs.Error("nodeclient.SetRuleset ERROR connection through http new Request: "+err.Error()); return err}

    }

    return nil
}

func PutSuricataServicesFromGroup(anode map[string]string)(err error){
    nodesID,err := ndb.GetGroupNodesByUUID(anode["uuid"])
    if err != nil {logs.Error("node/PutSuricataServicesFromGroup error getting all nodes for a groups: "+err.Error()); return err}
    
    for x := range nodesID {
        //get node data by uuid
        if ndb.Db == nil { logs.Error("node/PutSuricataServicesFromGroup -- Can't acces to database"); return err}
        
        err = ndb.GetTokenByUuid(nodesID[x]["nodesid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
        ipnid,portnid,err := ndb.ObtainPortIp(nodesID[x]["nodesid"])
        if err != nil { logs.Error("node/PutSuricataServicesFromGroup ERROR Obtaining Port and Ip: "+err.Error()); return err}

        //send Suricata services to node
        err = nodeclient.PutSuricataServicesFromGroup(ipnid,portnid,anode)
        if err != nil { logs.Error("node/PutSuricataServicesFromGroup ERROR http data request: "+err.Error()); return err}
    }  

    return nil
}

func SyncAnalyzerToAllGroupNodes(anode map[string]map[string]string)(log map[string]map[string]string, err error){
    loadDataValue := map[string]map[string]string{}
    loadDataValue["analyzer"] = map[string]string{}
    loadDataValue["analyzer"]["conf"] = ""
    loadDataValue, err = utils.GetConf(loadDataValue)
    if err != nil {logs.Error("GetNodeFile error getting path from main.conf"); return nil,err}

    logSync := make(map[string]map[string]string)
    var activeNode bool = true
    for x := range anode {
        //get node data by uuid
        if ndb.Db == nil { logs.Error("node/SyncAnalyzerToAllGroupNodes -- Can't acces to database"); return nil, err}

        err = ndb.GetTokenByUuid(anode[x]["nuuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
        ipnid,portnid,err := ndb.ObtainPortIp(anode[x]["nuuid"])
        if err != nil { 
            logs.Error("node/SyncAnalyzerToAllGroupNodes ERROR Obtaining Port and Ip: "+err.Error()); 
            activeNode = false
            //add to log
            if logSync[anode[x]["nuuid"]] == nil{ logSync[anode[x]["nuuid"]] = map[string]string{} }
            logSync[anode[x]["nuuid"]]["name"] = anode[x]["nname"]
            logSync[anode[x]["nuuid"]]["ip"] = anode[x]["nip"]
            logSync[anode[x]["nuuid"]]["status"] = "error"
        }

        if activeNode{
            //get analyzer file content
            // analyzerFile, err := ioutil.ReadFile("conf/analyzer.json")
            analyzerFile, err := ioutil.ReadFile(loadDataValue["analyzer"]["conf"])
            if err != nil { 
                logs.Error("node/SyncAnalyzerToAllGroupNodes ERROR getting analyzer file content: "+err.Error())
                if logSync[anode[x]["nuuid"]] == nil{ logSync[anode[x]["nuuid"]] = map[string]string{} }
                logSync[anode[x]["nuuid"]]["name"] = anode[x]["nname"]
                logSync[anode[x]["nuuid"]]["ip"] = anode[x]["nip"]
                logSync[anode[x]["nuuid"]]["status"] = "error"  
                logs.Warn("ERROR -> reading file")
            }else{
                //send Suricata services to node
                err = nodeclient.SyncAnalyzerToAllGroupNodes(ipnid,portnid,analyzerFile)
                if err != nil { 
                    logs.Error("node/SyncAnalyzerToAllGroupNodes ERROR http data request: "+err.Error())
                    //add to log
                    if logSync[anode[x]["nuuid"]] == nil{ logSync[anode[x]["nuuid"]] = map[string]string{} }
                    logSync[anode[x]["nuuid"]]["name"] = anode[x]["nname"]
                    logSync[anode[x]["nuuid"]]["ip"] = anode[x]["nip"]
                    logSync[anode[x]["nuuid"]]["status"] = "error"
                    logs.Warn("ERROR -> nodeclient")
                }else{
                    //add to log
                    if logSync[anode[x]["nuuid"]] == nil{ logSync[anode[x]["nuuid"]] = map[string]string{} }
                    logSync[anode[x]["nuuid"]]["name"] = anode[x]["nname"]
                    logSync[anode[x]["nuuid"]]["ip"] = anode[x]["nip"]
                    logSync[anode[x]["nuuid"]]["status"] = "success"
                    logs.Notice("SUCCESS")
                }     
            }        
        }
    }

    return logSync,nil
}

func SyncUsersToNode()(){
    for {                 
        masterID,err := ndb.LoadMasterID()
        if err != nil{logs.Error("node/SyncUsersToNode Error getting master ID: "+err.Error())}    
        //get all users
        users,err:= ndb.GetLoginData()
        if err != nil{logs.Error("node/SyncUsersToNode Error getting users: "+err.Error())}    
        userValues := make(map[string]map[string]string)
        for user := range users {
            userValues[user] = map[string]string{}
            userValues[user]["masterID"] = masterID
            userValues[user]["user"] = users[user]["user"]
            userValues[user]["type"] = "master"
            userValues[user]["status"] = "exists"
        }

        nodes,err:= ndb.GetAllNodes()
        if err != nil{logs.Error("node/SyncUsersToNode Error getting allNodes: "+err.Error())}    
        for id := range nodes {
            ipnid,portnid,err := ndb.ObtainPortIp(id)
            if err != nil{logs.Error("node/SyncUsersToNode Error getting Node ip and port: "+err.Error())}  

            err = ndb.GetTokenByUuid(id); if err!=nil{logs.Error("node/SyncUsersToNode Error loading node token: %s",err)}  
            err = nodeclient.SyncUsersToNode(ipnid,portnid,userValues)
            if err != nil{logs.Error("node/SyncUsersToNode Error: "+err.Error())}    
        }
        logs.Info("Users synchronized to nodes")
        time.Sleep(time.Minute*10)
    }
}

func ChangeRotationStatus(anode map[string]string)(err error){
    //get node data by uuid
    if ndb.Db == nil { logs.Error("node/ChangeRotationStatus -- Can't acces to database"); return err}
    
    //load token for this node
    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("node/ChangeRotationStatus Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/ChangeRotationStatus ERROR Obtaining Port and Ip: "+err.Error()); return err}

    //send Suricata services to node
    err = nodeclient.ChangeRotationStatus(ipnid,portnid,anode)
    if err != nil { logs.Error("node/ChangeRotationStatus ERROR http data request: "+err.Error()); return err}
    

    return nil
}

func EditRotation(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("node/EditRotation -- Can't acces to database"); return err}
    
    //load token for this node
    err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("node/EditRotation Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/EditRotation ERROR Obtaining Port and Ip: "+err.Error()); return err}

    err = nodeclient.EditRotation(ipnid,portnid,anode)
    if err != nil { logs.Error("node/EditRotation ERROR http data request: "+err.Error()); return err}
    

    return nil
}