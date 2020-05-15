package group

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "errors"
    "owlhmaster/utils"
    "owlhmaster/node"
    "owlhmaster/nodeclient"
    "encoding/json"    
    "path"
    "os"
    "io"
    "bytes"
    "strings"
    "bufio"
    "io/ioutil"
)

func AddGroupElement(n map[string]string) (err error) {
    uuid := utils.Generate()
    // if _, ok := n["name"]; !ok {
    //     logs.Error("name empty: "+err.Error())
    //     return errors.New("name empty")
    // }
    // if _, ok := n["desc"]; !ok {
    //     logs.Error("desc empty: "+err.Error())
    //     return errors.New("desc empty")
    // }

    if err := ndb.GroupExists(uuid); err != nil {logs.Error("Group uuid exist error: "+err.Error()); return err}
    
    for key, value := range n {        
        err = ndb.InsertGroup(uuid, key, value); if err != nil {logs.Error("AddGroupElement error: "+ err.Error()); return err}
    }
    
    if err != nil {logs.Error("AddGroupElement Error adding element to group table: "+err.Error()); return err}
    return nil
}

func DeleteGroup(uuid string) (err error) {
    err = ndb.DeleteGroup(uuid)
    if err != nil {logs.Error("DeleteGroup error: "+ err.Error()); return err}
    
    groupnodes,err := ndb.GetGroupNodesByValue(uuid)
    if err != nil {logs.Error("DeleteGroup error groupnodes: "+ err.Error()); return err}

    for x := range groupnodes{
        err = ndb.DeleteNodeGroupById(x)
        if err != nil {logs.Error("DeleteGroup error for uuid: "+x+": "+ err.Error()); return err}
    }

    //delete groupruleset database
    //grouprulesets
    allGR,err := ndb.GetAllGroupRulesets()
    if err != nil {logs.Error("DeleteGroup error getting all groupruleset data:" + err.Error()); return err}
    
    for x := range allGR {
        if allGR[x]["groupid"] == uuid {
            err = ndb.DeleteGroupRuleset(x)
            if err != nil {logs.Error("DeleteGroup error deleting groupruleset data: "+ err.Error()); return err}
        }        
    }

    return nil
}

func EditGroup(data map[string]string) (err error) { 
    err = ndb.UpdateGroupValue(data["uuid"], "name", data["name"]); if err != nil {logs.Error("UpdateGroupValue error name: "+ err.Error()); return err}
    err = ndb.UpdateGroupValue(data["uuid"], "desc", data["desc"]); if err != nil {logs.Error("UpdateGroupValue error desc: "+ err.Error()); return err}
    
    return nil
}

// var Groups []Group
type Group struct{
    Uuid             string      `json:"guuid"`
    Name             string      `json:"gname"`
    Ruleset          string      `json:"gruleset"`
    RulesetID        string      `json:"grulesetID"`
    Description      string      `json:"gdesc"`
    MasterSuricata   string      `json:"mastersuricata"`
    NodeSuricata     string      `json:"nodesuricata"`
    MasterZeek       string      `json:"masterzeek"`
    NodeZeek         string      `json:"nodezeek"`
    Interface        string      `json:"interface"`
    BPFFile          string      `json:"BPFfile"`
    BPFRule          string      `json:"BPFrule"`
    ConfigFile       string      `json:"configFile"`
    CommandLine      string      `json:"commandLine"`
    Nodes            []Node
}
type Node struct {
    Dbuuid           string      `json:"dbuuid"`
    Uuid             string      `json:"nuuid"`
    Name             string      `json:"nname"`
    Ip               string      `json:"nip"`
    Status           string      `json:"nstatus"`
}

