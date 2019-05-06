package controllers

import (
    "owlhmaster/models"
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego"
)

type CollectorController struct {
	beego.Controller
}

// @Title PlayCollector
// @Description Play collector
// @Success 200 {object} models.Collector
// @Failure 403 body is empty
// @router /play/:uuid [get]
func (n *CollectorController) PlayCollector() {
	uuid := n.GetString(":uuid")
	err := models.PlayCollector(uuid)
	n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Error("COLLECTOR CREATE -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title StopCollector
// @Description Stop collector
// @Success 200 {object} models.Collector
// @Failure 403 body is empty
// @router /stop/:uuid [get]
func (n *CollectorController) StopCollector() {
	uuid := n.GetString(":uuid")
	err := models.StopCollector(uuid)
	n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Error("COLLECTOR CREATE -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title ShowCollector
// @Description Show collector
// @Success 200 {object} models.Collector
// @Failure 403 body is empty
// @router /show/:uuid [get]
func (n *CollectorController) ShowCollector() {
	uuid := n.GetString(":uuid")
	data, err := models.ShowCollector(uuid)
	n.Data["json"] = data
    if err != nil {
        logs.Error("COLLECTOR CREATE -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}