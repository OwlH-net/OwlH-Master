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

func DeleteRulesetSource(groupId string) (err error) {
    err = rulesetSource.DeleteRulesetSource(groupId)
    return err
}

func DownloadFile(data map[string]string) (err error) {
    err = rulesetSource.DownloadFile(data)
    return err
}

func CompareFiles(data map[string]string) (err error) {
    err = rulesetSource.CompareFiles(data)
    return err
}
