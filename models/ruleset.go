package models 

import (
    "owlhmaster/ruleset"
    "owlhmaster/node"
)

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/default \
func GetRules()(rules map[string]map[string]string, err error) {
    rules,err = ruleset.ReadRuleset("")
    return rules,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/rule/:sid/:uuid  \
func GetRuleSID(ruleSidPath map[string]string) (rules map[string]string, err error) {
    rules,err = ruleset.ReadSID(ruleSidPath)
    return rules,err
}

// // curl -X POST \
// //   https://52.47.197.22:50002/v1/ruleset/new \
// //   -H 'Content-Type: application/json' \
// //   -d '{
// //     "nodeuuid": "d",
// //     "uuid": "v",
// //     "param": "v",
// //     "value": "v",
// // }
// // }
// func AddRuleset(n map[string]string)(err error){
//     err = ruleset.AddRuleset(n)
//     return err
// }

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/rule/:sid/:uuid  \
func GetAllRulesets()(rulesets map[string]map[string]string, err error){
    rulesets,err = ruleset.GetAllRulesets()
    return rulesets,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/rules/:uuid  \
func GetRulesetRules(nid string)(rulesets map[string]map[string]string, err error){
    rulesets,err = ruleset.GetRulesetRules(nid)
    return rulesets,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/set \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "type": "d",
//     "uuid": "v",
// }
// }
func SetRuleSelected(n map[string]string)(err error){
    err = ruleset.SetRuleSelected(n)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/get/:uuid  \
func GetRuleSelected(nidSelected string)(rulesetReturn string, err error){
    rulesetReturn, err = ruleset.GetRuleSelected(nidSelected)
    return rulesetReturn, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/get/name/:uuid  \
func GetRuleName(nidRule string)(rulesetReturn string, err error){
    rulesetReturn, err = ruleset.GetRuleName(nidRule)
    return rulesetReturn, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/clone \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "type": "d",
//     "uuid": "v",
// }
// }
// func SetClonedRuleset(ruleCloned map[string]string)(err error){
//     err = ruleset.SetClonedRuleset(ruleCloned)
//     return err
// }


// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/action \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "sid": "d",
//     "action": "d",
//     "uuid": "v"
// }
// }
func SetRulesetAction(ruleAction map[string]string)(err error){
    err = ruleset.SetRulesetAction(ruleAction)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/note \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "sid": "d",
//     "note": "d",
//     "uuid": "v"
// }
// }
func SetRuleNote(ruleNote map[string]string)(err error){
    err = ruleset.SetRuleNote(ruleNote)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/getnote/:uuid/:sid \
func GetRuleNote(ruleGetNote map[string]string)(note string, err error){
    note,err = ruleset.GetRuleNote(ruleGetNote)
    return note,err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/ruleset/deleteRuleset \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "name": "d",
//     "uuid": "v"
// }
// }
func DeleteRuleset(rulesetMap map[string]string)(err error){
    err = ruleset.DeleteRuleset(rulesetMap)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/synchronize \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v"
// }
// }
func SyncRulesetToAllNodes(anode map[string]string)(err error){
    err = node.SyncRulesetToAllNodes(anode)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/syncAllRulesets \
func SynchronizeAllRulesets()(err error){
    err = ruleset.SynchronizeAllRulesets()
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/getAllRuleData \
func GetAllRuleData()(data map[string]map[string]string,err error) {
    data,err = ruleset.GetAllRuleData()
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/addNewRuleset \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": {
//          "sourceName":"aaa",
//          "fileName":"aaa",
//          "filePath":"aaa",
//          "rulesetName":"aaa",
//          "rulesetDesc":"aaa",
//          "sourceType":"aaa",
//      }
// }
func AddNewRuleset(data map[string]map[string]string)(duplicated []byte, err error) {
    duplicated, err = ruleset.AddNewRuleset(data)
    return duplicated, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/custom \
func GetAllCustomRulesets()(data map[string]map[string]string,err error) {
    data,err = ruleset.GetAllCustomRulesets()
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/addRulesToCustom \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": {
//          "origin":"aaa",
//          "dest":"aaa",
//          "ruleset":"aaa",
//          "sids":"aaa",
//      }
// }
func AddRulesToCustomRuleset(anode map[string]string)(duplicatedRules map[string]string, err error) {
    duplicatedRules, err = ruleset.AddRulesToCustomRuleset(anode)
    return duplicatedRules, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/readRuleset/:uuid \
func ReadRulesetData(uuid string)(content map[string]string, err error) {
    content, err = ruleset.ReadRulesetData(uuid)
    return content, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/saveRuleset \
//   -H 'Content-Type: application/json' \
//   -d '{
//      {
//          "uuid":"aaa",
//          "content":"aaa",
//      }
// }
func SaveRulesetData(content map[string]string)(err error) {
    err = ruleset.SaveRulesetData(content)
    return err
}

// // curl -X PUT \
// //   https://52.47.197.22:50002/v1/ruleset/timeSchedule \
// //   -H 'Content-Type: application/json' \
// //   -d '{
// //     "uuid": {
// //          "origin":"aaa",
// //          "dest":"aaa",
// //          "ruleset":"aaa",
// //          "sids":"aaa",
// //      }
// // }
// func TimeSchedule(content map[string]string)(err error) {
//     // err = ruleset.TimeSchedule(content)
//     return err
// }

// // curl -X PUT \
// //   https://52.47.197.22:50002/v1/ruleset/stopTimeSchedule \
// //   -H 'Content-Type: application/json' \
// //   -d '{
// //     "uuid": {
// //          "origin":"aaa",
// //          "dest":"aaa",
// //          "ruleset":"aaa",
// //          "sids":"aaa",
// //      }
// // }
// func StopTimeSchedule(content map[string]string)(err error) {
//     // err = ruleset.StopTimeSchedule(content)
//     return err
// }

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/updateRule \
//   -H 'Content-Type: application/json' \
//   -d '{
//      "uuid":"aaa",
//      "line":"aaa",
// }
func UpdateRule(anode map[string]string)(err error) {
    err = ruleset.UpdateRule(anode)
    return err
}
