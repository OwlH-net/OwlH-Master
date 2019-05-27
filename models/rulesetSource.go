package models 

import (
    "owlhmaster/rulesetSource"
)

func CreateRulesetSource(data map[string]string) (err error) {
    err = rulesetSource.CreateRulesetSource(data)
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

func DeleteRulesetSource(uuid string) (err error) {
    err = rulesetSource.DeleteRulesetSource(uuid)
    return err
}

func DeleteRulesetFile(uuid string) (err error) {
    err = rulesetSource.DeleteRulesetFile(uuid)
    return err
}

func DownloadFile(data map[string]string) (err error) {
    err = rulesetSource.DownloadFile(data)
    return err
}

func CompareFiles(data map[string]string) (mapData map[string]map[string]string, err error) {
    mapData,err = rulesetSource.CompareFiles(data)
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
