package controllers

import (
    "owlhmaster/models"
    "owlhmaster/validation"
    "github.com/astaxie/beego"
)

type ChangecontrolController struct {
    beego.Controller
}

// @Title GetChangeControl
// @Description Get changeControl database values
// @Success 200 {object} models.ChangeControl
// @Failure 403 body is empty
// @router / [get]
func (n *ChangecontrolController) GetChangeControl() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetChangeControl"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.GetChangeControl(n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }

    }
    n.ServeJSON()
}