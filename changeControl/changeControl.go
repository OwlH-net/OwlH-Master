package changecontrol

import (
    "owlhmaster/database"
    "owlhmaster/utils"
    "time"
    "github.com/astaxie/beego/logs"
    // "strconv"
    // "owlhmaster/nodeclient"
)

func GetChangeControl()(data map[string]map[string]string, err error) {
    data, err = ndb.GetChangeControl(); if err != nil{logs.Error("Error getting ChangeControl database values: "+err.Error()); return nil,err}
    return data, err
}

func InsertChangeControl(values map[string]string)(err error){
    loadMasterName := map[string]map[string]string{}
    loadMasterName["master"] = map[string]string{}
    loadMasterName["master"]["name"] = ""
    loadMasterName,err = utils.GetConf(loadMasterName)
    if err != nil {logs.Error("changeChangeControl/InsertChangeControl error readding data from main.conf: "+err.Error()); return}
    deviceName := loadMasterName["master"]["name"]

    uuid:= utils.Generate()
    currentTime := time.Now()
    timeFormated := currentTime.Format("2006-01-02T15:04:05")

    err = ndb.InsertChangeControl(uuid, "deviceName", deviceName); if err != nil{logs.Error("Error inserting ChangeControl database value: "+err.Error()); return err}
    err = ndb.InsertChangeControl(uuid, "user", "admin"); if err != nil{logs.Error("Error inserting ChangeControl database value: "+err.Error()); return err}
    err = ndb.InsertChangeControl(uuid, "time", timeFormated); if err != nil{logs.Error("Error inserting ChangeControl database value: "+err.Error()); return err}
    for x := range values {
        err = ndb.InsertChangeControl(uuid, x, values[x]); if err != nil{logs.Error("Error inserting ChangeControl database value: "+err.Error()); return err}
    }

    return nil
}