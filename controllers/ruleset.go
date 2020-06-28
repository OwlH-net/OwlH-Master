package controllers

import (
    "encoding/json"
    "owlhmaster/models"
    "owlhmaster/validation"
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
func (n *RulesetController) GetRules() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetRules"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        mstatus, err := models.GetRules(n.Ctx.Input.Header("user"))
        n.Data["json"] = mstatus
        if err != nil {
            logs.Info("GetRules -> error: %s", err.Error())
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title GetRuleSID
// @Description Get Ruleset by SID
// @Success 200 {object} models.Ruleset
// @Failure 403 SID not exist
// @router /rule/:sid/:uuid [get]
func (n *RulesetController) GetRuleSID() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetRuleSID"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        sid := n.GetString(":sid")
        uuid := n.GetString(":uuid")
        ruleSidPath := make(map[string]string)
        ruleSidPath["sid"] = sid
        ruleSidPath["uuid"] = uuid
        mstatus, err := models.GetRuleSID(ruleSidPath, n.Ctx.Input.Header("user"))
        n.Data["json"] = mstatus
        if err != nil {
            logs.Info("GetRuleSID -> error: %s", err.Error())
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()

}

// // @Title AddRuleset
// // @Description Add Ruleset
// // @Success 200 {object} models.Ruleset
// // @Failure 403 AddRuleset can't add ruleset
// // @router /new [post]
// func (n *RulesetController) AddRuleset(){
//     var ruleset map[string]string
//     json.Unmarshal(n.Ctx.Input.RequestBody, &ruleset)
//     err := models.AddRuleset(ruleset)
//     n.Data["json"] = map[string]string{"ack": "true"}
//     if err != nil {
//         logs.Info("AddRuleset -> error: %s", err.Error())
//         n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
//     }
//     n.ServeJSON()
// }

// @Title GetAllRulesets
// @Description Get full list of rulesets
// @Success 200 {object} models.ruleset
// @router / [get]
func (n *RulesetController) GetAllRulesets() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetAllRulesets"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        rulesets, err := models.GetAllRulesets(n.Ctx.Input.Header("user"))
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        } else {
            n.Data["json"] = rulesets
        }
    }
    n.ServeJSON()
}

// @Title GetRulesetRules
// @Description Get rules from specific ruleset
// @Success 200 {object} models.ruleset
// @router /rules/:uuid [get]
// @router /:uuid/rules [get]
func (n *RulesetController) GetRulesetRules() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetRulesetRules"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        rulesets, err := models.GetRulesetRules(uuid, n.Ctx.Input.Header("user"))
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        } else {
            n.Data["json"] = rulesets
        }
    }

    n.ServeJSON()
}

// @Title SetRuleSelected
// @Description Set rules from specific ruleset
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /set [put]
func (n *RulesetController) SetRuleSelected() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SetRuleSelected"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var ruleSelected map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &ruleSelected)
        err := models.SetRuleSelected(ruleSelected, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            logs.Info("RulesetSelected -> error: %s", err.Error())
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetRuleSelected
// @Description Get rule uid from associated node
// @Success 200 {object} models.ruleset
// @router /get/:uuid [get]
// @router /:uuid/get [get]
func (n *RulesetController) GetRuleSelected() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetRuleSelected"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        rulesets, err := models.GetRuleSelected(uuid, n.Ctx.Input.Header("user"))
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        } else {
            n.Data["json"] = rulesets
        }
        logs.Info("GetRuleSelected: " + rulesets)
    }
    n.ServeJSON()
}

// @Title GetRuleName
// @Description Get rule name from uid rule
// @Success 200 {object} models.ruleset
// @router /get/name/:uuid [get]
// @router /:uuid/get/name [get]
// @router /get/:uuid/name [get]
func (n *RulesetController) GetRuleName() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetRuleName"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        name, err := models.GetRuleName(uuid, n.Ctx.Input.Header("user"))
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        } else {
            n.Data["json"] = name
        }
    }
    n.ServeJSON()
}

// // @Title SetClonedRuleset
// // @Description Create a copy of selected ruleset with a new custom name
// // @Success 200 {object} models.ruleset
// // @router /clone [put]
// func (n *RulesetController) SetClonedRuleset() {
//     var clonedMap map[string]string
//     json.Unmarshal(n.Ctx.Input.RequestBody, &clonedMap)
//     err := models.SetClonedRuleset(clonedMap)
//     n.Data["json"] = map[string]string{"ack": "true"}
//     if err != nil {
//         logs.Info("SetClonedRuleset -> error: %s", err.Error())
//         n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
//     }
//     n.ServeJSON()
// }

