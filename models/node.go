package models

import (
    "errors"
    "github.com/astaxie/beego/logs"
    "owlhmaster/changeControl"
    "owlhmaster/node"
    "owlhmaster/nodeclient"
    "owlhmaster/utils"
)

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/ \
// }
func GetAllNodes(username string) (anode map[string]map[string]string, err error) {
    anode, err = node.GetAllNodes()
    changecontrol.ChangeControlInsertData(err, "Get all nodes", username)
    return anode, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/GetAllNodesReact \
// }
func GetAllNodesReact(username string) (anode node.NodeList, err error) {
    anode, err = node.GetAllNodesReact()
    changecontrol.ChangeControlInsertData(err, "Get all nodes", username)
    return anode, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/getAllTags \
// }
func GetAllTags(username string) (tags map[string]map[string]string, err error) {
    tags, err = node.GetAllTags()
    changecontrol.ChangeControlInsertData(err, "Get all tags", username)
    return tags, err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/node/ \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "name": "v",
//     "port": "v",
//     "ip": "v"
//     "nodeuser": "admin"
//     "nodepass": "admin"
// }
func AddNode(n map[string]string, username string) (err error) {
    err = node.AddNode(n)
    changecontrol.ChangeControlInsertData(err, "Add node", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/ \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "name": "v",
//     "port": "v",
//     "ip": "v",
//     "id": "v"
// }
func UpdateNode(n map[string]string, username string) (err error) {
    err = node.UpdateNode(n)
    changecontrol.ChangeControlInsertData(err, "Update node", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/ping/:uuid \
// }
func PingNode(n string, username string) (nodeResp map[string]string, err error) {
    nodeResp, err = node.NodePing(n)
    changecontrol.ChangeControlInsertData(err, "Ping node", username)
    return nodeResp, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/suricata/:uuid \
func Suricata(n string, username string) (data map[string]bool, err error) {
    data, err = node.Suricata(n)
    changecontrol.ChangeControlInsertData(err, "Ping Suricata", username)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/zeek/:uuid \
func Zeek(n string, username string) (data nodeclient.ZeekData, err error) {
    data, err = node.Zeek(n)
    changecontrol.ChangeControlInsertData(err, "Ping Zeek", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/zeek/:uuid/:cmd \
func ZeekCommand(n, cmd, username string) (data map[string]string, err error) {
    data, err = node.ZeekCommand(n, cmd)
    changecontrol.ChangeControlInsertData(err, "Run Zeek Command", username)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/wazuh/:uuid \
func Wazuh(n string, username string) (data map[string]bool, err error) {
    data, err = node.Wazuh(n)
    changecontrol.ChangeControlInsertData(err, "Ping Wazuh", username)
    return data, err
}

// func GetSuricataBPF(n string) (data string, err error) {
//     data,err = node.GetSuricataBPF(n)
//     changecontrol.ChangeControlInsertData(err, "")
// return errdata,err
// }

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/suricata/bpf \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "value": "v",
//     "service": "v"
// }
func PutSuricataBPF(n map[string]string, username string) (err error) {
    err = node.PutSuricataBPF(n)
    changecontrol.ChangeControlInsertData(err, "Put Suricata BPF", username)
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/:uuid \
func DeleteNode(n string, username string) (err error) {
    err = node.DeleteNode(n)
    changecontrol.ChangeControlInsertData(err, "Delete node", username)
    return err
}

// func SetRuleset(nid string) (err error) {
//     err = node.SetRuleset(nid)
//     changecontrol.ChangeControlInsertData(err, "")
// return errerr
// }

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/loadfile/:uuid/:fileName \
func GetNodeFile(n map[string]string, username string) (data map[string]string, err error) {
    data, err = node.GetNodeFile(n)
    changecontrol.ChangeControlInsertData(err, "Get node File", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/savefile \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "file": "v",
//     "content": "v"
// }
func SetNodeFile(n map[string]string, username string) (err error) {
    err = node.SetNodeFile(n)
    changecontrol.ChangeControlInsertData(err, "Sert node file", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/getAllFiles/:uuid \
// }
func GetAllFiles(n string, username string) (data map[string]string, err error) {
    data, err = node.GetAllFiles(n)
    changecontrol.ChangeControlInsertData(err, "Get all files", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/RunSuricata/:uuid \
// }
func RunSuricata(uuid string, username string) (data string, err error) {
    data, err = node.RunSuricata(uuid)
    changecontrol.ChangeControlInsertData(err, "Run suricata", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/StopSuricata/:uuid \
// }
func StopSuricata(uuid string, username string) (data string, err error) {
    data, err = node.StopSuricata(uuid)
    changecontrol.ChangeControlInsertData(err, "Stop Suricata", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/RunZeek/:uuid \
// }
func RunZeek(uuid string, username string) (data string, err error) {
    data, err = node.RunZeek(uuid)
    changecontrol.ChangeControlInsertData(err, "Run Zeek", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/StopZeek/:uuid \
// }
func StopZeek(uuid string, username string) (data string, err error) {
    data, err = node.StopZeek(uuid)
    changecontrol.ChangeControlInsertData(err, "Stop Zeek", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/RunWazuh/:uuid \
// }
func RunWazuh(uuid string, username string) (data string, err error) {
    data, err = node.RunWazuh(uuid)
    changecontrol.ChangeControlInsertData(err, "Run Wazuh", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/StopWazuh/:uuid \
// }
func StopWazuh(uuid string, username string) (data string, err error) {
    data, err = node.StopWazuh(uuid)
    changecontrol.ChangeControlInsertData(err, "Stop Wazuh", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/ruleset/set  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "type": "v"
// }
// }
func SyncRulesetToNode(anode map[string]string, username string) (err error) {
    err = node.SyncRulesetToNode(anode)
    changecontrol.ChangeControlInsertData(err, "Sync ruleset to node", username)
    return err
}

// // curl -X GET \
// //   https://52.47.197.22:50002/v1/node/deploy/:nid  \
// // }
// func DeployZeek(uuid string)(err error){
//     err = node.DeployZeek(uuid)
//     changecontrol.ChangeControlInsertData(err, "Deploy Zeek")
//     return err
// }

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/ports/:nid \
// }
func ShowPorts(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = node.ShowPorts(uuid)
    changecontrol.ChangeControlInsertData(err, "Show ports", username)
    return data, err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/ports/delete/:uuid \
// }
func DeletePorts(anode map[string]string, uuid string, username string) (err error) {
    err = node.DeletePorts(anode, uuid)
    changecontrol.ChangeControlInsertData(err, "Delete ports", username)
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/ports/deleteAll/:uuid \
// }
func DeleteAllPorts(uuid string, username string) (err error) {
    err = node.DeleteAllPorts(uuid)
    changecontrol.ChangeControlInsertData(err, "Delete all ports", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/ports/PingPluginsNode/:uuid \
// }
func PingPluginsNode(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = node.PingPluginsNode(uuid)
    changecontrol.ChangeControlInsertData(err, "Ping plugins node", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/mode  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "mode": "v"
// }
// }
func ChangeMode(anode map[string]string, username string) (err error) {
    err = node.ChangeMode(anode)
    changecontrol.ChangeControlInsertData(err, "change mode", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/status  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "status": "v"
// }
// }
func ChangeStatus(anode map[string]string, username string) (err error) {
    err = node.ChangeStatus(anode)
    changecontrol.ChangeControlInsertData(err, "Change status", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/PingAnalyzer/:uuid \
// }
func PingAnalyzer(uuid string, username string) (data map[string]string, err error) {
    data, err = node.PingAnalyzer(uuid)
    changecontrol.ChangeControlInsertData(err, "Ping analyzer", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/analyzer  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "status": "v",
//     "type": "v"
// }
// }
func ChangeAnalyzerStatus(uuid map[string]string, username string) (err error) {
    err = node.ChangeAnalyzerStatus(uuid)
    changecontrol.ChangeControlInsertData(err, "Change analyzer status", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/deployNode  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "value": "v"
// }
// }
func DeployNode(anode map[string]string, username string) (err error) {
    err = node.DeployNode(anode)
    changecontrol.ChangeControlInsertData(err, "Deploy node", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/updateNetworkInterface  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "value": "v",
//     "service": "v"
// }
// }
func UpdateNetworkInterface(anode map[string]string, username string) (err error) {
    err = node.UpdateNetworkInterface(anode)
    changecontrol.ChangeControlInsertData(err, "Update network interface", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/PingAnalyzer/:uuid \
// }
func GetServiceStatus(uuid string, username string) (err error) {
    err = node.GetServiceStatus(uuid)
    changecontrol.ChangeControlInsertData(err, "Get service status", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/deployservice/:uuid  \
// }
func DeployService(uuid string, username string) (err error) {
    err = node.DeployService(uuid)
    changecontrol.ChangeControlInsertData(err, "Deploy service", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/changeDataflowValues  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "value": "v",
//     "FlowUUID": "v",
//     "param": "v"
// }
// }
func ChangeDataflowValues(anode map[string]string, username string) (err error) {
    err = node.ChangeDataflowValues(anode)
    changecontrol.ChangeControlInsertData(err, "ChangeDataflowValues", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/checkDeploy/:uuid \
// }
func CheckDeploy(uuid string, username string) (anode map[string]string) {
    anode = node.CheckDeploy(uuid)
    var err error
    if anode == nil {
        err = errors.New("Error checking deploy")
    } else {
        err = nil
    }
    changecontrol.ChangeControlInsertData(err, "CheckDeploy", username)
    return anode
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/loadDataflowValues/:uuid \
// }
func LoadDataflowValues(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = node.LoadDataflowValues(uuid)
    changecontrol.ChangeControlInsertData(err, "LoadDataflowValues", username)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/loadNetworkValues/:uuid \
// }
func LoadNetworkValues(uuid string, username string) (data map[string]string, err error) {
    data, err = node.LoadNetworkValues(uuid)
    changecontrol.ChangeControlInsertData(err, "LoadNetworkValues", username)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/loadNetworkValuesSelected/:uuid \
// }
func LoadNetworkValuesSelected(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = node.LoadNetworkValuesSelected(uuid)
    changecontrol.ChangeControlInsertData(err, "LoadNetworkValuesSelected", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/saveSocketToNetwork  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "name": "v",
//     "cert": "v",
//     "port": "v",
//     "interface": "v"
// }
// }
func SaveSocketToNetwork(anode map[string]string, username string) (err error) {
    err = node.SaveSocketToNetwork(anode)
    changecontrol.ChangeControlInsertData(err, "SaveSocketToNetwork", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/saveNewLocal  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "name": "v",
//     "mtu": "v",
// }
// }
func SaveNewLocal(anode map[string]string, username string) (err error) {
    err = node.SaveNewLocal(anode)
    changecontrol.ChangeControlInsertData(err, "SaveNewLocal", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/saveVxLAN  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "interface": "v",
//     "lanIp": "v",
//     "localIp": "v",
//     "portIp": "v",
//     "type": "v",
//     "baseInterface": "v"
// }
// }
func SaveVxLAN(anode map[string]string, username string) (err error) {
    err = node.SaveVxLAN(anode)
    changecontrol.ChangeControlInsertData(err, "SaveVxLAN", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/socketToNetworkList/:uuid \
// }
func SocketToNetworkList(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = node.SocketToNetworkList(uuid)
    changecontrol.ChangeControlInsertData(err, "SocketToNetworkList", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/saveSocketToNetworkSelected  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "uuidNode": "v"
// }
// }
func SaveSocketToNetworkSelected(anode map[string]string, username string) (err error) {
    err = node.SaveSocketToNetworkSelected(anode)
    changecontrol.ChangeControlInsertData(err, "SaveSocketToNetworkSelected", username)
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/deleteDataFlowValueSelected  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "uuidNode": "v"
// }
// }
func DeleteDataFlowValueSelected(anode map[string]string, username string) (err error) {
    err = node.DeleteDataFlowValueSelected(anode)
    changecontrol.ChangeControlInsertData(err, "DeleteDataFlowValueSelected", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/pingmonitor/:uuid \
// }
func GetNodeMonitor(uuid string, username string) (data map[string]interface{}, err error) {
    data, err = node.GetNodeMonitor(uuid)
    changecontrol.ChangeControlInsertData(err, "GetNodeMonitor", username)
    return data, err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/node/add  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "name": "v",
//     "type": "v"
// }
// }
func AddPluginService(anode map[string]string, username string) (err error) {
    err = node.AddPluginService(anode)
    changecontrol.ChangeControlInsertData(err, "AddPluginService", username)
    return err
}

// // curl -X GET \
// //   https://52.47.197.22:50002/v1/node/pingmonitor/:uuid \
// // }
// func GetSuricataServices(uuid string)(data map[string]map[string]string, err error){
//     data, err = node.GetSuricataServices(uuid)
//     changecontrol.ChangeControlInsertData(err, "")
// return errdata, err
// }

// // curl -X GET \
// //   https://52.47.197.22:50002/v1/node/getMainconfData/:uuid \
// // }
func GetMainconfData(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = node.GetMainconfData(uuid)
    changecontrol.ChangeControlInsertData(err, "GetMainconfData", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/ChangeServiceStatus  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "status": "v",
//     "param": "v",
//     "service": "v",
//     "type": "v",
//     "interface": "v"
// }
// }
func ChangeServiceStatus(anode map[string]string, username string) (err error) {
    err = node.ChangeServiceStatus(anode)
    changecontrol.ChangeControlInsertData(err, "ChangeServiceStatus", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/mainconfStatus  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "service": "v",
//     "param": "v",
//     "status": "v"
// }
// }
func ChangeMainServiceStatus(anode map[string]string, username string) (err error) {
    err = node.ChangeMainServiceStatus(anode)
    changecontrol.ChangeControlInsertData(err, "ChangeMainServiceStatus", username)
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/deleteService  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "service": "v"
// }
// }
func DeleteService(anode map[string]string, username string) (err error) {
    err = node.DeleteService(anode)
    changecontrol.ChangeControlInsertData(err, "DeleteService", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/updateSuricataValue  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "service": "v",
//     "interface": "v",
//     "param": "v"
// }
// }
func UpdateSuricataValue(anode map[string]string, username string) (err error) {
    err = node.UpdateSuricataValue(anode)
    changecontrol.ChangeControlInsertData(err, "UpdateSuricataValue", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/deployStapService  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "service": "v",
//     "type": "v",
//     "interface": "v",
//     "port": "v",
//     "colector": "v"
// }
// }
func DeployStapService(anode map[string]string, username string) (err error) {
    err = node.DeployStapService(anode)
    changecontrol.ChangeControlInsertData(err, "DeployStapService", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/stopStapService  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "service": "v",
//     "type": "v"
// }
// }
func StopStapService(anode map[string]string, username string) (err error) {
    err = node.StopStapService(anode)
    changecontrol.ChangeControlInsertData(err, "StopStapService", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/modifyNodeOptionValues  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "service": "v",
//     "type": "v",
//     "name": "v",
//     "port": "v",
//     "cert": "v",
//     "": ...
// }
// }
func ModifyNodeOptionValues(anode map[string]string, username string) (err error) {
    err = node.ModifyNodeOptionValues(anode)
    changecontrol.ChangeControlInsertData(err, "ModifyNodeOptionValues", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/PingPorts/:uuid  \
func PingPorts(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = node.PingPorts(uuid)
    changecontrol.ChangeControlInsertData(err, "PingPorts", username)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/pingWazuhFiles/:uuid  \
func PingWazuhFiles(uuid string, username string) (anode map[string]map[string]string, err error) {
    anode, err = node.PingWazuhFiles(uuid)
    changecontrol.ChangeControlInsertData(err, "PingWazuhFiles", username)
    return anode, err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/deleteWazuhFile  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "service": ["/tmp","/log"]
// }
// }
func DeleteWazuhFile(anode map[string]interface{}, username string) (err error) {
    err = node.DeleteWazuhFile(anode)
    changecontrol.ChangeControlInsertData(err, "DeleteWazuhFile", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/addWazuhFile  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "paths": ["/tmp","/log"]
// }
// }
func AddWazuhFile(anode map[string]interface{}, username string) (err error) {
    err = node.AddWazuhFile(anode)
    changecontrol.ChangeControlInsertData(err, "AddWazuhFile", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/wazuh/loadLines   \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "path": "/tmp/log.log",
//     "number": "100"
// }
// }
func LoadFileLastLines(anode map[string]string, username string) (data map[string]string, err error) {
    data, err = node.LoadFileLastLines(anode)
    changecontrol.ChangeControlInsertData(err, "LoadFileLastLines", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/wazuh/saveFileContentWazuh \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "path": "/tmp/log.log",
//     "content": "content"
// }
// }
func SaveFileContentWazuh(anode map[string]string, username string) (err error) {
    err = node.SaveFileContentWazuh(anode)
    changecontrol.ChangeControlInsertData(err, "SaveFileContentWazuh", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/reloadFilesData/:uuid \
func ReloadFilesData(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = node.ReloadFilesData(uuid)
    changecontrol.ChangeControlInsertData(err, "ReloadFilesData", username)
    return data, err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/node/monitor/addFile \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "path": "path"
// }
func AddMonitorFile(anode map[string]string, username string) (err error) {
    err = node.AddMonitorFile(anode)
    changecontrol.ChangeControlInsertData(err, "AddMonitorFile", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/pingMonitorFiles/:uuid \
func PingMonitorFiles(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = node.PingMonitorFiles(uuid)
    changecontrol.ChangeControlInsertData(err, "PingMonitorFiles", username)
    return data, err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/monitor/deleteFile \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "file": "file"
// }
func DeleteMonitorFile(anode map[string]string, username string) (err error) {
    err = node.DeleteMonitorFile(anode)
    changecontrol.ChangeControlInsertData(err, "DeleteMonitorFile", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/zeek/changeZeekMode \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "mode": "cluster"
// }
func ChangeZeekMode(anode map[string]string, username string) (err error) {
    err = node.ChangeZeekMode(anode)
    changecontrol.ChangeControlInsertData(err, "ChangeZeekMode", username)
    return err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/node/zeek/addClusterValue \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "type": "cluster"
//     "host": "localhost"
//     "interface": "eth0"
// }
func AddClusterValue(anode map[string]string, username string) (err error) {
    err = node.AddClusterValue(anode)
    changecontrol.ChangeControlInsertData(err, "AddClusterValue", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/zeek/pingCluster/:uuid \
// }
func PingCluster(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = node.PingCluster(uuid)
    changecontrol.ChangeControlInsertData(err, "PingCluster", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/zeek/editClusterValue \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "type": "v",
//     "host": "v",
//     "interface": "cluster"
// }
func EditClusterValue(anode map[string]string, username string) (err error) {
    err = node.EditClusterValue(anode)
    changecontrol.ChangeControlInsertData(err, "EditClusterValue", username)
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/zeek/deleteClusterValue \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "type": "v"
// }
func DeleteClusterValue(anode map[string]string, username string) (err error) {
    err = node.DeleteClusterValue(anode)
    changecontrol.ChangeControlInsertData(err, "DeleteClusterValue", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/zeek/syncCluster \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "type": "cluster"
// }
func SyncCluster(anode map[string]string, username string) (err error) {
    err = node.SyncCluster(anode)
    changecontrol.ChangeControlInsertData(err, "SyncCluster", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/getChangeControlNode/:uuid \
// }
func GetChangeControlNode(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = node.GetChangeControlNode(uuid)
    changecontrol.ChangeControlInsertData(err, "GetChangeControlNode", username)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/incidents \
// }
func GetIncidentsNode(uuid string, username string) (data map[string]map[string]string, err error) {
    data, err = node.GetIncidentsNode(uuid)
    changecontrol.ChangeControlInsertData(err, "GetIncidentsNode", username)
    return data, err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/node/incidents \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "nodeuuid": "d",
//     "uuid": "v",
//     "param": "v",
//     "value": "v",
// }
func PutIncidentNode(anode map[string]string, username string) (err error) {
    err = node.PutIncidentNode(anode)
    changecontrol.ChangeControlInsertData(err, "PutIncidentNode", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/plugin/changeSuricataTable \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "status": "v"
//  }
func ChangeSuricataTable(anode map[string]string, username string) (err error) {
    err = node.ChangeSuricataTable(anode)
    changecontrol.ChangeControlInsertData(err, "ChangeSuricataTable", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/ruleset/syncGroups \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v"
// }
func SyncRulesetToAllGroupNodes(anode map[string]string, username string) (err error) {
    err = node.SyncRulesetToAllGroupNodes(anode)
    changecontrol.ChangeControlInsertData(err, "SyncRulesetToAllGroupNodes", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/analyzer/sync \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v"
//     "nodes": {
//          "key":"value"
//     }
// }
func SyncAnalyzerToAllGroupNodes(anode map[string]string, username string) (log map[string]map[string]string, err error) {
    log, err = node.SyncAnalyzerToAllGroupNodes(anode)
    changecontrol.ChangeControlInsertData(err, "SyncAnalyzerToAllGroupNodes", username)
    return log, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/StartSuricataMain \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v"
// }
func StartSuricataMainConf(uuid map[string]string, username string) (err error) {
    err = node.StartSuricataMainConf(uuid)
    changecontrol.ChangeControlInsertData(err, "StartSuricataMainConf", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/StopSuricataMain \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v"
// }
func StopSuricataMainConf(uuid map[string]string, username string) (err error) {
    err = node.StopSuricataMainConf(uuid)
    changecontrol.ChangeControlInsertData(err, "StopSuricataMainConf", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/KillSuricataMain \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "pid": "v"
// }
func KillSuricataMainConf(uuid map[string]string, username string) (err error) {
    err = node.KillSuricataMainConf(uuid)
    changecontrol.ChangeControlInsertData(err, "KillSuricataMainConf", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/ReloadSuricataMain \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "pid": "v"
// }
func ReloadSuricataMainConf(uuid map[string]string, username string) (err error) {
    err = node.ReloadSuricataMainConf(uuid)
    changecontrol.ChangeControlInsertData(err, "ReloadSuricataMainConf", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/LaunchZeekMainConf \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "param": "v"
// }
func LaunchZeekMainConf(uuid map[string]string, username string) (err error) {
    err = node.LaunchZeekMainConf(uuid)
    changecontrol.ChangeControlInsertData(err, "LaunchZeekMainConf", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/zeek/syncZeekValues \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v"
// }
func SyncZeekValues(anode map[string]string, username string) (err error) {
    err = node.SyncZeekValues(anode)
    changecontrol.ChangeControlInsertData(err, "SyncZeekValues", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/monitor/changeRotationStatus \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "file": "v"}
func ChangeRotationStatus(n map[string]string, username string) (err error) {
    err = node.ChangeRotationStatus(n)
    changecontrol.ChangeControlInsertData(err, "ChangeRotationStatus", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/monitor/editRotation \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "file": "v"
//     "path": "v"
//     "size": "v"
//     "lines": "v"
//     "days": "v"
func EditRotation(n map[string]string, username string) (err error) {
    err = node.EditRotation(n)
    changecontrol.ChangeControlInsertData(err, "EditRotation", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/plugins/getCommands  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "service": "v"
// }
// }
func GetServiceCommands(anode map[string]string, username string) (data map[string]map[string]string, err error) {
    data, err = node.GetServiceCommands(anode)
    changecontrol.ChangeControlInsertData(err, "GetServiceCommands", username)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/setRuleset  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "rule_uuid": "v",
//     "service": "v"
//     "source": "v"
//     "name": "v"
// }
func SaveSurictaRulesetSelected(anode map[string]string, username string) (err error) {
    err = node.SaveSurictaRulesetSelected(anode)
    changecontrol.ChangeControlInsertData(err, "SaveSurictaRulesetSelected", username)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/registerNode/:uuid \
func RegisterNode(uuid string, username string) (err error) {
    err = node.RegisterNode(uuid)
    changecontrol.ChangeControlInsertData(err, "RegisterNode", username)
    return err
}

func NodeEnrollment(nodeDetails utils.NodeData) (uuid string, success bool, details map[string]string) {
    logs.Info("lets enroll node %+v", nodeDetails)
    uuid, enroled, details := node.EnrollNode(nodeDetails)
    return uuid, enroled, details
}

func AssignNodeToGroup(uuid string, groupDetails utils.GroupData) (guuid string, success bool, details map[string]string) {
    logs.Info("lets assign node %s to %s", uuid, groupDetails.UUID)
    guuid, assigned, details := node.AssignNodeToGroup(uuid, groupDetails)
    return guuid, assigned, details
}

func CreateSuricataService(guuid, uuid string, suricataDetails utils.SuricataData) (success bool, details map[string]string) {
    logs.Info("lets create suricata service to node %s to %s", uuid, suricataDetails.Name)
    assigned, details := node.CreateSuricataService(guuid, uuid, suricataDetails)
    return assigned, details
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/node/enrollNewNode  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "NODE": {},
//     "GROUPS": [],
//     "SURICATA": {}
// }
func EnrollNewNode(anode utils.EnrollNewNodeStruct, username string) (err error) {
    err = node.EnrollNewNode(anode)
    changecontrol.ChangeControlInsertData(err, "EnrollNewNode", username)
    return err
}