func GetAllGroups()(Groups []Group, err error){
    allGroups, err := ndb.GetAllGroups()
    if err != nil {logs.Error("GetAllGroups error: "+ err.Error()); return nil,err}

    groupNodes, err := ndb.GetAllGroupNodes()
    if err != nil {logs.Error("group/GetNodeValues ERROR getting all nodes: "+err.Error()); return nil,err}    

    for gid := range allGroups{
        gr := Group{}
        gr.Uuid = gid
        gr.Name = allGroups[gid]["name"]
        gr.Description = allGroups[gid]["desc"]
        gr.Ruleset = allGroups[gid]["ruleset"]
        gr.RulesetID = allGroups[gid]["rulesetID"]
        gr.MasterSuricata = allGroups[gid]["mastersuricata"]
        gr.NodeSuricata = allGroups[gid]["nodesuricata"]
        gr.MasterZeek = allGroups[gid]["masterzeek"]
        gr.NodeZeek = allGroups[gid]["nodezeek"]
        gr.Interface = allGroups[gid]["interface"]
        gr.BPFFile = allGroups[gid]["BPFfile"]
        gr.BPFRule = allGroups[gid]["BPFrule"]
        gr.ConfigFile = allGroups[gid]["configFile"]
        gr.CommandLine = allGroups[gid]["commandLine"]

        for nid := range groupNodes{
            if gid == groupNodes[nid]["groupid"]{
                nodeValues, err := ndb.GetNodeById(groupNodes[nid]["nodesid"]); if err != nil {logs.Error("group/GetNodeValues ERROR getting node values: "+err.Error()); return nil,err}    
                nd := Node{}
                nd.Dbuuid = nid
                for b := range nodeValues{
                    //ping analyzer
                    analyzerData,err := node.PingAnalyzer(b)
                    if err != nil {
                        logs.Error("group/PingAnalyzer ERROR getting all nodes analyzer : "+err.Error())
                        nd.Status = "N/A"
                    }else{
                        nd.Status = analyzerData["status"]
                    }    

                    nd.Uuid = b
                    nd.Name = nodeValues[b]["name"]
                    // nd.Sync = nodeValues[b]["sync"]                    
                    nd.Ip = nodeValues[b]["ip"]
                    gr.Nodes = append(gr.Nodes, nd)
                }                  
            }
        }
        Groups = append(Groups, gr)
    }

    return Groups, nil
}

func GetAllNodesGroup(uuid string)(data map[string]map[string]string, err error) {
    data, err = ndb.GetAllNodes()
    if err != nil {logs.Error("GetAllNodesGroup GetAllNodes error: "+ err.Error()); return nil,err}

    groupNodes, err := ndb.GetAllGroupNodes()
    if err != nil {logs.Error("GetAllNodesGroup GetAllGroupNodes error: "+ err.Error()); return nil,err}

    for y := range groupNodes {
        if groupNodes[y]["groupid"] == uuid {
            data[groupNodes[y]["nodesid"]]["checked"] = "true"
        }
    }

    return data, err
}

func GetGroupSelectedRulesets(uuid string)(data map[string]map[string]string, err error) {
    var localRulesets = map[string]map[string]string{}
    rulesets,err := ndb.GetAllRulesets()
    if err != nil {logs.Error("GetGroupSelectedRulesets GetAllRulesets error: "+ err.Error()); return nil,err}

    for x := range rulesets{
        if rulesets[x]["type"] == "local"{
            if localRulesets[x] == nil { localRulesets[x] = map[string]string{}}
            localRulesets[x]["name"] = rulesets[x]["name"]
            localRulesets[x]["desc"] = rulesets[x]["desc"]
            localRulesets[x]["type"] = rulesets[x]["type"]
        }
    }

    groupRulesets, err := ndb.GetAllGroupRulesets()
    if err != nil {logs.Error("GetAllNodesGroup GetAllGroupRulesets error: "+ err.Error()); return nil,err}

    for y := range groupRulesets {
        if groupRulesets[y]["groupid"] == uuid {
            localRulesets[groupRulesets[y]["rulesetid"]]["checked"] = "true"
        }
    }

    return localRulesets, nil
}

