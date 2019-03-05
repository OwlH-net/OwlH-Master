package models 

import (
    "github.com/astaxie/beego/logs"
    "errors"
    "owlhmaster/node"
//    "database/sql"
//    "fmt"
//   "time"
//    _ "github.com/mattn/go-sqlite3"
//    "myapi/database"
)



func init() {
    logs.Info("NODE -> Init")
}



func InitNode() string {
    return "go"
}


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

func PingNode (n string) (err error) {
    err = node.NodePing(n)
    return err
}

func Suricata(n string) (data map[string]bool, err error) {
    logs.Info("Suricata status Node - %s",n)
    data,err = node.Suricata(n)
    return data,err
}

func GetSuricataBPF(n string) (data string, err error) {
    data,err = node.GetSuricataBPF(n)
    return data,err
}

func PutSuricataBPF(n map[string]string) (data string, err error) {
    data,err = node.PutSuricataBPF(n)
    return data,err
}

func Zeek(n string) (data []byte, err error) {
    data,err = node.Zeek(n)
    return data,err
}

func Wazuh(n string) (data []byte, err error) {
    data,err = node.Wazuh(n)
    return data,err
}

func DeleteNode(n string) (err error) {
    err = node.DeleteNode(n)
    return err
}

func SetRuleset(nid string) (err error) {
    err = node.SetRuleset(nid)
    return err
}

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