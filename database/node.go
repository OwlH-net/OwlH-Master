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
            logs.Error("DB NODE -> Error en el prepare, se ha perdido la conn a bd? ")
            return errors.New("DB NODE -> Error en el prepare, se ha perdido la conn a bd?")
        }
        //Validate fields!!!
        res, err := stmt.Exec(node.NName, node.NIp, node.NPort, node.NType, node.NUUID )
        logs.Info("DB -> info", res)
        if err != nil {
            logs.Error("DB NODE -> La Query no ha funcionado bien: %s", err.Error())
            return errors.New("DB NODE -> La Query no ha funcionado bien: " + err.Error())
        }
        return nil
    } else {
        logs.Error("DB NODE -> No tenemos acceso a la base de datos")
        return errors.New("DB NODE -> No tenemos acceso a la base de datos")
    }
}

func GetNode(nid string) (n *Node, err error) {
    logs.Info("DB -> Get Node")
    var node Node
    if Db != nil {
//        rows, err := Db.Query("SELECT * FROM master WHERE master_id=1;")
        row := Db.QueryRow("SELECT * FROM node WHERE node_id=%s;",nid)
        logs.Info ("DB -> Row %s", row)
        err = row.Scan(&node.NId, &node.NName, &node.NIp, &node.NPort, &node.NType, &node.NUUID)
        if err == sql.ErrNoRows {
            logs.Warn("DB NODE -> No encuentro na, ese id %s parece no existir",nid)
            return nil, errors.New("DB NODE -> No encuentro na, ese id "+nid+" parece no existir")
        }
        if err != nil {
            logs.Warn("DB NODE -> no hemos leido bien los campos de scan")
            return nil, errors.New("DB NODE -> no hemos leido bien los campos de scan")
        }
        return &node, nil
    } else {
        logs.Info("DB NODE -> no hay base de datos")
        return nil, errors.New("DB NODE -> no hay base de datos")
    }
}
