package collector

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/nodeclient"
    "os/exec"
    "owlhmaster/database"
    "owlhmaster/utils"
)

func PlayCollector(uuid string) (err error) {
    ip,port,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/GetAllFiles ERROR getting node port/ip : "+err.Error())
        return err
    }

    err = nodeclient.PlayCollector(ip,port)
    if err != nil {
        logs.Error("nodeclient.playCollector ERROR connection through http new Request: "+err.Error())
        return err
    }
    return nil
}

func StopCollector(uuid string) (err error) {
    ip,port,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/GetAllFiles ERROR getting node port/ip : "+err.Error())
        return err
    }   
    err = nodeclient.StopCollector(ip,port)
    if err != nil {
        logs.Error("nodeclient.StopCollector ERROR connection through http new Request: "+err.Error())
        return err
    }
    return nil
}

func ShowCollector(uuid string) (data string, err error) {
    ip,port,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/GetAllFiles ERROR getting node port/ip : "+err.Error())
        return data, err
    }   
    data, err = nodeclient.ShowCollector(ip,port)
    if err != nil {
        logs.Error("nodeclient.ShowCollector ERROR connection through http new Request: "+err.Error())
        return "",err
    }
    
    return data,nil
}

func PlayMasterCollector() (err error) {
    _, err = exec.Command("bash","-c","ls -la").Output()
    if err != nil{
        logs.Error("PlayMasterCollector Error executing command in StopCollector function: "+err.Error())
        return err    
    }
    return nil
}

func StopMasterCollector() (err error) {
    _, err = exec.Command("bash","-c","ls -la").Output()
    if err != nil{
        logs.Error("StopMasterCollector Error executing command in StopCollector function: "+err.Error())
        return err    
    }
    return nil
}

func ShowMasterCollector() (data string, err error) {
    stapCollector := map[string]map[string]string{}
    stapCollector["stapCollector"] = map[string]string{}
    stapCollector["stapCollector"]["status"] = ""
    stapCollector["stapCollector"]["param"] = ""
    stapCollector["stapCollector"]["command"] = ""
    stapCollector,err = utils.GetConf(stapCollector)
    status := stapCollector["stapCollector"]["status"]
    param := stapCollector["stapCollector"]["param"]
    command := stapCollector["stapCollector"]["command"]

    output, err := exec.Command(command, param, status).Output()
    if err != nil{
        logs.Error("ShowMasterCollector Error executing command in ShowCollector function: "+err.Error())
        return "",err    
    }
    return string(output),nil
}