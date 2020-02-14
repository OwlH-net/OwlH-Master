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
    privileges,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !privileges{
        n.Data["json"] = map[string]string{"ack": "false","privileges":"none"}
    }else{
        data, err := models.GetChangeControl()
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }

    }
    n.ServeJSON()
}