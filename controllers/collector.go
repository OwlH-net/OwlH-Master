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
        logs.Error("PlayCollector ERROR -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title StopCollector
// @Description Stop Collector
// @Success 200 {object} models.Collector
// @Failure 403 body is empty
// @router /stop/:uuid [get]
func (n *CollectorController) StopCollector() {
    uuid := n.GetString(":uuid")
    err := models.StopCollector(uuid)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Error("StopCollector ERROR -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title ShowCollector
// @Description Show Collector
// @Success 200 {object} models.Collector
// @Failure 403 body is empty
// @router /show/:uuid [get]
func (n *CollectorController) ShowCollector() {
    uuid := n.GetString(":uuid")
    data, err := models.ShowCollector(uuid)
    n.Data["json"] = data
    if err != nil {
        logs.Error("ShowCollector ERROR -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title PlayMasterCollector
// @Description Play Master Collector
// @Success 200 {object} models.Collector
// @Failure 403 body is empty
// @router /playMasterCollector [put]
func (n *CollectorController) PlayMasterCollector() {
    err := models.PlayMasterCollector()
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Error("PlayMasterCollector ERROR -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    
    n.ServeJSON()
}

// @Title StopMasterCollector
// @Description Play Master Collector
// @Success 200 {object} models.Collector
// @Failure 403 body is empty
// @router /stopMasterCollector [put]
func (n *CollectorController) StopMasterCollector() {
    err := models.StopMasterCollector()
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Error("StopMasterCollector ERROR -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    
    n.ServeJSON()
}

// @Title ShowMasterCollector
// @Description Show Master Collector
// @Success 200 {object} models.Collector
// @Failure 403 body is empty
// @router /showMasterCollector [get]
func (n *CollectorController) ShowMasterCollector() {
    data, err := models.ShowMasterCollector()
    n.Data["json"] = data
    if err != nil {
        logs.Error("ShowMasterCollector Error -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    
    n.ServeJSON()
}