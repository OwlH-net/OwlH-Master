package controllers

import (
	"owlhmaster/models"
	"encoding/json"
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego"
)

type SchedulerController struct {
	beego.Controller
}

// @Title SchedulerTask
// @Description add scheduler task
// @Success 200 {object} models.scheduler
// @Failure 403 Connection Failure
// @router /add [put]
func (n *SchedulerController) SchedulerTask() { 
    var anode map[string]string
	json.Unmarshal(n.Ctx.Input.RequestBody, &anode)	
	logs.Warn(anode)
	err := models.SchedulerTask(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title StopTask
// @Description stop scheduler task
// @Success 200 {object} models.scheduler
// @Failure 403 Connection Failure
// @router /stop [put]
func (n *SchedulerController) StopTask() { 
    var anode map[string]string
	json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.StopTask(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetLog
// @Description get scheduler log
// @Success 200 {object} models.scheduler
// @Failure 403 Connection Failure
// @router /log/:uuid [get]
func (n *SchedulerController) GetLog() { 
    uuid := n.GetString(":uuid")
	logReg,err := models.GetLog(uuid)	
    n.Data["json"] = logReg
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}