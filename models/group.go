package models 

import (
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
func AddGroupElement(data map[string]string)(err error) {
    err = group.AddGroupElement(data)
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
func EditGroup(data map[string]string)(err error) {
    err = group.EditGroup(data)
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/group/DeleteGroup/:uuid \
// }
func DeleteGroup(groupId string)(err error) {
    err = group.DeleteGroup(groupId)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group \
// }
func GetAllGroups()(Groups []group.Group, err error) {
    Groups, err = group.GetAllGroups()
    return Groups, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group/getAllNodesGroup/:uuid \
// }
func GetAllNodesGroup(uuid string)(data map[string]map[string]string, err error) {
    data, err = group.GetAllNodesGroup(uuid)
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
func AddGroupNodes(data map[string]interface{})(err error) {
    err = group.AddGroupNodes(data)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group/pingGroupNodes \
// }
func PingGroupNodes()(data map[string]map[string]string, err error) {
    data, err = group.PingGroupNodes()
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group/getNodeValues/:uuid \
// }
func GetNodeValues(uuid string)(data map[string]map[string]string, err error) {
    data, err = group.GetNodeValues(uuid)
    return data, err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/group/deleteNodeGroup/:uuid \
// }
func DeleteNodeGroup(uuid string)(err error) {
    err = group.DeleteNodeGroup(uuid)
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
func ChangeGroupRuleset(data map[string]string)(err error) {
    err = group.ChangeGroupRuleset(data)
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
func ChangePathsGroups(data map[string]string)(err error) {
    err = group.ChangePathsGroups(data)
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
func UpdateGroupService(data map[string]string)(err error) {
    err = group.UpdateGroupService(data)
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
func SyncPathGroup(data map[string]string)(err error) {
    err = group.SyncPathGroup(data)
    return err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/group/syncAll/:uuid  \
//   -H 'Content-Type: application/json' \
//   -d '{
//       "suricata-rulesets": {
//          "ruleset-group":"value",
//       },"suricata-config": {
//          "mastersuricata":"value",
//          "nodesuricata":"value",
//       },"suricata-services": {
//          "interface":"value",
//          "BPFfile":"value",
//          "BPFrule":"value",
//          "configFile":"value",
//          "commandLine":"value",
//       },"zeek-policies": {
//          "masterzeek":"value",
//          "nodezeek":"value",
//       },
// }
// }
func SyncAll(uuid string, data map[string]map[string]string)(err error) {
    err = group.SyncAll(uuid, data)
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
func AddCluster(data map[string]string)(err error) {
    err = group.AddCluster(data)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group/getClusterFiles/:uuid \
// }
func GetClusterFiles(uuid string)(data map[string]map[string]string, err error) {
    data, err = group.GetClusterFiles(uuid)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/group/getClusterFileContent/:uuid \
// }
func GetClusterFileContent(uuid string)(data map[string]string, err error) {
    data, err = group.GetClusterFileContent(uuid)
    return data, err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/group/deleteCluster \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d"
// }
// }
func DeleteCluster(data map[string]string)(err error) {
    err = group.DeleteCluster(data)
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
func ChangeClusterValue(data map[string]string)(err error) {
    err = group.ChangeClusterValue(data)
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
func SaveClusterFileContent(data map[string]string)(err error) {
    err = group.SaveClusterFileContent(data)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/syncClusterFile \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d"
// }
// }
func SyncClusterFile(data map[string]string)(err error) {
    err = group.SyncClusterFile(data)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/group/syncAllGroupCluster \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "d"
// }
// }
func SyncAllGroupCluster(data map[string]string)(err error) {
    err = group.SyncAllGroupCluster(data)
    return err
}