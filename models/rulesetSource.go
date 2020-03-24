package models 

import (
    "owlhmaster/rulesetSource"
    "owlhmaster/changeControl"
    // "github.com/astaxie/beego/logs"
)

// curl -X POST \
//   https://52.47.197.22:50002/v1/rulesetSource/ \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "name": "v",
//     "desc": "v",
//     "url": "v",
//     "fileName": "v",
//     "type": "v",
//     "sourceType": "v",
// }
// }
func CreateRulesetSource(data map[string]string) (err error) {
    err = rulesetSource.CreateRulesetSource(data)
    changecontrol.ChangeControlInsertData(err, "CreateRulesetSource")
    return err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/rulesetSource/custom \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "name": "v",
//     "desc": "v",
//     "url": "v",
//     "fileName": "v",
//     "type": "v",
//     "sourceType": "v",
//     "isDownloaded": "v",
// }
func CreateCustomRulesetSource(data map[string]string) (err error) {
    err = rulesetSource.CreateCustomRulesetSource(data)
    changecontrol.ChangeControlInsertData(err, "CreateCustomRulesetSource")
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/rulesetSource/ \
func GetAllRulesetSource(hasPermissions bool) (data map[string]map[string]string, err error) {
    data, err = rulesetSource.GetAllRulesetSource(hasPermissions)
    changecontrol.ChangeControlInsertData(err, "GetAllRulesetSource")
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/rulesetSource/EditRulesetSource \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "name": "v",
//     "desc": "v",
//     "url": "v",
//     "uuid": "v",
// }
func EditRulesetSource(data map[string]string) (err error) {
    err = rulesetSource.EditRulesetSource(data)
    changecontrol.ChangeControlInsertData(err, "EditRulesetSource")
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/rulesetSource/DeleteRulesetSource \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "sourceType": "v",
//     "uuid": "v",
// }
func DeleteRulesetSource(anode map[string]string) (err error) {
    err = rulesetSource.DeleteRulesetSource(anode)
    changecontrol.ChangeControlInsertData(err, "DeleteRulesetSource")
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/rulesetSource/DeleteRulesetSource/:uuid \
func DeleteRulesetFile(uuid string) (err error) {
    err = rulesetSource.DeleteRulesetFile(uuid)
    changecontrol.ChangeControlInsertData(err, "DeleteRulesetFile")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/rulesetSource/downloadFile \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "url": "v",
//     "uuid": "v",
//     "name": "v",
//     "path": "v",
// }
func DownloadFile(data map[string]string) (err error) {
    err = rulesetSource.DownloadFile(data)
    changecontrol.ChangeControlInsertData(err, "DownloadFile")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/rulesetSource/overwriteDownload \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "url": "v",
//     "uuid": "v",
//     "name": "v",
//     "path": "v",
// }
func OverwriteDownload(data map[string]string) (err error) {
    err = rulesetSource.OverwriteDownload(data)
    changecontrol.ChangeControlInsertData(err, "OverwriteDownload")
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/rulesetSource/compareSourceFiles/:uuid \
func CompareFiles(uuid string) (mapData map[string]map[string]string, err error) {
    mapData,err = rulesetSource.CompareFiles(uuid)
    changecontrol.ChangeControlInsertData(err, "CompareFiles")
    return mapData,err
}

// // curl -X POST \
// //   https://52.47.197.22:50002/v1/rulesetSource/createNewFile \
// //   -H 'Content-Type: application/json' \
// //   -d '{
// //     "url": "v",
// //     "uuid": "v",
// //     "name": "v",
// //     "path": "v",
// // }
// func CreateNewFile(data map[string]string) (err error) {
//     err = rulesetSource.CreateNewFile(data)
//     changecontrol.ChangeControlInsertData(err, "")
    // return err
// }

// curl -X GET \
//   https://52.47.197.22:50002/v1/rulesetSource/getDetails/:uuid \
func GetDetails(uuid string) (files map[string]map[string]string, err error) {
    files, err = rulesetSource.GetDetails(uuid)
    changecontrol.ChangeControlInsertData(err, "GetDetails")
    return files ,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/rulesetSource/GetFileUUIDfromRulesetUUID/:uuid \
func GetFileUUIDfromRulesetUUID(value string)(uuid string, err error){
    uuid,err = rulesetSource.GetFileUUIDfromRulesetUUID(value)
    changecontrol.ChangeControlInsertData(err, "GetFileUUIDfromRulesetUUID")
    return uuid,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/rulesetSource/OverwriteRuleFile/:uuid \
func OverwriteRuleFile(uuid string)(err error){
    err = rulesetSource.OverwriteRuleFile(uuid)
    changecontrol.ChangeControlInsertData(err, "OverwriteRuleFile")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/rulesetSource/AddNewLinesToRuleset/:uuid \
func AddNewLinesToRuleset(uuid string)(err error){
    err = rulesetSource.AddNewLinesToRuleset(uuid)
    changecontrol.ChangeControlInsertData(err, "AddNewLinesToRuleset")
    return err
}