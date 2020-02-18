package controllers

import (
    "owlhmaster/models"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "owlhmaster/validation"
    "encoding/json"
)

type MasterController struct {
    beego.Controller
}


// @Title GetMasterTitle
// @Description Get title for master
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /getMasterTitle [get]
func (n *MasterController) GetMasterTitle() {
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATION
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATION
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATION
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATION
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATION
    data, err := models.GetMasterTitle()
    n.Data["json"] = data
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }

    n.ServeJSON()
}

// @Title GetFileContent
// @Description Get file content
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /editFile/:file [get]
func (n *MasterController) GetFileContent() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        file := n.GetString(":file")
        data, err := models.GetFileContent(file)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SaveFileContent
// @Description Get file content
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /savefile [put]
func (n *MasterController) SaveFileContent() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveFileContent(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title PingPlugins
// @Description Ping all plugins
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /pingPlugins [get]
func (n *MasterController) PingPlugins() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data,err := models.PingPlugins()
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title PingFlow
// @Description Ping flow data for master
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /pingFlow [get]
func (n *MasterController) PingFlow() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data,err := models.PingFlow()
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ChangePluginStatus
// @Description Change plugin status
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /changePluginStatus [put]
func (n *MasterController) ChangePluginStatus() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangePluginStatus(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SaveStapInterface
// @Description Save new interface from master stap or plugins
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /saveStapInterface [put]
func (n *MasterController) SaveStapInterface() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveStapInterface(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ChangeDataflowStatus
// @Description Change plugin status
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /changeDataflowStatus [put]
func (n *MasterController) ChangeDataflowStatus() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeDataflowStatus(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetNetworkInterface
// @Description Get network values for master
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /interface [get]
func (n *MasterController) GetNetworkInterface() { 
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.GetNetworkInterface()
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }   
    n.ServeJSON()
}

// @Title DeployMaster
// @Description Change plugin status
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /deployMaster [put]
func (n *MasterController) DeployMaster() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeployMaster(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
        }
        
        n.Data["json"] = map[string]string{"ack": "false", "error" : "Can't deploy master plugins"}
    }
    n.ServeJSON()
}

// @Title UpdateMasterNetworkInterface
// @Description Update Master interface value
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /updateMasterNetworkInterface [put]
func (n *MasterController) UpdateMasterNetworkInterface() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.UpdateMasterNetworkInterface(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
        }
        
        n.Data["json"] = map[string]string{"ack": "false", "error" : "Can't deploy master plugins"}
    }
    n.ServeJSON()
}

// @Title LoadMasterNetworkValuesSelected
// @Description Get interface selected by user for master
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /loadMasterNetworkValuesSelected [get]
func (n *MasterController) LoadMasterNetworkValuesSelected() { 
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.LoadMasterNetworkValuesSelected()
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }   
    n.ServeJSON()
}

// @Title PingServiceMaster
// @Description Get Master service status
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /pingservice [get]
func (n *MasterController) PingServiceMaster() { 
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        err := models.PingServiceMaster()
        n.Data["json"] =  map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false","error": err.Error()}
        }
    }   
    n.ServeJSON()
}

// @Title DeployServiceMaster
// @Description Deploy Master service
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /deployservice [put]
func (n *MasterController) DeployServiceMaster() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        err := models.DeployServiceMaster()
        n.Data["json"] =  map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false","error": err.Error()}
        }
    }    
    n.ServeJSON()
}

