package controllers

import (
    "owlhmaster/models"
    "encoding/json"
    "github.com/astaxie/beego"
    "owlhmaster/validation"
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
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"),"put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)    
        err := models.SchedulerTask(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title StopTask
// @Description stop scheduler task
// @Success 200 {object} models.scheduler
// @Failure 403 Connection Failure
// @router /stop [put]
func (n *SchedulerController) StopTask() { 
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"),"put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.StopTask(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetLog
// @Description get scheduler log
// @Success 200 {object} models.scheduler
// @Failure 403 Connection Failure
// @router /log/:uuid [get]
func (n *SchedulerController) GetLog() { 
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"),"get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        logReg,err := models.GetLog(uuid)    
        n.Data["json"] = logReg
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}