package models 

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/ruleset"
)


func GetRules () (rules map[string]map[string]string, err error) {
    logs.Info("MODEL===Lectura de Ruleset")
    rules,err = ruleset.Read()
    return rules,err
}