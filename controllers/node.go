package controllers

import (
    "owlhmaster/models"
    "encoding/json"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
)

type NodeController struct {
	beego.Controller
}



// @Title GetNode
// @Description Get Node detail
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /:nid [get]
func (n *NodeController) GetNode() { 
    nid := n.GetString(":nid")
    if nid != "" {
        nn, err := models.GetNode(nid)
        logs.Info ("Node Details ->  %s", nn)
        if err != nil {
            n.Data["json"] = err.Error()
        } else {
            n.Data["json"] = nn
            logs.Info ("Node Details after node ->  %s", n)
        }
    }
    n.ServeJSON()
}

// @Title CreateNode
// @Description Create Node
// @Success 200 {object} models.Node
// @Failure 403 body is empty
// @router / [post]
func (n *NodeController) CreateNode() {
    logs.Info("NODE CREATE -> In")
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.AddNode(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Info("NODE CREATE -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}


// @Title DeployNode
// @Description Deploy Node
// @Success 200 {string} node deployed
// @Failure 403 :nid is empty
// @router /deploy/:nid [post]
// @router /:nid/deploy [post]
func (n *NodeController) DeployNode() {
    logs.Info("NODE DEPLOY -> In")
    nid := n.GetString(":nid")
    if nid == "" {
        logs.Info("NODE DEPLOY -> error")
    }
    n.Data["json"] = map[string]string{"nid": nid, "state":"Success"}
    n.ServeJSON()
}


// @Title UpdateNode
// @Description Update Node
// @Success 200 {string} node updated
// @Failure 403 body is empty
// @router / [put]
func (n *NodeController) UpdateNode() {
    logs.Info("NODE Update -> In")
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    for key, value := range anode {
        logs.Info("Key: %s, Value: %s", key, value)
    }
    err := models.UpdateNode(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Info("NODE CREATE -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}


// @Title Ping Node
// @Description Get Pong from Node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /ping/:nid [get]
// @router /:nid/ping [get]
func (n *NodeController) GetPong() { 
    logs.Info("GET PONG -> In")
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
    if nid != "" {
        err := models.PingNode(nid)
        n.Data["json"] = map[string]string{"ping": "pong", "nid": nid}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "nid": nid, "error": err.Error()}
        }
    }
    logs.Info("GET PING -> OUT -> %s", n.Data["json"])
    n.ServeJSON()
}

// @Title Get Suricata
// @Description Get Suricate status from Node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /suricata/:nid [get]
// @router /:nid/suricata [get]
func (n *NodeController) GetSuricata() { 
    logs.Info("GET Suricata -> In")
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"status": "false", "error": "No hay NID"}
    if nid != "" {
        data,err := models.Suricata(nid)
        var anode map[string]string
        json.Unmarshal(data, &anode)
        n.Data["json"] = anode
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }
    }
    logs.Info("GET Suricata -> OUT -> %s", n.Data["json"])
    n.ServeJSON()
}


// @Title Get All Nodes
// @Description Get full list of nodes
// @Success 200 {object} models.Node
// @router / [get]
func (n *NodeController) GetAllNodes() { 
    nodes, err := models.GetAllNodes()
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.Data["json"] = nodes
    n.ServeJSON()
}

// @Title Delete a Node
// @Description Delete a node
// @Param   nid        path    string  true        "The nodeid you want to delete"
// @Failure 403 :nid is empty
// @Success 200 {string} node deleted
// @router /:nid [delete]
func (n *NodeController) DeleteNode() { 
    nid := n.Ctx.Input.Param(":nid")
    logs.Info("NODE DELETE -> node id: %s", nid)
    n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
    if nid != "" {
        err := models.DeleteNode(nid)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}
