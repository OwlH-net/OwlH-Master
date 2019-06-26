package controllers

import (
	"owlhmaster/models"
	"encoding/json"
    // "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego"
)

type SchedulerController struct {
	beego.Controller
}

// @Title SchedulerTask
// @Description Add schedule task
// @Success 200 {object} models.scheduler
// @Failure 403 Connection Failure
// @router /add [put]
func (n *SchedulerController) SchedulerTask() { 
    var anode map[string]string
	json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.SchedulerTask(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title StopTask
// @Description Add schedule task
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