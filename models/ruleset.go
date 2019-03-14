package models 

import (
    // "github.com/astaxie/beego/logs"
    "owlhmaster/ruleset"
)


func GetRules()(rules map[string]map[string]string, err error) {
    rules,err = ruleset.Read("")
    return rules,err
}

func GetRuleSID(ruleSidPath map[string]string) (rules map[string]string, err error) {
    rules,err = ruleset.ReadSID(ruleSidPath)
    return rules,err
}

func AddRuleset(n map[string]string)(err error){
    err = ruleset.AddRuleset(n)
    return err
}

func GetAllRulesets()(rulesets *map[string]map[string]string, err error){
    rulesets,err = ruleset.GetAllRulesets()
    return rulesets,err
}

func GetRulesetRules(nid string)(rulesets map[string]map[string]string, err error){
    rulesets,err = ruleset.GetRulesetRules(nid)
    return rulesets,err
}

func SetRuleSelected(n map[string]string)(err error){
    err = ruleset.SetRuleSelected(n)
    return err
}

func GetRuleSelected(nidSelected string)(rulesetReturn string, err error){
    rulesetReturn, err = ruleset.GetRuleSelected(nidSelected)
    return rulesetReturn, err
}

func GetRuleName(nidRule string)(rulesetReturn string, err error){
    rulesetReturn, err = ruleset.GetRuleName(nidRule)
    return rulesetReturn, err
}

func SetClonedRuleset(ruleCloned map[string]string)(err error){
    err = ruleset.SetClonedRuleset(ruleCloned)
    return err
}

func SetRulesetAction(ruleAction map[string]string)(err error){
    err = ruleset.SetRulesetAction(ruleAction)
    return err
}

func SetRuleNote(ruleNote map[string]string)(err error){
    err = ruleset.SetRuleNote(ruleNote)
    return err
}

func GetRuleNote(ruleGetNote map[string]string)(note string, err error){
    note,err = ruleset.GetRuleNote(ruleGetNote)
    return note,err
}