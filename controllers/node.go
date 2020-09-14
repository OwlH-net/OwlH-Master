package controllers

import (
    "encoding/json"
    "owlhmaster/models"
    "owlhmaster/utils"
    "owlhmaster/validation"
    // "jwt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
)

type NodeController struct {
    beego.Controller
}

type NodeEnroll struct {
    Node     NodeData     `json:"node"`
    Group    GroupData    `json:"group"`
    Suricata SuricataData `json:"suricata"`
}

type NodeData struct {
    IP           string `json:"ip"`
    Name         string `json:"name"`
    Port         int    `json:"port"`
    NodeUser     string `json:"nodeuser"`
    NodePassword string `json:"nodepassword"`
    Force        bool   `json:"force"`
}

type GroupData struct {
    UUID string `json:"uuid"`
}

type SuricataData struct {
    Interface  string `json:"interface"`
    Bpf        string `json:"bpf"`
    BpfFile    string `json:"bpffile"`
    ConfigFile string `json:"configfile"`
    Ruleset    string `json:"ruleset"`
    Name       string `json:"name"`
    Status     string `json:"status"`
}

// @Title CreateNode
// @Description Create Node
// @Success 200 {object} models.Node
// @Failure 403 body is empty
// @router / [post]
func (n *NodeController) CreateNode() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"CreateNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]interface{}
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        n.Data["json"] = map[string]string{"ack": "true"}

        if _, ok := anode["bulkmode"]; ok {
            var bulk map[string][]map[string]string
            json.Unmarshal(n.Ctx.Input.RequestBody, &bulk)

            for x := range bulk["newnodes"] {
                models.AddNode(bulk["newnodes"][x], n.Ctx.Input.Header("user"))
            }
        } else {
            var node map[string]string
            var err error
            json.Unmarshal(n.Ctx.Input.RequestBody, &node)

            err = models.AddNode(node, n.Ctx.Input.Header("user"))

            if err != nil {
                logs.Error("NODE CREATE -> error: %s", err.Error())
                n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
            }
        }
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"DeployNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        logs.Info("NODE DEPLOY -> In")
        nid := n.GetString(":nid")
        n.Data["json"] = map[string]string{"nid": nid, "state": "Success"}
        if nid == "" {
            n.Data["json"] = map[string]string{"nid": "", "state": "Failure"}
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"UpdateNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.UpdateNode(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            logs.Info("NODE UPDATE -> error: %s", err.Error())
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title PingNode
// @Description Get Pong from Node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /ping/:nid [get]
// @router /:nid/ping [get]
func (n *NodeController) GetPong() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"PingNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nid := n.GetString(":nid")
        // nodeResp, err := models.PingNode(nid)
        _, err := models.PingNode(nid, n.Ctx.Input.Header("user"))
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        } else {
            n.Data["json"] = map[string]string{"ping": "pong"}
            //     if nodeResp != nil{
            //         n.Data["json"] = map[string]string{"nodeToken": nodeResp["nodeToken"], "error": nodeResp["error"]}
            //     }
        }
    }

    n.ServeJSON()
}

// @Title GetSuricata
// @Description Get Suricate status from Node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /suricata/:nid [get]
// @router /:nid/suricata [get]
func (n *NodeController) GetSuricata() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetSuricata"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nid := n.GetString(":nid")
        data, err := models.Suricata(nid, n.Ctx.Input.Header("user"))
        logs.Warn("GetSuricata")
        logs.Warn(data)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title GetZeek
// @Description Get Zeek status from Node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /zeek/:nid [get]
// @router /:nid/zeek [get]
func (n *NodeController) GetZeek() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetZeek"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nid := n.GetString(":nid")
        data, err := models.Zeek(nid, n.Ctx.Input.Header("user"))
        logs.Warn("GetZeek")
        logs.Warn(data)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title Run Zeek Command
// @Description Run zeek command on node
// @Success 200 {object} models.Node
// @router /zeek/:nid/:cmd [put]
// @router /:nid/zeek/:cmd [put]
func (n *NodeController) ZeekCommand() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetZeek"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nid := n.GetString(":nid")
        cmd := n.GetString(":cmd")
        data, err := models.ZeekCommand(nid, cmd, n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"status": "false", "nid": nid, "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title GetWazuh
// @Description Get wazuh status from Node
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /wazuh/:nid [get]
// @router /:nid/wazuh [get]
func (n *NodeController) GetWazuh() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetWazuh"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nid := n.GetString(":nid")
        n.Data["json"] = map[string]string{"status": "false", "error": "No hay NID"}
        data, err := models.Wazuh(nid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"PutSuricataBPF"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.PutSuricataBPF(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetAllNodes"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nodes, err := models.GetAllNodes(n.Ctx.Input.Header("user"))
        returnResult := make(map[string]string)
        jsonData, _ := json.Marshal(nodes)
        returnResult["ack"] = "true"
        returnResult["result"] = string(jsonData)
        if err != nil {
            returnResult["ack"] = "false"
            returnResult["error"] = err.Error()
            // n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
        n.Data["json"] = returnResult
        // n.Data["json"] = nodes
    }

    n.ServeJSON()
}

// @Title GetAllNodes2
// @Description Get full list of nodes
// @Success 200 {object} models.Node
// @router /GetAllNodes2 [get]
func (n *NodeController) GetAllNodes2() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetAllNodes"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nodes, err := models.GetAllNodes(n.Ctx.Input.Header("user"))
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
        n.Data["json"] = nodes
    }

    n.ServeJSON()
}

