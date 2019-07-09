package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}


// @Title Home
// @Description Get master info
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router / [get]
func (m *HomeController) Home() {
	m.Data["json"] = map[string]string{"ack": "true"}
    m.ServeJSON()
}