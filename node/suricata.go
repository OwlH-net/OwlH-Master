package node

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "errors"
    "owlhmaster/nodeclient"
)

func Suricata(n string) (data map[string]bool, err error) {
    ip,port,err := ndb.ObtainPortIp(n)
    if err != nil {
        logs.Info("Suricata - get IP and PORT Error -> %s", err.Error())
        return nil,err
    }    
    logs.Info("Suricata IP and PORT -> %s, %s", ip, port)
    data, err = nodeclient.Suricata(ip,port)
    if err != nil {
        return nil,err
    }
    return data,nil
}

// func GetSuricataBPF(n string)(bpf string, err error) {
//     if ndb.Db == nil {
//         logs.Error("GetSuricataBPF -- Can't acces to database: "+err.Error())
//         return "", err
//     }
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

func PutSuricataBPF(n map[string]string)(err error) {
    if ndb.Db == nil { logs.Error("PutSuricataBPF -- Can't acces to database: "); return errors.New("PutSuricataBPF -- Can't acces to database")}

    ipnid,portnid,err := ndb.ObtainPortIp(n["uuid"])
    if err != nil {logs.Error("node/PutSuricataBPF ERROR Obtaining Port and Ip: "+err.Error()); return err}

    err = nodeclient.PutSuricataBPF(ipnid,portnid,n)
    if err != nil { logs.Error("nodeclient.PutSuricataBPF error HTTP data request: "+err.Error()); return err}

    return nil
}

func RunSuricata(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("RunSuricata -- Can't acces to database")
        return "", errors.New("RunSuricata -- Can't acces to database")
    }
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/RunSuricata ERROR Obtaining Port and Ip: "+err.Error())
        return "",err
    }
    data,err = nodeclient.RunSuricata(ipnid,portnid)
    if err != nil {
        logs.Error("node/RunSuricata ERROR http data request: "+err.Error())
        return "",err
    }
    return data,nil
}

func StopSuricata(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopSuricata -- Can't acces to database")
        return "", errors.New("StopSuricata -- Can't acces to database")
    }
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/StopSuricata ERROR Obtaining Port and Ip: "+err.Error())
        return "",err
    }
    data,err = nodeclient.StopSuricata(ipnid,portnid)
    if err != nil {
        logs.Error("node/StopSuricata ERROR http data request: "+err.Error())
        return "",err
    }
    return data,nil
}

func AddPluginService(anode map[string]string)( err error){
    if ndb.Db == nil {logs.Error("AddPluginService -- Can't acces to database: "); return errors.New("AddPluginService -- Can't acces to database")}
    
    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {logs.Error("node/AddPluginService ERROR Obtaining Port and Ip: "+err.Error()); return err}

    err = nodeclient.AddPluginService(ipnid,portnid,anode)
    if err != nil {logs.Error("nodeclient.AddPluginService error HTTP data request: "+err.Error()); return err}

    return nil
}

func GetSuricataServices(uuid string)(data map[string]map[string]string, err error){
    if ndb.Db == nil {logs.Error("GetSuricataServices -- Can't acces to database: "); return nil,errors.New("GetSuricataServices -- Can't acces to database")}
    
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil {logs.Error("node/GetSuricataServices ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}

    data,err = nodeclient.GetSuricataServices(ipnid,portnid)
    if err != nil {logs.Error("nodeclient.GetSuricataServices error HTTP data request: "+err.Error()); return nil,err}

    return data,nil
}

func SaveSuricataInterface(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("SaveSuricataInterface -- Can't acces to database"); return err}

    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/SaveSuricataInterface ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.SaveSuricataInterface(ipnid,portnid,anode)
    if err != nil { logs.Error("node/SaveSuricataInterface ERROR http data request: "+err.Error()); return err}

    return nil
}

func ChangeSuricataTable(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("ChangeSuricataTable -- Can't acces to database"); return err}

    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/ChangeSuricataTable ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.ChangeSuricataTable(ipnid,portnid,anode)
    if err != nil { logs.Error("node/ChangeSuricataTable ERROR http data request: "+err.Error()); return err}

    return nil
}