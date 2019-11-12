package controllers

import (
    "owlhmaster/models"
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
    data, err := models.GetChangeControl()
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}