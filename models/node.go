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

var (
        anode Node
)

func init() {
    logs.Info("NODE -> Init")
}


type Node struct {
    NId     string
    NName   string
    NIp     string
    NPort   int
    NType   string
    NUUID   string
}

func InitNode() string {
    return "go"
}

func GetNode(n string) (anode *Node, err error) {
    var tnode Node
    if n == "" {
        logs.Info("NODE GET -> IN")
        return nil, errors.New("no way to search node")
    }
    tnode.NId = "12"
    tnode.NName = "hola"
    tnode.NIp = "192.168.1.1"
    tnode.NPort = 50
    tnode.NType = "NODE"
    tnode.NUUID = "xxxxxx.xxx.xxxxx.xxxx"
    return &tnode, nil
}

func GetAllNodes() (anode *map[string]map[string]string, err error) {
    logs.Info("models - get all nodes - in")
    anode, err = node.GetAllNodes()
    return anode, err
}

func AddNode(n map[string]string) (err error) {
    err = node.AddNode(n)
    return err
}

func UpdateNode(n map[string]string) (err error) {
    logs.Info("Models - Update Node -> name -> %s", n["name"])
    err = node.UpdateNode(n)
    return err
}

func PingNode (n string) (err error) {
    logs.Info("Ping Node - %s",n)
    err = node.NodePing(n)
    return err
}

func Suricata(n string) (data []byte, err error) {
    logs.Info("Suricata status Node - %s",n)
    data,err = node.Suricata(n)
    return data,err
}

func GetSuricataBPF(n string) (data string, err error) {
    logs.Info("Get Suricata BPF Node - %s",n)
    data,err = node.GetSuricataBPF(n)
    return data,err
}

func PutSuricataBPF(n map[string]string) (data string, err error) {
    logs.Info("Put Suricata BPF Node - %s",n)
    data,err = node.PutSuricataBPF(n)
    return data,err
}

func Zeek(n string) (data []byte, err error) {
    logs.Info("Zeek status Node - %s",n)
    data,err = node.Zeek(n)
    return data,err
}

func Wazuh(n string) (data []byte, err error) {
    logs.Info("Wazuh status Node - %s",n)
    data,err = node.Wazuh(n)
    return data,err
}

func DeleteNode(n string) (err error) {
    logs.Info("Models - delete Node -> name -> %s", n)
    err = node.DeleteNode(n)
    return err
}

func SetRuleset(nid string) (err error) {
    logs.Info("Models - SetRuleset Node -> name -> %s", nid)
    err = node.SetRuleset(nid)
    return err
}

func GetNodeFile(n map[string]string) (data map[string]string, err error) {
    logs.Info("GetNodeFile status Node")
    data,err = node.GetNodeFile(n)
    return data,err
}

func SetNodeFile(n map[string]string) (err error) {
    logs.Info("SetNodeFile status Node")
    err = node.SetNodeFile(n)
    return err
}