package controllers

import (
    "owlhmaster/models"
    "encoding/json"
    //"strconv"
    "github.com/astaxie/beego"
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
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    data, err := models.GetRulesetsBySearch(anode)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}