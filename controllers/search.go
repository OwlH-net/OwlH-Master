package controllers

import (
  "encoding/json"

  "github.com/OwlH-net/OwlH-Master/models"

  //"strconv"
  "github.com/OwlH-net/OwlH-Master/validation"
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
  //this is a get. Used put because is needed to send json from UI
  errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
  if errToken != nil {
    n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
    n.ServeJSON()
    return
  }
  permissions := []string{"GetRulesetsBySearch"}
  hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
  if permissionsErr != nil || hasPermission == false {
    n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
  } else {
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    data, err := models.GetRulesetsBySearch(anode, n.Ctx.Input.Header("user"))
    n.Data["json"] = data
    if err != nil {
      n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
  }
  n.ServeJSON()
}
