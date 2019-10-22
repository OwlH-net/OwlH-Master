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
// @router /DeleteGroup/:groupID [put]
func (n *GroupController) DeleteGroup() { 
	groupID := n.GetString(":groupID") 
    err := models.DeleteGroup(groupID)
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
// @router /getAllNodesGroup [get]
func (n *GroupController) GetAllNodesGroup() { 
    groups, err := models.GetAllNodesGroup()
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