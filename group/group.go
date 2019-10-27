package group

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "errors"
    "owlhmaster/utils"
    "encoding/json"
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

    return nil
}

func EditGroup(data map[string]string) (err error) { 
    err = ndb.UpdateGroupData(data["uuid"], "name", data["name"]); if err != nil {logs.Error("UpdateGroupData error name: "+ err.Error()); return err}
    err = ndb.UpdateGroupData(data["uuid"], "desc", data["desc"]); if err != nil {logs.Error("UpdateGroupData error desc: "+ err.Error()); return err}
    
    return nil
}

// var Groups []Group
type Group struct{
    Uuid           string      `json:"guuid"`
    Name           string      `json:"gname"`
    Description    string      `json:"gdesc"`
    Nodes          []Node
}
type Node struct {
    Uuid    string      `json:"nuuid"`
    Name    string      `json:"nname"`
    Ip      string      `json:"nip"`
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

        for nid := range groupNodes{
            if gid == groupNodes[nid]["groupid"]{
                nodeValues, err := ndb.GetAllNodesById(groupNodes[nid]["nodesid"]); if err != nil {logs.Error("group/GetNodeValues ERROR getting node values: "+err.Error()); return nil,err}	
                nd := Node{}
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

    logs.Notice(Groups)

    return Groups, nil
}

func GetAllNodesGroup() (data map[string]map[string]string, err error) {
    data, err = ndb.GetAllNodes()
    if err != nil {logs.Error("GetAllNodesGroup error: "+ err.Error()); return nil,err}

    return data, err
}

type GroupNode struct {
	Uuid 	  	string		`json:"uuid"`
	Nodes		[]string	`json:"nodes"`
}

func AddGroupNodes(data map[string]interface{}) (err error) {
    var nodesList *GroupNode

    LinesOutput, err := json.Marshal(data)
    err = json.Unmarshal(LinesOutput, &nodesList)

    
    for x := range nodesList.Nodes{
        uuid := utils.Generate()
        err = ndb.InsertGroupNodes(uuid, "groupid", nodesList.Uuid); if err != nil {logs.Error("AddGroupNodes group uuid error: "+ err.Error()); return err}    
        err = ndb.InsertGroupNodes(uuid, "nodesid", nodesList.Nodes[x]); if err != nil {logs.Error("AddGroupNodes nodes uuid error: "+ err.Error()); return err}    
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