package controllers

import (
	"owlhmaster/models"
	"github.com/astaxie/beego"
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