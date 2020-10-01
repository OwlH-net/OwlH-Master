package node

import (
    "errors"
    "github.com/astaxie/beego/logs"
    "io/ioutil"
    "owlhmaster/database"
    "owlhmaster/nodeclient"
)

func Zeek(n string) (data nodeclient.ZeekData, err error) {
    err = ndb.GetTokenByUuid(n)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return data, err
    }
    ip, port, err := ndb.ObtainPortIp(n)
    if err != nil {
        logs.Info("Zeek - get IP and PORT Error -> %s", err.Error())
        return data, err
    }
    logs.Info("Zeek IP and PORT -> %s, %s", ip, port)
    data, err = nodeclient.Zeek(ip, port)
    if err != nil {
        return data, err
    }

    //get node name adn uuid
    nodes, err := ndb.GetAllNodes()
    // for x := range nodes {
    //     if nodes[x]["ip"] == data.ManagerIP {
    data.ManagerUUID = n
    data.ManagerName = nodes[n]["name"]
    //     }
    // }

    return data, nil
}

func RunZeek(uuid string) (data string, err error) {
    if ndb.Db == nil {
        logs.Error("RunZeek -- Can't acces to database")
        return "", errors.New("RunZeek -- Can't acces to database")
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return "", err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/RunZeek ERROR Obtaining Port and Ip: " + err.Error())
        return "", err
    }
    data, err = nodeclient.RunZeek(ipnid, portnid)
    if err != nil {
        logs.Error("node/RunZeek ERROR http data request: " + err.Error())
        return "", err
    }
    return data, nil
}

func StopZeek(uuid string) (data string, err error) {
    if ndb.Db == nil {
        logs.Error("StopZeek -- Can't acces to database")
        return "", errors.New("StopZeek -- Can't acces to database")
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return "", err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/StopZeek ERROR Obtaining Port and Ip: " + err.Error())
        return "", err
    }
    data, err = nodeclient.StopZeek(ipnid, portnid)
    if err != nil {
        logs.Error("node/StopZeek ERROR http data request: " + err.Error())
        return "", err
    }
    return data, nil
}

// func DeployZeek(uuid string)(err error){
//     if ndb.Db == nil {
//         logs.Error("DeployZeek -- Can't acces to database")
//         return errors.New("DeployZeek -- Can't acces to database")
//     }

//     err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
//     ipnid,portnid,err := ndb.ObtainPortIp(uuid)
//     if err != nil {
//         logs.Error("node/DeployZeek ERROR Obtaining Port and Ip: "+err.Error())
//         return err
//     }
//     err = nodeclient.DeployZeek(ipnid,portnid)
//     if err != nil {
//         logs.Error("node/DeployZeek ERROR http data request: "+err.Error())
//         return err
//     }
//     return nil
// }

func ChangeZeekMode(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("ChangeZeekMode -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/ChangeZeekMode ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.ChangeZeekMode(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/ChangeZeekMode ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func AddClusterValue(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("AddClusterValue -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/AddClusterValue ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.AddClusterValue(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/AddClusterValue ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func PingCluster(uuid string) (data map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("PingCluster -- Can't acces to database")
        return nil, err
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/PingCluster ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    data, err = nodeclient.PingCluster(ipnid, portnid)
    if err != nil {
        logs.Error("node/PingCluster ERROR http data request: " + err.Error())
        return nil, err
    }

    return data, nil
}

func ZeekCommand(uuid, cmd string) (data map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("ZeekCommand -- Can't acces to database")
        return nil, err
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/ZeekCommand ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    data, err = nodeclient.ZeekCommand(ipnid, portnid, cmd)
    if err != nil {
        logs.Error("node/ZeekCommand ERROR http data request: " + err.Error())
        return nil, err
    }

    return data, nil
}

func EditClusterValue(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("EditClusterValue -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/EditClusterValue ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.EditClusterValue(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/EditClusterValue ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func DeleteClusterValue(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("DeleteClusterValue -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/DeleteClusterValue ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.DeleteClusterValue(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/DeleteClusterValue ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func SyncCluster(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("SyncCluster -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/SyncCluster ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.SyncCluster(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/SyncCluster ERROR http data request: " + err.Error())
        return err
    }

    return nil
}

func LaunchZeekMainConf(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("LaunchZeekMainConf -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/LaunchZeekMainConf ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.LaunchZeekMainConf(ipnid, portnid, anode)
    if err != nil {
        logs.Error("node/LaunchZeekMainConf ERROR http data request PUT Zeek status: " + err.Error())
        return err
    }

    return nil
}

// func SaveZeekValues(anode map[string]string)(err error){
//     if ndb.Db == nil { logs.Error("SaveZeekValues -- Can't acces to database"); return err}

//     err = ndb.GetTokenByUuid(anode["uuid"]); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
//     ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
//     if err != nil { logs.Error("node/SaveZeekValues ERROR Obtaining Port and Ip: "+err.Error()); return err}

//     err = nodeclient.SaveZeekValues(ipnid,portnid,anode)
//     if err != nil { logs.Error("node/SaveZeekValues ERROR http data request: "+err.Error()); return err}

//     return nil
// }

func SyncZeekValues(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("SyncZeekValues -- Can't acces to database")
        return err
    }

    err = ndb.GetTokenByUuid(anode["uuid"])
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(anode["uuid"])
    if err != nil {
        logs.Error("node/SyncZeekValues ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    data, err := ndb.GetPlugins()
    for x, y := range data {
        for y := range y {
            if x == "zeek" {
                if y == "policiesNode" || y == "variables2" {
                    continue
                }
                fileReaded, err := ioutil.ReadFile(data[x][y])
                if err != nil {
                    logs.Error("zeek/SyncZeekValues Error reading file for path: " + data[x][y])
                    return err
                }
                syncValue := make(map[string]string)
                syncValue[y] = string(fileReaded)

                if y == "policiesMaster" || y == "variables1" {
                    if y == "policiesMaster" {
                        syncValue["dst"] = data[x]["policiesNode"]
                    }
                    if y == "variables1" {
                        syncValue["dst"] = data[x]["variables2"]
                    }
                }

                err = nodeclient.SyncZeekValues(ipnid, portnid, syncValue)
                if err != nil {
                    logs.Error("zeek/SyncZeekValues ERROR http data request: " + err.Error())
                    return err
                }
            }
        }
    }
    return err
}
