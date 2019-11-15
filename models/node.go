package models 

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/node"
    "owlhmaster/changeControl"
)

func GetAllNodes() (anode map[string]map[string]string, err error) {
    anode, err = node.GetAllNodes()
    return anode, err
}

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

func UpdateNode(n map[string]string) (err error) {
    err = node.UpdateNode(n)
    return err
}

func PingNode(n string) (err error) {
    err = node.NodePing(n)
    return err
}

func Suricata(n string) (data map[string]bool, err error) {
    logs.Info("Suricata status Node - %s",n)
    data,err = node.Suricata(n)
    return data,err
}

func Zeek(n string) (data  map[string]bool, err error) {
    data,err = node.Zeek(n)
    return data,err
}

func Wazuh(n string) (data  map[string]bool, err error) {
    data,err = node.Wazuh(n)
    return data,err
}

// func GetSuricataBPF(n string) (data string, err error) {
//     data,err = node.GetSuricataBPF(n)
//     return data,err
// }

func PutSuricataBPF(n map[string]string) (err error) {
    err = node.PutSuricataBPF(n)
    return err
}


func DeleteNode(n string) (err error) {
    err = node.DeleteNode(n)
    return err
}

// func SetRuleset(nid string) (err error) {
//     err = node.SetRuleset(nid)
//     return err
// }

func GetNodeFile(n map[string]string) (data map[string]string, err error) {
    data,err = node.GetNodeFile(n)
    return data,err
}

func SetNodeFile(n map[string]string) (err error) {
    err = node.SetNodeFile(n)
    return err
}

func GetAllFiles(n string) (data map[string]string, err error) {
    data,err = node.GetAllFiles(n)
    return data,err
}

func RunSuricata(uuid string) (data string, err error) {
    logs.Info("RunSuricata status Node")
    data,err = node.RunSuricata(uuid)
    return data,err
}

func StopSuricata(uuid string) (data string, err error) {
    logs.Info("StopSuricata status Node")
    data,err = node.StopSuricata(uuid)
    return data,err
}

func RunZeek(uuid string) (data string, err error) {
    logs.Info("RunZeek status Node")
    data,err = node.RunZeek(uuid)
    return data,err
}

func StopZeek(uuid string) (data string, err error) {
    logs.Info("StopZeek status Node")
    data,err = node.StopZeek(uuid)
    return data,err
}

func RunWazuh(uuid string) (data string, err error) {
    logs.Info("RunWazuh status Node")
    data,err = node.RunWazuh(uuid)
    return data,err
}

func StopWazuh(uuid string) (data string, err error) {
    logs.Info("StopWazuh status Node")
    data,err = node.StopWazuh(uuid)
    return data,err
}

func SyncRulesetToNode(anode map[string]string)(err error){
    err = node.SyncRulesetToNode(anode)
    return err
}

func DeployZeek(uuid string)(err error){
    err = node.DeployZeek(uuid)
    return err
}

func ShowPorts(uuid string)(data map[string]map[string]string, err error){
    data, err = node.ShowPorts(uuid)
    return data, err
}

func DeletePorts(anode map[string]string, uuid string)(err error){
    err = node.DeletePorts(anode, uuid)
    return err
}

func DeleteAllPorts(uuid string)(err error){
    err = node.DeleteAllPorts(uuid)
    return err
}

func PingPluginsNode(uuid string)(data map[string]map[string]string, err error){
    data, err = node.PingPluginsNode(uuid)
    return data, err
}

func ChangeMode(anode map[string]string)(err error){
    err = node.ChangeMode(anode)
    return err
}
func ChangeStatus(anode map[string]string)(err error){
    err = node.ChangeStatus(anode)
    return err
}

func PingAnalyzer(uuid string)(data map[string]string, err error){
    data, err = node.PingAnalyzer(uuid)
    return data, err
}

func ChangeAnalyzerStatus(uuid map[string]string)(err error){
    err = node.ChangeAnalyzerStatus(uuid)
    return err
}

func DeployNode(anode map[string]string)(err error){
    err = node.DeployNode(anode)
    return err
}

