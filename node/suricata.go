package node

import (
  "errors"

  ndb "github.com/OwlH-net/OwlH-Master/database"
  "github.com/OwlH-net/OwlH-Master/nodeclient"
  "github.com/OwlH-net/OwlH-Master/utils"
  "github.com/astaxie/beego/logs"
)

func Suricata(n string) (data map[string]bool, err error) {
  err = ndb.GetTokenByUuid(n)
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return nil, err
  }
  ip, port, err := ndb.ObtainPortIp(n)
  if err != nil {
    logs.Info("Suricata - get IP and PORT Error -> %s", err.Error())
    return nil, err
  }
  logs.Info("Suricata IP and PORT -> %s, %s", ip, port)
  data, err = nodeclient.Suricata(ip, port)
  if err != nil {
    return nil, err
  }
  return data, nil
}

// func GetSuricataBPF(n string)(bpf string, err error) {
//     if ndb.Db == nil {
//         logs.Error("GetSuricataBPF -- Can't acces to database: "+err.Error())
//         return "", err
//     }
//     err = ndb.GetTokenByUuid(n); if err!=nil{logs.Error("Error loading node token: %s",err); return "",err}
//     ipnid,portnid,err := ndb.ObtainPortIp(n)
//     if err != nil {
//         logs.Error("node/GetSuricataBPF ERROR Obtaining Port and Ip: "+err.Error())
//         return "",err
//     }
//     bpf,err = nodeclient.GetSuricataBPF(ipnid,portnid)
//     if err != nil {
//         logs.Error("nodeclient.GetSuricataBPF error HTTP data request: "+err.Error())
//         return "",err
//     }
//     return bpf, nil
// }

func PutSuricataBPF(n map[string]string) (err error) {
  if ndb.Db == nil {
    logs.Error("PutSuricataBPF -- Can't acces to database: ")
    return errors.New("PutSuricataBPF -- Can't acces to database")
  }

  err = ndb.GetTokenByUuid(n["uuid"])
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return err
  }
  ipnid, portnid, err := ndb.ObtainPortIp(n["uuid"])
  if err != nil {
    logs.Error("node/PutSuricataBPF ERROR Obtaining Port and Ip: " + err.Error())
    return err
  }

  err = nodeclient.PutSuricataBPF(ipnid, portnid, n)
  if err != nil {
    logs.Error("nodeclient.PutSuricataBPF error HTTP data request: " + err.Error())
    return err
  }

  return nil
}

func RunSuricata(uuid string) (data string, err error) {
  if ndb.Db == nil {
    logs.Error("RunSuricata -- Can't acces to database")
    return "", errors.New("RunSuricata -- Can't acces to database")
  }
  err = ndb.GetTokenByUuid(uuid)
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return "", err
  }
  ipnid, portnid, err := ndb.ObtainPortIp(uuid)
  if err != nil {
    logs.Error("node/RunSuricata ERROR Obtaining Port and Ip: " + err.Error())
    return "", err
  }
  data, err = nodeclient.RunSuricata(ipnid, portnid)
  if err != nil {
    logs.Error("node/RunSuricata ERROR http data request: " + err.Error())
    return "", err
  }
  return data, nil
}

func StopSuricata(uuid string) (data string, err error) {
  if ndb.Db == nil {
    logs.Error("StopSuricata -- Can't acces to database")
    return "", errors.New("StopSuricata -- Can't acces to database")
  }
  err = ndb.GetTokenByUuid(uuid)
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return "", err
  }
  ipnid, portnid, err := ndb.ObtainPortIp(uuid)
  if err != nil {
    logs.Error("node/StopSuricata ERROR Obtaining Port and Ip: " + err.Error())
    return "", err
  }
  data, err = nodeclient.StopSuricata(ipnid, portnid)
  if err != nil {
    logs.Error("node/StopSuricata ERROR http data request: " + err.Error())
    return "", err
  }
  return data, nil
}

func AddPluginService(anode map[string]string) (err error) {
  if ndb.Db == nil {
    logs.Error("AddPluginService -- Can't acces to database: ")
    return errors.New("AddPluginService -- Can't acces to database")
  }

  err = ndb.GetTokenByUuid(anode["uuid"])
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return err
  }
  ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
  if err != nil {
    logs.Error("node/AddPluginService ERROR Obtaining Port and Ip: " + err.Error())
    return err
  }

  err = nodeclient.AddPluginService(ipnid, portnid, anode)
  if err != nil {
    logs.Error("nodeclient.AddPluginService error HTTP data request: " + err.Error())
    return err
  }

  return nil
}

func GetSuricataServices(uuid string) (data map[string]map[string]string, err error) {
  if ndb.Db == nil {
    logs.Error("GetSuricataServices -- Can't acces to database: ")
    return nil, errors.New("GetSuricataServices -- Can't acces to database")
  }

  err = ndb.GetTokenByUuid(uuid)
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return nil, err
  }
  ipnid, portnid, err := ndb.ObtainPortIp(uuid)
  if err != nil {
    logs.Error("node/GetSuricataServices ERROR Obtaining Port and Ip: " + err.Error())
    return nil, err
  }

  data, err = nodeclient.GetSuricataServices(ipnid, portnid)
  if err != nil {
    logs.Error("nodeclient.GetSuricataServices error HTTP data request: " + err.Error())
    return nil, err
  }

  return data, nil
}

