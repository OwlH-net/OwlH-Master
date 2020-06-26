package models

import (
    "owlhmaster/changeControl"
    "owlhmaster/group"
)

// curl -X POST \
//   https://52.47.197.22:50001/v1/node/incidents \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "name": "d",
//     "desc": "v",
//     "type": "v",
//     "ruleset": "",
//     "rulesetID": "",
//     "mastersuricata": "",
//     "nodesuricata": "",
//     "masterzeek": "",
//     "nodezeek": "",
//     "interface": "",
//     "BPFfile": "",
//     "BPFrule": "",
//     "configFile": "",
//     "commandLine": "",
// }
// }
func AddGroupElement(data map[string]string, username string) (err error) {
    err = group.AddGroupElement(data)
    changecontrol.ChangeControlInsertData(err, "AddGroupElement", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/editGroup \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "name": "d",
//     "uuid": "d",
//     "desc": "d",
// }
// }
func EditGroup(data map[string]string, username string) (err error) {
    err = group.EditGroup(data)
    changecontrol.ChangeControlInsertData(err, "EditGroup", username)
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/group/DeleteGroup/:uuid \
// }
func DeleteGroup(groupId string, username string) (err error) {
    err = group.DeleteGroup(groupId)
    changecontrol.ChangeControlInsertData(err, "DeleteGroup", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group \
// }
func GetAllGroups(username string) (Groups []group.Group, err error) {
    Groups, err = group.GetAllGroups()
    changecontrol.ChangeControlInsertData(err, "GetAllGroups", username)
    return Groups, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group/getAllNodesGroup/:uuid \
// }
func GetAllNodesGroup(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = group.GetAllNodesGroup(uuid)
    changecontrol.ChangeControlInsertData(err, "GetAllNodesGroup", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/addGroupNodes \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d",
//     "nodes": "[uuid1,uuid2]",
// }
// }
func AddGroupNodes(data map[string]interface{}, username string) (err error) {
    err = group.AddGroupNodes(data)
    changecontrol.ChangeControlInsertData(err, "AddGroupNodes", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group/pingGroupNodes \
// }
func PingGroupNodes(username string) (data map[string]map[string]string, err error) {
    data, err = group.PingGroupNodes()
    changecontrol.ChangeControlInsertData(err, "PingGroupNodes", username)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group/getNodeValues/:uuid \
// }
func GetNodeValues(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = group.GetNodeValues(uuid)
    changecontrol.ChangeControlInsertData(err, "GetNodeValues", username)
    return data, err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/group/deleteNodeGroup/:uuid \
// }
func DeleteNodeGroup(uuid string, username string) (err error) {
    err = group.DeleteNodeGroup(uuid)
    changecontrol.ChangeControlInsertData(err, "DeleteNodeGroup", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/changeGroupRuleset \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d",
//     "ruleset": "ruleset",
//     "rulesetID": "rulesetID",
// }
// }
func ChangeGroupRuleset(data map[string]string, username string) (err error) {
    err = group.ChangeGroupRuleset(data)
    changecontrol.ChangeControlInsertData(err, "ChangeGroupRuleset", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/changePaths \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d",
//     "type": "suricata",
//     "mastersuricata": "/tmp/log.log",
//     "nodesuricata": "/tmp/log.log",
// }
// }
// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/changePaths \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d",
//     "type": "zeek",
//     "masterzeek": "/tmp/log.log",
//     "nodezeek": "/tmp/log.log",
// }
// }
func ChangePathsGroups(data map[string]string, username string) (err error) {
    err = group.ChangePathsGroups(data)
    changecontrol.ChangeControlInsertData(err, "ChangePathsGroups", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/updateGroupService \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d",
//     "param": "param",
//     "value": "value",
// }
// }
func UpdateGroupService(data map[string]string, username string) (err error) {
    err = group.UpdateGroupService(data)
    changecontrol.ChangeControlInsertData(err, "UpdateGroupService", username)
    return err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/group/syncPathGroup \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d",
//     "type": "zeek",
//     "masterzeek": "/tmp/log.log",
//     "nodezeek": "/tmp/log.log",
// }
// }
// curl -X POST \
//   https://52.47.197.22:50002/v1/group/syncPathGroup \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d",
//     "type": "suricata",
//     "mastersuricata": "/tmp/log.log",
//     "nodesuricata": "/tmp/log.log",
// }
// }
func SyncPathGroup(data map[string]string, username string) (err error) {
    err = group.SyncPathGroup(data)
    changecontrol.ChangeControlInsertData(err, "SyncPathGroup", username)
    return err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/group/syncAll/:uuid  \
//   -H 'Content-Type: application/json' \
func SyncAll(uuid string, username string) (err error) {
    err = group.SyncAll(uuid)
    changecontrol.ChangeControlInsertData(err, "SyncAll", username)
    return err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/group/addCluster \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d",
//     "path": "/tmp/log.log",
// }
// }
func AddCluster(data map[string]string, username string) (err error) {
    err = group.AddCluster(data)
    changecontrol.ChangeControlInsertData(err, "AddCluster", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group/getClusterFiles/:uuid \
// }
func GetClusterFiles(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = group.GetClusterFiles(uuid)
    changecontrol.ChangeControlInsertData(err, "GetClusterFiles", username)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group/getClusterFileContent/:uuid \
// }
func GetClusterFileContent(uuid string, username string) (data map[string]string, err error) {
    data, err = group.GetClusterFileContent(uuid)
    changecontrol.ChangeControlInsertData(err, "GetClusterFileContent", username)
    return data, err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/group/deleteCluster \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d"
// }
// }
func DeleteCluster(data map[string]string, username string) (err error) {
    err = group.DeleteCluster(data)
    changecontrol.ChangeControlInsertData(err, "DeleteCluster", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/changeClusterValue \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d"
//     "guuid": "d"
//     "path": "d"
// }
// }
func ChangeClusterValue(data map[string]string, username string) (err error) {
    err = group.ChangeClusterValue(data)
    changecontrol.ChangeControlInsertData(err, "ChangeClusterValue", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/saveClusterFileContent \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d"
//     "content": "d"
//     "path": "d"
// }
// }
func SaveClusterFileContent(data map[string]string, username string) (err error) {
    err = group.SaveClusterFileContent(data)
    changecontrol.ChangeControlInsertData(err, "SaveClusterFileContent", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/syncClusterFile \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d"
// }
// }
func SyncClusterFile(data map[string]string, username string) (err error) {
    err = group.SyncClusterFile(data)
    changecontrol.ChangeControlInsertData(err, "SyncClusterFile", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/syncAllGroupCluster \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d"
// }
// }
func SyncAllGroupCluster(data map[string]string, username string) (err error) {
    err = group.SyncAllGroupCluster(data)
    changecontrol.ChangeControlInsertData(err, "SyncAllGroupCluster", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/syncAllSuricataGroup \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d"
// }
func SyncAllSuricataGroup(data map[string]string, username string) (err error) {
    err = group.SyncAllSuricataGroup(data)
    changecontrol.ChangeControlInsertData(err, "SyncAllSuricataGroup", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/suricata \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d",
//     "action": "stop" //or start
// }
func SuricataGroupService(data map[string]string, username string) (err error) {
    err = group.SuricataGroupService(data)
    changecontrol.ChangeControlInsertData(err, "SuricataGroupService", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/GetMD5files \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "sadfasdfasdfasdfasdfasdf",
//     "type": "suricata",
//     "mastersuricata": "/path/master",
//     "nodesuricata": "/path/node"
// }
func GetMD5files(data map[string]string, username string) (allData map[string]map[string]map[string]string, err error) {
    allData, err = group.GetMD5files(data)
    changecontrol.ChangeControlInsertData(err, "GetMD5files", username)
    return allData, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group/suricata/status/:uuid \
// }
func SuricataNodesStatus(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = group.SuricataNodesStatus(uuid)
    changecontrol.ChangeControlInsertData(err, "SuricataNodesStatus", username)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group/getGroupSelectedRulesets/:uuid \
// }
func GetGroupSelectedRulesets(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = group.GetGroupSelectedRulesets(uuid)
    changecontrol.ChangeControlInsertData(err, "GetGroupSelectedRulesets", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/addRulesetsToGroup \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "sadfasdfasdfasdfasdfasdf",
//     "rulesets": "ruleset array",
// }
func AddRulesetsToGroup(data map[string]string, username string) (err error) {
    err = group.AddRulesetsToGroup(data)
    changecontrol.ChangeControlInsertData(err, "AddRulesetsToGroup", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/deleteExpertGroupRuleset \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "sadfasdfasdfasdfasdfasdf",
//     "ruleset": "sadfasdfasdfasdfasdfasdf",
// }
func DeleteExpertGroupRuleset(data map[string]string, username string) (err error) {
    err = group.DeleteExpertGroupRuleset(data)
    changecontrol.ChangeControlInsertData(err, "DeleteExpertGroupRuleset", username)
    return err
}

func SetDefault(uuid string) (err error) {
    group.SetDefaultGroup(uuid)
    return nil
}
