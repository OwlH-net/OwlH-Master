package hwaddmng

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "owlhmaster/nodeclient"
)

func AddMacIp(uuid string, anode map[string]string) (err error) {
	err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("AddMacIp Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/AddMacIp ERROR Obtaining Port and Ip: "+err.Error()); return err}

    err = nodeclient.AddMacIp(ipnid,portnid, anode)
    if err != nil { logs.Error("node/AddMacIp ERROR http data request: "+err.Error()); return err}
	
    return err
}

func LoadConfig(uuid string, anode map[string]string) (data map[string]string, err error) {
	err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("LoadConfig Error loading node token: %s",err); return nil, err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/LoadConfig ERROR Obtaining Port and Ip: "+err.Error()); return nil, err}

    data,err = nodeclient.LoadConfig(ipnid,portnid, anode)
	if err != nil { logs.Error("node/LoadConfig ERROR http data request: "+err.Error()); return nil, err}
	
	return data,err
}

func Config(uuid string, anode map[string]interface{}) (err error) {
	err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Config Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/Config ERROR Obtaining Port and Ip: "+err.Error()); return err}

    err = nodeclient.Config(ipnid,portnid, anode)
	if err != nil { logs.Error("node/Config ERROR http data request: "+err.Error()); return err}
	
	return err
}

func Db(uuid string, anode map[string]string) (err error) {
	err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Db Error loading node token: %s",err); return err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/Db ERROR Obtaining Port and Ip: "+err.Error()); return err}

    err = nodeclient.Db(ipnid,portnid, anode)
	if err != nil { logs.Error("node/Db ERROR http data request: "+err.Error()); return err}
	
	return err
}

func ConfigGet(uuid string) (values map[string]string, err error) {
	err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Db Error loading node token: %s",err); return nil,err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/Db ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}

    data, err := nodeclient.ConfigGet(ipnid,portnid)
	if err != nil { logs.Error("node/Db ERROR http data request: "+err.Error()); return nil,err}
	
	return data,err
}