// @Title SetRulesetAction
// @Description Set rules from specific ruleset
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /action [put]
func (n *RulesetController) SetRulesetAction() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SetRulesetAction"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var ruleAction map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &ruleAction)
        err := models.SetRulesetAction(ruleAction, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetRuleNote
// @Description Set note for specific rule
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /getnote/:uuid/:sid [get]
func (n *RulesetController) GetRuleNote() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetRuleNote"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        sid := n.GetString(":sid")
        uuid := n.GetString(":uuid")
        ruleGetNote := make(map[string]string)
        ruleGetNote["sid"] = sid
        ruleGetNote["uuid"] = uuid
        note, err := models.GetRuleNote(ruleGetNote, n.Ctx.Input.Header("user"))
        n.Data["json"] = note
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SetRuleNote
// @Description Set note for specific rule
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /note [put]
func (n *RulesetController) SetRuleNote() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SetRuleNote"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var ruleAction map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &ruleAction)
        err := models.SetRuleNote(ruleAction, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteRuleset
// @Description Delete a ruleset
// @Success 200 {string} ruleset deleted
// @Failure 403 Connection failure
// @router /deleteRuleset [delete]
func (n *RulesetController) DeleteNode() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"DeleteRuleset"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var rulesetDelete map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &rulesetDelete)
        err := models.DeleteRuleset(rulesetDelete, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SyncRulesetToAllNodes
// @Description synchronize Ruleset to all nodes who use it
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /synchronize [put]
func (n *RulesetController) SyncRulesetToAllNodes() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SyncRulesetToAllNodes"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.SyncRulesetToAllNodes(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title GetAllRuleData
// @Description Get all data from rule data
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /getAllRuleData [get]
func (n *RulesetController) GetAllRuleData() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetAllRuleData"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        data, err := models.GetAllRuleData(n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title AddNewRuleset
// @Description Add new ruleset
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /addNewRuleset [put]
func (n *RulesetController) AddNewRuleset() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"AddNewRuleset"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        duplicated, err := models.AddNewRuleset(anode, n.Ctx.Input.Header("user"))

        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        } else {
            if duplicated == nil {
                n.Data["json"] = map[string]string{"ack": "true"}
            } else {
                n.Data["json"] = string(duplicated)
            }
        }
    }
    n.ServeJSON()
}

// @Title ModifyRuleset
// @Description modify local ruleset
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /modify [put]
func (n *RulesetController) ModifyRuleset() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ModifyRuleset"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        duplicated, err := models.ModifyRuleset(anode, n.Ctx.Input.Header("user"))

        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        } else {
            if duplicated == nil {
                n.Data["json"] = map[string]string{"ack": "true"}
            } else {
                n.Data["json"] = string(duplicated)
            }
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetAllCustomRulesets"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        data, err := models.GetAllCustomRulesets(n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SynchronizeAllRulesets
// @Description Synchronize All Custom Rulesets
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /syncAllRulesets [put]
func (n *RulesetController) SynchronizeAllRulesets() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SynchronizeAllRulesets"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        err := models.SynchronizeAllRulesets(n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title AddRulesToCustomRuleset
// @Description Add rules to custom ruleset
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /addRulesToCustom [put]
func (n *RulesetController) AddRulesToCustomRuleset() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"AddRulesToCustomRuleset"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        data, err := models.AddRulesToCustomRuleset(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ReadRulesetData
// @Description Add rules to custom ruleset
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /readRuleset/:uuid [put]
func (n *RulesetController) ReadRulesetData() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ReadRulesetData"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.ReadRulesetData(uuid, n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SaveRulesetData
// @Description Add rules to custom ruleset
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /saveRuleset/ [put]
func (n *RulesetController) SaveRulesetData() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SaveRulesetData"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveRulesetData(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// // @Title TimeSchedule
// // @Description Add a time schedule for syncronize rulesets
// // @Success 200 {object} models.ruleset
// // @Failure 403 Connection Failure
// // @router /timeSchedule [put]
// func (n *RulesetController) TimeSchedule() {
//     var anode map[string]string
//     json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
//     err := models.TimeSchedule(anode)
//     n.Data["json"] = map[string]string{"ack": "true"}
//     if err != nil {
//         n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
//     }
//     n.ServeJSON()
// }

// // @Title StopTimeSchedule
// // @Description Stop a current time schedule for syncronize rulesets
// // @Success 200 {object} models.ruleset
// // @Failure 403 Connection Failure
// // @router /stopTimeSchedule [put]
// func (n *RulesetController) StopTimeSchedule() {
//     var anode map[string]string
//     json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
//     err := models.StopTimeSchedule(anode)
//     n.Data["json"] = map[string]string{"ack": "true"}
//     if err != nil {
//         n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
//     }
//     n.ServeJSON()
// }

// @Title UpdateRule
// @Description Update a specific rule
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /updateRule [put]
func (n *RulesetController) UpdateRule() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"UpdateRule"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.UpdateRule(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SyncToAll
// @Description synchronize Ruleset to all nodes and groups who use it
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /syncToAll [put]
func (n *RulesetController) SyncToAll() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SyncToAll"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.SyncToAll(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title Set Default Ruleset
// @Description Set Default Ruleset
// @Success 200 {object} models.ruleset
// @router /setdefault/:uuid [put]
func (n *RulesetController) SetDefaultRuleset() {

    n.Data["json"] = map[string]string{"ack": "true"}

    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"AddNewRuleset"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        logs.Info("lets set this %s as default ruleset", uuid)
        err := models.SetDefaultRuleset(uuid)
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}
