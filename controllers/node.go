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
    data,err := models.Suricata(nid)

    // data["path"]=false
    // data["bin"]=true
    // data["running"]=false

    logs.Warn(data)

    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
    }
    logs.Info("GET Suricata -> OUT -> %s", n.Data["json"])
    n.ServeJSON()
}


// @Title Get Suricata BPF
// @Description Get Suricata BPF from node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /suricata/:nid/bpf [get]
// @router /:nid/suricata/bpf [get]
// @router /bpf/:nid [get]
func (n *NodeController) GetSuricataBPF() { 
    logs.Info("GET SuricataBPF -> In")
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"status": "false", "error": "There is no BPF"}
    if nid != "" {
        data,err := models.GetSuricataBPF(nid)
        n.Data["json"] = map[string]string{"bpf": data}
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }
    }
    logs.Info("GET SuricataBPF -> OUT -> %s", n.Data["json"])
    n.ServeJSON()
}


// @Title Put Suricata BPF
// @Description Put Suricata BPF from web to node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /suricata/:nid/bpf [put]
// @router /:nid/suricata/bpf [put]
// @router /bpf/:nid [put]
func (n *NodeController) PutSuricataBPF() { 
    logs.Info("PUT SuricataBPF -> In")
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    //err := models.AddNode(anode)
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"status": "false", "error": "There is no BPF"}
    if nid != "" {
        data,err := models.PutSuricataBPF(anode)
        n.Data["json"] = map[string]string{"bpf": data}
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }
    }
    logs.Info("PUT SuricataBPF -> OUT -> %s", n.Data["json"])
    n.ServeJSON()
}



// @Title Get Zeek
// @Description Get Zeek status from Node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /zeek/:nid [get]
// @router /:nid/zeek [get]
func (n *NodeController) GetZeek() { 
    logs.Info("GET Zeek -> In")
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"status": "false", "error": "No hay NID"}
    if nid != "" {
        data,err := models.Zeek(nid)
        var anode map[string]string
        json.Unmarshal(data, &anode)
        n.Data["json"] = anode
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }
    }
    logs.Info("GET Zeek -> OUT -> %s", n.Data["json"])
    n.ServeJSON()
}

// @Title Get Wazuh
// @Description Get wazuh status from Node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /wazuh/:nid [get]
// @router /:nid/wazuh [get]
func (n *NodeController) GetWazuh() { 
    logs.Info("GET Wazuh -> In")
    nid := n.GetString(":nid")
    n.Data["json"] = map[string]string{"status": "false", "error": "No hay NID"}
    if nid != "" {
        data,err := models.Wazuh(nid)
        var anode map[string]string
        json.Unmarshal(data, &anode)
        n.Data["json"] = anode
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }
    }
    logs.Info("GET Wazuh -> OUT -> %s", n.Data["json"])
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
    logs.Info("Dentro de GetNodeFile")
    anode := make(map[string]string)
    anode["uuid"] = n.GetString(":uuid")
    anode["file"] = n.GetString(":fileName")

    logs.Error("JSON DESDE GET "+anode["uuid"]+"  ++++++++"+anode["file"]);
    returnData,err := models.GetNodeFile(anode)
    logs.Error(returnData);

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
    logs.Info("Dentro de SetNodeFile")
    anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    logs.Warn(anode["uuid"])
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
    logs.Info("Vuelto getAllFiles")
    n.Data["json"] = data
    logs.Info(data)
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
    logs.Info("Inside RunSuricata")
    uuid := n.GetString(":uuid")
    data, err := models.RunSuricata(uuid)
    logs.Info("Back RunSuricata")
    n.Data["json"] = data
    logs.Info(data)
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
    logs.Info("Back StopSuricata")
    n.Data["json"] = data
    logs.Info(data)
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}