package node

import (
    "bufio"
    "encoding/json"
    "errors"
    "github.com/astaxie/beego/logs"
    "io/ioutil"
    "os"
    "owlhmaster/database"
    "owlhmaster/nodeclient"
    "owlhmaster/utils"
    "regexp"
    "strconv"
    "time"
)

func AddNode(n map[string]string) (err error) {
    //historical log
    logUuid := utils.Generate()
    currentTime := time.Now()
    timeFormated := currentTime.Format("2006-01-02T15:04:05")
    _ = ndb.InsertPluginCommand(logUuid, "date", timeFormated)
    _ = ndb.InsertPluginCommand(logUuid, "action", "AddNode")
    _ = ndb.InsertPluginCommand(logUuid, "description", "Add new node")

    //insert node
    uuid := utils.Generate()

    if n["ip"] == "" {
        return errors.New("AddNode - IP should not be an empty field")
    }
    if n["port"] == "" {
        return errors.New("AddNode - Port should not be an empty field")
    }

    //check if exists a node with the same ip and port
    nodes, err := ndb.GetAllNodes()
    for id := range nodes {
        if nodes[id]["ip"] == n["ip"] {
            if nodes[id]["port"] == n["port"] {
                return errors.New("AddNode - There is already a node with the same IP and Port")
            }
        }
    }

    //get token
    login := make(map[string]string)
    masterid, err := ndb.LoadMasterID()
    if err != nil {
        return errors.New("AddNode ERROR getting data for get token")
    }
    login["user"] = n["nodeuser"]
    login["pass"] = n["nodepass"]
    login["master"] = masterid

    //Get token from node and insert data
    token, err := nodeclient.GetNodeToken(n["ip"], n["port"], login)
    if err != nil {
        if err.Error() == "CreateMasterToken Incorrect Login credentials" {
            _ = ndb.InsertPluginCommand(uuid, "status", "Error")
            _ = ndb.InsertPluginCommand(uuid, "output", "Add node error: Invalid login credentials")
            return errors.New("AddNode ERROR. Invalid login credentials.")
        }
        err = ndb.InsertNodeKey(uuid, "status", "n/a")
        if err != nil {
            logs.Error("AddNode Insert node token error: " + err.Error())
            return err
        }
        err = ndb.InsertNodeKey(uuid, "token", "wait")
        if err != nil {
            logs.Error("AddNode Insert node token error: " + err.Error())
            return err
        }
        if n["nodeuser"] != "" {
            err = ndb.InsertNodeKey(uuid, "nodeuser", n["nodeuser"])
            if err != nil {
                logs.Error("AddNode Insert node user error: " + err.Error())
                return err
            }
        } else {
            return errors.New("Empty node user field")
        }
        if n["nodepass"] != "" {
            err = ndb.InsertNodeKey(uuid, "nodepass", n["nodepass"])
            if err != nil {
                logs.Error("AddNode Insert node pass error: " + err.Error())
                return err
            }
        } else {
            return errors.New("Empty node pass field")
        }
        if n["name"] != "" {
            err = ndb.InsertNodeKey(uuid, "name", n["name"])
            if err != nil {
                logs.Error("AddNode Insert node name error: " + err.Error())
                return err
            }
        } else {
            return errors.New("Empty name field ")
        }
        if n["port"] != "" {
            err = ndb.InsertNodeKey(uuid, "port", n["port"])
            if err != nil {
                logs.Error("AddNode Insert node port error: " + err.Error())
                return err
            }
        } else {
            return errors.New("Empty port field")
        }
        if n["ip"] != "" {
            err = ndb.InsertNodeKey(uuid, "ip", n["ip"])
            if err != nil {
                logs.Error("AddNode Insert node ip error: " + err.Error())
                return err
            }
        } else {
            return errors.New("Empty ip form")
        }

        _ = ndb.InsertPluginCommand(uuid, "status", "Error")
        _ = ndb.InsertPluginCommand(uuid, "output", "Add node error: Invalid token")
    } else {
        err = ndb.InsertNodeKey(uuid, "status", "n/a")
        if err != nil {
            logs.Error("AddNode Insert node token error: " + err.Error())
            return err
        }
        err = ndb.InsertNodeKey(uuid, "token", token)
        if err != nil {
            logs.Error("AddNode Insert node token error: " + err.Error())
            return err
        }
        if n["nodeuser"] != "" {
            err = ndb.InsertNodeKey(uuid, "nodeuser", n["nodeuser"])
            if err != nil {
                logs.Error("AddNode Insert node user error: " + err.Error())
                return err
            }
        } else {
            return errors.New("Empty form data")
        }
        if n["nodepass"] != "" {
            err = ndb.InsertNodeKey(uuid, "nodepass", n["nodepass"])
            if err != nil {
                logs.Error("AddNode Insert node pass error: " + err.Error())
                return err
            }
        } else {
            return errors.New("Empty form data")
        }
        if n["name"] != "" {
            err = ndb.InsertNodeKey(uuid, "name", n["name"])
            if err != nil {
                logs.Error("AddNode Insert node name error: " + err.Error())
                return err
            }
        } else {
            return errors.New("Empty form data")
        }
        if n["port"] != "" {
            err = ndb.InsertNodeKey(uuid, "port", n["port"])
            if err != nil {
                logs.Error("AddNode Insert node port error: " + err.Error())
                return err
            }
        } else {
            return errors.New("Empty form data")
        }
        if n["ip"] != "" {
            err = ndb.InsertNodeKey(uuid, "ip", n["ip"])
            if err != nil {
                logs.Error("AddNode Insert node ip error: " + err.Error())
                return err
            }
        } else {
            return errors.New("Empty form data")
        }

        //Sync user, group, roles and their relations to the new node
        SyncUsersToNode()
        SyncUserGroupRolesToNode()
        SyncRolesToNode()
        SyncGroupsToNode()
        SyncRolePermissions()
        SyncPermissions()
        SyncRoleGroups()
    }

    //Load token
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("AddNode Error loading node token: %s", err)
        return err
    }

    //Save node values into node db
    nodeValues, err := ndb.GetNodeById(uuid)
    if err != nil {
        logs.Error("AddNode Error loading node values: %s", err)
        return err
    }

    //delete data for node
    delete(nodeValues[uuid], "nodeuser")
    delete(nodeValues[uuid], "nodepass")
    delete(nodeValues[uuid], "token")
    err = nodeclient.SaveNodeInformation(n["ip"], n["port"], nodeValues)
    if err != nil {
        _ = ndb.InsertPluginCommand(uuid, "status", "Error")
        _ = ndb.InsertPluginCommand(uuid, "output", "Add node error: Error saving node information into node")
        logs.Error("AddNode Error updating node data")
        return err
    }

    _ = ndb.InsertPluginCommand(logUuid, "status", "Success")
    _ = ndb.InsertPluginCommand(logUuid, "output", "Node added successfully")
    return nil
}

func GetAllNodesReact() (data NodeList, err error) {
    //historical log
    logUuid := utils.Generate()
    currentTime := time.Now()
    timeFormated := currentTime.Format("2006-01-02T15:04:05")
    _ = ndb.InsertPluginCommand(logUuid, "date", timeFormated)
    _ = ndb.InsertPluginCommand(logUuid, "action", "GetAllNodesReact")
    _ = ndb.InsertPluginCommand(logUuid, "description", "Get all nodes")

    allNodes, err := ndb.GetAllNodes()
    if err != nil {
        logs.Error("GetAllNodesReact error getting all nodes from db: " + err.Error())
        return data, err
    }

    for id := range allNodes {
        
        values,err := NodePing(id)
        if err != nil {
            _ = ndb.UpdateNode(id, "status", "offline")
            logs.Error("GetAllNodesReact error ping all nodes: " + err.Error())
        }else if values["ack"] == "true" {
            _ = ndb.UpdateNode(id, "status", "online")
        }
        //get token
        login := make(map[string]string)
        masterid, err := ndb.LoadMasterID()
        if err != nil {
            logs.Error("node/GetAllNodesReact ERROR getting master id: " + err.Error())
            return data, err
        }

        login["user"] = allNodes[id]["nodeuser"]
        login["pass"] = allNodes[id]["nodepass"]
        login["master"] = masterid

        token, err := nodeclient.GetNodeToken(allNodes[id]["ip"], allNodes[id]["port"], login)
        if err != nil {
            if err.Error() == "connection refused" {
                _ = ndb.InsertPluginCommand(logUuid, "status", "Error")
                _ = ndb.InsertPluginCommand(logUuid, "output", "Add node error: Connection refused")
                continue
            } else {
                _ = ndb.InsertPluginCommand(logUuid, "status", "Error")
                _ = ndb.InsertPluginCommand(logUuid, "output", "Add node error: "+err.Error())

                err = ndb.UpdateNode(id, "token", "wait")
                if err != nil {
                    logs.Error("node/GetAllNodesReact ERROR updating node token: " + err.Error())
                    return data, err
                }
                logs.Warn("node/GetAllNodesReact ERROR getting node id. Pending registering...")
                allNodes[id]["token"] = "wait"
            }
        } else {
            if token == "" || (allNodes[id]["token"] != "wait" && allNodes[id]["token"] != token) {
                err = ndb.UpdateNode(id, "token", "wait")
                if err != nil {
                    logs.Error("node/GetAllNodesReact ERROR updating node token: " + err.Error())
                    return data, err
                }
                logs.Warn("node/GetAllNodesReact ERROR getting node id. Pending registering...")
                allNodes[id]["token"] = "wait"
            }
        }
    }

    _ = ndb.InsertPluginCommand(logUuid, "status", "Success")
    _ = ndb.InsertPluginCommand(logUuid, "output", "All nodes added successfully")

    newnodeValues := setNodesToSlice(allNodes)

    return newnodeValues, nil
}


