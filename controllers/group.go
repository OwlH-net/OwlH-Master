package controllers

import (
    "owlhmaster/models"
    "encoding/json"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
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
	logs.Warn(anode["name"])
	logs.Warn(anode["desc"])
	n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Error("GROUP CREATE -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}