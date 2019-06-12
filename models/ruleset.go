package models 

import (
    "owlhmaster/ruleset"
    "owlhmaster/node"
)


func GetRules()(rules map[string]map[string]string, err error) {
    rules,err = ruleset.ReadRuleset("")
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

func GetAllRulesets()(rulesets map[string]map[string]string, err error){
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

func DeleteRuleset(rulesetMap map[string]string)(err error){
    err = ruleset.DeleteRuleset(rulesetMap)
    return err
}

func SyncRulesetToAllNodes(anode map[string]string)(err error){
    err = node.SyncRulesetToAllNodes(anode)
    return err
}

func GetAllRuleData()(data map[string]map[string]string,err error) {
    data,err = ruleset.GetAllRuleData()
    return data,err
}

func AddNewRuleset(data map[string]map[string]string)(duplicated []byte, err error) {
	duplicated, err = ruleset.AddNewRuleset(data)
    return duplicated, err
}

func GetAllCustomRulesets()(data map[string]map[string]string,err error) {
	data,err = ruleset.GetAllCustomRulesets()
	return data,err
}

func AddRulesToCustomRuleset(anode map[string]string)(duplicatedRules map[string]string, err error) {
	duplicatedRules, err = ruleset.AddRulesToCustomRuleset(anode)
	return duplicatedRules, err
}

func ReadRulesetData(uuid string)(content map[string]string, err error) {
	content, err = ruleset.ReadRulesetData(uuid)
	return content, err
}

func SaveRulesetData(content map[string]string)(err error) {
	err = ruleset.SaveRulesetData(content)
	return err
}