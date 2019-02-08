package controllers

import (
	"owlhmaster/models"
	//"encoding/json"
	//"strconv"
	"github.com/astaxie/beego"
    //"github.com/astaxie/beego/logs"
)

type RulesetController struct {
	beego.Controller
}

// @Title GetRules
// @Description Get Ruleset
// @Success 200 {object} models.Ruleset
// @Failure 403 ruleset is empty
// @router / [get]
func (n *RulesetController) GetRules(){ 
	mstatus, _ := models.GetRules()
	n.Data["json"] = mstatus
	n.ServeJSON()
}