// @Title AddPluginServiceMaster
// @Description Add new stap service at master
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /add [put]
func (n *MasterController) AddPluginServiceMaster() { 
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddPluginServiceMaster(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            logs.Error("AddPluginServiceMaster -> error: %s", err.Error())
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteServiceMaster
// @Description delete a specific plugin service
// @router /deleteService [delete]
func (n *MasterController) DeleteServiceMaster() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteServiceMaster(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ModifyStapValuesMaster
// @Description delete a specific plugin service
// @router /modifyStapValues [put]
func (n *MasterController) ModifyStapValuesMaster() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ModifyStapValuesMaster(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title UpdateMasterStapInterface
// @Description Update master STAP interfaces
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /updateMasterStapInterface [put]
func (n *MasterController) UpdateMasterStapInterface() { 
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.UpdateMasterStapInterface(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            logs.Error("UpdateMasterStapInterface -> error: %s", err.Error())
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title setBPF
// @Description Set new STAP BPF
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /setbpf [put]
func (n *MasterController) SetBPF() { 
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SetBPF(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            logs.Error("SetBPF -> error: %s", err.Error())
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeployStapServiceMaster
// @Description Deploy a Master STAP service
// @Success 200 {object} models.Master
// @router /deployStapServiceMaster [put]
func (n *MasterController) DeployStapServiceMaster() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.DeployStapServiceMaster(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title StopStapServiceMaster
// @Description Stop a Master STAP service
// @Success 200 {object} models.Master
// @router /stopStapServiceMaster [put]
func (n *MasterController) StopStapServiceMaster() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.StopStapServiceMaster(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetIncidents
// @Description Get incidents for master
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /incidents [get]
func (n *MasterController) GetIncidents() {  
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data,err := models.GetIncidents()
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false","error": err.Error()}
        }
    }  
    n.ServeJSON()
}

// // @Title PutIncident
// // @Description Add new incident
// // @Success 200 {object} models.Master
// // @router /incidents [post]
// func (n *MasterController) PutIncident() {
//     anode := make(map[string]string)
//     json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    
//     err := models.PutIncident(anode)
//     n.Data["json"] = map[string]string{"ack": "true"}

//     if err != nil {
//         n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
//     }
//     n.ServeJSON()
// }

// @Title SaveZeekValues
// @Description Edit Zeek expert values
// @Success 200 {object} models.Master
// @router /zeek/saveZeekValues [put]
func (n *MasterController) SaveZeekValues() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.SaveZeekValues(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
    
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}

// @Title PingPluginsMaster
// @Description Get Master plugins
// @Success 200 {object} models.Master
// @router /pingPlugins [get]
func (n *MasterController) PingPluginsMaster() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.PingPluginsMaster()
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetPathFileContent
// @Description Get file content from path
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /editPathFile/:param [get]
func (n *MasterController) GetPathFileContent() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        param := n.GetString(":param")
        data, err := models.GetPathFileContent(param)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SaveFilePathContent
// @Description Set file content to selected path file
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /savefilePath [put]
func (n *MasterController) SaveFilePathContent() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveFilePathContent(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title Login
// @Description Get Master plugins
// @Success 200 {object} models.Master
// @router /login [put]
func (n *MasterController) Login() {  
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    token, err := models.Login(anode)
    n.Data["json"] = token
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    
    n.ServeJSON()
}

// @Title AddUser
// @Description Add new user
// @Success 200 {object} models.Master
// @router /addUser [post]
func (n *MasterController) AddUser() {  
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "post")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddUser(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}

// @Title GetAllUsers
// @Description Get all users
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /getAllUsers [get]
func (n *MasterController) GetAllUsers() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.GetAllUsers()
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteUser
// @Description delete a specific user
// @router /deleteUser [delete]
func (n *MasterController) DeleteUser() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteUser(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title AddGroupUsers
// @Description Add group for users
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /addGroupUsers [put]
func (n *MasterController) AddGroupUsers() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.AddGroupUsers(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title AddRole
// @Description Add role for users
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /addRole [put]
func (n *MasterController) AddRole() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.AddRole(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetRolesForUser
// @Description Get all user roles
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /getRolesForUser/:uuid [get]
func (n *MasterController) GetRolesForUser() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        id := n.GetString(":uuid")
        data, err := models.GetRolesForUser(id)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetGroupsForUser
// @Description Get all user roles
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /getGroupsForUser/:uuid [get]
func (n *MasterController) GetGroupsForUser() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        id := n.GetString(":uuid")
        data, err := models.GetGroupsForUser(id)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title AddUsersTo
// @Description Add user to a group or role
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /addUsersTo [put]
func (n *MasterController) AddUsersTo() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddUsersTo(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title ChangePassword
// @Description Add user to a group or role
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /changePassword [put]
func (n *MasterController) ChangePassword() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangePassword(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteUserRole
// @Description delete a specific userrole
// @router /deleteUserRole [delete]
func (n *MasterController) DeleteUserRole() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteUserRole(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteUserGroup
// @Description delete a specific usergroup
// @router /deleteUserGroup [delete]
func (n *MasterController) DeleteUserGroup() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteUserGroup(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetAllRoles
// @Description Get all roles
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /getAllRoles [get]
func (n *MasterController) GetAllRoles() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.GetAllRoles()
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteRole
// @Description delete a specific role
// @router /deleteRole [delete]
func (n *MasterController) DeleteRole() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteRole(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title EditRole
// @Description Edit role
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /editRole [put]
func (n *MasterController) EditRole() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.EditRole(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetAllUserGroups
// @Description Get all user roles
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /getAllUserGroups [get]
func (n *MasterController) GetAllUserGroups() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.GetAllUserGroups()
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title EditUserGroup
// @Description Edit user group
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /editUserGroup [put]
func (n *MasterController) EditUserGroup() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.EditUserGroup(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// // @Title DeleteUserGroup
// // @Description delete a specific role
// // @router /deleteUserGroup [delete]
// func (n *MasterController) DeleteUserGroup() {
//     permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "delete")
//     if err != nil {
//         n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
//     }else if !permissions{
//         n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
//     }else{
//         var anode map[string]string
//         json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
//         err := models.DeleteUserGroup(anode)
//         n.Data["json"] = map[string]string{"ack": "true"}
//         if err != nil {
//             n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
//         }
//     }
//     n.ServeJSON()
// }


// @Title GetRolesForGroups
// @Description Get all user roles
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /getRolesForGroups/:uuid [get]
func (n *MasterController) GetRolesForGroups() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "get")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        id := n.GetString(":uuid")
        data, err := models.GetRolesForGroups(id)
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title AddRoleToGroup
// @Description Add role for users
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /addRoleToGroup [put]
func (n *MasterController) AddRoleToGroup() {
    permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"), "put")
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else if !permissions{
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.AddRoleToGroup(anode)
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}