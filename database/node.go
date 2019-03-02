package ndb

import (
    "github.com/astaxie/beego/logs"
    "database/sql"
//    "fmt"
//   "time"
    _ "github.com/mattn/go-sqlite3"
    "errors"
)


type Node struct {
    NId       string
    NName     string
    NIp       string
    NPort     int
    NType     string
    NUUID     string
}

func AddNode(node Node) (err error){
    logs.Info("DB -> Add Node")
    if Db != nil {
        stmt, err := Db.Prepare("INSERT INTO node(node_name, node_ip, node_port, node_type, node_UUID) values(?,?,?,?,?)")
        if err != nil {
            logs.Error("DB NODE -> Db.Prepare failure. ")
            return errors.New("DB NODE -> Db.Prepare failure")
        }
        //Validate fields!!!
        res, err := stmt.Exec(node.NName, node.NIp, node.NPort, node.NType, node.NUUID )
        logs.Info("DB -> info", res)
        if err != nil {
            logs.Error("DB NODE -> There was a problem with Query: %s", err.Error())
            return errors.New("DB NODE -> There was a problem with Query: " + err.Error())
        }
        return nil
    } else {
        logs.Error("DB NODE -> No access to database")
        return errors.New("DB NODE -> No access to database")
    }
}

func GetNode(nid string) (n *Node, err error) {
    var node Node
    if Db != nil {
        row := Db.QueryRow("SELECT * FROM node WHERE node_id=%s;",nid)
        logs.Info ("DB -> Row %s", row)
        err = row.Scan(&node.NId, &node.NName, &node.NIp, &node.NPort, &node.NType, &node.NUUID)
        if err == sql.ErrNoRows {
            logs.Warn("DB NODE -> There is no answer, node id %s doesn't exist",nid)
            return nil, errors.New("DB NODE -> There is no answer, node id "+nid+" doesn't exist")
        }
        if err != nil {
            logs.Warn("DB NODE -> row.Scan error -> %s", err.Error())
            return nil, errors.New("DB NODE -> row.Scan error -> " + err.Error())
        }
        return &node, nil
    } else {
        logs.Info("DB NODE -> No access to database")
        return nil, errors.New("DB NODE -> no access to database")
    }
}
