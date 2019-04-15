package controllers

import (
	"owlhmaster/models"
	"github.com/astaxie/beego"
)

type MasterController struct {
	beego.Controller
}


// @Title GetMasterTitle
// @Description Get title for master
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /getMasterTitle [get]
func (m *MasterController) GetMasterTitle() {
    data, err := models.GetMasterTitle()
    m.Data["json"] = data
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false: " + err.Error()}
    }
    m.ServeJSON()
}


