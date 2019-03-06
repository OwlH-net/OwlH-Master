package controllers

import (
	"owlhmaster/models"
	"encoding/json"
	//"strconv"
	"github.com/astaxie/beego"
    // "github.com/astaxie/beego/logs"
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
// @router /Stap/:nid [get]
// @router /:nid/Stap [get]
func (n *StapController) GetStap() { 
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"status": "false", "error": "No hay NID"}
    if nid != "" {
        data,err := models.Stap(nid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }
    }
    n.ServeJSON()
}