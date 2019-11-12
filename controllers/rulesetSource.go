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

// @Title CreateCustomRulesetSource
// @Description Create new RulesetSource
// @Success 200 {object} models.RulesetSource
// @Failure 403 body is empty
// @router /custom [post]
func (n *RulesetSourceController) CreateCustomRulesetSource() {
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.CreateCustomRulesetSource(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Error("CreateCustomRulesetSource CREATE -> error: %s", err.Error())
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
// @router /DeleteRulesetSource [delete]
func (n *RulesetSourceController) DeleteRulesetSource() { 
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.DeleteRulesetSource(anode)
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

// @Title OverwriteDownload
// @Description Overwrite the current download
// @Success 200 {object} models.RulesetSource
// @router /overwriteDownload [put]
func (n *RulesetSourceController) OverwriteDownload() { 
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    
    err := models.OverwriteDownload(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }

    n.ServeJSON()
}

// @Title CompareFiles
// @Description Compare local ruleset file with their downloaded source file
// @Success 200 {object} models.RulesetSource
// @router /compareSourceFiles/:uuid [get]
func (n *RulesetSourceController) CompareFiles() { 
    uuid := n.GetString(":uuid")
    mapData,err := models.CompareFiles(uuid)
    n.Data["json"] = mapData
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title AddNewLinesToRuleset
// @Description Add new downloaded rules to local ruleset
// @Success 200 {object} models.RulesetSource
// @router /AddNewLinesToRuleset/:uuid [put]
func (n *RulesetSourceController) AddNewLinesToRuleset() { 
    uuid := n.GetString(":uuid")
    err := models.AddNewLinesToRuleset(uuid)
    n.Data["json"] = map[string]string{"ack": "true"}
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

// @Title GetFileUUIDfromRulesetUUID
// @Description Get full list of RulesetSource
// @Success 200 {object} models.RulesetSource
// @router /GetFileUUIDfromRulesetUUID/:uuid [get]
func (n *RulesetSourceController) GetFileUUIDfromRulesetUUID() { 
    value := n.GetString(":uuid")
    rulesetSource, err := models.GetFileUUIDfromRulesetUUID(value)
    n.Data["json"] = rulesetSource
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title OverwriteRuleFile
// @Description delete a RulesetSource specific file
// @Success 200 {object} models.OverwriteRuleFile
// @router /OverwriteRuleFile/:uuid [put]
func (n *RulesetSourceController) OverwriteRuleFile() { 
    uuid := n.GetString(":uuid") 
    err := models.OverwriteRuleFile(uuid)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}