//return nodes as objects
type NodeList struct {
    Nodes        []NewNodeValue 
}
type NewNodeValue struct {
    IP           string     `json:"ip"`
    Token        string     `json:"token"`
    Name         string     `json:"name"`
    Port         string     `json:"port"`
    Status       string     `json:"status"`
    NodeUser     string     `json:"nodeuser"`
    NodePass     string     `json:"nodepass"`
    UUID         string     `json:"uuid"`
}
func setNodesToSlice(nodes map[string]map[string]string)(val NodeList) {
    var nodeList NodeList
    for x := range nodes{
        values := NewNodeValue{}
        values.IP = nodes[x]["ip"]
        values.Token = nodes[x]["token"]
        values.Name = nodes[x]["name"]
        values.Port = nodes[x]["port"]
        values.Status = nodes[x]["status"]
        values.NodeUser = nodes[x]["nodeuser"]
        values.NodePass = nodes[x]["nodepass"]
        values.UUID = x

        nodeList.Nodes = append(nodeList.Nodes, values)
    }

    return nodeList
}

func GetAllNodes() (data map[string]map[string]string, err error) {
    //historical log
    logUuid := utils.Generate()
    currentTime := time.Now()
    timeFormated := currentTime.Format("2006-01-02T15:04:05")
    _ = ndb.InsertPluginCommand(logUuid, "date", timeFormated)
    _ = ndb.InsertPluginCommand(logUuid, "action", "GetAllNodes")
    _ = ndb.InsertPluginCommand(logUuid, "description", "Get all nodes")

    allNodes, err := ndb.GetAllNodes()
    if err != nil {
        logs.Error("GetAllNodes error getting all nodes from db: " + err.Error())
        return nil, err
    }

    for id := range allNodes {
        //get token
        login := make(map[string]string)
        masterid, err := ndb.LoadMasterID()
        if err != nil {
            logs.Error("node/GetAllNodes ERROR getting master id: " + err.Error())
            return nil, err
        }

        login["user"] = allNodes[id]["nodeuser"]
        login["pass"] = allNodes[id]["nodepass"]
        login["master"] = masterid

        token, err := nodeclient.GetNodeToken(allNodes[id]["ip"], allNodes[id]["port"], login)
        if err != nil {
            if err.Error() == "connection refused" {
                _ = ndb.InsertPluginCommand(logUuid, "status", "Error")
                _ = ndb.InsertPluginCommand(logUuid, "output", "Add node error: Connection refused")
                continue
            } else {
                _ = ndb.InsertPluginCommand(logUuid, "status", "Error")
                _ = ndb.InsertPluginCommand(logUuid, "output", "Add node error: "+err.Error())

                err = ndb.UpdateNode(id, "token", "wait")
                if err != nil {
                    logs.Error("node/GetAllNodes ERROR updating node token: " + err.Error())
                    return nil, err
                }
                logs.Warn("node/GetAllNodes ERROR getting node id. Pending registering...")
                allNodes[id]["token"] = "wait"
            }
        } else {
            if token == "" || (allNodes[id]["token"] != "wait" && allNodes[id]["token"] != token) {
                err = ndb.UpdateNode(id, "token", "wait")
                if err != nil {
                    logs.Error("node/GetAllNodes ERROR updating node token: " + err.Error())
                    return nil, err
                }
                logs.Warn("node/GetAllNodes ERROR getting node id. Pending registering...")
                allNodes[id]["token"] = "wait"
            }
        }
    }

    _ = ndb.InsertPluginCommand(logUuid, "status", "Success")
    _ = ndb.InsertPluginCommand(logUuid, "output", "All nodes added successfully")

    return allNodes, nil
}

func RegisterNode(uuid string) (err error) {
    allNodes, err := ndb.GetAllNodes()
    if err != nil {
        logs.Error("node/RegisterNode ERROR getting master id: " + err.Error())
        return err
    }

    login := make(map[string]string)
    masterid, err := ndb.LoadMasterID()
    if err != nil {
        logs.Error("node/RegisterNode ERROR getting master id: " + err.Error())
        return err
    }

    login["user"] = allNodes[uuid]["nodeuser"]
    login["pass"] = allNodes[uuid]["nodepass"]
    login["master"] = masterid

    if err != nil {
        logs.Error("RegisterNode error getting all nodes from db: " + err.Error())
        return err
    }
    token, err := nodeclient.GetNodeToken(allNodes[uuid]["ip"], allNodes[uuid]["port"], login)

    err = ndb.UpdateNode(uuid, "token", token)
    if err != nil {
        logs.Error("node/RegisterNode ERROR updating node token: " + err.Error())
        return err
    }

    //Sync again all user, group, roles and their relations to the new node due to db deleted
    SyncUsersToNode()
    SyncUserGroupRolesToNode()
    SyncRolesToNode()
    SyncGroupsToNode()
    SyncRolePermissions()
    SyncPermissions()
    SyncRoleGroups()

    allNodes[uuid]["token"] = token
    //delete data for node
    delete(allNodes[uuid], "nodeuser")
    delete(allNodes[uuid], "nodepass")
    delete(allNodes[uuid], "token")
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("RegisterNode Error loading node token: %s", err)
        return err
    }
    if err != nil {
        logs.Error("RegisterNode Error getting new token from database: " + err.Error())
        return err
    }

    err = nodeclient.SaveNodeInformation(allNodes[uuid]["ip"], allNodes[uuid]["port"], allNodes)
    if err != nil {
        logs.Error("RegisterNode Error updating node data")
        return err
    }

    return nil
}

func findNode(s string) (id string, err error) {
    if ndb.Db == nil {
        logs.Error("Find Node -> no access to database")
        return "", errors.New("Find Node -> no access to database")
    }
    ip, _ := regexp.Compile(`\d+\.\d+\.\d+\.\d+`)
    uuid, _ := regexp.Compile(`\w{8}-\w{4}-\w{4}-\w{4}-\w{12}`)
    sql := "SELECT node_uniqueid FROM nodes where node_param = 'name' and node_value='" + s + "';"
    if ip.MatchString(s) {
        sql = "SELECT node_uniqueid FROM nodes where node_param = 'ip' and node_value='" + s + "';"
    } else if uuid.MatchString(s) {
        sql = "SELECT node_uniqueid FROM nodes where node_param = 'UUID' and node_value='" + s + "';"
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

func DeleteNode(nodeid string) (err error) {
    //get node ip and port
    ipData, portData, err := ndb.ObtainPortIp(nodeid)
    if err != nil {
        logs.Error("node/DeleteNode ERROR getting node port/ip: " + err.Error())
        return err
    }

    //delete node from database
    err = ndb.DeleteNode(nodeid)
    if err != nil {
        logs.Error("DeleteNode error for uuid: " + nodeid + ": " + err.Error())
        return err
    }

    //delete ruleset for this node
    err = ndb.DeleteRulesetNodeByNode(nodeid)
    if err != nil {
        logs.Error("DeleteNode DeleteRulesetNodeByNode error for uuid: " + nodeid + ": " + err.Error())
        return err
    }

    //delete node for group
    groupnodes, err := ndb.GetGroupNodesByValue(nodeid)
    if err != nil {
        logs.Error("DeleteNode Execute ruleset_node -> %s", err.Error())
        return err
    }
    for x := range groupnodes {
        err = ndb.DeleteNodeGroupById(x)
        if err != nil {
            logs.Error("DeleteNode error for uuid: " + x + ": " + err.Error())
            return err
        }
    }

    ///get master iu
    ndb.LoadMasterID()

    //delete node information at node db
    err = nodeclient.DeleteNode(ipData, portData)
    if err != nil {
        logs.Error("node/DeleteNode nodeclient ERROR: " + err.Error())
        return err
    }

    return nil
}

func nodeKeyExists(nodekey string, key string) (id int, err error) {
    nodesExists, err := ndb.NodeKeyExists(nodekey, key)
    if err != nil {
        logs.Error("Get all nodes error: " + err.Error())
        return nodesExists, err
    }
    return nodesExists, err
}

func nodeExists(nodeid string) (err error) {
    node, err := ndb.GetNodeById(nodeid)
    if err != nil {
        logs.Error("Get node error: " + err.Error())
        return err
    }
    if len(node) == 0 {
        logs.Error("Node not exists: " + err.Error())
        return errors.New("Node does not exists.")
    }
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
    nodes, err := ndb.GetAllNodes()
    for id := range nodes {
        if nodes[id]["ip"] == n["ip"] {
            if nodes[id]["port"] == n["port"] {
                if id != n["id"] {
                    return errors.New("There is already a node with the same IP and Port")
                }
            }
        }
    }

    //update node
    err = ndb.UpdateNode(n["id"], "name", n["name"])
    if err != nil {
        logs.Error("UpdateNode name error: " + err.Error())
        return err
    }
    err = ndb.UpdateNode(n["id"], "ip", n["ip"])
    if err != nil {
        logs.Error("UpdateNode ip error: " + err.Error())
        return err
    }
    err = ndb.UpdateNode(n["id"], "port", n["port"])
    if err != nil {
        logs.Error("UpdateNode port error: " + err.Error())
        return err
    }

    //update node
    nodeValues, err := ndb.GetNodeById(n["id"])
    if err != nil {
        logs.Error("node/NodePing ERROR getting node data for update : " + err.Error())
        return err
    }

    err = ndb.GetTokenByUuid(n["id"])
    if err != nil {
        logs.Error("UpdateNode Error loading node token: %s", err)
        return err
    }

    ipnid, portnid, err := ndb.ObtainPortIp(n["id"])
    if err != nil {
        logs.Error("node/GetChangeControlNode ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }
    err = nodeclient.UpdateNodeData(ipnid, portnid, nodeValues)
    if err != nil {
        logs.Error("Error updating node data")
    }

    return nil
}

func getNodeIpbyName(n string) (ip string, err error) {
    ip, err = ndb.GetNodeIpbyName(n)
    if err != nil {
        logs.Error("node/GetNodeIpbyName ERROR getting node port/ip: " + err.Error())
        return "", err
    }
    return ip, err
}

func NodePing(uuid string) (nodeResp map[string]string, err error) {
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("NodePing Error loading node token: %s", err)
        return nil, err
    }
    ipData, portData, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/NodePing ERROR getting node port/ip: " + err.Error())
        return nil, err
    }

    nodeResp, err = nodeclient.PingNode(ipData, portData)
    if err != nil {
        return nil, err
    }
    return nodeResp, err
}

func GetServiceStatus(uuid string) (err error) {
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("GetServiceStatus Error loading node token: %s", err)
        return err
    }
    ipData, portData, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/GetServiceStatus ERROR getting node port/ip : " + err.Error())
        return err
    }
    err = nodeclient.GetServiceStatus(ipData, portData)
    if err != nil {
        return err
    }
    return nil
}

