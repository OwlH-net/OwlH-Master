package controllers

import (
    "owlhmaster/models"
    "encoding/json"
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego"
)

type RulesetSourceController struct {
	beego.Controller
}

// @Title CreateRulesetSource
// @Description Create new RulesetSource
// @Success 200 {object} models.RulesetSource
// @Failure 403 body is empty
// @router / [post]
func (n *RulesetSourceController) CreateRulesetSource() {
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.CreateRulesetSource(anode)
	n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Error("CreateRulesetSource CREATE -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetAllRulesetSource
// @Description Get full list of RulesetSource
// @Success 200 {object} models.RulesetSource
// @router / [get]
func (n *RulesetSourceController) GetAllRulesetSource() { 
    rulesetSource, err := models.GetAllRulesetSource()
    n.Data["json"] = rulesetSource
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title DeleteRulesetSource
// @Description delete a RulesetSource
// @Success 200 {object} models.DeleteRulesetSource
// @router /DeleteRulesetSource/:uuid [delete]
func (n *RulesetSourceController) DeleteRulesetSource() { 
	uuid := n.GetString(":uuid") 
    err := models.DeleteRulesetSource(uuid)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title DeleteRulesetFile
// @Description delete a RulesetSource specific file
// @Success 200 {object} models.DeleteRulesetFile
// @router /DeleteRulesetFile/:uuid [delete]
func (n *RulesetSourceController) DeleteRulesetFile() { 
	uuid := n.GetString(":uuid") 
    err := models.DeleteRulesetFile(uuid)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title EditRulesetSource
// @Description edit a RulesetSource
// @Success 200 {object} models.RulesetSource
// @router /EditRulesetSource [put]
func (n *RulesetSourceController) EditRulesetSource() { 
	var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.EditRulesetSource(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title DownloadFile
// @Description edit a RulesetSource
// @Success 200 {object} models.RulesetSource
// @router /downloadFile [put]
func (n *RulesetSourceController) DownloadFile() { 
	var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.DownloadFile(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
	}
    n.ServeJSON()
}

// @Title CompareFiles
// @Description edit a RulesetSource
// @Success 200 {object} models.RulesetSource
// @router /compareFiles [put]
func (n *RulesetSourceController) CompareFiles() { 
	var anode map[string]string
	json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    mapData,err := models.CompareFiles(anode)
    n.Data["json"] = mapData
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title createNewFile
// @Description edit a RulesetSource
// @Success 200 {object} models.RulesetSource
// @router /createNewFile [put]
func (n *RulesetSourceController) CreateNewFile() { 
	var anode map[string]string
	json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.CreateNewFile(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetDetails
// @Description edit a RulesetSource
// @Success 200 {object} models.RulesetSource
// @router /getDetails/:uuid [get]
func (n *RulesetSourceController) GetDetails() { 
	uuid := n.GetString(":uuid")
	data, err := models.GetDetails(uuid)
	
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
	
    n.ServeJSON()
}