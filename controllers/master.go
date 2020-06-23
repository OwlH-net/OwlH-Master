package controllers

import (
    "owlhmaster/models"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "owlhmaster/validation"
    "encoding/json"
    "strings"
    "encoding/base64"
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
    data, err := models.GetMasterTitle(n.Ctx.Input.Header("user"))
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
    logs.Info("CHECKED FLAG!!")
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetFileContent"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)   
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        file := n.GetString(":file")
        data, err := models.GetFileContent(file, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"SaveFileContent"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveFileContent(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"PingPlugins"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data,err := models.PingPlugins(n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"PingFlow"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data,err := models.PingFlow(n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"ChangePluginStatus"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangePluginStatus(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"SaveStapInterface"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveStapInterface(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"ChangeDataflowStatus"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangeDataflowStatus(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetNetworkInterface"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.GetNetworkInterface(n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"DeployMaster"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeployMaster(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"UpdateMasterNetworkInterface"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.UpdateMasterNetworkInterface(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"LoadMasterNetworkValuesSelected"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.LoadMasterNetworkValuesSelected(n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"PingServiceMaster"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        err := models.PingServiceMaster(n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"DeployServiceMaster"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        err := models.DeployServiceMaster(n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"AddPluginServiceMaster"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddPluginServiceMaster(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"DeleteServiceMaster"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteServiceMaster(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"ModifyStapValuesMaster"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ModifyStapValuesMaster(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"UpdateMasterStapInterface"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.UpdateMasterStapInterface(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            logs.Error("UpdateMasterStapInterface -> error: %s", err.Error())
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title SetBPF
// @Description Set new STAP BPF
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /setbpf [put]
func (n *MasterController) SetBPF() { 
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"SetBPF"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SetBPF(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"DeployStapServiceMaster"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.DeployStapServiceMaster(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"StopStapServiceMaster"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.StopStapServiceMaster(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetIncidents"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data,err := models.GetIncidents(n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"SaveZeekValues"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.SaveZeekValues(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"PingPluginsMaster"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.PingPluginsMaster(n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetPathFileContent"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        param := n.GetString(":param")
        data, err := models.GetPathFileContent(param, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"SaveFilePathContent"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.SaveFilePathContent(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title Login
// @Description Get master token
// @Success 200 {object} models.Master
// @router /login [put]
func (n *MasterController) Login() {  
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    anode := make(map[string]string)
    
    if n.Ctx.Input.Header("Authorization") != "" {     

        code := strings.Replace(n.Ctx.Input.Header("Authorization"), "Basic ", "", -1)
        data, err := base64.StdEncoding.DecodeString(code)
        userPass := strings.Split(string(data), ":")
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }else{   
            anode["user"] = userPass[0]
            anode["password"] = userPass[1]
            token, err := models.Login(anode, anode["user"])
            n.Data["json"] = token
            if err != nil {
                n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
            }
        }
    }else{
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        token, err := models.Login(anode, anode["user"])
        n.Data["json"] = token
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}

// @Title Auth
// @Description Get master token
// @Success 200 {object} models.Master
// @router /auth [put]
func (n *MasterController) Auth() {  
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    //THIS FUNCTION DOES NOT REQUIRE TOKEN VALIDATE
    anode := make(map[string]string)
    
    if n.Ctx.Input.Header("Authorization") != "" {     

        code := strings.Replace(n.Ctx.Input.Header("Authorization"), "Basic ", "", -1)
        data, err := base64.StdEncoding.DecodeString(code)
        userPass := strings.Split(string(data), ":")
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }else{   
            anode["user"] = userPass[0]
            anode["password"] = userPass[1]
            token, err := models.Login(anode, anode["user"])
            // n.Data["json"] = token
            n.Data["json"] = map[string]string{"ack": "true", "token": token}
            if err != nil {
                n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
            }
        }
    }else{
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        token, err := models.Login(anode, anode["user"])
        // n.Data["json"] = token
        n.Data["json"] = map[string]string{"ack": "true", "token": token}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    
    n.ServeJSON()
}

// @Title AddUser
// @Description Add new user
// @Success 200 {object} models.Master
// @router /addUser [post]
func (n *MasterController) AddUser() {  
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"AddUser"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddUser(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetAllUsers"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.GetAllUsers(n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"DeleteUser"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteUser(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"AddGroupUsers"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        
        err := models.AddGroupUsers(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"AddRole"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.AddRole(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetRolesForUser"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        id := n.GetString(":uuid")
        data, err := models.GetRolesForUser(id, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetGroupsForUser"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        id := n.GetString(":uuid")
        data, err := models.GetGroupsForUser(id, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"AddUsersTo"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.AddUsersTo(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"ChangePassword"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.ChangePassword(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"DeleteUserRole"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteUserRole(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"DeleteUserGroup"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteUserGroup(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetAllRoles"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.GetAllRoles(n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"DeleteRole"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteRole(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"EditRole"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.EditRole(anode, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetAllUserGroups"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.GetAllUserGroups(n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"EditUserGroup"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.EditUserGroup(anode, n.Ctx.Input.Header("user"))
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
//     permissions,err := validation.CheckToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"), n.Ctx.Input.Header("uuid"),"delete")
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetRolesForGroups"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        id := n.GetString(":uuid")
        data, err := models.GetRolesForGroups(id, n.Ctx.Input.Header("user"))
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
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"AddRoleToGroup"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.AddRoleToGroup(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteRoleUser
// @Description delete a specific user at role
// @router /deleteRoleUser [delete]
func (n *MasterController) DeleteRoleUser() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"DeleteRoleUser"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteRoleUser(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteRoleGroup
// @Description delete a specific group at role
// @router /deleteRoleGroup [delete]
func (n *MasterController) DeleteRoleGroup() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"DeleteRoleGroup"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteRoleGroup(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteGroupUser
// @Description delete a specific user at group
// @router /deleteGroupUser [delete]
func (n *MasterController) DeleteGroupUser() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"DeleteGroupUser"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteGroupUser(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title DeleteGroupRole
// @Description delete a specific user at group
// @router /deleteGroupRole [delete]
func (n *MasterController) DeleteGroupRole() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"DeleteGroupRole"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
        err := models.DeleteGroupRole(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetPermissions
// @Description Get all role permissions
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /getPermissions [get]
func (n *MasterController) GetPermissions() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetPermissions"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.GetPermissions(n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetPermissionsByRole
// @Description Get all role permissions
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /getPermissionsByRole/:uuid [get]
func (n *MasterController) GetPermissionsByRole() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetPermissionsByRole"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        uuid := n.GetString(":uuid")
        data, err := models.GetPermissionsByRole(uuid, n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title AddNewRole
// @Description Add role for users
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /addNewRole [put]
func (n *MasterController) AddNewRole() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"AddNewRole"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.AddNewRole(anode, n.Ctx.Input.Header("user"))
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title GetAllGroupRulesetsForAllNodes
// @Description Get all role permissions
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /getAllGroupRulesetsForAllNodes [get]
func (n *MasterController) GetAllGroupRulesetsForAllNodes() {
    errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"GetAllGroupRulesetsForAllNodes"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{
        data, err := models.GetAllGroupRulesetsForAllNodes(n.Ctx.Input.Header("user"))
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}