package controllers

import (
    "owlhmaster/models"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
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
func (m *MasterController) GetMasterTitle() {
    data, err := models.GetMasterTitle()
    m.Data["json"] = data
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false: " + err.Error()}
    }
    m.ServeJSON()
}

// @Title GetFileContent
// @Description Get file content
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /editFile/:file [get]
func (m *MasterController) GetFileContent() {
    file := m.GetString(":file")
    data, err := models.GetFileContent(file)
    m.Data["json"] = data
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false: " + err.Error()}
    }
    m.ServeJSON()
}

// @Title SaveFileContent
// @Description Get file content
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /savefile [put]
func (m *MasterController) SaveFileContent() {
    anode := make(map[string]string)
    json.Unmarshal(m.Ctx.Input.RequestBody, &anode)
    err := models.SaveFileContent(anode)
    m.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false: " + err.Error()}
    }
    m.ServeJSON()
}

// @Title PingPlugins
// @Description Ping all plugins
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /pingPlugins [get]
func (m *MasterController) PingPlugins() {
    data,err := models.PingPlugins()
    m.Data["json"] = data
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
    }
    m.ServeJSON()
}

// @Title PingFlow
// @Description Ping flow data for master
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /pingFlow [get]
func (m *MasterController) PingFlow() {
    data,err := models.PingFlow()
    m.Data["json"] = data
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
    }
    m.ServeJSON()
}

// @Title ChangePluginStatus
// @Description Change plugin status
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /changePluginStatus [put]
func (m *MasterController) ChangePluginStatus() {
    anode := make(map[string]string)
    json.Unmarshal(m.Ctx.Input.RequestBody, &anode)
    err := models.ChangePluginStatus(anode)
    m.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
    }
    m.ServeJSON()
}

// @Title SaveStapInterface
// @Description Save new interface from master stap or plugins
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /saveStapInterface [put]
func (m *MasterController) SaveStapInterface() {
    anode := make(map[string]string)
    json.Unmarshal(m.Ctx.Input.RequestBody, &anode)
    err := models.SaveStapInterface(anode)
    m.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
    }
    m.ServeJSON()
}

// @Title ChangeDataflowStatus
// @Description Change plugin status
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /changeDataflowStatus [put]
func (m *MasterController) ChangeDataflowStatus() {
    anode := make(map[string]string)
    json.Unmarshal(m.Ctx.Input.RequestBody, &anode)
    err := models.ChangeDataflowStatus(anode)
    m.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
    }
    m.ServeJSON()
}

// @Title GetNetworkInterface
// @Description Get network values for master
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /interface [get]
func (m *MasterController) GetNetworkInterface() {    
    data, err := models.GetNetworkInterface()
    m.Data["json"] = data
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false: " + err.Error()}
    }
    m.ServeJSON()
}

// @Title DeployMaster
// @Description Change plugin status
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /deployMaster [put]
func (m *MasterController) DeployMaster() {
    anode := make(map[string]string)
    json.Unmarshal(m.Ctx.Input.RequestBody, &anode)
    err := models.DeployMaster(anode)
    m.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
    }
    
    m.Data["json"] = map[string]string{"ack": "false", "error" : "Can't deploy master plugins"}
    m.ServeJSON()
}

// @Title UpdateMasterNetworkInterface
// @Description Update Master interface value
// @Param body body models.Master true "body for master content"
// @Success 200 {object} models.Master
// @router /updateMasterNetworkInterface [put]
func (m *MasterController) UpdateMasterNetworkInterface() {
    anode := make(map[string]string)
    json.Unmarshal(m.Ctx.Input.RequestBody, &anode)
    err := models.UpdateMasterNetworkInterface(anode)
    m.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false", "error" : err.Error()}
    }
    
    m.Data["json"] = map[string]string{"ack": "false", "error" : "Can't deploy master plugins"}
    m.ServeJSON()
}

// @Title LoadMasterNetworkValuesSelected
// @Description Get interface selected by user for master
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /loadMasterNetworkValuesSelected [get]
func (m *MasterController) LoadMasterNetworkValuesSelected() {    
    data, err := models.LoadMasterNetworkValuesSelected()
    m.Data["json"] = data
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false: " + err.Error()}
    }
    m.ServeJSON()
}

// @Title PingServiceMaster
// @Description Get Master service status
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /pingservice [get]
func (m *MasterController) PingServiceMaster() {    
    err := models.PingServiceMaster()
    m.Data["json"] =  map[string]string{"ack": "true"}
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false","error": err.Error()}
    }
    m.ServeJSON()
}

// @Title DeployServiceMaster
// @Description Deploy Master service
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /deployservice [put]
func (m *MasterController) DeployServiceMaster() {    
    err := models.DeployServiceMaster()
    m.Data["json"] =  map[string]string{"ack": "true"}
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false","error": err.Error()}
    }
    m.ServeJSON()
}

// @Title AddPluginServiceMaster
// @Description Add new stap service at master
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /add [put]
func (n *MasterController) AddPluginServiceMaster() { 
    anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.AddPluginServiceMaster(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Error("AddPluginServiceMaster -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title DeleteServiceMaster
// @Description delete a specific plugin service
// @router /deleteService [delete]
func (n *MasterController) DeleteServiceMaster() {
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.DeleteServiceMaster(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title ModifyStapValuesMaster
// @Description delete a specific plugin service
// @router /modifyStapValues [put]
func (n *MasterController) ModifyStapValuesMaster() {
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.ModifyStapValuesMaster(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title UpdateMasterStapInterface
// @Description Update master STAP interfaces
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /updateMasterStapInterface [put]
func (n *MasterController) UpdateMasterStapInterface() { 
    anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.UpdateMasterStapInterface(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Error("UpdateMasterStapInterface -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title setBPF
// @Description Set new STAP BPF
// @Success 200 {object} models.ruleset
// @Failure 403 Connection Failure
// @router /setbpf [put]
func (n *MasterController) SetBPF() { 
    anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.SetBPF(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Error("SetBPF -> error: %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title DeployStapServiceMaster
// @Description Deploy a Master STAP service
// @Success 200 {object} models.Master
// @router /deployStapServiceMaster [put]
func (n *MasterController) DeployStapServiceMaster() {
    anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    
    err := models.DeployStapServiceMaster(anode)
    n.Data["json"] = map[string]string{"ack": "true"}

    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title StopStapServiceMaster
// @Description Stop a Master STAP service
// @Success 200 {object} models.Master
// @router /stopStapServiceMaster [put]
func (n *MasterController) StopStapServiceMaster() {
    anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    
    err := models.StopStapServiceMaster(anode)
    n.Data["json"] = map[string]string{"ack": "true"}

    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetIncidents
// @Description Get incidents for master
// @Param       body            body    models.Master     true            "body for master content"
// @Success 200 {object} models.Master
// @router /incidents [get]
func (m *MasterController) GetIncidents() {    
    data,err := models.GetIncidents()
    m.Data["json"] = data
    if err != nil {
        m.Data["json"] = map[string]string{"ack": "false","error": err.Error()}
    }
    m.ServeJSON()
}

// @Title PutIncident
// @Description Add new incident
// @Success 200 {object} models.Master
// @router /incidents [post]
func (n *MasterController) PutIncident() {
    anode := make(map[string]string)
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    
    err := models.PutIncident(anode)
    n.Data["json"] = map[string]string{"ack": "true"}

    if err != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}