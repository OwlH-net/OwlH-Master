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

// @Title GetAllNodes
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

// @Title GetServiceStatus
// @Description Get service status for a specific node
// @Success 200 {object} models.Node
// @router /pingservice/:uuid [get]
func (n *NodeController) GetServiceStatus() {
	uuid := n.GetString(":uuid")
    err := models.GetServiceStatus(uuid)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title DeployService
// @Description Get service status for a specific node
// @Success 200 {object} models.Node
// @router /deployservice/:uuid [put]
func (n *NodeController) DeployService() {
	uuid := n.GetString(":uuid")
    err := models.DeployService(uuid)
    n.Data["json"] = map[string]string{"ack": "true"}
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

// @Title SyncRulesetToNode
// @Description Send Ruleset file to node
// @Success 200 {object} models.Node
// @router /ruleset/set [put]
func (n *NodeController) SyncRulesetToNode() {
	anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.SyncRulesetToNode(anode)
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

// @Title DeletePorts
// @Description delete ports from knownports
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /ports/delete/:uuid [put]
func (n *NodeController) DeletePorts() {
	uuid := n.GetString(":uuid")
	anode := make(map[string]string)
	json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

	err := models.DeletePorts(anode,uuid)
	n.Data["json"] = map[string]string{"ack": "true"}
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "nid": uuid, "error": err.Error()}
	}

    n.ServeJSON()
}

// @Title DeleteAllPorts
// @Description delete all ports from knownports
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /ports/deleteAll/:uuid [put]
func (n *NodeController) DeleteAllPorts() {
	uuid := n.GetString(":uuid")
	err := models.DeleteAllPorts(uuid)
	n.Data["json"] = map[string]string{"ack": "true"}
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
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

// @Title PingAnalyzer
// @Description Ping node analyzer
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /PingAnalyzer/:uuid [get]
// @router /:nid/PingAnalyzer [get]
func (n *NodeController) PingAnalyzer() {
	uuid := n.GetString(":uuid")
	data, err := models.PingAnalyzer(uuid)
	n.Data["json"] = data
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
	}

    n.ServeJSON()
}

// @Title ChangeAnalyzerStatus
// @Description Ping node analyzer
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /analyzer [put]
func (n *NodeController) ChangeAnalyzerStatus() {
	anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.ChangeAnalyzerStatus(anode)
	n.Data["json"] = map[string]string{"ack": "true", "uuid": anode["uuid"]}
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "uuid": anode["uuid"], "error": err.Error()}
	}

    n.ServeJSON()
}

// @Title Deploy
// @Description deploy node elements
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /deployNode [put]
func (n *NodeController) Deploy() {
	anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.DeployNode(anode)
	n.Data["json"] = map[string]string{"ack": "true", "uuid": anode["uuid"]}
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "uuid": anode["uuid"], "error": err.Error()}
	}

    n.ServeJSON()
}

// @Title CheckDeploy
// @Description deploy node elements
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /checkDeploy/:uuid [get]
func (n *NodeController) CheckDeploy() {
	uuid := n.GetString(":uuid")
	anode := models.CheckDeploy(uuid)
	n.Data["json"] = anode
    n.ServeJSON()
}

// @Title ChangeDataflowValues
// @Description Change node data flow values
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /changeDataflowValues [put]
func (n *NodeController) ChangeDataflowValues() {
	anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.ChangeDataflowValues(anode)
	n.Data["json"] = map[string]string{"ack": "true", "uuid": anode["uuid"]}
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "uuid": anode["uuid"], "error": err.Error()}
	}

    n.ServeJSON()
}

// @Title LoadDataflowValues
// @Description Load node data flow values
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /loadDataflowValues/:uuid [get]
func (n *NodeController) LoadDataflowValues() {
	uuid := n.GetString(":uuid")
	data, err := models.LoadDataflowValues(uuid)
	n.Data["json"] = data
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
	}

    n.ServeJSON()
}

// @Title LoadNetworkValues
// @Description Load network data from network values
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /loadNetworkValues/:uuid [get]
func (n *NodeController) LoadNetworkValues() {
	uuid := n.GetString(":uuid")
	data, err := models.LoadNetworkValues(uuid)
	n.Data["json"] = data
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
	}

    n.ServeJSON()
}

// @Title LoadNetworkValuesSelected
// @Description Load network data from network values
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /loadNetworkValuesSelected/:uuid [get]
func (n *NodeController) LoadNetworkValuesSelected() {
	uuid := n.GetString(":uuid")
	data, err := models.LoadNetworkValuesSelected(uuid)
	n.Data["json"] = data
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
	}

    n.ServeJSON()
}

// @Title UpdateNetworkInterface
// @Description Get network values for master
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /updateNetworkInterface [put]
func (m *NodeController) UpdateNetworkInterface() {
	anode := make(map[string]string)
    json.Unmarshal(m.Ctx.Input.RequestBody, &anode)
    err := models.UpdateNetworkInterface(anode)
	m.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false: " + err.Error()}
    }
    m.ServeJSON()
}

// @Title SaveSocketToNetwork
// @Description save socket to network information
// @Success 200 {object} models.Node
// @router /saveSocketToNetwork [put]
func (n *NodeController) SaveSocketToNetwork() {
	anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.SaveSocketToNetwork(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}


// @Title SaveNewLocal
// @Description save New local into dataflow at node
// @Success 200 {object} models.Node
// @router /saveNewLocal [put]
func (n *NodeController) SaveNewLocal() {
	anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.SaveNewLocal(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title SaveVxLAN
// @Description save VxLAN into dataflow at node
// @Success 200 {object} models.Node
// @router /saveVxLAN [put]
func (n *NodeController) SaveVxLAN() {
	anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
	err := models.SaveVxLAN(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title SocketToNetworkList
// @Description Load node data flow values
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /socketToNetworkList/:uuid [get]
func (n *NodeController) SocketToNetworkList() {
	uuid := n.GetString(":uuid")
	data, err := models.SocketToNetworkList(uuid)
	n.Data["json"] = data
	if err != nil {
		n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
	}

    n.ServeJSON()
}