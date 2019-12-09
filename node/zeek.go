package node

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "errors"
    "owlhmaster/nodeclient"
)

func Zeek(n string) (data nodeclient.ZeekData, err error) {
    ip,port,err := ndb.ObtainPortIp(n)
    if err != nil {
        logs.Info("Zeek - get IP and PORT Error -> %s", err.Error())
        return data,err
    }    
    logs.Info("Zeek IP and PORT -> %s, %s", ip, port)
    data, err = nodeclient.Zeek(ip,port)
    if err != nil {
        return data,err
    }
    return data,nil
}

func RunZeek(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("RunZeek -- Can't acces to database")
        return "", errors.New("RunZeek -- Can't acces to database")
    }
    
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/RunZeek ERROR Obtaining Port and Ip: "+err.Error())
        return "",err
    }
    data,err = nodeclient.RunZeek(ipnid,portnid)
    if err != nil {
        logs.Error("node/RunZeek ERROR http data request: "+err.Error())
        return "",err
    }
    return data,nil
}

func StopZeek(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopZeek -- Can't acces to database")
        return "", errors.New("StopZeek -- Can't acces to database")
    }
    
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/StopZeek ERROR Obtaining Port and Ip: "+err.Error())
        return "",err
    }
    data,err = nodeclient.StopZeek(ipnid,portnid)
    if err != nil {
        logs.Error("node/StopZeek ERROR http data request: "+err.Error())
        return "",err
    }
    return data,nil
}

func DeployZeek(uuid string)(err error){
    if ndb.Db == nil {
        logs.Error("DeployZeek -- Can't acces to database")
        return errors.New("DeployZeek -- Can't acces to database")
    }
    
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/DeployZeek ERROR Obtaining Port and Ip: "+err.Error())
        return err
    }
    err = nodeclient.DeployZeek(ipnid,portnid)
    if err != nil {
        logs.Error("node/DeployZeek ERROR http data request: "+err.Error())
        return err
    }
    return nil
}

func ChangeZeekMode(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("ChangeZeekMode -- Can't acces to database"); return err}

    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/ChangeZeekMode ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.ChangeZeekMode(ipnid,portnid,anode)
    if err != nil { logs.Error("node/ChangeZeekMode ERROR http data request: "+err.Error()); return err}

    return nil
}

func AddClusterValue(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("AddClusterValue -- Can't acces to database"); return err}

    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/AddClusterValue ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.AddClusterValue(ipnid,portnid,anode)
    if err != nil { logs.Error("node/AddClusterValue ERROR http data request: "+err.Error()); return err}

    return nil
}

func PingCluster(uuid string)(data map[string]map[string]string, err error){
    if ndb.Db == nil { logs.Error("PingCluster -- Can't acces to database"); return nil,err}

    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil { logs.Error("node/PingCluster ERROR Obtaining Port and Ip: "+err.Error()); return nil,err}
    
    data,err = nodeclient.PingCluster(ipnid,portnid)
    if err != nil { logs.Error("node/PingCluster ERROR http data request: "+err.Error()); return nil,err}

    return data,nil
}

func EditClusterValue(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("EditClusterValue -- Can't acces to database"); return err}

    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/EditClusterValue ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.EditClusterValue(ipnid,portnid,anode)
    if err != nil { logs.Error("node/EditClusterValue ERROR http data request: "+err.Error()); return err}

    return nil
}

func DeleteClusterValue(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("DeleteClusterValue -- Can't acces to database"); return err}

    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/DeleteClusterValue ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.DeleteClusterValue(ipnid,portnid,anode)
    if err != nil { logs.Error("node/DeleteClusterValue ERROR http data request: "+err.Error()); return err}

    return nil
}

func SyncCluster(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("SyncCluster -- Can't acces to database"); return err}

    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/SyncCluster ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.SyncCluster(ipnid,portnid,anode)
    if err != nil { logs.Error("node/SyncCluster ERROR http data request: "+err.Error()); return err}

    return nil
}

func LaunchZeekMainConf(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("LaunchZeekMainConf -- Can't acces to database"); return err}

    ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil { logs.Error("node/LaunchZeekMainConf ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.LaunchZeekMainConf(ipnid,portnid,anode)
    if err != nil { logs.Error("node/LaunchZeekMainConf ERROR http data request: "+err.Error()); return err}

    return nil
}