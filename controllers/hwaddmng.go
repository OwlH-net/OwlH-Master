package controllers

import (
    "encoding/json"
    "github.com/astaxie/beego"
    // "github.com/astaxie/beego/logs"
    "owlhmaster/models"
    "owlhmaster/validation"
)

type HwaddmngController struct {
    beego.Controller
}

// @Title AddMacIp
// @Description Add MAC and IP to Knownmacs
// @router /:uuid [post]
func (n *HwaddmngController) AddMacIp() {
	errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"AddMacIp"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{      
		id := n.GetString(":uuid")   
        var anode map[string]string
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        err := models.AddMacIp(id, anode, n.Ctx.Input.Header("user"))
    
        n.Data["json"] = map[string]string{"ack": "true"}
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }
    n.ServeJSON()
}

// @Title LoadConfig
// @Description Locad ARPConfig
// @router /config/:uuid [put]
func (n *HwaddmngController) LoadConfig() {
	errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"LoadConfig"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{         
		id := n.GetString(":uuid")
		anode := make(map[string]string)
        json.Unmarshal(n.Ctx.Input.RequestBody, &anode)

        data,err := models.LoadConfig(id, anode, n.Ctx.Input.Header("user"))
		if err != nil {
			n.Data["json"] = map[string]string{"ack": "false","error":err.Error()}
		}else{
			n.Data["json"] = data
		}
    }
    n.ServeJSON()
}

// @Title Config
// @Description Config ARPConfig
// @router /config/:uuid [post]
func (n *HwaddmngController) Config() {
	errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"ConfigPost"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{         
        anodeIface := make(map[string]interface{})
        json.Unmarshal(n.Ctx.Input.RequestBody, &anodeIface)
		id := n.GetString(":uuid")
        
        err := models.Config(id, anodeIface, n.Ctx.Input.Header("user"))
		if err != nil {
			n.Data["json"] = map[string]string{"ack": "false","error":err.Error()}
		}else{
			n.Data["json"] = map[string]string{"ack": "true"}
		}
        
    }
    n.ServeJSON()
}

// @Title DbManagement
// @Description DB Management
// @router /db/:uuid [post]
func (n *HwaddmngController) Db() {
	errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"Db"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{         
        anodeIface := make(map[string]string)
		id := n.GetString(":uuid")
        json.Unmarshal(n.Ctx.Input.RequestBody, &anodeIface)
        
		err := models.Db(id,anodeIface, n.Ctx.Input.Header("user"))
		if err != nil {
			n.Data["json"] = map[string]string{"ack": "false","error":err.Error()}
		}else{
			n.Data["json"] = map[string]string{"ack": "true"}
		}
    }
    n.ServeJSON()
}

// @Title Config
// @Description Config ARPConfig
// @router /config/:uuid [get]
func (n *HwaddmngController) GetConfig() {
	errToken := validation.VerifyToken(n.Ctx.Input.Header("token"), n.Ctx.Input.Header("user"))
    if errToken != nil {
        n.Data["json"] = map[string]string{"ack": "false", "error": errToken.Error(), "token":"none"}
        n.ServeJSON()
        return
    }    
    permissions := []string{"ConfigGet"}
    hasPermission,permissionsErr := validation.VerifyPermissions(n.Ctx.Input.Header("user"), "any", permissions)    
    if permissionsErr != nil || hasPermission == false {
        n.Data["json"] = map[string]string{"ack": "false","permissions":"none"}
    }else{         
		id := n.GetString(":uuid")
        
        data,err := models.ConfigGet(id,n.Ctx.Input.Header("user"))
        if err != nil {
			n.Data["json"] = map[string]string{"ack": "false","error":err.Error()}
		}else{
			n.Data["json"] = data
		}
    }
    n.ServeJSON()
}