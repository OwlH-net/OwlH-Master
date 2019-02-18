package controllers

import (
	"owlhmaster/models"
	"encoding/json"
	//"strconv"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
)

type RulesetController struct {
	beego.Controller
}

// @Title GetRules
// @Description Get Ruleset
// @Success 200 {object} models.Ruleset
// @Failure 403 ruleset is empty
// @router /default [get]
func (n *RulesetController) GetRules(){ 
	mstatus, _ := models.GetRules()
	n.Data["json"] = mstatus
	n.ServeJSON()
}

// @Title GetRuleSID
// @Description Get Ruleset by SID
// @Success 200 {object} models.Ruleset
// @Failure 403 SID not exist
// @router /rule/:sid [get]
// @router /:sid/rule [get]
func (n *RulesetController) GetRuleSID(){ 
	sid := n.GetString(":sid")
	mstatus, _ := models.GetRuleSID(sid)
	n.Data["json"] = mstatus
	n.ServeJSON()
}

// @Title AddRuleset
// @Description Add Ruleset
// @Success 200 {object} models.Ruleset
// @Failure 403 AddRuleset can't add ruleset
// @router /new [post]
func (n *RulesetController) AddRuleset(){ 
    var ruleset map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &ruleset)
    err := models.AddRuleset(ruleset)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Info("AddRuleset -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetAllRulesets
// @Description Get full list of rulesets
// @Success 200 {object} models.ruleset
// @router / [get]
func (n *RulesetController) GetAllRulesets() { 
    rulesets, err := models.GetAllRulesets()
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }else{
		n.Data["json"] = rulesets
	}
    n.ServeJSON()
}

// @Title GetRulesetRules
// @Description Get rules from specific ruleset
// @Success 200 {object} models.ruleset
// @router /rules/:uuid [get]
// @router /:uuid/rules [get]
func (n *RulesetController) GetRulesetRules() { 
	uuid := n.GetString(":uuid")
	rulesets, err := models.GetRulesetRules(uuid)
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }else{
		n.Data["json"] = rulesets
	}
    n.ServeJSON()
}

// @Title SetRuleSelected
// @Description Get rules from specific ruleset
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /set [put]
func (n *RulesetController) SetRuleSelected() { 
    logs.Info("ROUTER SetRuleSelected --> iNSIDE")
    var ruleSelected map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &ruleSelected)
    err := models.SetRuleSelected(ruleSelected)
    logs.Info("vOLVIENDO")
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Info("RulesetSelected -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

