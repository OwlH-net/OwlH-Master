package models

import (
  changecontrol "github.com/OwlH-net/OwlH-Master/changeControl"
  "github.com/OwlH-net/OwlH-Master/rulesetSource"
  // "github.com/astaxie/beego/logs"
)

//  curl -X POST \
//    https://52.47.197.22:50002/v1/rulesetSource/ \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "name": "v",
//      "desc": "v",
//      "url": "v",
//      "fileName": "v",
//      "type": "v",
//      "sourceType": "v",
//  }
//
// }
func CreateRulesetSource(data map[string]string, username string) (err error) {
  err = rulesetSource.CreateRulesetSource(data)
  changecontrol.ChangeControlInsertData(err, "CreateRulesetSource", username)
  return err
}

//  curl -X POST \
//    https://52.47.197.22:50002/v1/rulesetSource/custom \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "name": "v",
//      "desc": "v",
//      "url": "v",
//      "fileName": "v",
//      "type": "v",
//      "sourceType": "v",
//      "isDownloaded": "v",
//  }
func CreateCustomRulesetSource(data map[string]string, username string) (err error) {
  err = rulesetSource.CreateCustomRulesetSource(data)
  changecontrol.ChangeControlInsertData(err, "CreateCustomRulesetSource", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/rulesetSource/ \
func GetAllRulesetSource(hasPermissions bool, username string) (data map[string]map[string]string, err error) {
  data, err = rulesetSource.GetAllRulesetSource(hasPermissions)
  changecontrol.ChangeControlInsertData(err, "GetAllRulesetSource", username)
  return data, err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/rulesetSource/EditRulesetSource \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "name": "v",
//      "desc": "v",
//      "url": "v",
//      "uuid": "v",
//  }
func EditRulesetSource(data map[string]string, username string) (err error) {
  err = rulesetSource.EditRulesetSource(data)
  changecontrol.ChangeControlInsertData(err, "EditRulesetSource", username)
  return err
}

//  curl -X DELETE \
//    https://52.47.197.22:50002/v1/rulesetSource/DeleteRulesetSource \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "sourceType": "v",
//      "uuid": "v",
//  }
func DeleteRulesetSource(anode map[string]string, username string) (err error) {
  err = rulesetSource.DeleteRulesetSource(anode)
  changecontrol.ChangeControlInsertData(err, "DeleteRulesetSource", username)
  return err
}

//  curl -X DELETE \
//    https://52.47.197.22:50002/v1/rulesetSource/DeleteRulesetSource/:uuid \
func DeleteRulesetFile(uuid string, username string) (err error) {
  err = rulesetSource.DeleteRulesetFile(uuid)
  changecontrol.ChangeControlInsertData(err, "DeleteRulesetFile", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/rulesetSource/downloadFile \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "url": "v",
//      "uuid": "v",
//      "name": "v",
//      "path": "v",
//  }
func DownloadFile(data map[string]string, username string) (err error) {
  err = rulesetSource.DownloadFile(data)
  changecontrol.ChangeControlInsertData(err, "DownloadFile", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/rulesetSource/overwriteDownload \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "url": "v",
//      "uuid": "v",
//      "name": "v",
//      "path": "v",
//  }
func OverwriteDownload(data map[string]string, username string) (err error) {
  err = rulesetSource.OverwriteDownload(data)
  changecontrol.ChangeControlInsertData(err, "OverwriteDownload", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/rulesetSource/compareSourceFiles/:uuid \
func CompareFiles(uuid string, username string) (mapData map[string]map[string]string, err error) {
  mapData, err = rulesetSource.CompareFiles(uuid)
  changecontrol.ChangeControlInsertData(err, "CompareFiles", username)
  return mapData, err
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

//  curl -X GET \
//    https://52.47.197.22:50002/v1/rulesetSource/getDetails/:uuid \
func GetDetails(uuid string, username string) (files map[string]map[string]string, err error) {
  files, err = rulesetSource.GetDetails(uuid)
  changecontrol.ChangeControlInsertData(err, "GetDetails", username)
  return files, err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/rulesetSource/GetFileUUIDfromRulesetUUID/:uuid \
func GetFileUUIDfromRulesetUUID(value string, username string) (uuid string, err error) {
  uuid, err = rulesetSource.GetFileUUIDfromRulesetUUID(value)
  changecontrol.ChangeControlInsertData(err, "GetFileUUIDfromRulesetUUID", username)
  return uuid, err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/rulesetSource/OverwriteRuleFile/:uuid \
func OverwriteRuleFile(uuid string, username string) (err error) {
  err = rulesetSource.OverwriteRuleFile(uuid)
  changecontrol.ChangeControlInsertData(err, "OverwriteRuleFile", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/rulesetSource/AddNewLinesToRuleset/:uuid \
func AddNewLinesToRuleset(uuid string, username string) (err error) {
  err = rulesetSource.AddNewLinesToRuleset(uuid)
  changecontrol.ChangeControlInsertData(err, "AddNewLinesToRuleset", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/rulesetSource/loadDefaultRulesets \
func LoadDefaultRulesets(username string) (mapData map[string]map[string]string, err error) {
  mapData, err = rulesetSource.LoadDefaultRulesets()
  changecontrol.ChangeControlInsertData(err, "LoadDefaultRulesets", username)
  return mapData, err
}
