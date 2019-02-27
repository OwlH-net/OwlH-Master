package controllers

import (
	"owlhmaster/models"
	"encoding/json"
	//"strconv"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
)

type StapController struct {
	beego.Controller
}

// @Title AddServer
// @Description Add server Software TAP
// @Success 200 {object} models.Stap
// @Failure 403 ruleset is empty
// @router / [post]
func (n *StapController) AddServer(){ 
	logs.Info("Dentro de stap AddServer()")
	var newServer map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &newServer)

	err := models.AddServer(newServer)
	n.Data["json"] = map[string]string{"ack": "true"}

	if err != nil {
        logs.Info("RulesetSelected -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}