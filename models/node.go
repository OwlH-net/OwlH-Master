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

func GetSuricataBPF(n string) (data string, err error) {
    data,err = node.GetSuricataBPF(n)
    return data,err
}

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

func PingPorts(uuid string)(data map[string]map[string]string, err error){
    data, err = node.PingPorts(uuid)
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