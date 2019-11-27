package controllers

import (
    "owlhmaster/models"
    "encoding/json"
    //"strconv"
    "github.com/astaxie/beego"
)

type StapController struct {
    beego.Controller
}

// @Title AddServer
// @Description Add server Software TAP
// @Success 200 {object} models.Stap
// @Failure 403 ruleset is empty
// @router / [post]
func (n *StapController) AddServer(){ 
    var newServer map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &newServer)
    err := models.AddServer(newServer)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetAllServers
// @Description Get all servers
// @Success 200 {object} models.stap
// @router /:uuid [get]
func (n *StapController) GetAllServers() {
    uuid := n.GetString(":uuid") 
    servers, err := models.GetAllServers(uuid)
    n.Data["json"] = servers
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}


// @Title GetServer
// @Description Get a server
// @Success 200 {object} models.stap
// @router /server/:uuid/:serveruuid [get]
func (n *StapController) GetServer() {
    uuid := n.GetString(":uuid") 
    serveruuid := n.GetString(":serveruuid")
    server, err := models.GetServer(uuid,serveruuid)
    n.Data["json"] = server
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title Get Stap
// @Description Get Stap status from Node
// @Success 200 {object} models.stap
// @Failure 403 :nid is empty
// @router /stap/:nid [get]
// @router /:nid/stap [get]
func (n *StapController) GetStap() { 
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"status": "false", "error": "No hay NID"}
    data,err := models.Stap(nid)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title RunStap
// @Description Run Stap on node
// @Success 200 {object} models.stap
// @router /RunStap/:uuid [put]
func (n *StapController) RunStap() { 
    uuid := n.GetString(":uuid")
    data, err := models.RunStap(uuid)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title StopStap
// @Description Stop Stap on node
// @Success 200 {object} models.stap
// @router /StopStap/:uuid [put]
func (n *StapController) StopStap() { 
    uuid := n.GetString(":uuid")
    data, err := models.StopStap(uuid)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title RunStapServer
// @Description Run Stap server
// @Success 200 {object} models.stap
// @router /RunStapServer/:uuid/:server [put]
func (n *StapController) RunStapServer() { 
    uuid := n.GetString(":uuid")
    server := n.GetString(":server")
    data, err := models.RunStapServer(uuid,server)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title StopStapServer
// @Description Stop Stap server
// @Success 200 {object} models.stap
// @router /StopStapServer/:uuid/:server [put]
func (n *StapController) StopStapServer() { 
    uuid := n.GetString(":uuid")
    server := n.GetString(":server")
    data, err := models.StopStapServer(uuid,server)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title PingServerStap
// @Description Get Pong from specific stap server
// @Success 200 {object} models.stap
// @Failure 403 :nid is empty
// @router /PingServerStap/:nid/:server [get]
func (n *StapController) PingServerStap() { 
    nid := n.GetString(":nid")
    server := n.GetString(":server")
    data,err := models.PingServerStap(nid,server)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "nid": nid, "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title DeleteStapServer
// @Description Stop Stap server
// @Success 200 {object} models.stap
// @router /DeleteStapServer/:uuid/:server [put]
func (n *StapController) DeleteStapServer() { 
    uuid := n.GetString(":uuid")
    server := n.GetString(":server")
    data, err := models.DeleteStapServer(uuid,server)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title EditStapServer
// @Description Edit Stap server
// @Success 200 {object} models.stap
// @router /EditStapServer [put]
func (n *StapController) EditStapServer() { 
    var editedMap map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &editedMap)
    err := models.EditStapServer(editedMap)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}