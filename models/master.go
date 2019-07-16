package models 

import (
    "owlhmaster/master"
    "owlhmaster/database"
)

func GetMasterTitle() (data string, err error) {
    data, err = master.GetMasterTitle()
    return data, err
}

func GetFileContent(file string) (data map[string]string, err error) {
    data, err = master.GetFileContent(file)
    return data, err
}

func SaveFileContent(anode map[string]string) (err error) {
    err = master.SaveFileContent(anode)
    return err
}

func PingPlugins() (data map[string]map[string]string, err error) {
    data,err = ndb.PingPlugins()
    return data,err
}

func ChangePluginStatus(anode map[string]string) (err error) {
    err = ndb.ChangePluginStatus(anode)
    return err
}