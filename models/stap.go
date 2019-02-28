package models 

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/stap"
)

func AddServer(data map[string]string)(err error) {
    logs.Info("MODEL===AddServer stap")
    err = stap.AddServer(data)
    return err
}

func GetAllServers(nodeuuid string)(data map[string]map[string]string, err error) {
    logs.Info("MODEL===GetAllServers stap")
    data,err = stap.GetAllServers(nodeuuid)
    return data,err
}