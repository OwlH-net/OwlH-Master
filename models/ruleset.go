package models 

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/ruleset"
)


func GetRules()(rules map[string]map[string]string, err error) {
    logs.Info("MODEL===Lectura de Ruleset")
    rules,err = ruleset.Read("")
    return rules,err
}

func GetRuleSID(sid string) (rules map[string]string, err error) {
    logs.Info("MODEL===Lectura de línea según SID")
    rules,err = ruleset.ReadSID(sid)
    return rules,err
}

func AddRuleset(n map[string]string)(err error){
    logs.Info("model/ruleset -- Addruleset")
    err = ruleset.AddRuleset(n)
    return err
}

func GetAllRulesets()(rulesets *map[string]map[string]string, err error){
    logs.Info("model/ruleset -- GetAllRulesets")
    rulesets,err = ruleset.GetAllRulesets()
    return rulesets,err
}

func GetRulesetRules(nid string)(rulesets map[string]map[string]string, err error){
    logs.Info("model/ruleset -- GetRulesetRules")
    rulesets,err = ruleset.GetRulesetRules(nid)
    return rulesets,err
}