package controllers

import (
    "owlhmaster/models"
    "encoding/json"
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego"
)

type GroupController struct {
	beego.Controller
}

// @Title CreateGroup
// @Description Create new group
// @Success 200 {object} models.Group
// @Failure 403 body is empty
// @router / [post]
func (n *GroupController) CreateGroup() {
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.CreateGroup(anode)
	n.Data["json"] = map[string]string{"ack": "true!"}
    if err != nil {
        logs.Error("GROUP CREATE -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetAllGroups
// @Description Get full list of groups
// @Success 200 {object} models.Groups
// @router / [get]
func (n *GroupController) GetAllGroups() { 
    groups, err := models.GetAllGroups()
    n.Data["json"] = groups
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title DeleteGroup
// @Description delete a group
// @Success 200 {object} models.Groups
// @router /DeleteGroup/:uuid [put]
func (n *GroupController) DeleteGroup() { 
	uuid := n.GetString(":uuid") 
    err := models.DeleteGroup(uuid)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title EditGroup
// @Description edit a group information
// @Success 200 {object} models.Groups
// @router /editGroup [put]
func (n *GroupController) EditGroup() { 
	var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.EditGroup(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetAllNodesGroup
// @Description Get full list of nodes
// @Success 200 {object} models.Groups
// @router /getAllNodesGroup/:uuid [get]
func (n *GroupController) GetAllNodesGroup() { 
    uuid := n.GetString(":uuid") 
    groups, err := models.GetAllNodesGroup(uuid)
    n.Data["json"] = groups
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title AddGroupNodes
// @Description add nodes to a group
// @Success 200 {object} models.Groups
// @router /addGroupNodes [put]
func (n *GroupController) AddGroupNodes() { 
	var anode map[string]interface{}
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.AddGroupNodes(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}   

// @Title PingGroupNodes
// @Description Ping all group nodes
// @Success 200 {object} models.Groups
// @router /pingGroupNodes [get]
func (n *GroupController) PingGroupNodes() { 
    groups, err := models.PingGroupNodes()
    n.Data["json"] = groups
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetNodeValues
// @Description Get values for specific node
// @Success 200 {object} models.Groups
// @router /getNodeValues/:uuid [get]
func (n *GroupController) GetNodeValues() { 
    uuid := n.GetString(":uuid") 
    groups, err := models.GetNodeValues(uuid)
    n.Data["json"] = groups
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title DeleteNodeGroup
// @Description add nodes to a group
// @Success 200 {object} models.Groups
// @router /deleteNodeGroup/:uuid [put]
func (n *GroupController) DeleteNodeGroup() { 
	uuid := n.GetString(":uuid") 
    err := models.DeleteNodeGroup(uuid)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
} 

// @Title ChangeGroupRuleset
// @Description Change group ruleset
// @Success 200 {object} models.Groups
// @router /changeGroupRuleset [put]
func (n *GroupController) ChangeGroupRuleset() { 
	var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.ChangeGroupRuleset(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title ChangePaths
// @Description Change group paths
// @Success 200 {object} models.Groups
// @router /changePaths [put]
func (n *GroupController) ChangePaths() { 
	var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.ChangePaths(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}