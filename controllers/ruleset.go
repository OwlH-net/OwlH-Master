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
	mstatus, err:= models.GetRules()
	n.Data["json"] = mstatus
	if err != nil {
        logs.Info("GetRules -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
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
    mstatus, err := models.GetRuleSID(ruleSidPath)
	n.Data["json"] = mstatus
	if err != nil {
        logs.Info("GetRuleSID -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
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
    var clonedMap map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &clonedMap)
    err := models.SetClonedRuleset(clonedMap)
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
    var ruleAction map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &ruleAction)
    err := models.SetRulesetAction(ruleAction)
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
    sid := n.GetString(":sid")
    uuid := n.GetString(":uuid")
    ruleGetNote := make(map[string]string)
    ruleGetNote["sid"] = sid
    ruleGetNote["uuid"] = uuid
    note,err := models.GetRuleNote(ruleGetNote)
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
    var ruleAction map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &ruleAction)
    err := models.SetRuleNote(ruleAction)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title DeleteRuleset
// @Description Delete a ruleset
// @Success 200 {string} ruleset deleted
// @Failure 403 Connection failure
// @router /deleteRuleset [delete]
func (n *RulesetController) DeleteNode() { 
	var rulesetDelete map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &rulesetDelete)
    err := models.DeleteRuleset(rulesetDelete)
	n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title SyncRulesetToAllNodes
// @Description synchronize Ruleset to all nodes using it
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /synchronize [put]
func (n *RulesetController) SyncRulesetToAllNodes() { 
	var anode map[string]string
	json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	
    err := models.SyncRulesetToAllNodes(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
	}
	
    n.ServeJSON()
}

// @Title GetAllRuleData
// @Description Get all data from rule data
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /getAllRuleData [get]
func (n *RulesetController) GetAllRuleData() { 
    data,err := models.GetAllRuleData()
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title AddNewRuleset
// @Description Add new custom ruleset
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /addNewRuleset [put]
func (n *RulesetController) AddNewRuleset() { 
    var anode map[string]map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	duplicated,err := models.AddNewRuleset(anode)
	
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
	}else {
		if duplicated == nil {
			n.Data["json"] = map[string]string{"ack": "true"}	
		}else{
			n.Data["json"] = string(duplicated)
		}
	}
    n.ServeJSON()
}

// @Title GetAllCustomRulesets
// @Description Get All Custom Rulesets
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /custom [get]
func (n *RulesetController) GetAllCustomRulesets() { 
    data,err := models.GetAllCustomRulesets()
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}