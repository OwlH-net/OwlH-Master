package ndb

import (
    "github.com/astaxie/beego/logs"
    "database/sql"
    "os"
    "errors"
    "owlhmaster/utils"
    _ "github.com/mattn/go-sqlite3"
)

var (
    Db *sql.DB
)

func Close() {
    Db.Close()
    Rdb.Close()
    Gdb.Close()
    RSdb.Close()
    Mdb.Close()
}

func Conn() {
    var err error  
    path, err := utils.GetKeyValueString("dbsConn", "path")
    if err != nil {logs.Error("Conn Error getting data from main.conf at master: "+err.Error())}
    cmd, err := utils.GetKeyValueString("dbsConn", "cmd")
    if err != nil {logs.Error("Conn Error getting data from main.conf at master: "+err.Error())}
    
    _, err = os.Stat(path) 
    if err != nil {
        panic("Error: dbs/node DB -- DB Open Failed: "+err.Error())
    }    
    Db, err = sql.Open(cmd,path)
    if err != nil {
        logs.Error("dbs/node DB -- DB Open Failed: "+err.Error())
    }else{
        logs.Info("dbs/node DB -- DB -> sql.Open, DB Ready") 
    }
}

func DeleteNode(uuid string)(err error){
    deleteNodeQuery, err := Db.Prepare("delete from nodes where node_uniqueid = ?;")
    _, err = deleteNodeQuery.Exec(&uuid)
    defer deleteNodeQuery.Close()
    if err != nil {
        logs.Error("DB DeleteRulese/deleteNodeQuery -> ERROR on table Ruleset...")
        return errors.New("DB DeleteRuleset/deleteNodeQuery -> ERROR on table Ruleset...")
    }
    return nil
}

func NodeKeyExists(nodekey string, key string) (id int, err error) {
    if Db == nil {
        logs.Error("no access to database")
        return 0, errors.New("no access to database")
    }
    sql := "SELECT node_id FROM nodes where node_uniqueid = '"+nodekey+"' and node_param = '"+key+"';"
    rows, err := Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return 0, err
    }
    defer rows.Close()
    if rows.Next() {
        if err = rows.Scan(&id); err == nil {
            return id, err
        }
    }
    return 0, nil
}

func InsertNodeKey(nkey string, key string, value string) (err error) {
    if Db == nil {logs.Error("no access to database"); return errors.New("no access to database")}
    stmt, err := Db.Prepare("insert into nodes (node_uniqueid, node_param, node_value) values(?,?,?);")
    if err != nil {logs.Error("InsertNodeKey Prepare -> %s", err.Error()); return err}

    _, err = stmt.Exec(&nkey, &key, &value)
    if err != nil {logs.Error("InsertNodeKey Execute -> %s", err.Error()); return err}
    return nil
}

func UpdateNode(uuid string, param string, value string)(err error){
    updateNode, err := Db.Prepare("update nodes set node_value = ? where node_uniqueid = ? and node_param = ?;")
    if (err != nil){logs.Error("updateNode UPDATE prepare error for update-- "+err.Error()); return err}

    _, err = updateNode.Exec(&value, &uuid, &param)
    defer updateNode.Close()
    if (err != nil){logs.Error("updateNode UPDATE error -- "+err.Error()); return err}
    return nil
}

func GetNodeIpbyName(n string)(ip string, err error) {
    if Db == nil {logs.Error("no access to database"); return "", errors.New("no access to database")}
    sql := "select node_value from nodes where node_uniqueid like '%"+n+"%' and node_param = 'ip';"
    rows, err := Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return "", err
    }
    defer rows.Close()
    if rows.Next() {
        if err = rows.Scan(&ip); err == nil {
            return ip, err
        }
    }
    return "", errors.New("There is no IP for given node name")
}