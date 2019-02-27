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