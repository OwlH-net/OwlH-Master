package models 

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/node"
)

func GetAllNodes() (anode *map[string]map[string]string, err error) {
    anode, err = node.GetAllNodes()
    return anode, err
}

func AddNode(n map[string]string) (err error) {
    err = node.AddNode(n)
    return err
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

func ModifyStapValues(anode map[string]string)(err error){
    err = node.ModifyStapValues(anode)
    return err
}