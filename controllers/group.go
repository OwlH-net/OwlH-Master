package controllers

import (
    "owlhmaster/models"
    "encoding/json"
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego"
    "owlhmaster/validation"
)

type GroupController struct {
    beego.Controller
}

// @Title AddGroupElement
// @Description Add new group element
// @Success 200 {object} models.Group
// @Failure 403 body is empty
// @router / [post]
func (n *GroupController) AddGroupElement() {
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "post")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddGroupElement(anode)
        n.Data["json"] = map[string]string{"ack": "true!"}
        if err != nil {
            logs.Error("GROUP ADD ELEMENT -> error: %s", err.Error())
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title GetAllGroups
// @Description Get full list of groups
// @Success 200 {object} models.Groups
// @router / [get]
func (n *GroupController) GetAllGroups() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        groups, err := models.GetAllGroups()
        n.Data["json"] = groups
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteGroup
// @Description delete a group
// @Success 200 {object} models.Groups
// @router /DeleteGroup/:uuid [delete]
func (n *GroupController) DeleteGroup() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        uuid := n.GetString(":uuid") 
        err := models.DeleteGroup(uuid)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title EditGroup
// @Description edit a group information
// @Success 200 {object} models.Groups
// @router /editGroup [put]
func (n *GroupController) EditGroup() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.EditGroup(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetAllNodesGroup
// @Description Get full list of nodes
// @Success 200 {object} models.Groups
// @router /getAllNodesGroup/:uuid [get]
func (n *GroupController) GetAllNodesGroup() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        uuid := n.GetString(":uuid") 
        groups, err := models.GetAllNodesGroup(uuid)
        n.Data["json"] = groups
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title AddGroupNodes
// @Description add nodes to a group
// @Success 200 {object} models.Groups
// @router /addGroupNodes [put]
func (n *GroupController) AddGroupNodes() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]interface{}
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddGroupNodes(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}   

// @Title PingGroupNodes
// @Description Ping all group nodes
// @Success 200 {object} models.Groups
// @router /pingGroupNodes [get]
func (n *GroupController) PingGroupNodes() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        groups, err := models.PingGroupNodes()
        n.Data["json"] = groups
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetNodeValues
// @Description Get values for specific node
// @Success 200 {object} models.Groups
// @router /getNodeValues/:uuid [get]
func (n *GroupController) GetNodeValues() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        uuid := n.GetString(":uuid") 
        groups, err := models.GetNodeValues(uuid)
        n.Data["json"] = groups
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteNodeGroup
// @Description add nodes to a group
// @Success 200 {object} models.Groups
// @router /deleteNodeGroup/:uuid [delete]
func (n *GroupController) DeleteNodeGroup() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        uuid := n.GetString(":uuid") 
        err := models.DeleteNodeGroup(uuid)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
} 

// @Title ChangeGroupRuleset
// @Description Change group ruleset
// @Success 200 {object} models.Groups
// @router /changeGroupRuleset [put]
func (n *GroupController) ChangeGroupRuleset() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeGroupRuleset(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ChangePathsGroups
// @Description Change group paths
// @Success 200 {object} models.Groups
// @router /changePaths [put]
func (n *GroupController) ChangePathsGroups() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangePathsGroups(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SyncPathGroup
// @Description Change group paths
// @Success 200 {object} models.Groups
// @router /syncPathGroup [post]
func (n *GroupController) SyncPathGroup() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "post")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SyncPathGroup(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title UpdateGroupService
// @Description Update gorup service value
// @Success 200 {object} models.Groups
// @router /updateGroupService [put]
func (n *GroupController) UpdateGroupService() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.UpdateGroupService(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SyncAll
// @Description Synchronize all group elements
// @Success 200 {object} models.Groups
// @router /syncAll/:uuid [put]
func (n *GroupController) SyncAll() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        uuid := n.GetString(":uuid") 
        var anode map[string]map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    
        err := models.SyncAll(uuid, anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title AddCluster
// @Description Add cluster for Zeek group
// @Success 200 {object} models.Groups
// @router /addCluster [post]
func (n *GroupController) AddCluster() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "post")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddCluster(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetClusterFiles
// @Description Get all cluster elemenst for Zeek group
// @Success 200 {object} models.Groups
// @router /getClusterFiles/:uuid [get]
func (n *GroupController) GetClusterFiles() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        uuid := n.GetString(":uuid") 
    
        data,err := models.GetClusterFiles(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteCluster
// @Description delete cluster for Zeek group
// @Success 200 {object} models.Groups
// @router /deleteCluster [delete]
func (n *GroupController) DeleteCluster() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteCluster(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ChangeClusterValue
// @Description delete cluster for Zeek group
// @Success 200 {object} models.Groups
// @router /changeClusterValue [put]
func (n *GroupController) ChangeClusterValue() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeClusterValue(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetClusterFileContent
// @Description Get cluster file content from Zeek group
// @Success 200 {object} models.Groups
// @router /getClusterFileContent/:uuid [get]
func (n *GroupController) GetClusterFileContent() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        uuid := n.GetString(":uuid") 
    
        data,err := models.GetClusterFileContent(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SaveClusterFileContent
// @Description delete cluster for Zeek group
// @Success 200 {object} models.Groups
// @router /saveClusterFileContent [put]
func (n *GroupController) SaveClusterFileContent() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveClusterFileContent(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SyncClusterFile
// @Description Sync cluster for Zeek group
// @Success 200 {object} models.Groups
// @router /syncClusterFile [put]
func (n *GroupController) SyncClusterFile() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SyncClusterFile(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SyncAllGroupCluster
// @Description sync all clusters for Zeek group
// @Success 200 {object} models.Groups
// @router /syncAllGroupCluster [put]
func (n *GroupController) SyncAllGroupCluster() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SyncAllGroupCluster(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SyncAllSuricataGroup
// @Description Synchronize all group elements
// @Success 200 {object} models.Groups
// @router /syncAllSuricataGroup [put]
func (n *GroupController) SyncAllSuricataGroup() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    
        err := models.SyncAllSuricataGroup(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SuricataGroupService
// @Description Suricata stazrt/stop for all group nodes.
// @Success 200 {object} models.Groups
// @router /suricata [put]
func (n *GroupController) SuricataGroupService() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    
        err := models.SuricataGroupService(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SuricataNodesStatus
// @Description Get cluster file content from Zeek group
// @Success 200 {object} models.Groups
// @router /suricata/status/:uuid [get]
func (n *GroupController) SuricataNodesStatus() { 
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        uuid := n.GetString(":uuid") 
    
        data,err := models.SuricataNodesStatus(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}