func DeployService(uuid string) (err error) {
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("DeployService Error loading node token: %s", err)
        return err
    }
    ipData, portData, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/DeployService ERROR getting node port/ip : " + err.Error())
        return err
    }
    err = nodeclient.DeployService(ipData, portData)
    if err != nil {
        return err
    }
    return nil
}

//Get specific file from node files
func GetNodeFile(loadFile map[string]string) (values map[string]string, err error) {
    conf, err := utils.GetKeyValueString("analyzer", "conf")
    if err != nil {
        logs.Error("GetNodeFile error getting path from main.conf")
        return nil, err
    }

    rData := make(map[string]string)
    if loadFile["file"] == "group-analyzer" {
        fileReaded, err := ioutil.ReadFile(conf)
        if err != nil {
            logs.Error("node/GetNodeFile ERROR getting analyzer from master: " + err.Error())
            return nil, err
        }

        rData["fileContent"] = string(fileReaded)
        rData["fileName"] = loadFile["file"]
    } else {
        err = ndb.GetTokenByUuid(loadFile["uuid"])
        if err != nil {
            logs.Error("GetNodeFile Error loading node token: %s", err)
            return nil, err
        }
        ipData, portData, err := ndb.ObtainPortIp(loadFile["uuid"])
        if err != nil {
            logs.Error("node/GetNodeFile ERROR getting node port/ip: " + err.Error())
            return nil, err
        }

        rData, err = nodeclient.GetNodeFile(ipData, portData, loadFile)
        if err != nil {
            logs.Error("node/GetNodeFile ERROR reading file: " + err.Error())
            return nil, err
        }
    }

    return rData, nil
}

//Get specific file from node files
func SetNodeFile(saveFile map[string]string) (err error) {
    conf, err := utils.GetKeyValueString("analyzer", "conf")
    if err != nil {
        logs.Error("GetNodeFile error getting path from main.conf")
        return err
    }

    if saveFile["uuid"] == "local" {
        bytearray := []byte(saveFile["content"])
        err = utils.WriteNewDataOnFile(conf, bytearray)
    } else {
        err = ndb.GetTokenByUuid(saveFile["uuid"])
        if err != nil {
            logs.Error("SetNodeFile Error loading node token: %s", err)
            return err
        }
        ipData, portData, err := ndb.ObtainPortIp(saveFile["uuid"])
        if err != nil {
            logs.Error("node/SetNodeFile ERROR getting node port/ip : " + err.Error())
            return err
        }

        err = nodeclient.SetNodeFile(ipData, portData, saveFile)
        if err != nil {
            logs.Error("node/SetNodeFile ERROR request HTTP: " + err.Error())
            return err
        }
    }
    return nil
}

func GetAllFiles(uuid string) (data map[string]string, err error) {
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("GetAllFiles Error loading node token: %s", err)
        return nil, err
    }
    ipData, portData, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/GetAllFiles ERROR getting node port/ip : " + err.Error())
        return data, err
    }

    resp, err := nodeclient.GetAllFiles(ipData, portData, uuid)
    if err != nil {
        logs.Error("node/GetAllFiles ERROR connection through http new Request: " + err.Error())
        return data, err
    }

    return resp, err
}

func ShowPorts(uuid string) (data map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("ShowPorts -- Can't acces to database")
        return data, errors.New("ShowPorts -- Can't acces to database")
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("ShowPorts Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/ShowPorts ERROR Obtaining Port and Ip: " + err.Error())
        return data, err
    }
    data, err = nodeclient.ShowPorts(ipnid, portnid)
    if err != nil {
        logs.Error("node/ShowPorts ERROR http data request: " + err.Error())
        return data, err
    }
    return data, nil
}

func PingPluginsNode(uuid string) (data map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("PingPluginsNode -- Can't acces to database")
        return data, errors.New("PingPluginsNode -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("PingPluginsNode Error loading node token: %s", err)
        return nil, err
    }

    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/PingPluginsNode ERROR Obtaining Port and Ip: " + err.Error())
        return data, err
    }

    data, err = nodeclient.PingPluginsNode(ipnid, portnid)
    if err != nil {
        logs.Error("node/PingPluginsNode ERROR http data request: " + err.Error())
        return data, err
    }

    return data, nil
}