func UpdateNetworkInterface(anode map[string]string)(err error){
    err = node.UpdateNetworkInterface(anode)
    return err
}

func GetServiceStatus(uuid string)(err error){
    err = node.GetServiceStatus(uuid)
    return err
}

func DeployService(uuid string)(err error){
    err = node.DeployService(uuid)
    return err
}

func ChangeDataflowValues(anode map[string]string)(err error){
    err = node.ChangeDataflowValues(anode)
    return err
}

func CheckDeploy(uuid string)(anode map[string]string){
    anode = node.CheckDeploy(uuid)
    return anode
}

func LoadDataflowValues(uuid string)(data map[string]map[string]string, err error){
    data, err = node.LoadDataflowValues(uuid)
    return data, err
}

func LoadNetworkValues(uuid string)(data map[string]string, err error){
    data, err = node.LoadNetworkValues(uuid)
    return data, err
}

func LoadNetworkValuesSelected(uuid string)(data map[string]map[string]string, err error){
    data, err = node.LoadNetworkValuesSelected(uuid)
    return data, err
}

func SaveSocketToNetwork(anode map[string]string)(err error){
    err = node.SaveSocketToNetwork(anode)
    return err
}

func SaveNewLocal(anode map[string]string)(err error){
    err = node.SaveNewLocal(anode)
    return err
}

func SaveVxLAN(anode map[string]string)(err error){
    err = node.SaveVxLAN(anode)
    return err
}

func SocketToNetworkList(uuid string)(data map[string]map[string]string, err error){
    data, err = node.SocketToNetworkList(uuid)
    return data, err
}

func SaveSocketToNetworkSelected(anode map[string]string)(err error){
    err = node.SaveSocketToNetworkSelected(anode)
    return err
}

func DeleteDataFlowValueSelected(anode map[string]string)(err error){
    err = node.DeleteDataFlowValueSelected(anode)
    return err
}

func GetNodeMonitor(uuid string)(data map[string]interface{}, err error){
    data, err = node.GetNodeMonitor(uuid)
    return data, err
}

func AddPluginService(anode map[string]string)(err error){
    err = node.AddPluginService(anode)
    return err
}

func GetSuricataServices(uuid string)(data map[string]map[string]string, err error){
    data, err = node.GetSuricataServices(uuid)
    return data, err
}

func GetMainconfData(uuid string)(data map[string]map[string]string, err error){
    data, err = node.GetMainconfData(uuid)
    return data, err
}

func ChangeServiceStatus(anode map[string]string)(err error){
    err = node.ChangeServiceStatus(anode)
    return err
}

func ChangeMainServiceStatus(anode map[string]string)(err error){
    err = node.ChangeMainServiceStatus(anode)
    return err
}

func DeleteService(anode map[string]string)(err error){
    err = node.DeleteService(anode)
    return err
}

func SaveSuricataInterface(anode map[string]string)(err error){
    err = node.SaveSuricataInterface(anode)
    return err
}

func DeployStapService(anode map[string]string)(err error){
    err = node.DeployStapService(anode)
    return err
}

func StopStapService(anode map[string]string)(err error){
    err = node.StopStapService(anode)
    return err
}

func ModifyStapValues(anode map[string]string)(err error){
    err = node.ModifyStapValues(anode)
    return err
}

func PingPorts(uuid string)(data map[string]map[string]string, err error){
    data, err = node.PingPorts(uuid)
    return data, err
}

func PingWazuhFiles(uuid string)(anode map[int]map[string]string, err error){
    anode,err = node.PingWazuhFiles(uuid)
    return anode,err
}

func DeleteWazuhFile(anode map[string]interface{})(err error){
    err = node.DeleteWazuhFile(anode)
    return err
}
func AddWazuhFile(anode map[string]interface{})(err error){
    err = node.AddWazuhFile(anode)
    return err
}

func LoadFileLastLines(anode map[string]string)(data map[string]string, err error){
    data,err = node.LoadFileLastLines(anode)
    return data,err
}

func SaveFileContentWazuh(anode map[string]string)(err error){
    err = node.SaveFileContentWazuh(anode)
    return err
}

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
// }
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