func DeleteExpertGroupRuleset(data map[string]string)(err error) {
    groupRulesets, err := ndb.GetAllGroupRulesets()
    if err != nil {logs.Error("DeleteExpertGroupRuleset GetAllGroupRulesets error: "+ err.Error()); return err}
    
    for x := range groupRulesets {
        if groupRulesets[x]["groupid"] == data["uuid"] && groupRulesets[x]["rulesetid"] == data["ruleset"] {
            err = ndb.DeleteGroupRuleset(x)
            if err != nil {logs.Error("DeleteExpertGroupRuleset error deleting database value: "+ err.Error()); return  err}
        }
    }

    return nil
}

func AddRulesetsToGroup(data map[string]string)(err error) {
    //split array
    groupRulesets, err := ndb.GetAllGroupRulesets()
    if err != nil {logs.Error("DeleteExpertGroupRuleset GetAllGroupRulesets error: "+ err.Error()); return err}

    
    rulesets := strings.Split(data["rulesets"], ",")
    
    for x := range rulesets {
        exists := false
        for y := range groupRulesets{
            if groupRulesets[y]["groupid"] == data["uuid"] && groupRulesets[y]["rulesetid"] == rulesets[x] {
                exists = true                
            } 
        }
        if !exists{
            uuid := utils.Generate()
            err = ndb.InsertGroupRulesets(uuid, "groupid", data["uuid"])
            if err != nil {logs.Error("AddRulesetsToGroup error inserting data into grouprulesets: "+ err.Error()); return err}
            err = ndb.InsertGroupRulesets(uuid, "rulesetid", rulesets[x])
            if err != nil {logs.Error("AddRulesetsToGroup error inserting data into grouprulesets: "+ err.Error()); return err}
        }
    }

    return nil
}

type GroupNode struct {
    Uuid           string     `json:"uuid"`
    Nodes        []string    `json:"nodes"`
}

func AddGroupNodes(data map[string]interface{}) (err error) {
    var nodesList *GroupNode
    nodeExists := false;
    LinesOutput, err := json.Marshal(data)
    err = json.Unmarshal(LinesOutput, &nodesList)

    groupNodes, err := ndb.GetAllGroupNodes()
    if err != nil {logs.Error("GetAllGroupNodes GetAllGroupNodes error: "+ err.Error()); return err}

    for x := range nodesList.Nodes{           
        for y := range groupNodes{
            if nodesList.Nodes[x] == groupNodes[y]["nodesid"] && nodesList.Uuid == groupNodes[y]["groupid"]{
                nodeExists = true;
            }
        }
        if !nodeExists{
            uuid := utils.Generate()
            err = ndb.InsertGroupNodes(uuid, "groupid", nodesList.Uuid); if err != nil {logs.Error("AddGroupNodes group uuid error: "+ err.Error()); return err}    
            err = ndb.InsertGroupNodes(uuid, "nodesid", nodesList.Nodes[x]); if err != nil {logs.Error("AddGroupNodes nodes uuid error: "+ err.Error()); return err}    
        }
        nodeExists = false;
    }
    return nil
}

func PingGroupNodes()(data map[string]map[string]string, err error) {
    groups, err := ndb.GetAllGroupNodes()
    if err != nil {logs.Error("PingGroupNodes GetAllGroupNodes error: "+ err.Error()); return nil,err}

    return groups, nil
}

func GetNodeValues(uuid string)(data map[string]map[string]string, err error) {
    data, err = ndb.GetNodeById(uuid)
    if err != nil {logs.Error("group/GetNodeValues ERROR getting node data: "+err.Error()); return nil,err}    

    return data, nil
}

func DeleteNodeGroup(uuid string)(err error) {
    err = ndb.DeleteNodeGroupById(uuid)
    if err != nil {logs.Error("group/GetNodeValues ERROR getting node data: "+err.Error()); return err}    

    return nil
}

func ChangeGroupRuleset(data map[string]string)(err error) {
    err = ndb.UpdateGroupValue(data["uuid"], "ruleset", data["ruleset"])
    if err != nil {logs.Error("group/ChangeGroupRuleset ERROR updating group data for ruleset: "+err.Error()); return err}    

    err = ndb.UpdateGroupValue(data["uuid"], "rulesetID", data["rulesetID"])
    if err != nil {logs.Error("group/ChangeGroupRuleset ERROR updating group data for rulesetID: "+err.Error()); return err}    

    return err
}