func ChangeMode(anode map[string]string) (err error) {
    uuid := anode["uuid"]
    mode := anode["mode"]
    if ndb.Db == nil {
        logs.Error("ChangeMode -- Can't acces to database")
        return errors.New("ChangeMode -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("ChangeMode Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/ChangeMode ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }
    err = nodeclient.ChangeMode(ipnid, portnid, mode)
    if err != nil {
        logs.Error("node/ChangeMode ERROR http data request: " + err.Error())
        return err
    }
    return nil
}

func ChangeStatus(anode map[string]string) (err error) {
    uuid := anode["uuid"]
    status := anode["status"]
    if ndb.Db == nil {
        logs.Error("ChangeStatus -- Can't acces to database")
        return errors.New("ChangeStatus -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("ChangeStatus Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/ChangeStatus ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }
    err = nodeclient.ChangeStatus(ipnid, portnid, status)
    if err != nil {
        logs.Error("node/ChangeStatus ERROR http data request: " + err.Error())
        return err
    }
    return nil
}

func DeletePorts(anode map[string]string, uuid string) (err error) {
    if ndb.Db == nil {
        logs.Error("DeletePorts -- Can't acces to database")
        return errors.New("DeletePorts -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("DeletePorts Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/DeletePorts ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }
    err = nodeclient.DeletePorts(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/DeletePorts ERROR http data request: " + err.Error())
        return err
    }
    return nil
}

func DeleteAllPorts(uuid string) (err error) {
    if ndb.Db == nil {
        logs.Error("DeleteAllPorts -- Can't acces to database")
        return errors.New("DeleteAllPorts -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("DeleteAllPorts Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/DeleteAllPorts ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }
    err = nodeclient.DeleteAllPorts(ipnid, portnid)
    if err != nil {
        logs.Error("node/DeleteAllPorts ERROR http data request: " + err.Error())
        return err
    }
    return nil
}

func PingPorts(uuid string) (data map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("PingPorts -- Can't acces to database")
        return data, errors.New("PingPorts -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("PingPorts Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/PingPorts ERROR Obtaining Port and Ip: " + err.Error())
        return data, err
    }
    data, err = nodeclient.PingPorts(ipnid, portnid)
    if err != nil {
        logs.Error("node/PingPorts ERROR http data request: " + err.Error())
        return data, err
    }
    return data, nil
}

func SyncRulesetToNode(anode map[string]string) (err error) {
    //read lines by ruleset uuid
    data, err := CreateNewRuleFile(anode["ruleset"])
    if err != nil {
        logs.Error("SyncRulesetToNode ERROR creating a nunique ruleset file: " + err.Error())
        return err
    }

    //get ruleset name
    rsetName, err := ndb.GetRulesetSourceValue(anode["ruleset"], "name")
    if err != nil {
        logs.Error("SyncRulesetToNode Error getting ruleset name: %s", err)
        return err
    }

    //send lines to node
    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("SyncRulesetToNode Error loading node token: %s", err)
        return err
    }

    ipData, portData, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/GetAllFiles ERROR getting node port/ip : " + err.Error())
        return err
    }

    err = nodeclient.SyncRulesetToNode(ipData, portData, data, anode["service"], rsetName)
    if err != nil {
        logs.Error("SyncRulesetToNode ERROR connection through http new Request: " + err.Error())
        return err
    }

    return nil
}

//create new file with all de ruleset rules
func CreateNewRuleFile(uuid string) (data []byte, err error) {
    var uniqueid string
    var rulePath string
    var uuidArray []string
    var validID = regexp.MustCompile(`sid:\s?(\d+);`)

    //read rule uuid
    uuidRules, err := ndb.Rdb.Query("select rule_uniqueid from rule_files where rule_value='" + uuid + "'")
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
    for x := range uuidArray {
        rules, err := ndb.Rdb.Query("select rule_value from rule_files where rule_param = 'path' and rule_uniqueid= '" + uuidArray[x] + "'")
        if err != nil {
            logs.Error("CreateNewRuleFile ndb.Rdb.Query Error loading files paths: %s", err.Error())
            return nil, err
        }
        defer rules.Close()
        for rules.Next() {
            if err = rules.Scan(&rulePath); err != nil {
                logs.Error("CreateNewRuleFile rows.Scan: %s", err.Error())
                return nil, err
            }
            file, err := os.Open(rulePath)
            if err != nil {
                logs.Error("File reading error: %s .Skipping file.", err.Error())
                continue
                // return nil, err
            }
            scanner := bufio.NewScanner(file)
            for scanner.Scan() {
                if validID.MatchString(scanner.Text()) {
                    data = append(data, scanner.Bytes()...)
                    data = append(data, "\n"...)
                }
            }
        }
    }
    return data, nil
}

func SyncRulesetToAllNodes(anode map[string]string) (err error) {
    uuid := anode["uuid"]

    if ndb.Rdb == nil {
        logs.Error("SyncRulesetToAllNodes -- Can't access to database")
        return errors.New("SyncRulesetToAllNodes -- Can't access to database")
    }

    sqlQuery := "SELECT node_uniqueid FROM ruleset_node WHERE ruleset_uniqueid = \"" + uuid + "\" ;"
    rows, err := ndb.Rdb.Query(sqlQuery)
    if err != nil {
        logs.Error("SyncRulesetToAllNodes query error %s", err.Error())
        return err
    }
    defer rows.Close()
    for rows.Next() {
        var nodeID string
        err = rows.Scan(&nodeID)
        err = ndb.GetTokenByUuid(nodeID)
        if err != nil {
            logs.Error("SyncRulesetToAllNodes Error loading node token: %s", err)
            return err
        }
        ipData, portData, err := ndb.ObtainPortIp(nodeID)
        if err != nil {
            logs.Error("node/GetAllFiles ERROR getting node port/ip : " + err.Error())
            return err
        }

        data, err := CreateNewRuleFile(uuid)
        if err != nil {
            logs.Error("SyncRulesetToAllNodes node.CreateNewRuleFile query error %s", err.Error())
            return err
        }

        //get ruleset name
        rsetName, err := ndb.GetRulesetSourceValue(uuid, "name")
        if err != nil {
            logs.Error("SyncRulesetToAllNodes Error getting ruleset name: %s", err)
            return err
        }

        //send lines to node
        err = nodeclient.SyncRulesetToNode(ipData, portData, data, uuid, rsetName)
        if err != nil {
            logs.Error("SyncRulesetToAllNodes ERROR connection through http new Request: " + err.Error())
            return err
        }
    }
    return nil
}

func PingAnalyzer(uuid string) (data map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("PingAnalyzer -- Can't acces to database")
        return data, errors.New("PingAnalyzer -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("PingAnalyzer Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/PingAnalyzer ERROR Obtaining Port and Ip: " + err.Error())
        return data, err
    }
    data, err = nodeclient.PingAnalyzer(ipnid, portnid)
    if err != nil {
        logs.Error("node/PingAnalyzer ERROR http data request: " + err.Error())
        return data, err
    }
    return data, nil
}

func ChangeAnalyzerStatus(anode map[string]string) (err error) {
    var nodeExists bool = true
    if ndb.Db == nil {
        logs.Error("ChangeAnalyzerStatus -- Can't acces to database")
        return errors.New("ChangeAnalyzerStatus -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("ChangeAnalyzerStatus Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        if anode["type"] != "groups" {
            logs.Error("node/ChangeAnalyzerStatus ERROR Obtaining Port and Ip: " + err.Error())
            return err
        } else {
            nodeExists = false
            logs.Error("node/ChangeAnalyzerStatus ERROR Obtaining Port and Ip for groups: " + err.Error())
        }
    }

    if nodeExists {
        err = nodeclient.ChangeAnalyzerStatus(ipnid, portnid, anode)
        if err != nil {
            logs.Error("node/ChangeAnalyzerStatus ERROR http data request: " + err.Error())
            return err
        }
    }

    return nil
}

func DeployNode(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("Deploy -- Can't acces to database")
        return errors.New("Deploy -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("DeployNode Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/Deploy ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }
    err = nodeclient.DeployNode(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/Deploy ERROR http data request: " + err.Error())
        return err
    }
    return nil

}

func CheckDeploy(uuid string) (anode map[string]string) {
    if ndb.Db == nil {
        logs.Error("CheckDeploy -- Can't acces to database")
        return nil
    }
    err := ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("CheckDeploy Error loading node token: %s", err)
        return nil
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/CheckDeploy ERROR Obtaining Port and Ip: " + err.Error())
        return nil
    }
    anode = nodeclient.CheckDeploy(ipnid, portnid)
    if err != nil {
        logs.Error("node/CheckDeploy ERROR http data request: " + err.Error())
        return nil
    }
    return anode
}

func ChangeDataflowValues(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("ChangeDataflowValues -- Can't acces to database")
        return errors.New("ChangeDataflowValues -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("ChangeDataflowValues Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/ChangeDataflowValues ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }
    err = nodeclient.ChangeDataflowValues(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/ChangeDataflowValues ERROR http data request: " + err.Error())
        return err
    }
    return nil
}

func UpdateNetworkInterface(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("UpdateNetworkInterface -- Can't acces to database")
        return errors.New("UpdateNetworkInterface -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("UpdateNetworkInterface Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/UpdateNetworkInterface ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }
    err = nodeclient.UpdateNetworkInterface(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/UpdateNetworkInterface ERROR http data request: " + err.Error())
        return err
    }
    return nil
}

func LoadDataflowValues(uuid string) (anode map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("LoadDataflowValues -- Can't acces to database")
        return nil, err
    }
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("LoadDataflowValues Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/LoadDataflowValues ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    anode, err = nodeclient.LoadDataflowValues(ipnid, portnid)
    if err != nil {
        logs.Error("node/LoadDataflowValues ERROR http data request: " + err.Error())
        return nil, err
    }

    return anode, nil
}

func LoadNetworkValues(uuid string) (anode map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("LoadNetworkValues -- Can't acces to database")
        return nil, err
    }
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("LoadNetworkValues Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/LoadNetworkValues ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    anode, err = nodeclient.LoadNetworkValues(ipnid, portnid)
    if err != nil {
        logs.Error("node/LoadNetworkValues ERROR http data request: " + err.Error())
        return nil, err
    }

    return anode, nil
}

func LoadNetworkValuesSelected(uuid string) (anode map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("LoadNetworkValuesSelected -- Can't acces to database")
        return nil, err
    }
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("LoadNetworkValuesSelected Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/LoadNetworkValuesSelected ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    anode, err = nodeclient.LoadNetworkValuesSelected(ipnid, portnid)
    if err != nil {
        logs.Error("node/LoadNetworkValuesSelected ERROR http data request: " + err.Error())
        return nil, err
    }

    return anode, nil
}

func SaveSocketToNetwork(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("SaveSocketToNetwork -- Can't acces to database")
        return err
    }
    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("SaveSocketToNetwork Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/SaveSocketToNetwork ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.SaveSocketToNetwork(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/SaveSocketToNetwork ERROR http data request: " + err.Error())
        return err
    }

    return err
}

