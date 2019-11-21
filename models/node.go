package models 

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/node"
    "owlhmaster/changeControl"
)

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/ \
// }
func GetAllNodes() (anode map[string]map[string]string, err error) {
    anode, err = node.GetAllNodes()
    return anode, err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/node/ \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "name": "v",
//     "port": "v",
//     "ip": "v"
// }
// }
func AddNode(n map[string]string) (err error) {
    err = node.AddNode(n)
    if err!=nil { 
        n["actionStatus"] = "error"
        n["errorDescription"] = err.Error()
    }else{
        n["actionStatus"] = "success"
    }
    n["action"] = "POST"
    n["actionDescription"] = "Add node"
    
    //add incident
    var controlError error
    controlError = changecontrol.InsertChangeControl(n)
    if controlError!=nil { logs.Error("AddNode controlError: "+controlError.Error()) }

    if err != nil {return err}

    return nil
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
// }
func UpdateNode(n map[string]string) (err error) {
    err = node.UpdateNode(n)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/ping/:uuid \
// }
func PingNode(n string) (err error) {
    err = node.NodePing(n)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/suricata/:uuid \
// }
func Suricata(n string) (data map[string]bool, err error) {
    logs.Info("Suricata status Node - %s",n)
    data,err = node.Suricata(n)
    return data,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/xeek/:uuid \
// }
func Zeek(n string) (data  map[string]bool, err error) {
    data,err = node.Zeek(n)
    return data,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/wazuh/:uuid \
// }
func Wazuh(n string) (data  map[string]bool, err error) {
    data,err = node.Wazuh(n)
    return data,err
}

// func GetSuricataBPF(n string) (data string, err error) {
//     data,err = node.GetSuricataBPF(n)
//     return data,err
// }

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/suricata/bpf \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "value": "v",
//     "service": "v"
// }
// }
func PutSuricataBPF(n map[string]string) (err error) {
    err = node.PutSuricataBPF(n)
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/:uuid \
// }
func DeleteNode(n string) (err error) {
    err = node.DeleteNode(n)
    return err
}

// func SetRuleset(nid string) (err error) {
//     err = node.SetRuleset(nid)
//     return err
// }

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/loadfile/:uuid/:fileName \
// }
func GetNodeFile(n map[string]string) (data map[string]string, err error) {
    data,err = node.GetNodeFile(n)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/savefile \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "file": "v",
//     "content": "v"
// }
// }
func SetNodeFile(n map[string]string) (err error) {
    err = node.SetNodeFile(n)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/getAllFiles/:uuid \
// }
func GetAllFiles(n string) (data map[string]string, err error) {
    data,err = node.GetAllFiles(n)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/RunSuricata/:uuid \
// }
func RunSuricata(uuid string) (data string, err error) {
    logs.Info("RunSuricata status Node")
    data,err = node.RunSuricata(uuid)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/StopSuricata/:uuid \
// }
func StopSuricata(uuid string) (data string, err error) {
    logs.Info("StopSuricata status Node")
    data,err = node.StopSuricata(uuid)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/RunZeek/:uuid \
// }
func RunZeek(uuid string) (data string, err error) {
    logs.Info("RunZeek status Node")
    data,err = node.RunZeek(uuid)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/StopZeek/:uuid \
// }
func StopZeek(uuid string) (data string, err error) {
    logs.Info("StopZeek status Node")
    data,err = node.StopZeek(uuid)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/RunWazuh/:uuid \
// }
func RunWazuh(uuid string) (data string, err error) {
    logs.Info("RunWazuh status Node")
    data,err = node.RunWazuh(uuid)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/StopWazuh/:uuid \
// }
func StopWazuh(uuid string) (data string, err error) {
    logs.Info("StopWazuh status Node")
    data,err = node.StopWazuh(uuid)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/ruleset/set  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "type": "v"
// }
// }
func SyncRulesetToNode(anode map[string]string)(err error){
    err = node.SyncRulesetToNode(anode)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/deploy/:nid  \
// }
func DeployZeek(uuid string)(err error){
    err = node.DeployZeek(uuid)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/ports/:nid \
// }
func ShowPorts(uuid string)(data map[string]map[string]string, err error){
    data, err = node.ShowPorts(uuid)
    return data, err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/ports/delete/:uuid \
// }
func DeletePorts(anode map[string]string, uuid string)(err error){
    err = node.DeletePorts(anode, uuid)
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/ports/deleteAll/:uuid \
// }
func DeleteAllPorts(uuid string)(err error){
    err = node.DeleteAllPorts(uuid)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/ports/PingPluginsNode/:uuid \
// }
func PingPluginsNode(uuid string)(data map[string]map[string]string, err error){
    data, err = node.PingPluginsNode(uuid)
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
func ChangeMode(anode map[string]string)(err error){
    err = node.ChangeMode(anode)
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
func ChangeStatus(anode map[string]string)(err error){
    err = node.ChangeStatus(anode)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/PingAnalyzer/:uuid \
// }
func PingAnalyzer(uuid string)(data map[string]string, err error){
    data, err = node.PingAnalyzer(uuid)
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
func ChangeAnalyzerStatus(uuid map[string]string)(err error){
    err = node.ChangeAnalyzerStatus(uuid)
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
func DeployNode(anode map[string]string)(err error){
    err = node.DeployNode(anode)
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
func UpdateNetworkInterface(anode map[string]string)(err error){
    err = node.UpdateNetworkInterface(anode)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/PingAnalyzer/:uuid \
// }
func GetServiceStatus(uuid string)(err error){
    err = node.GetServiceStatus(uuid)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/deployservice/:uuid  \
// }
func DeployService(uuid string)(err error){
    err = node.DeployService(uuid)
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
func ChangeDataflowValues(anode map[string]string)(err error){
    err = node.ChangeDataflowValues(anode)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/checkDeploy/:uuid \
// }
func CheckDeploy(uuid string)(anode map[string]string){
    anode = node.CheckDeploy(uuid)
    return anode
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/loadDataflowValues/:uuid \
// }
func LoadDataflowValues(uuid string)(data map[string]map[string]string, err error){
    data, err = node.LoadDataflowValues(uuid)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/loadNetworkValues/:uuid \
// }
func LoadNetworkValues(uuid string)(data map[string]string, err error){
    data, err = node.LoadNetworkValues(uuid)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/loadNetworkValuesSelected/:uuid \
// }
func LoadNetworkValuesSelected(uuid string)(data map[string]map[string]string, err error){
    data, err = node.LoadNetworkValuesSelected(uuid)
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
func SaveSocketToNetwork(anode map[string]string)(err error){
    err = node.SaveSocketToNetwork(anode)
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
func SaveNewLocal(anode map[string]string)(err error){
    err = node.SaveNewLocal(anode)
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
func SaveVxLAN(anode map[string]string)(err error){
    err = node.SaveVxLAN(anode)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/socketToNetworkList/:uuid \
// }
func SocketToNetworkList(uuid string)(data map[string]map[string]string, err error){
    data, err = node.SocketToNetworkList(uuid)
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
func SaveSocketToNetworkSelected(anode map[string]string)(err error){
    err = node.SaveSocketToNetworkSelected(anode)
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
func DeleteDataFlowValueSelected(anode map[string]string)(err error){
    err = node.DeleteDataFlowValueSelected(anode)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/pingmonitor/:uuid \
// }
func GetNodeMonitor(uuid string)(data map[string]interface{}, err error){
    data, err = node.GetNodeMonitor(uuid)
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
func AddPluginService(anode map[string]string)(err error){
    err = node.AddPluginService(anode)
    return err
}

// // curl -X GET \
// //   https://52.47.197.22:50002/v1/node/pingmonitor/:uuid \
// // }
// func GetSuricataServices(uuid string)(data map[string]map[string]string, err error){
//     data, err = node.GetSuricataServices(uuid)
//     return data, err
// }

// // curl -X GET \
// //   https://52.47.197.22:50002/v1/node/getMainconfData/:uuid \
// // }
func GetMainconfData(uuid string)(data map[string]map[string]string, err error){
    data, err = node.GetMainconfData(uuid)
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
func ChangeServiceStatus(anode map[string]string)(err error){
    err = node.ChangeServiceStatus(anode)
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
func ChangeMainServiceStatus(anode map[string]string)(err error){
    err = node.ChangeMainServiceStatus(anode)
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
func DeleteService(anode map[string]string)(err error){
    err = node.DeleteService(anode)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/saveSuricataInterface  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "service": "v",
//     "interface": "v",
//     "param": "v"
// }
// }
func SaveSuricataInterface(anode map[string]string)(err error){
    err = node.SaveSuricataInterface(anode)
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
func DeployStapService(anode map[string]string)(err error){
    err = node.DeployStapService(anode)
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
func StopStapService(anode map[string]string)(err error){
    err = node.StopStapService(anode)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/modifyStapValues  \
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
func ModifyStapValues(anode map[string]string)(err error){
    err = node.ModifyStapValues(anode)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/PingPorts/:uuid  \
func PingPorts(uuid string)(data map[string]map[string]string, err error){
    data, err = node.PingPorts(uuid)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/pingWazuhFiles/:uuid  \
func PingWazuhFiles(uuid string)(anode map[int]map[string]string, err error){
    anode,err = node.PingWazuhFiles(uuid)
    return anode,err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/deleteWazuhFile  \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "service": ["/tmp","/log"]
// }
// }
func DeleteWazuhFile(anode map[string]interface{})(err error){
    err = node.DeleteWazuhFile(anode)
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
func AddWazuhFile(anode map[string]interface{})(err error){
    err = node.AddWazuhFile(anode)
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
func LoadFileLastLines(anode map[string]string)(data map[string]string, err error){
    data,err = node.LoadFileLastLines(anode)
    return data,err
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
func SaveFileContentWazuh(anode map[string]string)(err error){
    err = node.SaveFileContentWazuh(anode)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/reloadFilesData/:uuid \
func ReloadFilesData(uuid string)(data map[string]map[string]string, err error){
    data, err = node.ReloadFilesData(uuid)
    return data, err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/node/monitor/addFile \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "path": "path"
// }
func AddMonitorFile(anode map[string]string)(err error){
    err = node.AddMonitorFile(anode)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/pingMonitorFiles/:uuid \
func PingMonitorFiles(uuid string)(data map[string]map[string]string, err error){
    data, err = node.PingMonitorFiles(uuid)
    return data, err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/monitor/deleteFile \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "file": "file"
// }
func DeleteMonitorFile(anode map[string]string)(err error){
    err = node.DeleteMonitorFile(anode)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/zeek/changeZeekMode \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "mode": "cluster"
// }
func ChangeZeekMode(anode map[string]string)(err error){
    err = node.ChangeZeekMode(anode)
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
func AddClusterValue(anode map[string]string)(err error){
    err = node.AddClusterValue(anode)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/zeek/pingCluster/:uuid \
// }
func PingCluster(uuid string)(data map[string]map[string]string, err error){
    data, err = node.PingCluster(uuid)
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
func EditClusterValue(anode map[string]string)(err error){
    err = node.EditClusterValue(anode)
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/node/zeek/deleteClusterValue \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "type": "v"
// }
func DeleteClusterValue(anode map[string]string)(err error){
    err = node.DeleteClusterValue(anode)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/zeek/syncCluster \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "type": "cluster"
// }
func SyncCluster(anode map[string]string)(err error){
    err = node.SyncCluster(anode)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/getChangeControlNode/:uuid \
// }
func GetChangeControlNode(uuid string)(data map[string]map[string]string, err error) {
    data, err = node.GetChangeControlNode(uuid)
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/incidents \
// }
func GetIncidentsNode(uuid string)(data map[string]map[string]string, err error){
    data,err = node.GetIncidentsNode(uuid)
    return data,err
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
func PutIncidentNode(anode map[string]string)(err error){
    err = node.PutIncidentNode(anode)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/plugin/changeSuricataTable \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "status": "v"
//  }
func ChangeSuricataTable(anode map[string]string)(err error){
    err = node.ChangeSuricataTable(anode)
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/ruleset/syncGroups \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v"
// }
func SyncRulesetToAllGroupNodes(anode map[string]string)(err error){
    err = node.SyncRulesetToAllGroupNodes(anode)
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
func SyncAnalyzerToAllGroupNodes(anode map[string]map[string]string)(log map[string]map[string]string, err error){
    log,err = node.SyncAnalyzerToAllGroupNodes(anode)
    return log,err
}