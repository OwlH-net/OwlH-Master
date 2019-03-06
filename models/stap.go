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

func Stap(n string) (data  map[string]bool, err error) {
    data,err = stap.Stap(n)
    return data,err
}