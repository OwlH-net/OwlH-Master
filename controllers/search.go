package controllers

import (
    "owlhmaster/models"
    "encoding/json"
    //"strconv"
    "github.com/astaxie/beego"
    "owlhmaster/validation"
    // "github.com/astaxie/beego/logs"
)

type SearchController struct {
    beego.Controller
}

// @Title GetRulesetsBySearch
// @Description Get rule files from ruleset by search
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /getRulesetsBySearch [put]
func (n *SearchController) GetRulesetsBySearch() { 
    //this is a get. Used put because is needed to send json from UI
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get") 
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        data, err := models.GetRulesetsBySearch(anode)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}