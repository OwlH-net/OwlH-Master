package models 

import (
    // "github.com/astaxie/beego/logs"
    "owlhmaster/stap"
)

func AddServer(data map[string]string)(err error) {
    err = stap.AddServer(data)
    return err
}

func GetAllServers(nodeuuid string)(data map[string]map[string]string, err error) {
    data,err = stap.GetAllServers(nodeuuid)
    return data,err
}

func GetServer(uuid string, serveruuid string)(data map[string]map[string]string, err error) {
    data,err = stap.GetServer(uuid, serveruuid)
    return data,err
}

func Stap(n string) (data map[string]bool, err error) {
    data,err = stap.Stap(n)
    return data,err
}

func RunStap(n string) (data string, err error) {
    data,err = stap.RunStap(n)
    return data,err
}

func StopStap(n string) (data string, err error) {
    data,err = stap.StopStap(n)
    return data,err
}

func RunStapServer(uuid string, server string) (data string, err error) {
    data,err = stap.RunStapServer(uuid, server)
    return data,err
}

func StopStapServer(uuid string, server string) (data string, err error) {
    data,err = stap.StopStapServer(uuid, server)
    return data,err
}

func PingServerStap(uuid string, server string) (data map[string]string, err error) {
    data,err = stap.PingServerStap(uuid, server)
    return data,err
}