func ChangePathsGroups(data map[string]string)(err error) {
    if data["type"] == "suricata"{
        if _, err := os.Stat(data["mastersuricata"]); os.IsNotExist(err) {logs.Error("Suricata master path doesn't exists: "+err.Error()); return errors.New("Suricata master path doesn't exists: "+err.Error())}
        err = ndb.UpdateGroupValue(data["uuid"], "mastersuricata", data["mastersuricata"])
        if err != nil {logs.Error("group/ChangePaths ERROR updating suricata master path: "+err.Error()); return err}    
        
        err = ndb.UpdateGroupValue(data["uuid"], "nodesuricata", data["nodesuricata"])
        if err != nil {logs.Error("group/ChangePaths ERROR updating suricata node path: "+err.Error()); return err}        
    }else{
        if _, err := os.Stat(data["masterzeek"]); os.IsNotExist(err) {logs.Error("Zeek node path doesn't exists: "+err.Error()); return errors.New("Zeek master path doesn't exists: "+err.Error())}
        err = ndb.UpdateGroupValue(data["uuid"], "masterzeek", data["masterzeek"])
        if err != nil {logs.Error("group/ChangePaths ERROR updating zeek master path: "+err.Error()); return err}    

        err = ndb.UpdateGroupValue(data["uuid"], "nodezeek", data["nodezeek"])
        if err != nil {logs.Error("group/ChangePaths ERROR updating zeek node path: "+err.Error()); return err}    
    }    

    return err
}

