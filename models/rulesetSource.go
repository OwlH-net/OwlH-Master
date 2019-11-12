package models 

import (
    "owlhmaster/rulesetSource"
    "github.com/astaxie/beego/logs"
)

func CreateRulesetSource(data map[string]string) (err error) {
    err = rulesetSource.CreateRulesetSource(data)
    return err
}

func CreateCustomRulesetSource(data map[string]string) (err error) {
    err = rulesetSource.CreateCustomRulesetSource(data)
    return err
}

func GetAllRulesetSource() (data map[string]map[string]string, err error) {
    data, err = rulesetSource.GetAllRulesetSource()
    return data, err
}

func EditRulesetSource(data map[string]string) (err error) {
    err = rulesetSource.EditRulesetSource(data)
    return err
}

func DeleteRulesetSource(anode map[string]string) (err error) {
    err = rulesetSource.DeleteRulesetSource(anode)
    return err
}

func DeleteRulesetFile(uuid string) (err error) {
    err = rulesetSource.DeleteRulesetFile(uuid)
    return err
}

func DownloadFile(data map[string]string) (err error) {
    err = rulesetSource.DownloadFile(data)
    logs.Notice(err)
    logs.Notice(err)
    logs.Notice(err)
    return err
}

func OverwriteDownload(data map[string]string) (err error) {
    err = rulesetSource.OverwriteDownload(data)
    logs.Notice(err)
    logs.Notice(err)
    logs.Notice(err)
    return err
}

func CompareFiles(uuid string) (mapData map[string]map[string]string, err error) {
    mapData,err = rulesetSource.CompareFiles(uuid)
    return mapData,err
}

func CreateNewFile(data map[string]string) (err error) {
    err = rulesetSource.CreateNewFile(data)
    return err
}

func GetDetails(uuid string) (files map[string]map[string]string, err error) {
    files, err = rulesetSource.GetDetails(uuid)
    return files ,err
}

func GetFileUUIDfromRulesetUUID(value string)(uuid string, err error){
    uuid,err = rulesetSource.GetFileUUIDfromRulesetUUID(value)
    return uuid,err
}

func OverwriteRuleFile(uuid string)(err error){
    err = rulesetSource.OverwriteRuleFile(uuid)
    return err
}

func AddNewLinesToRuleset(uuid string)(err error){
    err = rulesetSource.AddNewLinesToRuleset(uuid)
    return err
}