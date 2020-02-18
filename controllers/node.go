package controllers

import (
    "owlhmaster/models"    
    "owlhmaster/validation"    
    // "owlhmaster/utils"    
    "encoding/json"
    // "jwt"
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
//             logs.Info("NODE -> Get Node Detail -> error: %s", err.Error())
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
    // permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "post")
    // if err != nil {
    //     n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    // }else{
        var anode map[string]interface{}
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if _,ok := anode["bulkmode"]; ok{
            var bulk map[string][]map[string]string
            json.Unmarshal(n.Ctx.Input.RequestBody, &bulk)
    
            for x := range bulk["newnodes"]{
                models.AddNode(bulk["newnodes"][x])
            }
        }else{
            var node map[string]string
            var err error
            json.Unmarshal(n.Ctx.Input.RequestBody, &node)
        
            err = models.AddNode(node)
    
            if err != nil {
                logs.Error("NODE CREATE -> error: %s", err.Error())
                n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
            }
        }
    // }
    
    n.ServeJSON()
}


// @Title DeployNode
// @Description Deploy Node
// @Success 200 {string} node deployed
// @Failure 403 :nid is empty
// @router /deploy/:nid [post]
// @router /:nid/deploy [post]
func (n *NodeController) DeployNode() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "post")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        logs.Info("NODE DEPLOY -> In")
        nid := n.GetString(":nid")
        n.Data["json"] = map[string]string{"nid": nid, "state":"Success"}
        if nid == "" {
            n.Data["json"] = map[string]string{"nid": "", "state":"Failure"}
            logs.Error("NODE DEPLOY -> error -> No Node ID")
        }
    }
    n.ServeJSON()
}


// @Title UpdateNode
// @Description Update Node
// @Success 200 {string} node updated
// @Failure 403 body is empty
// @router / [put]
func (n *NodeController) UpdateNode() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.UpdateNode(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            logs.Info("NODE UPDATE -> error: %s", err.Error())
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
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
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        logs.Error("Error PingNode Master token: %s",err)
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        nid := n.GetString(":nid")
        // nodeResp, err := models.PingNode(nid)
        _, err := models.PingNode(nid)
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }else{
            n.Data["json"] = map[string]string{"ping": "pong"}
        //     if nodeResp != nil{
        //         n.Data["json"] = map[string]string{"nodeToken": nodeResp["nodeToken"], "error": nodeResp["error"]}
        //     }
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
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        nid := n.GetString(":nid")
        data,err := models.Suricata(nid)
        logs.Warn("GetSuricata")
        logs.Warn(data)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }

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
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        nid := n.GetString(":nid")
        data,err := models.Zeek(nid)
        logs.Warn("GetZeek")
        logs.Warn(data)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }

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
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        nid := n.GetString(":nid")
        n.Data["json"] = map[string]string{"status": "false", "error": "No hay NID"}
        data,err := models.Wazuh(nid)
        logs.Warn("GetWazuh")
        logs.Warn(data)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// // @Title GetSuricata BPF
// // @Description Get Suricata BPF from node
// // @Success 200 {object} models.Node
// // @Failure 403 :nid is empty
// // @router /suricata/:nid/bpf [get]
// // @router /:nid/suricata/bpf [get]
// func (n *NodeController) GetSuricataBPF() {
//     nid := n.GetString(":nid")
//     data,err := models.GetSuricataBPF(nid)
//     n.Data["json"] = data
    
//     if err != nil {
//         logs.Error("Can't get Suricata status" + err.Error())
//         n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
//     }
    
//     n.ServeJSON()
// }


// @Title PutSuricataBPF
// @Description Set BPF to node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /suricata/bpf [put]
func (n *NodeController) PutSuricataBPF() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.PutSuricataBPF(anode)
        n.Data["json"] = map[string]string{"status": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "error": err.Error()}
        }
        
    }
    n.ServeJSON()
}

// @Title GetAllNodes
// @Description Get full list of nodes
// @Success 200 {object} models.Node
// @router / [get]
func (n *NodeController) GetAllNodes() {
    //check token
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        nodes, err := models.GetAllNodes()
        n.Data["json"] = nodes
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title GetServiceStatus
// @Description Get service status for a specific node
// @Success 200 {object} models.Node
// @router /pingservice/:uuid [get]
func (n *NodeController) GetServiceStatus() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        err := models.GetServiceStatus(uuid)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
        
    }
    n.ServeJSON()
}

