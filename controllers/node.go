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



// // @Title GetNode
// // @Description Get Node detail
// // @Success 200 {object} models.Node
// // @Failure 403 :nid is empty
// // @router /:nid [get]
// func (n *NodeController) GetNode() { 
//     nid := n.GetString(":nid")
//     if nid != "" {
//         nn, err := models.GetNode(nid)
//         if err != nil {
// 	        logs.Info("NODE -> Get Node Detail -> error: %s", err.Error())
//             n.Data["json"] = err.Error()
//         } else {
//             n.Data["json"] = nn
//         }
//     }
//     n.ServeJSON()
// }

// @Title CreateNode
// @Description Create Node
// @Success 200 {object} models.Node
// @Failure 403 body is empty
// @router / [post]
func (n *NodeController) CreateNode() {
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.AddNode(anode)
	n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Error("NODE CREATE -> error: %s", err.Error())
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
    n.Data["json"] = map[string]string{"nid": nid, "state":"Success"}
    if nid == "" {
    	n.Data["json"] = map[string]string{"nid": "", "state":"Failure"}
        logs.Error("NODE DEPLOY -> error -> No Node ID")
    }
    n.ServeJSON()
}


// @Title UpdateNode
// @Description Update Node
// @Success 200 {string} node updated
// @Failure 403 body is empty
// @router / [put]
func (n *NodeController) UpdateNode() {
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.UpdateNode(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Info("NODE UPDATE -> error: %s", err.Error())
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
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
    if nid != "" {
        err := models.PingNode(nid)
        n.Data["json"] = map[string]string{"ping": "pong", "nid": nid}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "nid": nid, "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title Get Suricata
// @Description Get Suricate status from Node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /suricata/:nid [get]
// @router /:nid/suricata [get]
func (n *NodeController) GetSuricata() { 
    nid := n.GetString(":nid")
    data,err := models.Suricata(nid)   
    logs.Warn("GetSuricata")
    logs.Warn(data)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title Get Zeek
// @Description Get Zeek status from Node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /zeek/:nid [get]
// @router /:nid/zeek [get]
func (n *NodeController) GetZeek() { 
    nid := n.GetString(":nid")
    data,err := models.Zeek(nid)

    logs.Warn("GetZeek")
    logs.Warn(data)
    // var anode map[string]string
    // json.Unmarshal(data, &anode)
    // n.Data["json"] = anode
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title Get Wazuh
// @Description Get wazuh status from Node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /wazuh/:nid [get]
// @router /:nid/wazuh [get]
func (n *NodeController) GetWazuh() { 
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"status": "false", "error": "No hay NID"}
    data,err := models.Wazuh(nid)
    logs.Warn("GetWazuh")
    logs.Warn(data)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetSuricata BPF
// @Description Get Suricata BPF from node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /suricata/:nid/bpf [get]
// @router /:nid/suricata/bpf [get]
// @router /bpf/:nid [get]
func (n *NodeController) GetSuricataBPF() { 
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"status": "false", "error": "There is no BPF"}
    if nid != "" {
        data,err := models.GetSuricataBPF(nid)
        n.Data["json"] = map[string]string{"bpf": data}
        if err != nil {
			logs.Error("Can't get Suricata status" + err.Error())
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }
    }
    n.ServeJSON()
}


// @Title PutSuricataBPF
// @Description Set BPF to node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /suricata/:nid/bpf [put]
// @router /:nid/suricata/bpf [put]
// @router /bpf/:nid [put]
func (n *NodeController) PutSuricataBPF() { 
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"status": "false", "error": "There is no BPF"}
    if nid != "" {
        err := models.PutSuricataBPF(anode)
        n.Data["json"] = map[string]string{"status": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title Get All Nodes
// @Description Get full list of nodes
// @Success 200 {object} models.Node
// @router / [get]
func (n *NodeController) GetAllNodes() { 
    nodes, err := models.GetAllNodes()
    n.Data["json"] = nodes
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
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

// @Title SetRuleset
// @Description Send Ruleset file to node
// @Success 200 {object} models.Node
// @router /ruleset/set/:nid [get]
func (n *NodeController) SetRuleset() { 
    nid := n.GetString(":nid")
    err := models.SetRuleset(nid)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetNodeFile
// @Description Get specific file from node and send to webpage 
// @Success 200 {object} models.Node
// @router /loadfile/:uuid/:fileName [get]
func (n *NodeController) GetNodeFile() { 
    logs.Info("Inside GetNodeFile")
    anode := make(map[string]string)
    anode["uuid"] = n.GetString(":uuid")
    anode["file"] = n.GetString(":fileName")
    returnData,err := models.GetNodeFile(anode)
    n.Data["json"] = returnData
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title SetNodeFile
// @Description Set changed data on webpage to specific file in node.
// @Success 200 {object} models.Node
// @router /savefile [put]
func (n *NodeController) SetNodeFile() { 
    anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.SetNodeFile(anode)

    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetAllFiles
// @Description Send Ruleset file to node
// @Success 200 {object} models.Node
// @router /getAllFiles/:uuid [get]
func (n *NodeController) GetAllFiles() { 
    uuid := n.GetString(":uuid")
    data, err := models.GetAllFiles(uuid)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title RunSuricata
// @Description Run suricata server
// @Success 200 {object} models.Node
// @router /RunSuricata/:uuid [put]
func (n *NodeController) RunSuricata() { 
	uuid := n.GetString(":uuid")
	data, err := models.RunSuricata(uuid)

    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title StopSuricata
// @Description Stop suricata server
// @Success 200 {object} models.Node
// @router /StopSuricata/:uuid [put]
func (n *NodeController) StopSuricata() { 
    uuid := n.GetString(":uuid")
    data, err := models.StopSuricata(uuid)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title RunZeek
// @Description Run zeek server
// @Success 200 {object} models.Node
// @router /RunZeek/:uuid [put]
func (n *NodeController) RunZeek() { 
    uuid := n.GetString(":uuid")
    data, err := models.RunZeek(uuid)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title StopZeek
// @Description Stop zeek server
// @Success 200 {object} models.Node
// @router /StopZeek/:uuid [put]
func (n *NodeController) StopZeek() { 
    uuid := n.GetString(":uuid")
    data, err := models.StopZeek(uuid)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title RunWazuh
// @Description Run wazuh server
// @Success 200 {object} models.Node
// @router /RunWazuh/:uuid [put]
func (n *NodeController) RunWazuh() { 
    uuid := n.GetString(":uuid")
    data, err := models.RunWazuh(uuid)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title StopWazuh
// @Description Stop wazuh server
// @Success 200 {object} models.Node
// @router /StopWazuh/:uuid [put]
func (n *NodeController) StopWazuh() { 
    uuid := n.GetString(":uuid")
    data, err := models.StopWazuh(uuid)
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title SyncRulesetToAllNodes
// @Description synchronize Ruleset to all nodes using it
// @Success 200 {object} models.Node
// @Failure 403 Connection Failure
// @router /synchronize/:uuid [put]
func (n *NodeController) SyncRulesetToAllNodes() { 
	uuid := n.GetString(":uuid")
    err := models.SyncRulesetToAllNodes(uuid)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title DeployZeek
// @Description Get Pong from Node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /deploy/:nid [get]
// @router /:nid/deploy [get]
func (n *NodeController) DeployZeek() { 
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
    if nid != "" {
        err := models.DeployZeek(nid)
        n.Data["json"] = map[string]string{"ping": "pong", "nid": nid}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "nid": nid, "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ShowPorts
// @Description Get ports from knownports
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /ports/:nid [get]
// @router /:nid/ports [get]
func (n *NodeController) ShowPorts() { 
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
    if nid != "" {
        data, err := models.ShowPorts(nid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "nid": nid, "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ChangeMode
// @Description put Pong from Node
// @Success 200 {object} models.Node
// @Failure 403 Error
// @router /mode [put]
func (n *NodeController) ChangeMode() { 
	anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.ChangeMode(anode)
	n.Data["json"] = map[string]string{"ack": "true"}
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
	}
    n.ServeJSON()
}

// @Title ChangeStatus
// @Description put Pong from Node
// @Success 200 {object} models.Node
// @Failure 403 Error
// @router /status [put]
func (n *NodeController) ChangeStatus() { 
	anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
	err := models.ChangeStatus(anode)
	n.Data["json"] = map[string]string{"ack": "true"}
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
	}
    
    n.ServeJSON()
}

// @Title Ping Ports
// @Description Get Ping from ports
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /PingPorts/:nid [get]
// @router /:nid/PingPorts [get]
func (n *NodeController) PingPorts() { 
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
    if nid != "" {
        data, err := models.PingPorts(nid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "nid": nid, "error": err.Error()}
        }
	}

    n.ServeJSON()
}