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
// @router /rule/:sid/:uuid [get]
func (n *RulesetController) GetRuleSID(){ 
    sid := n.GetString(":sid")
    uuid := n.GetString(":uuid")
    ruleSidPath := make(map[string]string)
    ruleSidPath["sid"] = sid
    ruleSidPath["uuid"] = uuid
    mstatus, _ := models.GetRuleSID(ruleSidPath)
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
// @Description Set rules from specific ruleset
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /set [put]
func (n *RulesetController) SetRuleSelected() { 
    logs.Info("ROUTER SetRuleSelected --> iNSIDE")
    var ruleSelected map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &ruleSelected)
    err := models.SetRuleSelected(ruleSelected)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Info("RulesetSelected -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetRuleSelected
// @Description Get rule uid from associated node
// @Success 200 {object} models.ruleset
// @router /get/:uuid [get]
// @router /:uuid/get [get]
func (n *RulesetController) GetRuleSelected() { 
	uuid := n.GetString(":uuid")
	rulesets, err := models.GetRuleSelected(uuid)
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }else{
		n.Data["json"] = rulesets
    }
    logs.Info("GetRuleSelected: "+rulesets)
    n.ServeJSON()
}

// @Title GetRuleName
// @Description Get rule name from uid rule
// @Success 200 {object} models.ruleset
// @router /get/name/:uuid [get]
// @router /:uuid/get/name [get]
// @router /get/:uuid/name [get]
func (n *RulesetController) GetRuleName() { 
	uuid := n.GetString(":uuid")
	name, err := models.GetRuleName(uuid)
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }else{
		n.Data["json"] = name
    }
    logs.Info("GetRuleName: "+name)
    n.ServeJSON()
}

// @Title SetClonedRuleset
// @Description Create a copy of selected ruleset with a new custom name 
// @Success 200 {object} models.ruleset
// @router /clone [put]
func (n *RulesetController) SetClonedRuleset() { 
    logs.Info("Entrando a SetClonedRuleset")
    var clonedMap map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &clonedMap)
    err := models.SetClonedRuleset(clonedMap)
    logs.Info("Saliendo a SetClonedRuleset")
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Info("SetClonedRuleset -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title SetRulesetAction
// @Description Set rules from specific ruleset
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /action [put]
func (n *RulesetController) SetRulesetAction() { 
    logs.Info("Entrando a SetRulesetAction")
    var ruleAction map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &ruleAction)
    err := models.SetRulesetAction(ruleAction)
    logs.Info("Saliendo a SetRulesetAction")
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetRuleNote
// @Description Set note for specific rule
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /getnote/:uuid/:sid [get]
func (n *RulesetController) GetRuleNote() { 
    logs.Info("Entrando a GetRuleNote")
    sid := n.GetString(":sid")
    uuid := n.GetString(":uuid")
    ruleGetNote := make(map[string]string)
    ruleGetNote["sid"] = sid
    ruleGetNote["uuid"] = uuid
    note,err := models.GetRuleNote(ruleGetNote)
    logs.Info("Saliendo a GetRuleNote")
    n.Data["json"] = note
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title SetRuleNote
// @Description Set note for specific rule
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /note [put]
func (n *RulesetController) SetRuleNote() { 
    logs.Info("Entrando a SetRuleNote")
    var ruleAction map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &ruleAction)
    err := models.SetRuleNote(ruleAction)
    logs.Info("Saliendo a SetRuleNote")
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}