func UpdateSuricataValue(anode map[string]string) (err error) {
  if ndb.Db == nil {
    logs.Error("UpdateSuricataValue -- Can't acces to database")
    return err
  }

  err = ndb.GetTokenByUuid(anode["uuid"])
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return err
  }
  ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
  if err != nil {
    logs.Error("node/UpdateSuricataValue ERROR Obtaining Port and Ip: " + err.Error())
    return err
  }

  err = nodeclient.UpdateSuricataValue(ipnid, portnid, anode)
  if err != nil {
    logs.Error("node/UpdateSuricataValue ERROR http data request: " + err.Error())
    return err
  }

  return nil
}

func ChangeSuricataTable(anode map[string]string) (err error) {
  if ndb.Db == nil {
    logs.Error("ChangeSuricataTable -- Can't acces to database")
    return err
  }

  err = ndb.GetTokenByUuid(anode["uuid"])
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return err
  }
  ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
  if err != nil {
    logs.Error("node/ChangeSuricataTable ERROR Obtaining Port and Ip: " + err.Error())
    return err
  }

  err = nodeclient.ChangeSuricataTable(ipnid, portnid, anode)
  if err != nil {
    logs.Error("node/ChangeSuricataTable ERROR http data request: " + err.Error())
    return err
  }

  return nil
}

func StartSuricataMainConf(anode map[string]string) (err error) {
  if ndb.Db == nil {
    logs.Error("StartSuricataMainConf -- Can't acces to database: ")
    return errors.New("StartSuricataMainConf -- Can't acces to database")
  }

  err = ndb.GetTokenByUuid(anode["uuid"])
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return err
  }
  ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
  if err != nil {
    logs.Error("node/StartSuricataMainConf ERROR Obtaining Port and Ip: " + err.Error())
    return err
  }

  err = nodeclient.StartSuricataMainConf(ipnid, portnid, anode)
  if err != nil {
    logs.Error("nodeclient.StartSuricataMainConf error HTTP data request: " + err.Error())
    return err
  }

  return nil
}
func StopSuricataMainConf(anode map[string]string) (err error) {
  if ndb.Db == nil {
    logs.Error("StopSuricataMainConf -- Can't acces to database: ")
    return errors.New("StartSuricataMainConf -- Can't acces to database")
  }

  err = ndb.GetTokenByUuid(anode["uuid"])
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return err
  }
  ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
  if err != nil {
    logs.Error("node/StopSuricataMainConf ERROR Obtaining Port and Ip: " + err.Error())
    return err
  }

  err = nodeclient.StopSuricataMainConf(ipnid, portnid, anode)
  if err != nil {
    logs.Error("nodeclient.StopSuricataMainConf error HTTP data request: " + err.Error())
    return err
  }

  return nil
}
func KillSuricataMainConf(anode map[string]string) (err error) {
  if ndb.Db == nil {
    logs.Error("KillSuricataMainConf -- Can't acces to database: ")
    return errors.New("StartSuricataMainConf -- Can't acces to database")
  }

  err = ndb.GetTokenByUuid(anode["uuid"])
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return err
  }
  ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
  if err != nil {
    logs.Error("node/KillSuricataMainConf ERROR Obtaining Port and Ip: " + err.Error())
    return err
  }

  err = nodeclient.KillSuricataMainConf(ipnid, portnid, anode)
  if err != nil {
    logs.Error("nodeclient.KillSuricataMainConf error HTTP data request: " + err.Error())
    return err
  }

  return nil
}
func ReloadSuricataMainConf(anode map[string]string) (err error) {
  if ndb.Db == nil {
    logs.Error("ReloadSuricataMainConf -- Can't acces to database: ")
    return errors.New("StartSuricataMainConf -- Can't acces to database")
  }

  err = ndb.GetTokenByUuid(anode["uuid"])
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return err
  }
  ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
  if err != nil {
    logs.Error("node/ReloadSuricataMainConf ERROR Obtaining Port and Ip: " + err.Error())
    return err
  }

  err = nodeclient.ReloadSuricataMainConf(ipnid, portnid, anode)
  if err != nil {
    logs.Error("nodeclient.ReloadSuricataMainConf error HTTP data request: " + err.Error())
    return err
  }

  return nil
}

func CreateSuricataService(guuid, nuuid string, suricataData utils.SuricataData) (assigned bool, details map[string]string) {
  if nuuid == "" || guuid == "" {
    return true, nil
  }
  if suricataData.Ruleset == "default" {
    suricataData.Ruleset = ndb.GetDefaultRuleset()
  } else if suricataData.Ruleset == "group" {
    suricataData.Ruleset = ndb.GetGroupRulesets(guuid)
  }

  sData := make(map[string]string)
  sData["name"] = suricataData.Name
  sData["type"] = "suricata"
  sData["status"] = suricataData.Status
  sData["interface"] = suricataData.Interface
  sData["bpfFile"] = suricataData.BpfFile
  sData["bpf"] = suricataData.Bpf
  sData["ruleset"] = suricataData.Ruleset

  rname, _ := ndb.GetRuleName(suricataData.Ruleset)
  sData["rulesetName"] = rname
  sData["localRulesetName"] = rname
  sData["configFile"] = suricataData.ConfigFile

  logs.Info("CSS -> %+v", sData)
  nodeclient.AddSuricataService(nuuid, sData)

  return true, nil
}

func CreateStapService(nuuid string, stapData utils.StapData) (assigned bool, details map[string]string) {
  if nuuid == "" {
    return true, nil
  }

  sData := make(map[string]string)
  sData["name"] = stapData.Name
  sData["type"] = stapData.Type
  sData["status"] = stapData.Status
  sData["interface"] = stapData.Interface
  sData["cert"] = stapData.Cert
  sData["port"] = stapData.Port

  logs.Info("Create STAP Service -> %+v", sData)
  nodeclient.AddStapService(nuuid, sData)

  if stapData.Status == "enable" {
    logs.Info("start STAP service")
  }

  return true, nil
}
