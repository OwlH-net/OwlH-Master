package group

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "errors"
    "owlhmaster/utils"
    "owlhmaster/nodeclient"
    "encoding/json"
    "os"
)

func CreateGroup(n map[string]string) (err error) {
    uuid := utils.Generate()
    if _, ok := n["name"]; !ok {
        logs.Error("name empty: "+err.Error())
        return errors.New("name empty")
    }
    if _, ok := n["desc"]; !ok {
        logs.Error("desc empty: "+err.Error())
        return errors.New("desc empty")
    }

    if err := ndb.GroupExists(uuid); err != nil {logs.Error("Group exist error: "+err.Error()); return err}
    
    for key, value := range n {        
        err = ndb.InsertGroup(uuid, key, value); if err != nil {logs.Error("InsertGroup error: "+ err.Error()); return err}
    }
    
    if err != nil {logs.Error("Error creating group: "+err.Error()); return err}
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
                nodeValues, err := ndb.GetAllNodesById(groupNodes[nid]["nodesid"]); if err != nil {logs.Error("group/GetNodeValues ERROR getting node values: "+err.Error()); return nil,err}    
                nd := Node{}
                nd.Dbuuid = nid
                for b := range nodeValues{
                    nd.Uuid = b
                    nd.Name = nodeValues[b]["name"]
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

type GroupNode struct {
    Uuid           string        `json:"uuid"`
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
    if err != nil {logs.Error("GetAllGroups GetAllGroupNodes error: "+ err.Error()); return nil,err}

    return groups, nil
}

func GetNodeValues(uuid string)(data map[string]map[string]string, err error) {
    data, err = ndb.GetAllNodesById(uuid)
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

func SyncPathGroup(data map[string]string)(err error) {
    fileList := make(map[string]map[string][]byte)    
    filesMap := make(map[string][]byte)
    if data["type"] == "suricata"{
        filesMap,err = utils.ListFilepath(data["mastersuricata"])
        if err != nil {logs.Error("group/ChangePaths ERROR getting Suricata path and files for send to node: "+err.Error()); return err}    
        if fileList[data["nodesuricata"]] == nil { fileList[data["nodesuricata"]] = map[string][]byte{}}
        fileList[data["nodesuricata"]] = filesMap
    }else{
        filesMap,err = utils.ListFilepath(data["masterzeek"])
        if err != nil {logs.Error("group/ChangePaths ERROR getting Zeek path and files for send to node: "+err.Error()); return err}    
        if fileList[data["nodezeek"]] == nil { fileList[data["nodezeek"]] = map[string][]byte{}}
        fileList[data["nodezeek"]] = filesMap
    }

    //get all node uuid
    nodesID,err := ndb.GetGroupNodesByUUID(data["uuid"])
    if err != nil {logs.Error("SyncPathGroup error getting all nodes for a groups: "+err.Error()); return err}

    for uuid := range nodesID{
        //get ip and port for node
        if ndb.Db == nil { logs.Error("SyncPathGroup -- Can't acces to database"); return err}
        ipnid,portnid,err := ndb.ObtainPortIp(nodesID[uuid]["nodesid"])
        if err != nil { logs.Error("node/SyncPathGroup ERROR Obtaining Port and Ip: "+err.Error()); return err}
        
        //send to nodeclient all data
        if data["type"] == "suricata"{
            err = nodeclient.ChangeSuricataPathsGroups(ipnid,portnid,fileList)
            if err != nil { logs.Error("node/SyncPathGroup ChangeSuricataPathsGroups ERROR http data request: "+err.Error()); return err}    
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