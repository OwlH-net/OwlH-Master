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

func Suricata (n string) (data []byte, err error) {
    logs.Info("Suricata status Node - %s",n)
    data,err = node.Suricata(n)
    return data,err
}

func DeleteNode (n string) (err error) {
    logs.Info("Models - delete Node -> name -> %s", n)
    err = node.DeleteNode(n)
    return err
}

