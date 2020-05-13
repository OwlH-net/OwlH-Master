package models 

import (
    "owlhmaster/ruleset"
    "owlhmaster/node"
    "owlhmaster/changeControl"
)

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/default \
func GetRules(username string)(rules map[string]map[string]string, err error) {
    rules,err = ruleset.ReadRuleset("")
    changecontrol.ChangeControlInsertData(err, "GetRules", username)
    return rules,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/rule/:sid/:uuid  \
func GetRuleSID(ruleSidPath map[string]string, username string) (rules map[string]string, err error) {
    rules,err = ruleset.ReadSID(ruleSidPath)
    changecontrol.ChangeControlInsertData(err, "GetRuleSID", username)
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
//     changecontrol.ChangeControlInsertData(err, "")
    // return err
// }

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/rule/:sid/:uuid  \
func GetAllRulesets(username string)(rulesets map[string]map[string]string, err error){
    rulesets,err = ruleset.GetAllRulesets()
    changecontrol.ChangeControlInsertData(err, "GetAllRulesets", username)
    return rulesets,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/rules/:uuid  \
func GetRulesetRules(nid string, username string)(rulesets map[string]map[string]string, err error){
    rulesets,err = ruleset.GetRulesetRules(nid)
    changecontrol.ChangeControlInsertData(err, "GetRulesetRules", username)
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
func SetRuleSelected(n map[string]string, username string)(err error){
    err = ruleset.SetRuleSelected(n)
    changecontrol.ChangeControlInsertData(err, "SetRuleSelected", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/get/:uuid  \
func GetRuleSelected(nidSelected string, username string)(rulesetBack string, err error){
    rulesetBack, err = ruleset.GetRuleSelected(nidSelected)
    changecontrol.ChangeControlInsertData(err, "GetRuleSelected", username)
    return rulesetBack, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/get/name/:uuid  \
func GetRuleName(nidRule string, username string)(rulesetBack string, err error){
    rulesetBack, err = ruleset.GetRuleName(nidRule)
    changecontrol.ChangeControlInsertData(err, "GetRuleName", username)
    return rulesetBack, err
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
//     changecontrol.ChangeControlInsertData(err, "")
    // return err
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
func SetRulesetAction(ruleAction map[string]string, username string)(err error){
    err = ruleset.SetRulesetAction(ruleAction)
    changecontrol.ChangeControlInsertData(err, "SetRulesetAction", username)
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
func SetRuleNote(ruleNote map[string]string, username string)(err error){
    err = ruleset.SetRuleNote(ruleNote)
    changecontrol.ChangeControlInsertData(err, "SetRuleNote", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/getnote/:uuid/:sid \
func GetRuleNote(ruleGetNote map[string]string, username string)(note string, err error){
    note,err = ruleset.GetRuleNote(ruleGetNote)
    changecontrol.ChangeControlInsertData(err, "GetRuleNote", username)
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
func DeleteRuleset(rulesetMap map[string]string, username string)(err error){
    err = ruleset.DeleteRuleset(rulesetMap)
    changecontrol.ChangeControlInsertData(err, "DeleteRuleset", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/synchronize \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v"
// }
// }
func SyncRulesetToAllNodes(anode map[string]string, username string)(err error){
    err = node.SyncRulesetToAllNodes(anode)
    changecontrol.ChangeControlInsertData(err, "SyncRulesetToAllNodes", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/syncAllRulesets \
func SynchronizeAllRulesets(username string)(err error){
    err = ruleset.SynchronizeAllRulesets()
    changecontrol.ChangeControlInsertData(err, "SynchronizeAllRulesets", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/getAllRuleData \
func GetAllRuleData(username string)(data map[string]map[string]string,err error) {
    data,err = ruleset.GetAllRuleData()
    changecontrol.ChangeControlInsertData(err, "GetAllRuleData", username)
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
func AddNewRuleset(data map[string]map[string]string, username string)(duplicated []byte, err error) {
    duplicated, err = ruleset.AddNewRuleset(data)
    changecontrol.ChangeControlInsertData(err, "AddNewRuleset", username)
    return duplicated, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/modify \
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
func ModifyRuleset(data map[string]map[string]string, username string)(duplicated []byte, err error) {
    duplicated, err = ruleset.ModifyRuleset(data)
    changecontrol.ChangeControlInsertData(err, "ModifyRuleset", username)
    return duplicated, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/custom \
func GetAllCustomRulesets(username string)(data map[string]map[string]string,err error) {
    data,err = ruleset.GetAllCustomRulesets()
    changecontrol.ChangeControlInsertData(err, "GetAllCustomRulesets", username)
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
func AddRulesToCustomRuleset(anode map[string]string, username string)(duplicatedRules map[string]string, err error) {
    duplicatedRules, err = ruleset.AddRulesToCustomRuleset(anode)
    changecontrol.ChangeControlInsertData(err, "AddRulesToCustomRuleset", username)
    return duplicatedRules, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/ruleset/readRuleset/:uuid \
func ReadRulesetData(uuid string, username string)(content map[string]string, err error) {
    content, err = ruleset.ReadRulesetData(uuid)
    changecontrol.ChangeControlInsertData(err, "ReadRulesetData", username)
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
func SaveRulesetData(content map[string]string, username string)(err error) {
    err = ruleset.SaveRulesetData(content)
    changecontrol.ChangeControlInsertData(err, "SaveRulesetData", username)
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
//     changecontrol.ChangeControlInsertData(err, "")
    // return err
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
//     changecontrol.ChangeControlInsertData(err, "")
    // return err
// }

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/updateRule \
//   -H 'Content-Type: application/json' \
//   -d '{
//      "uuid":"aaa",
//      "line":"aaa",
// }
func UpdateRule(anode map[string]string, username string)(err error) {
    err = ruleset.UpdateRule(anode)
    changecontrol.ChangeControlInsertData(err, "UpdateRule", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/ruleset/syncToAll \
//   -H 'Content-Type: application/json' \
//   -d '{
//      "uuid":"aaa"
// }
func SyncToAll(anode map[string]string, username string)(err error) {
    err = ruleset.SyncToAll(anode)
    changecontrol.ChangeControlInsertData(err, "SyncToAll", username)
    return err
}