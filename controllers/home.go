package controllers

import (
    "github.com/astaxie/beego"
    "owlhmaster/validation"
)

type HomeController struct {
    beego.Controller
}


// @Title Home
// @Description Get master info
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router / [get]
func (n *HomeController) Home() {
    err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"))
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }else{
        n.Data["json"] = map[string]string{"ack": "true"}
    }
    n.ServeJSON()
}