func SaveNewLocal(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("SaveNewLocal -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("SaveNewLocal Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/SaveNewLocal ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.SaveNewLocal(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/SaveNewLocal ERROR http data request: " + err.Error())
        return err
    }

    return err
}

func SaveVxLAN(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("SaveVxLAN -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("SaveVxLAN Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/SaveVxLAN ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.SaveVxLAN(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/SaveVxLAN ERROR http data request: " + err.Error())
        return err
    }

    return err
}

func SocketToNetworkList(uuid string) (data map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("SocketToNetworkList -- Can't acces to database")
        return nil, err
    }
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("SocketToNetworkList Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/SocketToNetworkList ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    anode, err := nodeclient.SocketToNetworkList(ipnid, portnid)
    if err != nil {
        logs.Error("node/SocketToNetworkList ERROR http data request: " + err.Error())
        return nil, err
    }

    return anode, nil
}

func SaveSocketToNetworkSelected(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("SaveSocketToNetworkSelected -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("SaveSocketToNetworkSelected Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/SaveSocketToNetworkSelected ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.SaveSocketToNetworkSelected(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/SaveSocketToNetworkSelected ERROR http data request: " + err.Error())
        return err
    }

    return err
}

func DeleteDataFlowValueSelected(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("DeleteDataFlowValueSelected -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("DeleteDataFlowValueSelected Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/DeleteDataFlowValueSelected ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.DeleteDataFlowValueSelected(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/DeleteDataFlowValueSelected ERROR http data request: " + err.Error())
        return err
    }

    return err
}

func GetNodeMonitor(uuid string) (data map[string]interface{}, err error) {
    if ndb.Db == nil {
        logs.Error("GetNodeMonitor -- Can't acces to database")
        return data, err
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Warn("GetNodeMonitor Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/GetNodeMonitor ERROR Obtaining Port and Ip: " + err.Error())
        return data, err
    }

    data, err = nodeclient.GetNodeMonitor(ipnid, portnid)
    if err != nil {
        logs.Error("node/GetNodeMonitor ERROR http data request: " + err.Error())
        return data, err
    }

    return data, nil
}

func GetMainconfData(uuid string) (data map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("GetMainconfData -- Can't acces to database")
        return data, err
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("GetMainconfData Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/GetMainconfData ERROR Obtaining Port and Ip: " + err.Error())
        return data, err
    }

    data, err = nodeclient.GetMainconfData(ipnid, portnid)
    if err != nil {
        logs.Error("node/GetMainconfData ERROR http data request: " + err.Error())
        return data, err
    }

    return data, nil
}

func ChangeServiceStatus(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("ChangeServiceStatus -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("ChangeServiceStatus Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/ChangeServiceStatus ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.ChangeServiceStatus(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/ChangeServiceStatus ERROR http data request: " + err.Error())
        return err
    }

    return err
}

func ChangeMainServiceStatus(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("ChangeMainServiceStatus -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("ChangeMainServiceStatus Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/ChangeMainServiceStatus ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.ChangeMainServiceStatus(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/ChangeMainServiceStatus ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func DeleteService(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("DeleteService -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("DeleteService Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/DeleteService ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.DeleteService(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/DeleteService ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func DeployStapService(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("DeployStapService -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("DeployStapService Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/DeployStapService ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.DeployStapService(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/DeployStapService ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func StopStapService(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("StopStapService -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("StopStapService Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/StopStapService ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.StopStapService(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/StopStapService ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func ModifyNodeOptionValues(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("ModifyNodeOptionValues -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("ModifyNodeOptionValues Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/ModifyNodeOptionValues ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.ModifyNodeOptionValues(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/ModifyNodeOptionValues ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func ReloadFilesData(uuid string) (data map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("ReloadFilesData -- Can't acces to database")
        return nil, err
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("ReloadFilesData Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/ReloadFilesData ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    data, err = nodeclient.ReloadFilesData(ipnid, portnid)
    if err != nil {
        logs.Error("node/ReloadFilesData ERROR http data request: " + err.Error())
        return nil, err
    }

    return data, nil
}

func AddMonitorFile(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("AddMonitorFile -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("AddMonitorFile Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/AddMonitorFile ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.AddMonitorFile(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/AddMonitorFile ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func PingMonitorFiles(uuid string) (data map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("PingMonitorFiles -- Can't acces to database")
        return nil, err
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("PingMonitorFiles Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/PingMonitorFiles ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    data, err = nodeclient.PingMonitorFiles(ipnid, portnid)
    if err != nil {
        logs.Error("node/PingMonitorFiles ERROR http data request: " + err.Error())
        return nil, err
    }

    return data, nil
}

func DeleteMonitorFile(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("DeleteMonitorFile -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("DeleteMonitorFile Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/DeleteMonitorFile ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.DeleteMonitorFile(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/DeleteMonitorFile ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func GetChangeControlNode(uuid string) (data map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("GetChangeControlNode -- Can't acces to database")
        return nil, err
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("GetChangeControlNode Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/GetChangeControlNode ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    data, err = nodeclient.GetChangeControlNode(ipnid, portnid)
    if err != nil {
        logs.Error("node/GetChangeControlNode ERROR http data request: " + err.Error())
        return nil, err
    }

    return data, nil
}

func GetIncidentsNode(uuid string) (data map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("GetIncidentsNode -- Can't acces to database")
        return nil, err
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("GetIncidentsNode Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/GetIncidentsNode ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    data, err = nodeclient.GetIncidentsNode(ipnid, portnid)
    if err != nil {
        logs.Error("node/GetIncidentsNode ERROR http data request: " + err.Error())
        return nil, err
    }

    return data, nil
}

func PutIncidentNode(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("PutIncidentNode -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("PutIncidentNode Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/PutIncidentNode ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.PutIncidentNode(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/PutIncidentNode ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

type reportitem struct {
    Node    string `json:"node"`
    Success bool   `json:"success"`
    Desc    string `json:"desc"`
    Result  string `json:"result"`
}

func SyncRulesetToAllGroupNodes(anode map[string]string) (err error) {
    reportItemArray := []reportitem{}
    syncReport := reportitem{}
    var pingNodes = map[string]map[string]string{}

    //get all rulesets for a group

    nodesID, err := ndb.GetGroupNodesByUUID(anode["uuid"])
    if err != nil {
        logs.Error("SyncRulesetToAllGroupNodes error getting all nodes for a groups: " + err.Error())
        return err
    }

    for x := range nodesID {
        node, err := ndb.GetNodeById(nodesID[x]["nodesid"])
        if err != nil {
            syncReport.Node = nodesID[x]["nodesid"]
            syncReport.Success = false
            syncReport.Result = err.Error()
        } else {
            syncReport.Node = nodesID[x]["nodesid"]
            syncReport.Success = true
            syncReport.Result = "Success"
            syncReport.Desc = "Get node for sync group ruleset: " + node[nodesID[x]["nodesid"]]["name"] + " - " + node[nodesID[x]["nodesid"]]["ip"] + " - " + node[nodesID[x]["nodesid"]]["port"]
            //create ready nodes hashmap
            if pingNodes[nodesID[x]["nodesid"]] == nil {
                pingNodes[nodesID[x]["nodesid"]] = map[string]string{}
            }
            for param := range node[nodesID[x]["nodesid"]] {
                pingNodes[nodesID[x]["nodesid"]][param] = node[nodesID[x]["nodesid"]][param]
            }
        }

        reportItemArray = append(reportItemArray, syncReport)

    }
    if pingNodes == nil {
        for x := range reportItemArray {
            logs.Info("Node: " + reportItemArray[x].Node)
            logs.Info("Success: " + strconv.FormatBool(reportItemArray[x].Success))
            logs.Info("Result: " + reportItemArray[x].Result)
            if reportItemArray[x].Success {
                logs.Info("Description: " + reportItemArray[x].Desc)
            }
        }
    }

    //Get ruleset data for this group

    // allGroups,err := ndb.GetAllGroups()
    // if err != nil {logs.Error("SyncRulesetToAllGroupNodes error getting all groups: "+err.Error()); return err}
    // data,err := CreateNewRuleFile(allGroups[anode["uuid"]]["rulesetID"])
    // if err != nil {logs.Error("SyncRulesetToAllGroupNodes error creating ruleset file: "+err.Error()); return err}

    for x := range pingNodes {
        logs.Warn("NODE --> " + pingNodes[x]["name"])
        //get all groupRulesets
        groupRulesets, err := ndb.GetAllGroupRulesets()
        if err != nil {
            logs.Error("SyncRulesetToAllGroupNodes GetAllGroupRulesets error: " + err.Error())
            return err
        }

        for f := range groupRulesets {
            if groupRulesets[f]["groupid"] == anode["uuid"] {
                //get ruleset content by ruleset id
                data, err := CreateNewRuleFile(groupRulesets[f]["rulesetid"])
                if err != nil {
                    logs.Error("SyncRulesetToAllGroupNodes error creating ruleset file: " + err.Error())
                    return err
                }

                //get ruleset name
                name, err := ndb.GetRulesetSourceValue(groupRulesets[f]["rulesetid"], "name")
                if err != nil {
                    logs.Error("SyncRulesetToAllGroupNodes error creating ruleset file: " + err.Error())
                    return err
                }

                values := make(map[string][]byte)
                values["data"] = data
                values["name"] = []byte(name)

                logs.Notice(string(len(data)) + " Ruleset for this group --> " + name)

                //get node token
                err = ndb.GetTokenByUuid(x)
                if err != nil {
                    logs.Error("SyncRulesetToAllGroupNodes Error loading node token: %s", err)
                    continue
                }

                //sync ruleset to node
                err = nodeclient.SyncGroupRulesetToNode(pingNodes[x]["ip"], pingNodes[x]["port"], values)
                if err != nil {
                    syncReport.Node = pingNodes[x]["name"]
                    syncReport.Success = false
                    syncReport.Result = err.Error()
                } else {
                    syncReport.Success = true
                    syncReport.Result = "Success"
                    syncReport.Node = pingNodes[x]["name"]
                    syncReport.Desc = "Ruleset sync to node " + pingNodes[x]["name"] + " success!"
                }
                reportItemArray = append(reportItemArray, syncReport)

            }
        }

    }

    //PRINT REPORT HERE
    for x := range reportItemArray {
        logs.Info("Node: " + reportItemArray[x].Node)
        logs.Info("Success: " + strconv.FormatBool(reportItemArray[x].Success))
        logs.Info("Result: " + reportItemArray[x].Result)
        if reportItemArray[x].Success {
            logs.Info("Description: " + reportItemArray[x].Desc)
        }
    }

    return nil
}

func PutSuricataServicesFromGroup(anode map[string]string) (err error) {
    nodesID, err := ndb.GetGroupNodesByUUID(anode["uuid"])
    if err != nil {
        logs.Error("node/PutSuricataServicesFromGroup error getting all nodes for a groups: " + err.Error())
        return err
    }

    for x := range nodesID {
        //get node data by uuid
        if ndb.Db == nil {
            logs.Error("node/PutSuricataServicesFromGroup -- Can't acces to database")
            return err
        }

        err = ndb.GetTokenByUuid(nodesID[x]["nodesid"])
        if err != nil {
            logs.Error("PutSuricataServicesFromGroup Error loading node token: %s", err)
            return err
        }
        ipnid, portnid, err := ndb.ObtainPortIp(nodesID[x]["nodesid"])
        if err != nil {
            logs.Error("node/PutSuricataServicesFromGroup ERROR Obtaining Port and Ip: " + err.Error())
            return err
        }

        //send Suricata services to node
        err = nodeclient.PutSuricataServicesFromGroup(ipnid, portnid, anode)
        if err != nil {
            logs.Error("node/PutSuricataServicesFromGroup ERROR http data request: " + err.Error())
            return err
        }
    }

    return nil
}

func SyncAnalyzerToAllGroupNodes(anode map[string]string) (log map[string]map[string]string, err error) {
    conf, err := utils.GetKeyValueString("analyzer", "conf")
    if err != nil {
        logs.Error("GetNodeFile error getting path from main.conf")
        return nil, err
    }

    logSync := make(map[string]map[string]string)
    var activeNode bool = true

    //get all group nodes
    nodesID, err := ndb.GetGroupNodesByUUID(anode["uuid"])
    if err != nil {
        logs.Error("SyncAnalyzerToAllGroupNodes error getting all nodes for a groups: " + err.Error())
        return nil, err
    }
    //get all nodes
    allNodes, err := ndb.GetAllNodes()
    if err != nil {
        logs.Error("SyncAnalyzerToAllGroupNodes error getting all nodes: " + err.Error())
        return nil, err
    }

    //loop over all group nodes
    for x := range nodesID {
        //get node data by uuid
        if ndb.Db == nil {
            logs.Error("node/SyncAnalyzerToAllGroupNodes -- Can't acces to database")
            return nil, err
        }

        err = ndb.GetTokenByUuid(nodesID[x]["nodesid"])
        if err != nil {
            logs.Error("SyncAnalyzerToAllGroupNodes Error loading node token: %s", err)
            return nil, err
        }
        ipnid, portnid, err := ndb.ObtainPortIp(nodesID[x]["nodesid"])
        if err != nil {
            logs.Error("node/SyncAnalyzerToAllGroupNodes ERROR Obtaining Port and Ip: " + err.Error())
            activeNode = false
            //add to log
            if logSync[nodesID[x]["nodesid"]] == nil {
                logSync[nodesID[x]["nodesid"]] = map[string]string{}
            }
            logSync[nodesID[x]["nodesid"]]["name"] = allNodes[nodesID[x]["nodesid"]]["name"]
            logSync[nodesID[x]["nodesid"]]["ip"] = allNodes[nodesID[x]["nodesid"]]["ip"]
            logSync[nodesID[x]["nodesid"]]["status"] = "error"
        }
        if activeNode {
            //get analyzer file content
            analyzerFile, err := ioutil.ReadFile(conf)
            if err != nil {
                logs.Error("node/SyncAnalyzerToAllGroupNodes ERROR getting analyzer file content: " + err.Error())
                if logSync[nodesID[x]["nodesid"]] == nil {
                    logSync[nodesID[x]["nodesid"]] = map[string]string{}
                }
                logSync[nodesID[x]["nodesid"]]["name"] = allNodes[nodesID[x]["nodesid"]]["name"]
                logSync[nodesID[x]["nodesid"]]["ip"] = allNodes[nodesID[x]["nodesid"]]["ip"]
                logSync[nodesID[x]["nodesid"]]["status"] = "error"
                logs.Warn("ERROR -> reading file")
            } else {
                //send Suricata services to node
                err = nodeclient.SyncAnalyzerToAllGroupNodes(ipnid, portnid, analyzerFile)
                if err != nil {
                    logs.Error("node/SyncAnalyzerToAllGroupNodes ERROR http data request: " + err.Error())
                    //add to log
                    if logSync[nodesID[x]["nodesid"]] == nil {
                        logSync[nodesID[x]["nodesid"]] = map[string]string{}
                    }
                    logSync[nodesID[x]["nodesid"]]["name"] = allNodes[nodesID[x]["nodesid"]]["name"]
                    logSync[nodesID[x]["nodesid"]]["ip"] = allNodes[nodesID[x]["nodesid"]]["ip"]
                    logSync[nodesID[x]["nodesid"]]["status"] = "error"
                    logs.Warn("ERROR -> nodeclient")
                } else {
                    //add to log
                    if logSync[nodesID[x]["nodesid"]] == nil {
                        logSync[nodesID[x]["nodesid"]] = map[string]string{}
                    }
                    logSync[nodesID[x]["nodesid"]]["name"] = allNodes[nodesID[x]["nodesid"]]["name"]
                    logSync[nodesID[x]["nodesid"]]["ip"] = allNodes[nodesID[x]["nodesid"]]["ip"]
                    logSync[nodesID[x]["nodesid"]]["status"] = "success"
                    logs.Notice("SUCCESS")
                }
            }
        }
    }

    return logSync, nil
}

func SyncUsersToNode() {
    masterID, err := ndb.LoadMasterID()
    if err != nil {
        logs.Error("node/SyncUsersToNode Error getting master ID: " + err.Error())
    }
    //get all users
    users, err := ndb.GetLoginData()
    if err != nil {
        logs.Error("node/SyncUsersToNode Error getting users: " + err.Error())
    }
    userValues := make(map[string]map[string]string)
    for user := range users {
        userValues[user] = map[string]string{}
        userValues[user]["masterID"] = masterID
        userValues[user]["user"] = users[user]["user"]
        userValues[user]["type"] = "master"
        userValues[user]["status"] = "exists"
    }

    nodes, err := ndb.GetAllNodes()
    if err != nil {
        logs.Error("node/SyncUsersToNode Error getting allNodes: " + err.Error())
    }
    for id := range nodes {
        ipnid, portnid, err := ndb.ObtainPortIp(id)
        if err != nil {
            logs.Error("node/SyncUsersToNode Error getting Node ip and port: " + err.Error())
        }

        err = ndb.GetTokenByUuid(id)
        if err != nil {
            logs.Error("node/SyncUsersToNode Error loading node token: %s", err)
        }
        err = nodeclient.SyncUsersToNode(ipnid, portnid, userValues)
        if err != nil {
            if err.Error() == "VerifyToken - Incorrect Token" {
                ndb.UpdateNode(id, "token", "wait")
            }
            logs.Error("Error synchronizing users to node '" + nodes[id]["name"] + "': " + err.Error())
        } else {
            logs.Info("Users synchronized to node '" + nodes[id]["name"] + "'")
        }
    }
}

func SyncRolesToNode() {
    masterID, err := ndb.LoadMasterID()
    if err != nil {
        logs.Error("node/SyncRolesToNode Error getting master ID: " + err.Error())
    }
    //get all roles
    roles, err := ndb.GetUserRoles()
    if err != nil {
        logs.Error("node/SyncRolesToNode Error getting roles: " + err.Error())
    }
    roleValues := make(map[string]map[string]string)
    for role := range roles {
        roleValues[role] = map[string]string{}
        roleValues[role]["masterID"] = masterID
        roleValues[role]["role"] = roles[role]["role"]
        // roleValues[role]["permissions"] = roles[role]["permissions"]
        roleValues[role]["type"] = "master"
        roleValues[role]["status"] = "exists"
    }

    nodes, err := ndb.GetAllNodes()
    if err != nil {
        logs.Error("node/SyncRolesToNode Error getting allNodes: " + err.Error())
    }
    for id := range nodes {
        ipnid, portnid, err := ndb.ObtainPortIp(id)
        if err != nil {
            logs.Error("node/SyncRolesToNode Error getting Node ip and port: " + err.Error())
        }

        err = ndb.GetTokenByUuid(id)
        if err != nil {
            logs.Error("node/SyncRolesToNode Error loading node token: %s", err)
        }
        //get user uuid
        err = nodeclient.SyncRolesToNode(ipnid, portnid, roleValues)
        if err != nil {
            if err.Error() == "VerifyToken - Incorrect Token" {
                ndb.UpdateNode(id, "token", "wait")
            }
            logs.Error("Error synchronizing roles to node '" + nodes[id]["name"] + "': " + err.Error())
        } else {
            logs.Info("Roles synchronized to node '" + nodes[id]["name"] + "'")
        }
    }
}

func SyncGroupsToNode() {
    masterID, err := ndb.LoadMasterID()
    if err != nil {
        logs.Error("node/SyncGroupsToNode Error getting master ID: " + err.Error())
    }
    //get all groups
    groups, err := ndb.GetUserGroups()
    if err != nil {
        logs.Error("node/SyncGroupsToNode Error getting groups: " + err.Error())
    }
    groupValues := make(map[string]map[string]string)
    for group := range groups {
        groupValues[group] = map[string]string{}
        groupValues[group]["masterID"] = masterID
        groupValues[group]["group"] = groups[group]["group"]
        groupValues[group]["type"] = "master"
        groupValues[group]["status"] = "exists"
    }

    nodes, err := ndb.GetAllNodes()
    if err != nil {
        logs.Error("node/SyncGroupsToNode Error getting allNodes: " + err.Error())
    }
    for id := range nodes {
        ipnid, portnid, err := ndb.ObtainPortIp(id)
        if err != nil {
            logs.Error("node/SyncGroupsToNode Error getting Node ip and port: " + err.Error())
        }

        err = ndb.GetTokenByUuid(id)
        if err != nil {
            logs.Error("node/SyncGroupsToNode Error loading node token: %s", err)
        }
        //get user uuid
        err = nodeclient.SyncGroupsToNode(ipnid, portnid, groupValues)
        if err != nil {
            if err.Error() == "VerifyToken - Incorrect Token" {
                ndb.UpdateNode(id, "token", "wait")
            }
            logs.Error("Error synchronizing groups to node '" + nodes[id]["name"] + "': " + err.Error())
        } else {
            logs.Info("Groups synchronized to node '" + nodes[id]["name"] + "'")
        }
    }
}

func SyncUserGroupRolesToNode() {
    masterID, err := ndb.LoadMasterID()
    if err != nil {
        logs.Error("node/SyncUserGroupRolesToNode Error getting master ID: " + err.Error())
    }
    //get all ugr
    ugr, err := ndb.GetUserGroupRoles()
    if err != nil {
        logs.Error("node/SyncUserGroupRolesToNode Error getting groups: " + err.Error())
    }
    ugrValues := make(map[string]map[string]string)
    for id := range ugr {
        ugrValues[id] = map[string]string{}
        ugrValues[id]["masterID"] = masterID
        ugrValues[id]["type"] = "master"
        ugrValues[id]["status"] = "exists"
        if ugr[id]["user"] != "" {
            ugrValues[id]["user"] = ugr[id]["user"]
        }
        if ugr[id]["group"] != "" {
            ugrValues[id]["group"] = ugr[id]["group"]
        }
        if ugr[id]["role"] != "" {
            ugrValues[id]["role"] = ugr[id]["role"]
        }
    }

    nodes, err := ndb.GetAllNodes()
    if err != nil {
        logs.Error("node/SyncUserGroupRolesToNode Error getting allNodes: " + err.Error())
    }
    for id := range nodes {
        ipnid, portnid, err := ndb.ObtainPortIp(id)
        if err != nil {
            logs.Error("node/SyncUserGroupRolesToNode Error getting Node ip and port: " + err.Error())
        }

        err = ndb.GetTokenByUuid(id)
        if err != nil {
            logs.Error("node/SyncUserGroupRolesToNode Error loading node token: %s", err)
        }
        //get user uuid
        err = nodeclient.SyncUserGroupRolesToNode(ipnid, portnid, ugrValues)
        if err != nil {
            if err.Error() == "VerifyToken - Incorrect Token" {
                ndb.UpdateNode(id, "token", "wait")
            }
            logs.Error("Error synchronizing users,groups and roles relationship to node '" + nodes[id]["name"] + "': " + err.Error())
        } else {
            logs.Info("Users, groups and roles relationship synchronized to node '" + nodes[id]["name"] + "'")
        }
    }
}

func ChangeRotationStatus(anode map[string]string) (err error) {
    //get node data by uuid
    if ndb.Db == nil {
        logs.Error("node/ChangeRotationStatus -- Can't acces to database")
        return err
    }

    //load token for this node
    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("node/ChangeRotationStatus Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/ChangeRotationStatus ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    //send Suricata services to node
    err = nodeclient.ChangeRotationStatus(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/ChangeRotationStatus ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func EditRotation(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("node/EditRotation -- Can't acces to database")
        return err
    }

    //load token for this node
    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("node/EditRotation Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/EditRotation ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.EditRotation(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/EditRotation ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func GetServiceCommands(anode map[string]string) (values map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("node/GetServiceCommands -- Can't acces to database")
        return nil, err
    }

    //load token for this node
    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("node/GetServiceCommands Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/GetServiceCommands ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    values, err = nodeclient.GetServiceCommands(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/GetServiceCommands ERROR http data request: " + err.Error())
        return nil, err
    }

    return values, nil
}

func SaveSurictaRulesetSelected(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("node/SaveSurictaRulesetSelected -- Can't acces to database")
        return err
    }

    //load token for this node
    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("node/SaveSurictaRulesetSelected Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/SaveSurictaRulesetSelected ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.SaveSurictaRulesetSelected(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/SaveSurictaRulesetSelected ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func SyncRolePermissions() {
    masterID, err := ndb.LoadMasterID()
    if err != nil {
        logs.Error("node/SyncRolePermissions Error getting master ID: " + err.Error())
    }
    //get all roles
    rolePerm, err := ndb.GetRolePermissions()
    if err != nil {
        logs.Error("node/SyncRolePermissions Error getting roles: " + err.Error())
    }
    values := make(map[string]map[string]string)
    for id := range rolePerm {
        values[id] = map[string]string{}
        values[id]["masterID"] = masterID
        values[id]["type"] = "master"
        values[id]["status"] = "exists"
        values[id]["role"] = rolePerm[id]["role"]
        values[id]["permissions"] = rolePerm[id]["permissions"]
        values[id]["object"] = rolePerm[id]["object"]
    }

    nodes, err := ndb.GetAllNodes()
    if err != nil {
        logs.Error("node/SyncRolePermissions Error getting allNodes: " + err.Error())
    }
    for id := range nodes {
        ipnid, portnid, err := ndb.ObtainPortIp(id)
        if err != nil {
            logs.Error("node/SyncRolePermissions Error getting Node ip and port: " + err.Error())
        }

        err = ndb.GetTokenByUuid(id)
        if err != nil {
            logs.Error("node/SyncRolePermissions Error loading node token: %s", err)
        }
        //get user uuid
        err = nodeclient.SyncRolePermissions(ipnid, portnid, values)
        if err != nil {
            if err.Error() == "VerifyToken - Incorrect Token" {
                ndb.UpdateNode(id, "token", "wait")
            }
            logs.Error("Error synchronizing role permissions to node '" + nodes[id]["name"] + "': " + err.Error())
        } else {
            logs.Info("Role permissions synchronized to node '" + nodes[id]["name"] + "'")
        }
    }
}

func SyncRoleGroups() {
    masterID, err := ndb.LoadMasterID()
    if err != nil {
        logs.Error("node/SyncRoleGroups Error getting master ID: " + err.Error())
    }
    //get all roles
    roleGroups, err := ndb.GetRoleGroups()
    if err != nil {
        logs.Error("node/SyncRoleGroups Error getting roles: " + err.Error())
    }
    values := make(map[string]map[string]string)
    for id := range roleGroups {
        values[id] = map[string]string{}
        values[id]["masterID"] = masterID
        values[id]["type"] = "master"
        values[id]["status"] = "exists"
        values[id]["group"] = roleGroups[id]["group"]
        values[id]["desc"] = roleGroups[id]["desc"]
        values[id]["object"] = roleGroups[id]["object"]
    }

    nodes, err := ndb.GetAllNodes()
    if err != nil {
        logs.Error("node/SyncRoleGroups Error getting allNodes: " + err.Error())
    }
    for id := range nodes {
        ipnid, portnid, err := ndb.ObtainPortIp(id)
        if err != nil {
            logs.Error("node/SyncRoleGroups Error getting Node ip and port: " + err.Error())
        }

        err = ndb.GetTokenByUuid(id)
        if err != nil {
            logs.Error("node/SyncRoleGroups Error loading node token: %s", err)
        }
        //get user uuid
        err = nodeclient.SyncRoleGroups(ipnid, portnid, values)
        if err != nil {
            if err.Error() == "VerifyToken - Incorrect Token" {
                ndb.UpdateNode(id, "token", "wait")
            }
            logs.Error("Error synchronizing group roles to node '" + nodes[id]["name"] + "': " + err.Error())
        } else {
            logs.Info("Group roles synchronized to node '" + nodes[id]["name"] + "'")
        }
    }
}

func SyncPermissions() {
    masterID, err := ndb.LoadMasterID()
    if err != nil {
        logs.Error("node/SyncPermissions Error getting master ID: " + err.Error())
    }

    perms, err := ndb.GetPermissions()
    if err != nil {
        logs.Error("node/SyncPermissions Error getting groups: " + err.Error())
    }
    values := make(map[string]map[string]string)
    for id := range perms {
        values[id] = map[string]string{}
        values[id]["masterID"] = masterID
        values[id]["desc"] = perms[id]["desc"]
        values[id]["permisionGroup"] = perms[id]["permisionGroup"]
    }

    nodes, err := ndb.GetAllNodes()
    if err != nil {
        logs.Error("node/SyncPermissions Error getting allNodes: " + err.Error())
    }
    for id := range nodes {
        ipnid, portnid, err := ndb.ObtainPortIp(id)
        if err != nil {
            logs.Error("node/SyncPermissions Error getting Node ip and port: " + err.Error())
        }

        err = ndb.GetTokenByUuid(id)
        if err != nil {
            logs.Error("node/SyncPermissions Error loading node token: %s", err)
        }
        //get user uuid
        err = nodeclient.SyncPermissions(ipnid, portnid, values)
        if err != nil {
            if err.Error() == "VerifyToken - Incorrect Token" {
                ndb.UpdateNode(id, "token", "wait")
            }
            logs.Error("Error synchronizing permissions to node '" + nodes[id]["name"] + "': " + err.Error())
        } else {
            logs.Info("Permissions synchronized to node '" + nodes[id]["name"] + "'")
        }
    }
}

func SyncAllUserData() {
    for {
        t, err := utils.GetKeyValueString("loop", "usergrouproles")
        if err != nil {
            logs.Error("Search Error: Cannot load node information.")
        }
        tDuration, err := strconv.Atoi(t)

        SyncUsersToNode()
        SyncRolesToNode()
        SyncGroupsToNode()
        SyncUserGroupRolesToNode()
        SyncRolePermissions()
        SyncPermissions()
        SyncRoleGroups()
        time.Sleep(time.Minute * time.Duration(tDuration))
    }
}

func EnrollNode(nodeDetail utils.NodeData) (uuid string, enrolled bool, details map[string]string) {
    logs.Info("lets do it - enroll node -> %+v", nodeDetail)
    utils.VerifyField(nodeDetail.IP, "ip")
    utils.VerifyField(nodeDetail.Port, "number")
    utils.VerifyField(nodeDetail.Name, "string", "not null")
    utils.VerifyField(nodeDetail.NodeUser, "not null")
    utils.VerifyField(nodeDetail.NodePassword, "not null")

    var node = map[string]string{}

    node["ip"] = nodeDetail.IP
    node["name"] = nodeDetail.Name
    node["port"] = strconv.Itoa(nodeDetail.Port)
    node["nodeuser"] = nodeDetail.NodeUser
    node["nodepass"] = nodeDetail.NodePassword
    if nodeDetail.Force {
        node["force"] = "true"
    } else {
        node["force"] = "false"
    }

    uuid, err := InsertNode(node)
    if err != nil {
        var ret = map[string]string{}
        ret["enrollNode"] = err.Error()
        return "", false, ret
    }
    return uuid, true, nil
}

func InsertNode(n map[string]string) (uuid string, err error) {

    if n["ip"] == "" {
        return "", errors.New("AddNode - IP should not be an empty field")
    }
    if n["port"] == "" {
        return "", errors.New("AddNode - Port should not be an empty field")
    }
    if n["name"] == "" {
        return "", errors.New("AddNode - Name should not be an empty field")
    }
    if n["nodeuser"] == "" {
        return "", errors.New("AddNode - Node User should not be an empty field")
    }
    if n["nodepass"] == "" {
        return "", errors.New("AddNode - Node Pass should not be an empty field")
    }

    //check if exists a node with the same ip and port
    nodes, err := ndb.GetAllNodes()
    for id := range nodes {
        if nodes[id]["ip"] == n["ip"] {
            if nodes[id]["port"] == n["port"] {
                if n["force"] == "true" {
                    //register anyway and change other node ip
                } else {
                    return "", errors.New("AddNode - There is already a node with the same IP and Port")
                }
            }
        }
    }

    uuid = utils.Generate()
    ndb.InsertNodeKey(uuid, "nodeuser", n["nodeuser"])
    ndb.InsertNodeKey(uuid, "status", "n/a")
    ndb.InsertNodeKey(uuid, "nodepass", n["nodepass"])
    ndb.InsertNodeKey(uuid, "name", n["name"])
    ndb.InsertNodeKey(uuid, "port", n["port"])
    ndb.InsertNodeKey(uuid, "ip", n["ip"])

    masterid, err := ndb.LoadMasterID()
    if err != nil {
        return "", errors.New("InsertNode ERROR can get masterID")
    }

    login := make(map[string]string)
    login["user"] = n["nodeuser"]
    login["pass"] = n["nodepass"]
    login["master"] = masterid

    token, err := nodeclient.GetNodeToken(n["ip"], n["port"], login)
    if err != nil {
        err = ndb.InsertNodeKey(uuid, "token", "wait")

        if err.Error() == "CreateMasterToken Incorrect Login credentials" {
            _ = ndb.InsertPluginCommand(uuid, "status", "Error")
            _ = ndb.InsertPluginCommand(uuid, "output", "Add node error: Invalid login credentials")
            return "", errors.New("AddNode ERROR. Invalid login credentials.")
        }

        _ = ndb.InsertPluginCommand(uuid, "status", "Error")
        _ = ndb.InsertPluginCommand(uuid, "output", "Add node error: Invalid token")

    } else {
        err = ndb.InsertNodeKey(uuid, "token", token)
        //Sync user, group, roles and their relations to the new node
        SyncUsersToNode()
        SyncUserGroupRolesToNode()
        SyncRolesToNode()
        SyncGroupsToNode()
        SyncRolePermissions()
        SyncPermissions()
        SyncRoleGroups()
    }

    nodeValues, err := ndb.GetNodeById(uuid)
    delete(nodeValues[uuid], "nodeuser")
    delete(nodeValues[uuid], "nodepass")
    delete(nodeValues[uuid], "token")
    err = nodeclient.SaveNodeInformation(n["ip"], n["port"], nodeValues)
    if err != nil {
        ndb.InsertPluginCommand(uuid, "status", "Error")
        ndb.InsertPluginCommand(uuid, "output", "Add node error: Error saving node information into node")
        logs.Error("AddNode Error updating node data")
        return "", err
    }

    ndb.InsertPluginCommand(uuid, "status", "Success")
    ndb.InsertPluginCommand(uuid, "output", "Node added successfully")
    return uuid, nil
}

func GetDefaultGroupUUID() (uuid string) {
    return ndb.GetDefaultGroupUUID()
}

func AssignNodeToGroup(uuid string, groupDetail utils.GroupData) (guuid string, assigned bool, details map[string]string) {
    if uuid == "" {
        return "", true, nil
    }
    logs.Info("%+v", groupDetail)
    logs.Info("asign node %s to group %s", uuid, groupDetail.UUID)
    assignTo := ""
    if groupDetail.UUID == "default" {
        logs.Info("default group will be assigned")
        assignTo = GetDefaultGroupUUID()
    } else {
        logs.Info("group %s will be assigned", groupDetail.UUID)
        assignTo = groupDetail.UUID
    }
    logs.Info(assignTo)
    data := make(map[string]interface{})

    data["uuid"] = assignTo
    logs.Info("data uuid set")

    data["nodes"] = []string{uuid}
    logs.Info(data)
    AddGroupNodes(data)

    return assignTo, true, nil
}

type GroupNode struct {
    Uuid  string   `json:"uuid"`
    Nodes []string `json:"nodes"`
}

func AddGroupNodes(data map[string]interface{}) (err error) {
    var nodesList *GroupNode
    nodeExists := false
    LinesOutput, err := json.Marshal(data)
    err = json.Unmarshal(LinesOutput, &nodesList)

    groupNodes, err := ndb.GetAllGroupNodes()
    if err != nil {
        logs.Error("GetAllGroupNodes GetAllGroupNodes error: " + err.Error())
        return err
    }

    for x := range nodesList.Nodes {
        for y := range groupNodes {
            if nodesList.Nodes[x] == groupNodes[y]["nodesid"] && nodesList.Uuid == groupNodes[y]["groupid"] {
                nodeExists = true
            }
        }
        if !nodeExists {
            uuid := utils.Generate()
            err = ndb.InsertGroupNodes(uuid, "groupid", nodesList.Uuid)
            if err != nil {
                logs.Error("AddGroupNodes group uuid error: " + err.Error())
                return err
            }
            err = ndb.InsertGroupNodes(uuid, "nodesid", nodesList.Nodes[x])
            if err != nil {
                logs.Error("AddGroupNodes nodes uuid error: " + err.Error())
                return err
            }
        }
        nodeExists = false
    }
    return nil
}

func EnrollNewNode(anode utils.EnrollNewNodeStruct) (err error) {  
    logs.Notice(anode)     
    nodeExists := false
    //add new node
    newNode := make(map[string]string)
    newNode["ip"] = anode.Node.IP
    newNode["name"] = anode.Node.Name
    newNode["port"] = anode.Node.Port
    newNode["nodeuser"] = anode.Node.NodeUser
    newNode["nodepass"] = anode.Node.NodePass
    
    nodeUUID,err := InsertNode(newNode)
    if err != nil {logs.Error("EnrollNewNode ERROR addingnew node: " + err.Error()); return err}
    //loop groups
    groupNodes, err := ndb.GetAllGroupNodes()
    for x := range anode.Group {
        logs.Debug(anode.Group[x])
        if err != nil {logs.Error("EnrollNewNode GetAllGroupNodes error: " + err.Error()); return err}
        for y := range groupNodes {
            if nodeUUID == groupNodes[y]["nodesid"] && anode.Group[x] == groupNodes[y]["groupid"] {
                nodeExists = true
            }
        }
        if !nodeExists {
            uuid := utils.Generate()
            err = ndb.InsertGroupNodes(uuid, "groupid", anode.Group[x])
            if err != nil {
                logs.Error("AddGroupNodes group uuid error: " + err.Error())
                return err
            }
            err = ndb.InsertGroupNodes(uuid, "nodesid", nodeUUID)
            if err != nil {
                logs.Error("AddGroupNodes nodes uuid error: " + err.Error())
                return err
            }
        }
        nodeExists = false
    }

    //¿suricata?


    return err
}