// @Title DeployService
// @Description Get service status for a specific node
// @Success 200 {object} models.Node
// @router /deployservice/:uuid [put]
func (n *NodeController) DeployService() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        err := models.DeployService(uuid)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
        
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
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        nid := n.Ctx.Input.Param(":nid")
        n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
        if nid != "" {
            err := models.DeleteNode(nid)
            n.Data["json"] = map[string]string{"ack": "true"}
            if err != nil {
                n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
            }
        }

    }
    n.ServeJSON()
}

// @Title SyncRulesetToNode
// @Description Send Ruleset file to node
// @Success 200 {object} models.Node
// @router /ruleset/set [put]
func (n *NodeController) SyncRulesetToNode() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SyncRulesetToNode(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title GetNodeFile
// @Description Get specific file from node and send to webpage
// @Success 200 {object} models.Node
// @router /loadfile/:uuid/:fileName [get]
func (n *NodeController) GetNodeFile() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        logs.Info("Inside GetNodeFile")
        anode := make(map[string]string)
        anode["uuid"] = n.GetString(":uuid")
        anode["file"] = n.GetString(":fileName")
        returnData,err := models.GetNodeFile(anode)
        n.Data["json"] = returnData
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title SetNodeFile
// @Description Set changed data on webpage to specific file in node.
// @Success 200 {object} models.Node
// @router /savefile [put]
func (n *NodeController) SetNodeFile() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SetNodeFile(anode)
    
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title GetAllFiles
// @Description Send Ruleset file to node
// @Success 200 {object} models.Node
// @router /getAllFiles/:uuid [get]
func (n *NodeController) GetAllFiles() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {        
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.GetAllFiles(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title RunSuricata
// @Description Run suricata server
// @Success 200 {object} models.Node
// @router /RunSuricata/:uuid [put]
func (n *NodeController) RunSuricata() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.RunSuricata(uuid)
    
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title StopSuricata
// @Description Stop suricata server
// @Success 200 {object} models.Node
// @router /StopSuricata/:uuid [put]
func (n *NodeController) StopSuricata() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.StopSuricata(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title RunZeek
// @Description Run zeek server
// @Success 200 {object} models.Node
// @router /RunZeek/:uuid [put]
func (n *NodeController) RunZeek() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.RunZeek(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title StopZeek
// @Description Stop zeek server
// @Success 200 {object} models.Node
// @router /StopZeek/:uuid [put]
func (n *NodeController) StopZeek() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.StopZeek(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title RunWazuh
// @Description Run wazuh server
// @Success 200 {object} models.Node
// @router /RunWazuh/:uuid [put]
func (n *NodeController) RunWazuh() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.RunWazuh(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title StopWazuh
// @Description Stop wazuh server
// @Success 200 {object} models.Node
// @router /StopWazuh/:uuid [put]
func (n *NodeController) StopWazuh() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.StopWazuh(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
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
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        nid := n.GetString(":nid")
        n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
        if nid != "" {
            err := models.DeployZeek(nid)
            n.Data["json"] = map[string]string{"ping": "pong", "nid": nid}
            if err != nil {
                n.Data["json"] = map[string]string{"ack": "false", "nid": nid, "error": err.Error()}
            }
        }
    }
    n.ServeJSON()
}

// @Title PingPorts
// @Description Get Ping from ports
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /PingPorts/:nid [get]
// @router /:nid/PingPorts [get]
func (n *NodeController) PingPorts() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        nid := n.GetString(":nid")
        n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
        if nid != "" {
            data, err := models.PingPorts(nid)
            n.Data["json"] = data
            if err != nil {
                n.Data["json"] = map[string]string{"ack": "false", "nid": nid, "error": err.Error()}
            }
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
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        nid := n.GetString(":nid")
        n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
        if nid != "" {
            data, err := models.ShowPorts(nid)
            n.Data["json"] = data
            if err != nil {
                n.Data["json"] = map[string]string{"ack": "false", "nid": nid, "error": err.Error()}
            }
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
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    
        err := models.DeletePorts(anode,uuid)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "nid": uuid, "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title DeleteAllPorts
// @Description delete all ports from knownports
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /ports/deleteAll/:uuid [put]
func (n *NodeController) DeleteAllPorts() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        err := models.DeleteAllPorts(uuid)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
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
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeMode(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ChangeStatus
// @Description put Pong from Node
// @Success 200 {object} models.Node
// @Failure 403 Error
// @router /status [put]
func (n *NodeController) ChangeStatus() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
        err := models.ChangeStatus(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title PingPluginsNode
// @Description Get Ping from plugins
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /PingPluginsNode/:nid [get]
// @router /:nid/PingPluginsNode [get]
func (n *NodeController) PingPluginsNode() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        nid := n.GetString(":nid")
        data, err := models.PingPluginsNode(nid)
        n.Data["json"] = data
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "nid": nid, "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}

// @Title GetMainconfData
// @Description Get Ping from ports
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /getMainconfData/:nid [get]
// @router /:nid/getMainconfData [get]
func (n *NodeController) GetMainconfData() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        nid := n.GetString(":nid")
        data, err := models.GetMainconfData(nid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
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
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.PingAnalyzer(uuid)
    
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title ChangeAnalyzerStatus
// @Description Ping node analyzer
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /analyzer [put]
func (n *NodeController) ChangeAnalyzerStatus() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeAnalyzerStatus(anode)
        n.Data["json"] = map[string]string{"ack": "true", "uuid": anode["uuid"]}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": anode["uuid"], "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title Deploy
// @Description deploy node elements
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /deployNode [put]
func (n *NodeController) Deploy() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeployNode(anode)
        n.Data["json"] = map[string]string{"ack": "true", "uuid": anode["uuid"]}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": anode["uuid"], "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title CheckDeploy
// @Description deploy node elements
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /checkDeploy/:uuid [get]
func (n *NodeController) CheckDeploy() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        anode := models.CheckDeploy(uuid)
        n.Data["json"] = anode
    }
    n.ServeJSON()
}

// @Title ChangeDataflowValues
// @Description Change node data flow values
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /changeDataflowValues [put]
func (n *NodeController) ChangeDataflowValues() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeDataflowValues(anode)
        n.Data["json"] = map[string]string{"ack": "true", "uuid": anode["uuid"]}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": anode["uuid"], "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title LoadDataflowValues
// @Description Load node data flow values
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /loadDataflowValues/:uuid [get]
func (n *NodeController) LoadDataflowValues() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.LoadDataflowValues(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title LoadNetworkValues
// @Description Load network data from network values
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /loadNetworkValues/:uuid [get]
func (n *NodeController) LoadNetworkValues() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.LoadNetworkValues(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title LoadNetworkValuesSelected
// @Description Load network data from network values
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /loadNetworkValuesSelected/:uuid [get]
func (n *NodeController) LoadNetworkValuesSelected() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.LoadNetworkValuesSelected(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title UpdateNetworkInterface
// @Description Get network values for master
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /updateNetworkInterface [put]
func (n *NodeController) UpdateNetworkInterface() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.UpdateNetworkInterface(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false: " + err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SaveSocketToNetwork
// @Description save socket to network information
// @Success 200 {object} models.Node
// @router /saveSocketToNetwork [put]
func (n *NodeController) SaveSocketToNetwork() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveSocketToNetwork(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}


// @Title SaveNewLocal
// @Description save New local into dataflow at node
// @Success 200 {object} models.Node
// @router /saveNewLocal [put]
func (n *NodeController) SaveNewLocal() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveNewLocal(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SaveVxLAN
// @Description save VxLAN into dataflow at node
// @Success 200 {object} models.Node
// @router /saveVxLAN [put]
func (n *NodeController) SaveVxLAN() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveVxLAN(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SocketToNetworkList
// @Description Load node data flow values
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /socketToNetworkList/:uuid [get]
func (n *NodeController) SocketToNetworkList() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.SocketToNetworkList(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}

// @Title SaveSocketToNetworkSelected
// @Description Save socket to network selected
// @Success 200 {object} models.Node
// @router /saveSocketToNetworkSelected [put]
func (n *NodeController) SaveSocketToNetworkSelected() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveSocketToNetworkSelected(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteDataFlowValueSelected
// @Description Delete dataflow value selected
// @Success 200 {object} models.Node
// @router /deleteDataFlowValueSelected [delete]
func (n *NodeController) DeleteDataFlowValueSelected() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteDataFlowValueSelected(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetNodeMonitor
// @Description Get node monitor status
// @Success 200 {object} models.Node
// @router /pingmonitor/:uuid [get]
func (n *NodeController) GetNodeMonitor() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.GetNodeMonitor(uuid)
        n.Data["json"] = data
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title AddPluginService
// @Description Add new Suricata service
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /add [post]
func (n *NodeController) AddPluginService() { 
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "post")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddPluginService(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            logs.Info("AddPluginService -> error: %s", err.Error())
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ChangeServiceStatus
// @Description Change a service status
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /ChangeServiceStatus [put]
func (n *NodeController) ChangeServiceStatus() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeServiceStatus(anode)
        n.Data["json"] = map[string]string{"ack": "true", "uuid": anode["uuid"]}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": anode["uuid"], "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title ChangeMainServiceStatus
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /mainconfStatus [put]
func (n *NodeController) ChangeMainServiceStatus() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.ChangeMainServiceStatus(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteService
// @Description delete service
// @Success 200 {object} models.Node
// @router /deleteService [delete]
func (n *NodeController) DeleteService() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.DeleteService(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SaveSuricataInterface
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /saveSuricataInterface [put]
func (n *NodeController) SaveSuricataInterface() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.SaveSuricataInterface(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeployStapService
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /deployStapService [put]
func (n *NodeController) DeployStapService() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.DeployStapService(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title StopStapService
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /stopStapService [put]
func (n *NodeController) StopStapService() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.StopStapService(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ModifyStapValues
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /modifyStapValues [put]
func (n *NodeController) ModifyStapValues() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.ModifyStapValues(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title PingWazuhFiles
// @Description Get Wazuh files
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /pingWazuhFiles/:uuid [get]
func (n *NodeController) PingWazuhFiles() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data,err := models.PingWazuhFiles(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[int]map[string]string{0:{"ack": "false", "uuid": uuid, "error": err.Error()}}
        }
    }
    n.ServeJSON()
}

// @Title DeleteWazuhFile
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /deleteWazuhFile [delete]
func (n *NodeController) DeleteWazuhFile() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]interface{})
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.DeleteWazuhFile(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title AddWazuhFile
// @Description Add wazuh files
// @Success 200 {object} models.Node
// @router /addWazuhFile [put]
func (n *NodeController) AddWazuhFile() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]interface{})
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.AddWazuhFile(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title LoadFileLastLines
// @Description Add wazuh files
// @Success 200 {object} models.Node
// @router /wazuh/loadLines [put]
func (n *NodeController) LoadFileLastLines() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        data,err := models.LoadFileLastLines(anode)
        n.Data["json"] = data
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
        
    }
    n.ServeJSON()
}

// @Title SaveFileContentWazuh
// @Description save wazuh file content
// @Success 200 {object} models.Node
// @router /wazuh/saveFileContentWazuh [put]
func (n *NodeController) SaveFileContentWazuh() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.SaveFileContentWazuh(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}

// @Title ReloadFilesData
// @Description get files data for wazuh and analyzer
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /reloadFilesData/:uuid [get]
// @router /:nid/reloadFilesData [get]
func (n *NodeController) ReloadFilesData() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.ReloadFilesData(uuid)
    
        n.Data["json"] = data
        if err != nil {n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}}
    }

    n.ServeJSON()
}

// @Title AddMonitorFile
// @Description Add file to node monitor
// @Success 200 {object} models.Node
// @Failure 403 body is empty
// @router /monitor/addFile [post]
func (n *NodeController) AddMonitorFile() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "post")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddMonitorFile(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title PingMonitorFiles
// @Description get files and their data from monitor
// @Success 200 {object} models.Node
// @Failure 403 :uuid is empty
// @router /monitor/pingMonitorFiles/:uuid [get]
// @router /monitor/:nid/pingMonitorFiles [get]
func (n *NodeController) PingMonitorFiles() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.PingMonitorFiles(uuid)
    
        n.Data["json"] = data
        if err != nil {n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}}
    }

    n.ServeJSON()
}

// @Title DeleteMonitorFile
// @Description delete file from node monitor
// @Success 200 {object} models.Node
// @Failure 403 body is empty
// @router /monitor/deleteFile [delete]
func (n *NodeController) DeleteMonitorFile() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteMonitorFile(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ChangeZeekMode
// @Description Change Zeek mode
// @Success 200 {object} models.Node
// @router /zeek/changeZeekMode [put]
func (n *NodeController) ChangeZeekMode() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.ChangeZeekMode(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}

// @Title AddClusterValue
// @Description Add cluster value for Zeek service
// @Success 200 {object} models.Node
// @Failure 403 body is empty
// @router /addClusterValue [post]
func (n *NodeController) AddClusterValue() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "post")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddClusterValue(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title PingCluster
// @Description Get all Zeek cluster elements
// @Success 200 {object} models.Node
// @Failure 403 body is empty
// @router /pingCluster/:uuid [get]
func (n *NodeController) PingCluster() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data,err := models.PingCluster(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title EditClusterValue
// @Description Edit cluster value
// @Success 200 {object} models.Node
// @router /zeek/editClusterValue [put]
func (n *NodeController) EditClusterValue() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.EditClusterValue(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}

// @Title DeleteClusterValue
// @Description Delete cluster value
// @Success 200 {object} models.Node
// @router /zeek/deleteClusterValue [delete]
func (n *NodeController) DeleteClusterValue() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.DeleteClusterValue(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}

// @Title SyncCluster
// @Description Sync Zeek cluster
// @Success 200 {object} models.Node
// @router /zeek/syncCluster [put]
func (n *NodeController) SyncCluster() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SyncCluster(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetChangeControlNode
// @Description Get changeControl database values from node
// @Success 200 {object} models.Collector
// @Failure 403 body is empty
// @router /changecontrol/:uuid [get]
func (n *NodeController) GetChangeControlNode() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.GetChangeControlNode(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetIncidentsNode
// @Description Get incidents for node
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /incidents/:uuid [get]
func (n *NodeController) GetIncidentsNode() {   
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data,err := models.GetIncidentsNode(uuid)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false","error": err.Error()}
        }
    } 
    n.ServeJSON()
}

// // @Title PutIncidentNode
// // @Description Add new incident at node
// // @Success 200 {object} models.Master
// // @router /incidents [post]
// func (n *NodeController) PutIncidentNode() {
//     anode := make(map[string]string)
//     json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    
//     err := models.PutIncidentNode(anode)
//     n.Data["json"] = map[string]string{"ack": "true"}

//     if err != nil {
//         n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
//     }
//     n.ServeJSON()
// }

// @Title ChangeSuricataTable
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /plugin/changeSuricataTable [put]
func (n *NodeController) ChangeSuricataTable() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.ChangeSuricataTable(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SyncRulesetToAllGroupNodes
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /ruleset/syncGroups [put]
func (n *NodeController) SyncRulesetToAllGroupNodes() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.SyncRulesetToAllGroupNodes(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SyncAnalyzerToAllGroupNodes
// @Description synchronize analyzer to all nodes
// @Success 200 {object} models.Node
// @router /analyzer/sync [put]
func (n *NodeController) SyncAnalyzerToAllGroupNodes() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        log,err := models.SyncAnalyzerToAllGroupNodes(anode)
        n.Data["json"] = log
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title StartSuricataMainConf
// @Description Start Suricata Main Conf
// @Success 200 {object} models.Node
// @router /StartSuricataMain [put]
func (n *NodeController) StartSuricataMainConf() { 
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.StartSuricataMainConf(anode)
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}
// @Title StopSuricataMainConf
// @Description Stop Suricata Main Conf
// @Success 200 {object} models.Node
// @router /StopSuricataMain [put]
func (n *NodeController) StopSuricataMainConf() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.StopSuricataMainConf(anode)
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}
// @Title KillSuricataMainConf
// @Description Kill Suricata Main Conf
// @Success 200 {object} models.Node
// @router /KillSuricataMain [put]
func (n *NodeController) KillSuricataMainConf() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.KillSuricataMainConf(anode)
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}
// @Title ReloadSuricataMainConf
// @Description Kill Suricata Main Conf
// @Success 200 {object} models.Node
// @router /ReloadSuricataMain [put]
func (n *NodeController) ReloadSuricataMainConf() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ReloadSuricataMainConf(anode)
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title LaunchZeekMainConf
// @Description Launch Zeek option through Main.conf
// @Success 200 {object} models.Node
// @router /LaunchZeekMainConf [put]
func (n *NodeController) LaunchZeekMainConf() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.LaunchZeekMainConf(anode)
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SyncZeekValues
// @Description Sync Zeek cluster
// @Success 200 {object} models.Node
// @router /zeek/syncZeekValues [put]
func (n *NodeController) SyncZeekValues() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SyncZeekValues(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}