func GetMD5files(data map[string]string)(allData map[string]map[string]map[string]string, err error) {
    fileList := map[string]map[string]map[string]string{}
    
    //get all node uuid
    nodesID,err := ndb.GetGroupNodesByUUID(data["uuid"])
    if err != nil {logs.Error("GetMD5files error getting all nodes for a groups: "+err.Error()); return nil,err}

    //get all files synchronized
    hashedFiles,err := utils.FolderMapMD5(data["mastersuricata"], data["nodesuricata"])    
    if err != nil { return nil, err }

    for uuid := range nodesID{
        //get ip and port for node
        if ndb.Db == nil { logs.Error("SyncPathGroup -- Can't acces to database"); return nil,err}
        err = ndb.GetTokenByUuid(nodesID[uuid]["nodesid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
        ipnid,portnid,err := ndb.ObtainPortIp(nodesID[uuid]["nodesid"])
        if err != nil { logs.Error("node/SyncPathGroup ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}
        
        //send to nodeclient all data
        valuesMD5, err := nodeclient.GetMD5files(ipnid,portnid,hashedFiles)
        if err != nil { logs.Error("node/SyncPathGroup ChangeSuricataPathsGroups ERROR http data request: "+err.Error()); return nil,err}    
        
        //check if files in Master and Node has the same MD5 for every file
        filesHashed := utils.CompareFolderMapMD5(hashedFiles,valuesMD5)

        if fileList[nodesID[uuid]["nodesid"]] == nil { fileList[nodesID[uuid]["nodesid"]] = map[string]map[string]string{} }
        for x := range filesHashed {
            returnUUID := utils.Generate()
            if fileList[nodesID[uuid]["nodesid"]][returnUUID] == nil { fileList[nodesID[uuid]["nodesid"]][returnUUID] = map[string]string{} }
            fileList[nodesID[uuid]["nodesid"]][returnUUID]["masterPath"] = hashedFiles[x]["path"]
            fileList[nodesID[uuid]["nodesid"]][returnUUID]["nodePath"] = data["nodesuricata"]+"/"+filesHashed[x]["path"]
            fileList[nodesID[uuid]["nodesid"]][returnUUID]["nodeMD5"] = filesHashed[x]["md5"]
            fileList[nodesID[uuid]["nodesid"]][returnUUID]["masterMD5"] = hashedFiles[x]["md5"]
            fileList[nodesID[uuid]["nodesid"]][returnUUID]["equals"] = filesHashed[x]["equals"]
        }
        
    }   

    return fileList,nil
}

func SyncPathGroup(data map[string]string)(err error) {
    fileList := make(map[string]map[string][]byte)    
    filesMap := make(map[string][]byte)
    if data["type"] == "suricata"{
        var buf bytes.Buffer
        err = utils.Compress(data["mastersuricata"], &buf)
        if err != nil {logs.Error("group/SyncPathGroup ERROR compressing data: "+err.Error()); return err}
    
        // write the .tar.gzip
        fileToWrite, err := os.OpenFile("/tmp/file.tar.gzip", os.O_CREATE|os.O_RDWR, os.FileMode(600))
        if err != nil {logs.Error("group/SyncPathGroup ERROR creating tar file: "+err.Error()); return err}
    
        if _, err := io.Copy(fileToWrite, &buf); err != nil {
            logs.Error("group/SyncPathGroup ERROR copying data to tar file: "+err.Error()); return err
        }

        r,_ := os.Open("/tmp/file.tar.gzip")
        bytesFromTar, err := ioutil.ReadAll(r)
        
        filesMap[data["nodesuricata"]] = bytesFromTar
    }else{
        filesMap,err = utils.ListFilepath(data["masterzeek"])
        if err != nil {logs.Error("group/SyncPathGroup ERROR getting Zeek path and files for send to node: "+err.Error()); return err}    
        if fileList[data["nodezeek"]] == nil { fileList[data["nodezeek"]] = map[string][]byte{}}
        fileList[data["nodezeek"]] = filesMap
    }

    //get all node uuid
    nodesID,err := ndb.GetGroupNodesByUUID(data["uuid"])
    if err != nil {logs.Error("SyncPathGroup error getting all nodes for a groups: "+err.Error()); return err}

    for uuid := range nodesID{
        //get ip and port for node
        if ndb.Db == nil { logs.Error("SyncPathGroup -- Can't acces to database"); return err}
        err = ndb.GetTokenByUuid(nodesID[uuid]["nodesid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
        ipnid,portnid,err := ndb.ObtainPortIp(nodesID[uuid]["nodesid"])
        if err != nil { logs.Error("node/SyncPathGroup ERROR Obtaining Port and Ip: "+err.Error()); return err}
        
        //send to nodeclient all data
        if data["type"] == "suricata"{            
            err = nodeclient.ChangeSuricataPathsGroups(ipnid,portnid,filesMap)
            if err != nil { logs.Error("node/SyncPathGroup ChangeSuricataPathsGroups ERROR http data request: "+err.Error()); return err}    
            //remove tar.gz file
            os.Remove("/tmp/file.tar.gzip")
        }else{
            err = nodeclient.ChangeZeekPathsGroups(ipnid,portnid,fileList)
            if err != nil { logs.Error("node/SyncPathGroup ChangeZeekPathsGroups ERROR http data request: "+err.Error()); return err}    
        }
    }   

    return nil
}

func UpdateGroupService(data map[string]string)(err error) {
    err = ndb.UpdateGroupValue(data["uuid"], data["param"], data["value"])
    if err != nil {logs.Error("group/UpdateGroupService ERROR updating group data: "+err.Error()); return err}    

    return nil
}

func SyncAll(uuid string)(err error) {
    anode := make(map[string]string)
    anode["uuid"] = uuid
    //sync analyzer
    _,err = node.SyncAnalyzerToAllGroupNodes(anode)
    if err != nil {logs.Error("group/SyncAll ERROR synchronizing analyzer to all group nodes: "+err.Error()); return err}    
    //sync expert Rulesets
    err = node.SyncRulesetToAllGroupNodes(anode)
    if err != nil {logs.Error("group/SyncAll ERROR synchronizing all group rulesets to all group nodes: "+err.Error()); return err}    


    //sync expert configs



    // for key, _ := range data{
    //     filesMap := make(map[string]string)
    //     if key == "suricata-services"{
    //         filesMap["uuid"] = uuid
    //         filesMap["interface"] = data[key]["interface"]
    //         filesMap["BPFfile"] = data[key]["BPFfile"]
    //         filesMap["BPFrule"] = data[key]["BPFrule"]
    //         filesMap["configFile"] = data[key]["configFile"]
    //         filesMap["commandLine"] = data[key]["commandLine"]
    //         logs.Info("Synchronizing Suricata services")
    //         err = node.PutSuricataServicesFromGroup(filesMap)
    //         if err!=nil {logs.Error("group/SyncAll -- Error synchronizing Suricata services to node: "+err.Error()); return err}
    //         logs.Notice("Suricata services synchronized")
    //     // }else if key == "zeek-policies"{
    //     //     filesMap["uuid"] = uuid
    //     //     filesMap["type"] = data[key]["type"]
    //     //     filesMap["masterzeek"] = data[key]["masterzeek"]
    //     //     filesMap["nodezeek"] = data[key]["nodezeek"]
    //     //     logs.Info("Synchronizing Zeek policies")
    //     //     err = SyncPathGroup(filesMap)
    //     //     if err!=nil {logs.Error("group/SyncAll -- Error synchronizing Zeek configuration: "+err.Error()); return err}
    //     //     logs.Notice("Zeek policies synchronized")
    //     }else if key == "suricata-rulesets"{
    //         filesMap["uuid"] = uuid
    //         logs.Info("Synchronizing Suricata ruleset")
    //         err = node.SyncRulesetToAllGroupNodes(filesMap)
    //         if err!=nil {logs.Error("group/SyncAll -- Error synchronizing ruleset: "+err.Error()); return err}
    //         logs.Notice("Suricata ruleset synchronized")
    //     }else if key == "suricata-config"{
    //         filesMap["uuid"] = uuid
    //         filesMap["type"] = data[key]["type"]
    //         filesMap["mastersuricata"] = data[key]["mastersuricata"]
    //         filesMap["nodesuricata"] = data[key]["nodesuricata"]
    //         logs.Info("Synchronizing Suricata configuration")
    //         err = SyncPathGroup(filesMap)
    //         if err!=nil {logs.Error("group/SyncAll -- Error synchronizing Suricata configuration: "+err.Error()); return err}
    //         logs.Notice("Suricata configuration synchronized")
    //     }
    // }

    return nil
}

func AddCluster(anode map[string]string)(err error) {
    conf, err := utils.GetKeyValueString("group", "conf")
    if err != nil {logs.Error("AddCluster error getting path from main.conf"); return err}

    //check if exists path
    if _, err := os.Stat(anode["path"]); os.IsNotExist(err) {
        logs.Warn("Cluster path doesn't exists. Creating..."); 
        
        err = os.MkdirAll(path.Dir(anode["path"]), os.ModePerm)
        if err != nil{logs.Error("Error creating cluster path: "+err.Error()); return err }                

        _, err := os.Create(anode["path"])
        if err != nil{logs.Error("Error creating cluster file: "+err.Error()); return err }                
    }
    //read file data
    data, err := ioutil.ReadFile(conf)
    if err != nil {logs.Error("Error opening node.cfg file: "+err.Error()); return err }        
    
    //write file data
    err = ioutil.WriteFile(anode["path"], data, 0644)
    if err != nil {logs.Error("Error writting node.cfg data into new file: "+err.Error()); return err }        
    
    //add to db
    uuid := utils.Generate()
    err = ndb.InsertCluster(uuid, "guuid", anode["uuid"]); if err != nil {logs.Error("Error adding cluster guuid to database: "+err.Error()); return err }        
    err = ndb.InsertCluster(uuid, "path", anode["path"]); if err != nil {logs.Error("Error adding cluster path to database: "+err.Error()); return err }        

    return nil
}

func GetClusterFiles(uuid string)(data map[string]map[string]string, err error) {
    data, err = ndb.GetClusterByValue(uuid)
    if err != nil {logs.Error("Error GetClusterFiles: "+err.Error()); return nil,err }        
    return data, err
}

func DeleteCluster(data map[string]string)(err error) {
    err = ndb.DeleteCluster(data["uuid"])
    if err != nil {logs.Error("Error DeleteCluster: "+err.Error()); return err }        
    return nil
}

func ChangeClusterValue(anode map[string]string)(err error) {
    //check if exists path
    if _, err := os.Stat(anode["path"]); os.IsNotExist(err) {
        logs.Error("New cluster path doesn't exists: "+err.Error()); 
        return err             
    }
    err = ndb.UpdateGroupClusterValue(anode["uuid"], "path", anode["path"])
    if err != nil {logs.Error("Error ChangeClusterValue: "+err.Error()); return err }        
    return err
}

func GetClusterFileContent(uuid string)(values map[string]string, err error) {
    fileContent := make(map[string]string)
    data, err := ndb.GetClusterByUUID(uuid)
    
    for x := range data{
        //read file data
        content, err := ioutil.ReadFile(data[x]["path"])
        if err != nil {logs.Error("GetClusterFileContent Error opening file: "+err.Error()); return nil, err }
        fileContent[uuid] = string(content)
    }

    return fileContent, err
}

func SaveClusterFileContent(anode map[string]string)(err error) {
    //check if exists path
    if _, err := os.Stat(anode["path"]); os.IsNotExist(err) {
        logs.Error("SaveClusterFileContent path doesn't exists: "+err.Error()); 
        return err             
    }
    //write file data
    err = ioutil.WriteFile(anode["path"], []byte(anode["content"]), 0644)
    if err != nil {logs.Error("SaveClusterFileContent Error writting data into file: "+err.Error()); return err }

    return nil
}

func SyncClusterFile(data map[string]string)(err error) {
    var managerIP string
    //get file for this uuid
    cluster,err := ndb.GetClusterByUUID(data["uuid"])
    if err != nil {logs.Error("SyncClusterFile Error getting data: "+err.Error()); return err }
    for x := range cluster{
        var managerFound = false
        file, err := os.Open(cluster[x]["path"])
        if err != nil {logs.Error("SyncClusterFile Error opening path: "+err.Error()); return err }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            if scanner.Text() == "[manager]"{
                managerFound = true
            }
            if managerFound{
                if strings.Contains(scanner.Text(), "host="){
                    managerIP = strings.Trim(scanner.Text(), "host=")
                    break
                }
            }
        }

        //get uuid by manger ip
        nodes,err := ndb.GetAllNodes()
        for z := range nodes{
            if nodes[z]["ip"] == managerIP{
                // read file content
                content, err := ioutil.ReadFile(cluster[x]["path"])
                if err != nil {logs.Error("SyncClusterFile Error opening file: "+err.Error()); return err }
                //get ip and port for node
                err = ndb.GetTokenByUuid(z); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
                if ndb.Db == nil { logs.Error("SyncClusterFile -- Can't acces to database"); return err}
                ipnid,portnid,err := ndb.ObtainPortIp(z)
                if err != nil { logs.Error("node/SyncClusterFile ERROR Obtaining Port and Ip: "+err.Error())}            

                err = nodeclient.SyncClusterFileNode(ipnid,portnid,content)
                if err != nil { logs.Error("node/SyncClusterFile ERROR http data request: "+err.Error()); return err}
            }
        }
    }

    return nil
}

func SyncAllGroupCluster(data map[string]string)(err error) {
    groupnodes,err := ndb.GetAllCluster()
    if err != nil { logs.Error("node/SyncAllGroupCluster ERROR getting all group nodes: "+err.Error()); return err}
    for x := range groupnodes{
        if groupnodes[x]["guuid"] == data["uuid"]{
            anode := make(map[string]string)
            anode["uuid"] = x
            err = SyncClusterFile(anode)
            if err != nil { logs.Error("node/SyncAllGroupCluster ERROR SyncClusterFile: "+err.Error()); return err}
        }
    }
    
    return nil
}

func SyncAllSuricataGroup(data map[string]string)(err error) {
    groups,err := ndb.GetAllGroups()
    if err != nil { logs.Error("group/SyncAllSuricataGroup ERROR getting all group nodes: "+err.Error()); return err}
    gr,err := ndb.GetGroupNodesByUUID(data["uuid"])
    if err != nil { logs.Error("group/SyncAllSuricataGroup ERROR getting a group by id: "+err.Error()); return err}

    for x,y := range groups{
        anode := make(map[string]string)
        if x == data["uuid"]{
            for y := range y{
                anode[y] = groups[x][y]
            }
            for w := range gr{
                //get ip and port for node
                if ndb.Db == nil { logs.Error("group/SyncAllSuricataGroup -- Can't access to database"); return err}
                err = ndb.GetTokenByUuid(gr[w]["nodesid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
                ipnid,portnid,err := ndb.ObtainPortIp(gr[w]["nodesid"])
                if err != nil { logs.Error("group/SyncAllSuricataGroup ERROR Obtaining Port and Ip: "+err.Error()); return err}
                
                //send to nodeclient all data
                err = nodeclient.SyncAllSuricataGroup(ipnid,portnid,anode)
                if err != nil { logs.Error("group/SyncAllSuricataGroup ERROR http data request: "+err.Error()); return err}    
        
            }
        }
    }

    return nil
}

func SuricataGroupService(data map[string]string)(err error) {
    gr,err := ndb.GetGroupNodesByUUID(data["uuid"])
    if err != nil { logs.Error("group/SuricataGroupService ERROR getting a group by id: "+err.Error()); return err}

    for w := range gr{
        //get ip and port for node
        if ndb.Db == nil { logs.Error("group/SuricataGroupService -- Can't access to database"); return err}
        err = ndb.GetTokenByUuid(gr[w]["nodesid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
        ipnid,portnid,err := ndb.ObtainPortIp(gr[w]["nodesid"])
        if err != nil { logs.Error("group/SuricataGroupService ERROR Obtaining Port and Ip: "+err.Error()); return err}
        
        //send to nodeclient all data
        err = nodeclient.SuricataGroupService(ipnid,portnid,data)
        if err != nil { logs.Error("group/SuricataGroupService ERROR http data request: "+err.Error()); return err}    

    }

    return nil
}


func SuricataNodesStatus(uuid string)(data map[string]map[string]string, err error) {
    nodePlugins := map[string]map[string]string{}

    gr,err := ndb.GetGroupNodesByUUID(uuid)
    if err != nil { logs.Error("group/SuricataNodesStatus ERROR getting a group by id: "+err.Error()); return nil,err}

    for w := range gr{
        //get ip and port for node
        if ndb.Db == nil { logs.Error("group/SuricataNodesStatus -- Can't access to database"); return nil,err}
        err = ndb.GetTokenByUuid(gr[w]["nodesid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
        ipnid,portnid,err := ndb.ObtainPortIp(gr[w]["nodesid"])
        if err != nil { logs.Error("group/SuricataNodesStatus ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}
        
        //send to nodeclient all data
        plugins,err := nodeclient.PingPluginsNode(ipnid,portnid)
        if err != nil { logs.Error("group/SuricataNodesStatus ERROR http data request: "+err.Error()); return nil,err}    

        for x,y := range plugins{
            nodePlugins[x] = map[string]string{}
            for y := range y{
                nodePlugins[x][y] = plugins[x][y]
                nodePlugins[x]["ip"] = ipnid
                nodePlugins[x]["port"] = portnid
                nodeName,err := ndb.ObtainNodeName(gr[w]["nodesid"])
                if err != nil { logs.Error("group/SuricataNodesStatus ERROR getting node name: "+err.Error()); return nil,err}   
                nodePlugins[x]["nodeName"] = nodeName
            }
        }
    }
    
    return nodePlugins, err
}