// @Title GetServiceStatus
// @Description Get service status for a specific node
// @Success 200 {object} models.Node
// @router /pingservice/:uuid [get]
func (n *NodeController) GetServiceStatus() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetServiceStatus"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        err := models.GetServiceStatus(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"DeployService"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        err := models.DeployService(uuid, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }

    }
    n.ServeJSON()
}

// @Title DeleteNode
// @Description Delete a node
// @Param nid path    string  true "The nodeid you want to delete"
// @Failure 403 :nid is empty
// @Success 200 {string} node deleted
// @router /:nid [delete]
func (n *NodeController) DeleteNode() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"DeleteNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nid := n.Ctx.Input.Param(":nid")
        n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
        if nid != "" {
            err := models.DeleteNode(nid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SyncRulesetToNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SyncRulesetToNode(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetNodeFile"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        logs.Info("Inside GetNodeFile")
        anode := make(map[string]string)
        anode["uuid"] = n.GetString(":uuid")
        anode["file"] = n.GetString(":fileName")
        returnData, err := models.GetNodeFile(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SetNodeFile"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SetNodeFile(anode, n.Ctx.Input.Header("user"))

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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetAllFiles"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.GetAllFiles(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"RunSuricata"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.RunSuricata(uuid, n.Ctx.Input.Header("user"))

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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"StopSuricata"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.StopSuricata(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"RunZeek"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.RunZeek(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"StopZeek"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.StopZeek(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"RunWazuh"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.RunWazuh(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"StopWazuh"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.StopWazuh(uuid, n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// // @Title DeployZeek
// // @Description Get Pong from Node
// // @Success 200 {object} models.Node
// // @Failure 403 :nid is empty
// // @router /deploy/:nid [get]
// // @router /:nid/deploy [get]
// func (n *NodeController) DeployZeek() {
//     permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"),"get")
//     if err != nil {
//         n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
//     }else if !permissions{
//         n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
//     }else{
//         nid := n.GetString(":nid")
//         n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
//         if nid != "" {
//             err := models.DeployZeek(nid)
//             n.Data["json"] = map[string]string{"ping": "pong", "nid": nid}
//             if err != nil {
//                 n.Data["json"] = map[string]string{"ack": "false", "nid": nid, "error": err.Error()}
//             }
//         }
//     }
//     n.ServeJSON()
// }

// @Title PingPorts
// @Description Get Ping from ports
// @Success 200 {object} models.Node
// @Failure 403 :nid is empty
// @router /PingPorts/:nid [get]
// @router /:nid/PingPorts [get]
func (n *NodeController) PingPorts() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"PingPorts"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nid := n.GetString(":nid")
        n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
        if nid != "" {
            data, err := models.PingPorts(nid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ShowPorts"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nid := n.GetString(":nid")
        n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
        if nid != "" {
            data, err := models.ShowPorts(nid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"DeletePorts"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.DeletePorts(anode, uuid, n.Ctx.Input.Header("user"))
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
// @router /ports/deleteAll/:uuid [delete]
func (n *NodeController) DeleteAllPorts() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"DeleteAllPorts"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        err := models.DeleteAllPorts(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ChangeMode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeMode(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ChangeStatus"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
        err := models.ChangeStatus(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"PingPluginsNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nid := n.GetString(":nid")
        data, err := models.PingPluginsNode(nid, n.Ctx.Input.Header("user"))
        // data["ack"] = map[string]string{}
        // data["ack"]["ack"] = "true"

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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetMainconfData"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nid := n.GetString(":nid")
        data, err := models.GetMainconfData(nid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"PingAnalyzer"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.PingAnalyzer(uuid, n.Ctx.Input.Header("user"))

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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ChangeAnalyzerStatus"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeAnalyzerStatus(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"Deploy"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeployNode(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"CheckDeploy"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        anode := models.CheckDeploy(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ChangeDataflowValues"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeDataflowValues(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"LoadDataflowValues"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.LoadDataflowValues(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"LoadNetworkValues"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.LoadNetworkValues(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"LoadNetworkValuesSelected"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.LoadNetworkValuesSelected(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"UpdateNetworkInterface"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.UpdateNetworkInterface(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SaveSocketToNetwork"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveSocketToNetwork(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SaveNewLocal"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveNewLocal(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SaveVxLAN"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveVxLAN(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SocketToNetworkList"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.SocketToNetworkList(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SaveSocketToNetworkSelected"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveSocketToNetworkSelected(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"DeleteDataFlowValueSelected"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteDataFlowValueSelected(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetNodeMonitor"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.GetNodeMonitor(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"AddPluginService"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddPluginService(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ChangeServiceStatus"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeServiceStatus(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ChangeMainServiceStatus"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.ChangeMainServiceStatus(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"DeleteService"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.DeleteService(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}

        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title UpdateSuricataValue
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /updateSuricataValue [put]
func (n *NodeController) UpdateSuricataValue() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"UpdateSuricataValue"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.UpdateSuricataValue(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"DeployStapService"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.DeployStapService(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}

        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title StopStapServiceNode
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /stopStapService [put]
func (n *NodeController) StopStapService() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"StopStapServiceNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.StopStapService(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}

        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ModifyNodeOptionValues
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /modifyNodeOptionValues [put]
func (n *NodeController) ModifyNodeOptionValues() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ModifyNodeOptionValues"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.ModifyNodeOptionValues(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"PingWazuhFiles"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.PingWazuhFiles(uuid, n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteWazuhFile
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /deleteWazuhFile [delete]
func (n *NodeController) DeleteWazuhFile() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"DeleteWazuhFile"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]interface{})
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.DeleteWazuhFile(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"AddWazuhFile"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]interface{})
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.AddWazuhFile(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"LoadFileLastLines"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        data, err := models.LoadFileLastLines(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SaveFileContentWazuh"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.SaveFileContentWazuh(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ReloadFilesData"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.ReloadFilesData(uuid, n.Ctx.Input.Header("user"))

        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title AddMonitorFile
// @Description Add file to node monitor
// @Success 200 {object} models.Node
// @Failure 403 body is empty
// @router /monitor/addFile [post]
func (n *NodeController) AddMonitorFile() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"AddMonitorFile"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddMonitorFile(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"PingMonitorFiles"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.PingMonitorFiles(uuid, n.Ctx.Input.Header("user"))

        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "uuid": uuid, "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title DeleteMonitorFile
// @Description delete file from node monitor
// @Success 200 {object} models.Node
// @Failure 403 body is empty
// @router /monitor/deleteFile [delete]
func (n *NodeController) DeleteMonitorFile() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"DeleteMonitorFile"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteMonitorFile(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ChangeZeekMode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.ChangeZeekMode(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"AddClusterValue"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddClusterValue(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"PingCluster"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.PingCluster(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"EditClusterValue"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.EditClusterValue(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"DeleteClusterValue"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.DeleteClusterValue(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SyncCluster"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SyncCluster(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetChangeControlNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.GetChangeControlNode(uuid, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetIncidentsNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")
        data, err := models.GetIncidentsNode(uuid, n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ChangeSuricataTable"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.ChangeSuricataTable(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SyncRulesetToAllGroupNodes"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.SyncRulesetToAllGroupNodes(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SyncAnalyzerToAllGroupNodes"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        log, err := models.SyncAnalyzerToAllGroupNodes(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"StartSuricataMainConf"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.StartSuricataMainConf(anode, n.Ctx.Input.Header("user"))

        n.Data["json"] = map[string]string{"ack": "true"}
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"StopSuricataMainConf"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.StopSuricataMainConf(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"KillSuricataMainConf"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.KillSuricataMainConf(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ReloadSuricataMainConf"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ReloadSuricataMainConf(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"LaunchZeekMainConf"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.LaunchZeekMainConf(anode, n.Ctx.Input.Header("user"))

        n.Data["json"] = map[string]string{"ack": "true"}
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SyncZeekValues"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SyncZeekValues(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ChangeRotationStatus
// @Description Change rotation file status at node
// @Success 200 {object} models.Node
// @router /monitor/changeRotationStatus [put]
func (n *NodeController) ChangeRotationStatus() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"ChangeRotationStatus"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeRotationStatus(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title EditRotation
// @Description Edit rotation parameters at node
// @Success 200 {object} models.Node
// @router /monitor/editRotation [put]
func (n *NodeController) EditRotation() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"EditRotation"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.EditRotation(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetServiceCommands
// @Description Get a servide commands
// @Success 200 {object} models.Node
// @Failure 403 Error
// @router /plugins/getCommands [put]
func (n *NodeController) GetServiceCommands() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetServiceCommands"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        n.Data["json"] = map[string]string{"ack": "false", "error": "No hay NID"}
        data, err := models.GetServiceCommands(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }

    n.ServeJSON()
}

// @Title SaveSurictaRulesetSelected
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /setRuleset [put]
func (n *NodeController) SaveSurictaRulesetSelected() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"SaveSurictaRulesetSelected"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.SaveSurictaRulesetSelected(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}

        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title RegisterNode
// @Description Change mainconf db values
// @Success 200 {object} models.Node
// @router /registerNode/:uuid [put]
func (n *NodeController) RegisterNode() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"RegisterNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        uuid := n.GetString(":uuid")

        err := models.RegisterNode(uuid, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}

        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title AutoEnroll
// @Description auto enroll a Node with extra data like group, suricata service and ruleset
// @Success 200 {object} models.Node
// @Failure 403 body is empty
// @router /enroll [post]
func (n *NodeController) AutoEnroll() {

    n.Data["json"] = map[string]string{"ack": "true"}

    permissions := []string{"CreateNode"}
    canContinue, details := validation.CanContinue(n.Ctx.Input.Header("token"), "any", permissions)
    if !canContinue {
        logs.Error("NODE enrollment -> error: %+v", details)
        n.Data["json"] = map[string]string{"ack": "false", "error": "token invalid or user doesn't have permission"}
    }

    var nodeDetails utils.NodeEnroll
    json.Unmarshal(n.Ctx.Input.RequestBody, &nodeDetails)
    logs.Info("%+v", nodeDetails)

    uuid, enrolled, details := models.NodeEnrollment(nodeDetails.Node)
    if !enrolled {
        logs.Error("NODE enrollment -> error: %+v", details)
        n.Data["json"] = map[string]string{"ack": "false", "error": "There were problems with node enrollment"}
        n.ServeJSON()
        return
    }

    logs.Info("group uuid -> %s", nodeDetails.Group.UUID)
    guuid := ""
    assignedToGroup := false
    if nodeDetails.Group.UUID != "" {
        guuid, assignedToGroup, details = models.AssignNodeToGroup(uuid, nodeDetails.Group)
        if !assignedToGroup {
            logs.Error("NODE assign node to group -> error: %+v", details)
            n.Data["json"] = map[string]string{"ack": "false", "error": "There were problems with node enrollment"}
        }
    }

    logs.Info("create Suricata service for node -> %s and group -> %s", uuid, guuid)
    if nodeDetails.Suricata.Name != "" {
        suricataCreated, details := models.CreateSuricataService(guuid, uuid, nodeDetails.Suricata)
        if !suricataCreated {
            logs.Error("NODE create Suricata Service -> error: %+v", details)
            n.Data["json"] = map[string]string{"ack": "false", "error": "There were problems with node enrollment"}
        }
    }

    n.ServeJSON()
}

// @Title GetAllNodesReact
// @Description Get full list of nodes
// @Success 200 {object} models.Node
// @router /getAllNodesReact [get]
func (n *NodeController) GetAllNodesReact() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetAllNodes"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nodes, err := models.GetAllNodesReact(n.Ctx.Input.Header("user"))
        n.Data["json"] = nodes
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}

// @Title GetAllTags
// @Description Get full list of node tags
// @Success 200 {object} models.Node
// @router /getAllTags [get]
func (n *NodeController) GetAllTags() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetAllTags"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nodes, err := models.GetAllTags(n.Ctx.Input.Header("user"))
        n.Data["json"] = nodes
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}

// @Title EnrollNewNode
// @Description Enroll new Node with extra data
// @Success 200 {string} node deployed
// @Failure 403 :nid is empty
// @router /enrollNewNode [post]
func (n *NodeController) EnrollNewNode() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"CreateNode"}
    // permissions := []string{"EnrollNewNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var nodeDetails utils.EnrollNewNodeStruct
        json.Unmarshal(n.Ctx.Input.RequestBody, &nodeDetails)
        
        err := models.EnrollNewNode(nodeDetails, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title UpdateNodeReact
// @Description Update Node
// @Success 200 {string} node updated
// @Failure 403 body is empty
// @router /updateNodeReact [put]
func (n *NodeController) UpdateNodeReact() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"UpdateNode"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        var anode utils.EnrollNewNodeStruct
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    
        err := models.UpdateNodeReact(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            logs.Info("NODE UPDATE -> error: %s", err.Error())
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetAllOrganizations
// @Description Get full list of node tags
// @Success 200 {object} models.Node
// @router /getAllOrganizations [get]
func (n *NodeController) GetAllOrganizations() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token": "none"}
        n.ServeJSON()
        return
    }
    permissions := []string{"GetAllOrganizations"}
    hasPermission, permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false", "permissions": "none"}
    } else {
        nodes, err := models.GetAllOrganizations(n.Ctx.Input.Header("user"))
        n.Data["json"] = nodes
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}